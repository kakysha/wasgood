package models

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/fiam/gounidecode/unidecode"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"wasgood/app"
)

// Model interface to handle db stores and pulls of any struct type that has ID
type Model interface {
	getID() int
	setID(int)
}

// Item model
type Item struct {
	ID          int       `json:"id,omitempty"`
	Name        string    `json:"name"`
	Brand       *Tag      `json:"brand"`
	Description string    `json:"description"`
	Images      []string  `json:"images,omitempty"`
	Hidden      bool      `json:"hidden"`
	Rating      int       `json:"-"`
	PlusCount   int       `json:"-"`
	UserVoice   int       `json:"-"`
	Reviews     []*Review `json:"-"`
	slug        string
}

// NewItem constructor, initializes all pointers
func NewItem() (i *Item) {
	i = &Item{Brand: NewTag()}
	return
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func (i *Item) getID() int {
	return i.ID
}

func (i *Item) setID(id int) {
	i.ID = id
}

// Slug method returns slug generated from item's Name
func (i *Item) Slug() string {
	if i.slug == "" {
		i.slug = strings.ToLower(unidecode.Unidecode(i.Name))
		reg, err := regexp.Compile("[^A-Za-z0-9]+")
		check(err)
		i.slug = reg.ReplaceAllString(i.slug, "-")
		i.slug = strings.ToLower(strings.Trim(i.slug, "-"))
	}
	return i.slug
}

// GetItem tries to obtain Item or any descendant type from DB by its id, joining Tags to fill reference struct fields
func GetItem(itemID int, m Model) bool {
	var data []byte
	err := app.DB.QueryRow(`SELECT data FROM items WHERE id = $1`, itemID).Scan(&data)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		check(err)
	}
	err = json.Unmarshal(data, m)
	check(err)
	m.setID(itemID)

	// reflection to get all struct fields of type Tag and query DB for them, assigning to the struct after
	// at this point, all Tag fields already have ID set up (using smart unmarshalling of Tag type), so just query DB for them
	var tagIDs = make([]string, 0, 2)
	reflectAndHandleTags(m, func(f reflect.Value) {
		tagIDs = append(tagIDs, strconv.FormatInt(f.Elem().FieldByName("ID").Int(), 10))
	})
	tags := GetTags(tagIDs)
	reflectAndHandleTags(m, func(f reflect.Value) {
		id := f.Elem().FieldByName("ID").Int()
		f.Set(reflect.ValueOf(tags[int(id)]))
	})
	return true
}

func reflectAndHandleTags(m Model, handler func(reflect.Value)) {
	v := reflect.ValueOf(m)
	t := reflect.TypeOf(m)
	for i := 0; i < v.Elem().NumField(); i++ {
		fv := v.Elem().Field(i)
		ft := t.Elem().Field(i)
		//log.Println(f.Type().String())
		switch {
		case fv.Kind() == reflect.Struct && ft.Anonymous: //  "models.Item"
			reflectAndHandleTags(fv.Addr().Interface().(Model), handler)
		case fv.Type().String() == "*models.Tag":
			handler(fv)
		case fv.Type().String() == "[]*models.Tag":
			var tag reflect.Value
			for j := 0; j < fv.Len(); j++ {
				tag = fv.Index(j)
				handler(tag)
			}
		}
	}
}

// GetRatingForItem fills item's rating fields
func GetRatingForItem(item *Item, user *User) {
	var (
		query bytes.Buffer
		err   error
	)
	query.WriteString(`
		SELECT
			COALESCE(SUM(all_voices.voice), 0) AS rating,
			COALESCE(COUNT(CASE WHEN all_voices.voice > 0 THEN 1 END), 0) AS plus_count`)
	if user != nil {
		query.WriteString(`,
			COALESCE(MAX(user_voices.voice), 0) AS user_voice`)
	}
	query.WriteString(`
		FROM voices AS all_voices`)
	if user != nil {
		query.WriteString(`
		LEFT JOIN voices AS user_voices ON all_voices.item_id = user_voices.item_id AND user_voices.user_id = $2`)
	}
	query.WriteString(`
		WHERE all_voices.item_id = $1`)
	if user != nil {
		err = app.DB.QueryRow(query.String(), item.ID, user.ID).Scan(&item.Rating, &item.PlusCount, &item.UserVoice)
	} else {
		err = app.DB.QueryRow(query.String(), item.ID).Scan(&item.Rating, &item.PlusCount)
	}
	check(err)
}

