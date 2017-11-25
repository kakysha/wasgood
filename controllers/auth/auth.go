package auth

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"wasgood/app"
	"wasgood/models"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Login stores current user id in session
func Login(u *models.User, c *gin.Context) {
	session := sessions.Default(c)

	session.Set("userID", u.ID)
	session.Save()

	c.Set("User", u) // update in context right now
}

// Logout clears user from session
func Logout(c *gin.Context) {
	session := sessions.Default(c)

	session.Delete("userID")
	session.Save()

	c.Redirect(http.StatusTemporaryRedirect, app.RootURL+c.DefaultQuery("redirect", "/"))
}

// LoginFromSession middleware checks for user_id in session and
// if found fills c.User with that user from DB
func LoginFromSession(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("userID")
	if v != nil {
		userID := v.(int)
		if user, ok := models.GetUser(userID); ok {
			c.Set("User", user)
		}
	}
}

// AdminCheck middleware checks for admin grants on currently logged in user
func AdminCheck(c *gin.Context) {
	if u := models.GetUserFromContext(c); !u.IsAdmin {
		c.AbortWithStatus(http.StatusForbidden)
	}
}

// UserCheck middleware allows only logged in users
func UserCheck(c *gin.Context) {
	if u := models.GetUserFromContext(c); u == nil {
		c.AbortWithStatus(http.StatusForbidden)
	}
}

// ACLPresenceCheck just checks if user has any ACL rules set or is admin, real access check will be checked in controllers
func ACLPresenceCheck(c *gin.Context) {
	if u := models.GetUserFromContext(c); u.ACL == nil {
		AdminCheck(c)
	}
}

// HasACLAccess checks if user has rights to access item
func HasACLRights(c *gin.Context, item models.Model) bool {
	u := models.GetUserFromContext(c)
	if u.IsAdmin {
		return true
	}
	v := reflect.ValueOf(item).Elem()
	t := v.Type()
	for itemType, fields := range u.ACL {
		if t.Name() != "Item" && t.Name() != itemType { // Item is parent type for every other types,
			// can be provided instead of specific type (in case of image upload, for example)
			continue
		}
		for field, ids := range fields {
			fv := v.FieldByName(field)
			if !fv.IsValid() {
				continue
			}
			if field == "ID" {
				for _, id := range ids {
					if v.FieldByName("ID").Int() == id {
						return true
					}
				}
			}
			if fv.Type().String() == "*models.Tag" {
				for _, id := range ids {
					if fv.Elem().FieldByName("ID").Int() == id {
						return true
					}
				}
			}
			if fv.Type().String() == "[]*models.Tag" {
				var tag reflect.Value
				for j := 0; j < fv.Len(); j++ {
					tag = fv.Index(j)
					for _, id := range ids {
						if tag.Elem().FieldByName("ID").Int() == id {
							return true
						}
					}
				}
			}
		}
	}
	return false
}

// CheckACLAccess allows or denies user access to item based on ACL rights
func CheckACLAccess(c *gin.Context, item models.Model) bool {
	if HasACLRights(c, item) {
		return true
	}
	c.AbortWithStatus(http.StatusForbidden)
	return false
}
