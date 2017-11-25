package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
	"wasgood/models"
	"wasgood/templates"
)

func escapeString(input string) string {
	return fmt.Sprintf("E'%s'", strings.Replace(input, "'", "\\'", -1))
}

// Search searchs the DB for items or brand tags having Name like q
func Search(c *gin.Context) {
	q := c.Query("q")
	ajax := c.DefaultQuery("ajax", "0")
	if utf8.RuneCountInString(q) < 3 {
		c.String(http.StatusBadRequest, "Too short search query")
		return
	}
	pageStr := c.DefaultQuery("page", "1")
	var page int
	num, err := strconv.Atoi(pageStr)
	if err != nil || num <= 0 {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	page = num

	var (
		brands  map[int]*models.Tag
		flavors map[int]*models.Tag
	)
	if page == 1 {
		brands = models.GetTagsByNameAndField(escapeString(q), "brand")
		flavors = models.GetTagsByNameAndField(escapeString(q), "flavors")
	}

	user := models.GetUserFromContext(c)
	items := models.GetItemsByRating("similarity(items.data->>'name', "+escapeString(q)+") DESC, rating DESC", itemsPerPage, itemsPerPage*(page-1), user, "(items.data->>'name') % "+escapeString(q))
	if ajax == "1" {
		templates.WriteAjaxSearchResults(c.Writer.(io.Writer), c, &templates.SearchResultsPage{
			Brands:  brands,
			Flavors: flavors,
			Items:   items,
		})
	} else {
		templates.WriteSearchResults(c.Writer.(io.Writer), c, &templates.SearchResultsPage{
			templates.Page{
				Title: q,
			},
			brands,
			flavors,
			items,
			page,
			(len(items) < itemsPerPage),
		})
	}
}