// GetRelativeItemsForItem returns n items of the same brand
func GetRelativeItemsForItem(item *Item, user *User, n int) []*Item {
	var (
		query bytes.Buffer
		err   error
		rows  *sql.Rows
	)
	// WARNING: OUTER QUERY MUST NOT MULTIPLY NUMBER OF ITEM ROWs, otherwise SUM of voices will be multiplied too
	query.WriteString(`
		SELECT
			items.id,
			items.data,
			brands.data,
			COALESCE(SUM(all_voices.voice), 0) AS rating,
			COALESCE(COUNT(CASE WHEN all_voices.voice > 0 THEN 1 END), 0) AS plus_count,
			COALESCE(reviews.reviews_count, 0) AS reviews_count`)
	if user != nil {
		query.WriteString(`,
			COALESCE(MAX(user_voices.voice), 0) AS user_voice`) // MAX to get any value inside the group, they all are the same
	}
	query.WriteString(`
		FROM items
		JOIN tags AS brands ON items.data->'brand' @> to_jsonb(brands.id)
		LEFT JOIN voices AS all_voices ON items.id = all_voices.item_id
		LEFT JOIN (SELECT item_id, COUNT(id) AS reviews_count FROM reviews GROUP BY item_id) AS reviews ON items.id = reviews.item_id`)
	if user != nil {
		query.WriteString(`
		LEFT JOIN voices AS user_voices ON items.id = user_voices.item_id AND user_voices.user_id = $4`)
	}
	query.WriteString(`
		WHERE NOT items.data->'hidden' @> 'true'
		AND NOT items.id = $2
		AND items.data->'brand' @> $1`)
	query.WriteString(`
		GROUP BY items.id, reviews.reviews_count, brands.id
		ORDER BY rating DESC
		LIMIT $3`)

	if user != nil {
		rows, err = app.DB.Query(query.String(), item.Brand.ID, item.ID, n, user.ID)
	} else {
		rows, err = app.DB.Query(query.String(), item.Brand.ID, item.ID, n)
	}
	check(err)
	defer rows.Close()

	var items []*Item

	for rows.Next() {
		var (
			item         = NewItem()
			brand        = NewTag()
			itemData     []byte
			brandData    []byte
			reviewsCount int
		)
		if user != nil {
			err = rows.Scan(&item.ID, &itemData, &brandData, &item.Rating, &item.PlusCount, &reviewsCount, &item.UserVoice)
		} else {
			err = rows.Scan(&item.ID, &itemData, &brandData, &item.Rating, &item.PlusCount, &reviewsCount)
		}
		check(err)

		item.Reviews = make([]*Review, reviewsCount) // transmit only count (by slice length), not reviews themselves

		err = json.Unmarshal(itemData, item)
		check(err)
		err = brand.Unmarshal(brandData)
		check(err)
		item.Brand = brand

		items = append(items, item)
	}
	check(rows.Err())

	return items
}

// SaveItem updates or creates (if ID is zero) Item (or any type that fits into 'items' table) in DB
// forceRemoveFields enumerates which json fields must be erased if they are empty, even if they were not empty before
func SaveItem(i Model, forceRemoveFields []string) error {
	data, err := json.Marshal(i)
	var id int
	check(err)
	if i.getID() == 0 {
		err = app.DB.QueryRow(`INSERT INTO items(data) VALUES ($1) ON CONFLICT DO NOTHING RETURNING id`, string(data)).Scan(&id)
		if err != nil { // sql.ErrNoRows = conflict, can return constraint error
			return err
		}
		i.setID(id)
	} else {
		// concatenate to not lose images, as they are not sent on item post
		// and remove 'id' from data previously bound by form
		concatenateOldFields := "data"
		for _, field := range forceRemoveFields {
			concatenateOldFields = concatenateOldFields + " - '" + field + "'"
		}
		_, err = app.DB.Exec(`UPDATE items SET data = `+concatenateOldFields+` || $1::jsonb - 'id' WHERE id=$2`, string(data), i.getID())
	}
	check(err)
	return nil
}

