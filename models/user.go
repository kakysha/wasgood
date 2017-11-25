package models

import (
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"wasgood/app"
)

// ProviderType is oAuth provider name
type ProviderType uint8

const (
	// VK .com oAuth
	VK ProviderType = iota
	// FB .com oAuth
	FB
)

// User model
type User struct {
	ID         int                           `json:"-"`
	Provider   ProviderType                  `json:"-"`
	ProviderID int                           `json:"-"`
	Name       string                        `json:"name"`
	LastName   string                        `json:"last_name"`
	Photo      string                        `json:"photo"`
	Email      string                        `json:"email,omitempty"`
	IsAdmin    bool                          `json:"is_admin,omitempty"`
	ACL        map[string]map[string][]int64 `json:"acl,omitempty"` // {"Liquid": {"Brand": [1,3], "ID": [1111]}}
}

// GetUserFromContext helper returns pointer to User model stored in context if any, or nil otherwise
func GetUserFromContext(c *gin.Context) (u *User) {
	if v, ok := c.Get("User"); ok {
		u = v.(*User)
	}
	return
}

// Save updates or creates User (if pair provider + provider_id not existed before) filling it's ID field in DB
func (u *User) Save() {
	data, err := json.Marshal(u)
	check(err)
	err = app.DB.QueryRow(`INSERT INTO 
		users(provider, provider_id, data) 
		VALUES ((SELECT (enum_range(null::authprovider))[$1]), $2, $3) 
		ON CONFLICT (provider, provider_id) DO UPDATE SET
		data = users.data || excluded.data, -- concatenate to not lose fields not present in provider response (like 'is_admin')
		last_login = now()
		RETURNING id`, u.Provider+1, u.ProviderID, string(data)).Scan(&u.ID)
	check(err)
}

// GetUser tries to obtain User from DB by its id
func GetUser(userID int) (*User, bool) {
	var (
		user = &User{}
		data []byte
	)
	err := app.DB.QueryRow(`SELECT
		id, array_length(enum_range(NULL, provider), 1)-1, provider_id, data
		FROM users
		WHERE id = $1`, userID).Scan(&user.ID, &user.Provider, &user.ProviderID, &data)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false
		}
		check(err)
	}
	err = json.Unmarshal(data, &user)
	check(err)
	return user, true
}
