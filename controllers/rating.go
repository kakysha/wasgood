package controllers

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
	"wasgood/app"
	"wasgood/models"
	"wasgood/templates"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

const (
	itemsPerPage = 20
)

// Index outputs the index page
func Index(c *gin.Context) {
	user := models.GetUserFromContext(c)
	bestItems := models.GetItemsByRating("rating DESC", 10, 0, user, "")
	worstItems := models.GetItemsByRating("rating ASC", 10, 0, user, "")
	bestBrands := models.GetTagsByRating("brand", "rating DESC", 10, 0)
	worstBrands := models.GetTagsByRating("brand", "rating ASC", 10, 0)
	bestFlavors := models.GetTagsByRating("flavors", "rating DESC", 10, 0)
	worstFlavors := models.GetTagsByRating("flavors", "rating ASC", 10, 0)
	c.Status(http.StatusOK)
	templates.WriteIndex(c.Writer.(io.Writer), c, &templates.IndexPage{
		templates.Page{
			Title: "Рейтинг и отзывы на жидкости для электронных сигарет",
		},
		bestItems,
		worstItems,
		bestBrands,
		worstBrands,
		bestFlavors,
		worstFlavors,
	})
}

// Rating outputs paginated rating of items
func Rating(c *gin.Context, sort string) {
	pageStr := c.DefaultQuery("page", "1")
	var page int
	num, err := strconv.Atoi(pageStr)
	if err != nil || num <= 0 {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	page = num
	user := models.GetUserFromContext(c)
	items := models.GetItemsByRating("rating "+sort+", id DESC", itemsPerPage, itemsPerPage*(page-1), user, "")
	if len(items) == 0 {
		c.String(http.StatusNotFound, "404 Not Found")
		return
	}
	templates.WriteRating(c.Writer.(io.Writer), c, &templates.RatingPage{
		templates.Page{
			Title: "Рейтинг жиж",
		},
		items,
		page,
		(len(items) < itemsPerPage),
	})
}

// Vote controller to record user voices on items
func Vote(c *gin.Context) {
	id := c.Param("id")
	voice := c.Param("voice")
	user := models.GetUserFromContext(c)
	item := models.NewItem()
	num, err := strconv.Atoi(id)
	if err != nil || (voice != "1" && voice != "-1") {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	if ok := models.GetItem(num, item); !ok {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	_, err = app.DB.Exec(`INSERT INTO 
		voices(user_id, item_id, voice) 
		VALUES ($1, $2, $3) 
		ON CONFLICT (user_id, item_id) DO UPDATE SET
		voice = excluded.voice, -- concatenate to not lose fields not present in provider response (like 'is_admin')
		vote_time = now()`, user.ID, item.ID, voice)
	check(err)
	c.String(http.StatusOK, "OK")
}