// GetItemsByRating returns slice of Items, ordered by their total rating, paginated
// If user != nil, then also acquire his voices on selected items
func GetItemsByRating(sort string, limit int, offset int, user *User, condition string) []*Item {
	var (
		query bytes.Buffer
		err   error
		rows  *sql.Rows
	)
	// WARNING: OUTER QUERY MUST NOT MULTIPLY NUMBER OF ITEM ROWs, otherwise SUM of voices will be multiplied too
	query.WriteString(`
		SELECT
			items.id,
			items.data,
			brands.data,
			COALESCE(SUM(all_voices.voice), 0) AS rating,
			COALESCE(COUNT(CASE WHEN all_voices.voice > 0 THEN 1 END), 0) AS plus_count,
			COALESCE(reviews.reviews_count, 0) AS reviews_count`)
	if user != nil {
		query.WriteString(`,
			COALESCE(MAX(user_voices.voice), 0) AS user_voice`) // MAX to get any value inside the group, they all are the same
	}
	query.WriteString(`
		FROM items
		JOIN tags AS brands ON items.data->'brand' @> to_jsonb(brands.id)
		LEFT JOIN voices AS all_voices ON items.id = all_voices.item_id
		LEFT JOIN (SELECT item_id, COUNT(id) AS reviews_count FROM reviews GROUP BY item_id) AS reviews ON items.id = reviews.item_id`)
	if user != nil {
		query.WriteString(`
		LEFT JOIN voices AS user_voices ON items.id = user_voices.item_id AND user_voices.user_id = $3`)
	}
	query.WriteString(`
		WHERE 1=1`)
	if user == nil || !user.IsAdmin {
		query.WriteString(`
			AND NOT items.data->'hidden' @> 'true'`)
	}
	if condition != "" {
		query.WriteString(`
			AND ` + condition)
	}
	query.WriteString(`
		GROUP BY items.id, reviews.reviews_count, brands.id
		ORDER BY ` + sort + `
		LIMIT $1 OFFSET $2`)

	params := []interface{}{limit, offset}
	if user != nil {
		params = append(params, user.ID)
	}
	rows, err = app.DB.Query(query.String(), params...)
	check(err)
	defer rows.Close()

	var items []*Item

	for rows.Next() {
		var (
			item         = NewItem()
			brand        = NewTag()
			itemData     []byte
			brandData    []byte
			reviewsCount int
		)
		if user != nil {
			err = rows.Scan(&item.ID, &itemData, &brandData, &item.Rating, &item.PlusCount, &reviewsCount, &item.UserVoice)
		} else {
			err = rows.Scan(&item.ID, &itemData, &brandData, &item.Rating, &item.PlusCount, &reviewsCount)
		}
		check(err)

		item.Reviews = make([]*Review, reviewsCount) // transmit only count (by slice length), not reviews themselves

		err = json.Unmarshal(itemData, item)
		check(err)
		err = brand.Unmarshal(brandData)
		check(err)
		item.Brand = brand

		items = append(items, item)
	}
	check(rows.Err())

	return items
}

// GetTagsByRating returns slice of tags, ordered by their total rating, paginated
func GetTagsByRating(field string, sort string, limit int, offset int) []*Tag {
	var (
		err  error
		rows *sql.Rows
	)
	rows, err = app.DB.Query(`
		SELECT
			tags.data,
			COALESCE(SUM(all_voices.voice), 0) AS rating,
			COALESCE(COUNT(CASE WHEN all_voices.voice > 0 THEN 1 END), 0) AS plus_count
		FROM items
		JOIN tags ON items.data->'`+field+`' @> to_jsonb(tags.id)
		LEFT JOIN voices AS all_voices ON items.id = all_voices.item_id
		WHERE NOT items.data->'hidden' @> 'true'
		GROUP BY tags.id
		ORDER BY `+sort+`
		LIMIT $1 OFFSET $2`, limit, offset)
	check(err)
	defer rows.Close()

	var tags []*Tag

	for rows.Next() {
		var (
			tag     = NewTag()
			tagData []byte
		)
		err = rows.Scan(&tagData, &tag.Rating, &tag.PlusCount)
		check(err)

		err = tag.Unmarshal(tagData)
		check(err)

		tags = append(tags, tag)
	}
	check(rows.Err())

	return tags
}

// GetAllValuesForField returns slice of distinct values in data column for 'field' key in whole items table
func GetAllValuesForField(field string) []string {
	rows, err := app.DB.Query(`SELECT DISTINCT data->'` + field + `' FROM items`)
	check(err)
	defer rows.Close()

	return parseStringRows(rows)
}

// GetAllValuesForArray returns slice of distinct values in data column for 'field' array key in whole items table
func GetAllValuesForArray(field string) []string {
	rows, err := app.DB.Query(`SELECT DISTINCT jsonb_array_elements_text(data->'` + field + `') FROM items`)
	check(err)
	defer rows.Close()

	return parseStringRows(rows)
}

func parseStringRows(rows *sql.Rows) []string {
	values := make([]string, 0, 8)

	for rows.Next() {
		var value *string

		err := rows.Scan(&value)
		check(err)

		if value == nil {
			continue
		}

		values = append(values, *value)
	}
	check(rows.Err())
	return values
}
