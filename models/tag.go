package models

import (
	"database/sql"
	"encoding/json"
	"github.com/fiam/gounidecode/unidecode"
	"github.com/gin-gonic/gin"
	"regexp"
	"strings"
	"wasgood/app"
)

// Tag model
type Tag struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Slug        string `json:"slug"` // unique
	Description string `json:"description,omitempty"`
	Logo        string `json:"logo,omitempty"`
	Rating      int    `json:"-"`
	PlusCount   int    `json:"-"`
}

// NewTag constructor
func NewTag() (i *Tag) {
	i = &Tag{}
	return
}

// MarshalJSON converts Tag struct into int ID reference to it
func (t *Tag) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.ID)
}

// UnmarshalJSON assigns Tag ID field
func (t *Tag) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, &t.ID) // int received => id
	if err != nil {
		err := json.Unmarshal(data, &t.Name) // string received => name, create Tag
		if err == nil {
			t.Save(nil)
			return nil
			// will be re-read from DB by ID after
		}
	}
	return err
}

// TagAlias is an alias for Tag type to skip smart UnmarshalJSON on Tag type
type tagAlias Tag

// GetTags tries to obtain Tags for supplied slice of string IDs
func GetTags(IDs []string) map[int]*Tag {
	rows, err := app.DB.Query(`SELECT id, data FROM tags WHERE id IN (` + strings.Join(IDs, ",") + `)`)
	check(err)
	defer rows.Close()

	return parseTagRows(rows)
}

// GetTag tries to obtain Tag with supplied ID
func GetTag(tagID int) *Tag {
	var (
		data []byte
		tag  = NewTag()
	)
	err := app.DB.QueryRow(`SELECT data FROM tags WHERE id = $1`, tagID).Scan(&data)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		check(err)
	}
	tag.ID = tagID
	err = json.Unmarshal(data, (*tagAlias)(tag))
	check(err)
	return tag
}

// GetAllTags retrieves all Tags from DB
func GetAllTags() map[int]*Tag {
	rows, err := app.DB.Query(`SELECT id, data FROM tags ORDER BY id ASC`)
	check(err)
	defer rows.Close()

	return parseTagRows(rows)
}

// GetTagsByNameAndField search for tags having name similar to q and that are present as item's field value
func GetTagsByNameAndField(q string, field string) map[int]*Tag {
	rows, err := app.DB.Query(`
		SELECT tags.id, tags.data
		FROM tags
		JOIN items ON items.data->'` + field + `' @> to_jsonb(tags.id)
		WHERE (tags.data->>'name') % ` + q + `
		ORDER BY similarity(tags.data->>'name', ` + q + `) DESC
	`)
	check(err)
	defer rows.Close()

	return parseTagRows(rows)
}

// GetTagBySlug tries to obtain Tag by slug
func GetTagBySlug(slug string) *Tag {
	var (
		data []byte
		tag  = NewTag()
	)
	err := app.DB.QueryRow(`SELECT id, data FROM tags WHERE data->'slug' @> to_jsonb($1::text)`, slug).Scan(&tag.ID, &data)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		check(err)
	}
	err = json.Unmarshal(data, (*tagAlias)(tag))
	check(err)
	return tag
}

// Save method updates or inserts new Tag into DB, generating slug from name if empty
// forceRemoveFields enumerates which json fields must be erased if they are empty, even if they were not empty before
func (t *Tag) Save(forceRemoveFields []string) {
	if t.Slug == "" {
		t.Slug = strings.ToLower(unidecode.Unidecode(t.Name))
		reg, err := regexp.Compile("[^A-Za-z0-9]+")
		check(err)
		t.Slug = reg.ReplaceAllString(t.Slug, "-")
		t.Slug = strings.ToLower(strings.Trim(t.Slug, "-"))
	}
	data, err := json.Marshal((*tagAlias)(t))
	check(err)

	if t.ID == 0 {
		err = app.DB.QueryRow(`INSERT INTO tags(data) VALUES ($1) ON CONFLICT ((data->>'slug')) DO UPDATE SET id=tags.id RETURNING id, data`, string(data)).Scan(&t.ID, &data)
		err = json.Unmarshal(data, (*tagAlias)(t))
		check(err)
	} else {
		concatenateOldFields := "data"
		for _, field := range forceRemoveFields {
			concatenateOldFields = concatenateOldFields + " - '" + field + "'"
		}
		_, err = app.DB.Exec(`UPDATE tags SET data = `+concatenateOldFields+` || $1::jsonb - 'id' WHERE id=$2`, string(data), t.ID)
	}
	check(err)
}

// Unmarshal unmarshals full tag data from json, skipping smart id unmarshalling
func (t *Tag) Unmarshal(data []byte) error {
	return json.Unmarshal(data, (*tagAlias)(t))
}

// BindFromForm binds json form values to tag struct, skipping smart marshalling
func (t *Tag) BindFromForm(c *gin.Context) error {
	return c.Bind((*tagAlias)(t))
}

// GetAllTagsForField returns map of tags that are referenced at least once by all items 'field' key in whole items table
func GetAllTagsForField(field string) map[int]*Tag {
	rows, err := app.DB.Query(`SELECT
			DISTINCT tags.id,
			tags.data
		FROM tags
		JOIN items ON items.data->'` + field + `' @> to_jsonb(tags.id)`)
	check(err)
	defer rows.Close()

	return parseTagRows(rows)
}

func parseTagRows(rows *sql.Rows) map[int]*Tag {
	tags := make(map[int]*Tag)

	for rows.Next() {
		var (
			tag  = NewTag()
			data []byte
		)
		err := rows.Scan(&tag.ID, &data)
		check(err)

		err = json.Unmarshal(data, (*tagAlias)(tag))
		check(err)

		tags[tag.ID] = tag
	}
	check(rows.Err())

	return tags
}
