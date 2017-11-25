package controllers

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
	"wasgood/models"
	"wasgood/templates"
)

// Tag renders list of items which have provided tag in a field
func Tag(c *gin.Context, field string) {
	pageStr := c.DefaultQuery("page", "1")
	var page int
	num, err := strconv.Atoi(pageStr)
	if err != nil || num <= 0 {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	page = num

	user := models.GetUserFromContext(c)
	slug := c.Param("slug")
	tag := models.GetTagBySlug(slug)
	if tag == nil {
		c.String(http.StatusNotFound, "Not Found")
		return
	}

	items := models.GetItemsByRating("rating DESC", itemsPerPage, itemsPerPage*(page-1), user, "items.data->'"+field+"' @> to_jsonb("+strconv.Itoa(tag.ID)+")")
	var (
		title       string
		description string
		keywords    string
	)
	switch field {
	case "brand":
		title = tag.Name + " - отзывы на жидкости " + tag.Name + " для электронных сигарет"
		description = tag.Name + " жидкость - " + tag.Name + " отзывы на всю линейку жидкостей от " + tag.Name
		keywords = tag.Name
	case "flavors":
		title = "Жидкость для электронных сигарет со вкусом " + tag.Name
	}
	templates.WriteTag(c.Writer.(io.Writer), c, &templates.TagPage{
		templates.Page{
			Title:       title,
			Description: description,
			Keywords:    keywords,
		},
		tag,
		field,
		items,
		page,
		(len(items) < itemsPerPage),
	})
}
