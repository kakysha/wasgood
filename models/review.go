package models

import (
	"encoding/json"
	"time"
	"wasgood/app"
)

// Review model
type Review struct {
	ID        int
	Author    *User
	Text      string
	Timestamp time.Time
	Item      *Item
	Approved  bool
}

// Save method just insert review into DB
func (r Review) Save() {
	err := app.DB.QueryRow(`INSERT INTO reviews(user_id, item_id, text) VALUES ($1, $2, $3) RETURNING id`, r.Author.ID, r.Item.ID, r.Text).Scan(&r.ID)
	check(err)
}

// GetReviewsForItem fills item's reviews
func GetReviewsForItem(item *Item, user *User) {
	var (
		userID       int
		onlyApproved = true
	)
	if user != nil {
		userID = user.ID
		onlyApproved = !user.IsAdmin
	}
	rows, err := app.DB.Query(`
		SELECT
			reviews.id,
			users.id,
			array_length(enum_range(NULL, users.provider), 1)-1 AS provider,
			users.provider_id,
			users.data,
			reviews.text,
			reviews.approved,
			reviews.timestamp
		FROM reviews
		JOIN users ON reviews.user_id = users.id
		WHERE
			(reviews.approved IN (TRUE, $1) OR reviews.user_id = $2)
			AND
			reviews.item_id = $3
		ORDER BY reviews.timestamp ASC`, onlyApproved, userID, item.ID)
	check(err)
	defer rows.Close()

	for rows.Next() {
		var (
			review = &Review{}
			author = &User{}
			data   []byte
		)

		err = rows.Scan(&review.ID, &author.ID, &author.Provider, &author.ProviderID, &data, &review.Text, &review.Approved, &review.Timestamp)
		check(err)

		err = json.Unmarshal(data, &author)
		check(err)

		review.Author = author

		item.Reviews = append(item.Reviews, review)
	}
	check(rows.Err())
}

// ReviewDelete deletes review by id
func ReviewDelete(id int) {
	_, err := app.DB.Exec(`DELETE FROM reviews WHERE id=$1`, id)
	check(err)
}

// ReviewApprove deletes review by id
func ReviewApprove(id int) {
	_, err := app.DB.Exec(`UPDATE reviews SET approved=TRUE WHERE id=$1`, id)
	check(err)
}

// GetUnapprovedReviews returns unapproved reviews ids grouped by item
func GetUnapprovedReviews() map[int][]int {
	result := make(map[int][]int)
	rows, err := app.DB.Query(`SELECT id, item_id FROM reviews WHERE approved=FALSE`)
	check(err)
	defer rows.Close()

	for rows.Next() {
		var (
			itemID int
			ID     int
		)
		err = rows.Scan(&ID, &itemID)
		check(err)
		if result[itemID] == nil {
			result[itemID] = make([]int, 0, 1)
		}
		result[itemID] = append(result[itemID], ID)
	}
	check(rows.Err())
	return result
}
