package controllers

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
	"strings"
	"wasgood/models"
	"wasgood/templates"
)

// Item renders single item page
func Item(c *gin.Context) {
	id := c.Param("id")
	id = strings.SplitN(id, "-", 2)[0]
	item := models.NewItem()
	user := models.GetUserFromContext(c)
	num, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	if ok := models.GetItem(num, item); !ok || item.Hidden {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	models.GetRatingForItem(item, user)
	models.GetReviewsForItem(item, user)

	session := sessions.Default(c)
	flashes := session.Flashes()
	session.Save()

	templates.WriteItem(c.Writer.(io.Writer), c, &templates.ItemPage{
		templates.Page{
			Title:   item.Brand.Name + " " + item.Name,
			Flashes: flashes,
		},
		item,
	})
}
