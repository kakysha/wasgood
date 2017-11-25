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

// Liquid renders single liquid page
func Liquid(c *gin.Context) {
	id := c.Param("id")
	id = strings.SplitN(id, "-", 2)[0]
	item := models.NewLiquid()
	user := models.GetUserFromContext(c)
	num, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	if ok := models.GetItem(num, item); !ok || (item.Hidden && (user == nil || !user.IsAdmin)) {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	models.GetRatingForItem(&item.Item, user)
	models.GetReviewsForItem(&item.Item, user)
	relativeItems := models.GetRelativeItemsForItem(&item.Item, user, 5)

	session := sessions.Default(c)
	flashes := session.Flashes()
	session.Save()

	firstImage := ""
	if len(item.Images) > 0 {
		firstImage = item.Images[0]
	}

	templates.WriteLiquid(c.Writer.(io.Writer), c, &templates.LiquidPage{
		templates.Page{
			Title:       "Жидкость " + item.Brand.Name + " " + item.Name + " ОТЗЫВЫ",
			Description: item.Brand.Name + " " + item.Name + " - отзывы, описание жидкость " + item.Name,
			Keywords:    item.Brand.Name + ", " + item.Name,
			OgImage:     firstImage,
			Flashes:     flashes,
		},
		item,
		relativeItems,
	})
}
