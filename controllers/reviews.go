package controllers

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wasgood/app"
	"wasgood/models"
)

// PostReview saves review to DB
func PostReview(c *gin.Context) {
	id := c.Param("id")
	text := c.PostForm("text")
	user := models.GetUserFromContext(c)
	num, err := strconv.Atoi(id)
	if err != nil || text == "" {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	item := models.NewItem()
	if ok := models.GetItem(num, item); !ok {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	review := models.Review{Author: user, Item: item, Text: text}
	review.Save()
	session := sessions.Default(c)
	session.AddFlash("Обзор добавлен. Он появится после одобрения модератором.")
	session.Save()
	c.Redirect(http.StatusSeeOther, app.RootURL+c.DefaultQuery("redirect", "/"))
}

// ReviewDelete deletes review by id
func ReviewDelete(c *gin.Context) {
	id := c.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	models.ReviewDelete(num)
	c.String(http.StatusOK, "OK")
}

// ReviewApprove approves review by id
func ReviewApprove(c *gin.Context) {
	id := c.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	models.ReviewApprove(num)
	c.String(http.StatusOK, "OK")
}
