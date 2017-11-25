package admin

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
	"wasgood/models"
	"wasgood/templates/admintemplates"
)

// TagGet renders form for tag with id
func TagGet(c *gin.Context) {
	id := c.Param("id")
	num, _ := strconv.Atoi(id)
	tag := models.GetTag(num)
	if tag == nil {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	admintemplates.WriteTag(c.Writer.(io.Writer), c, tag)
}

// TagPost updates tag
func TagPost(c *gin.Context) {
	tag := models.NewTag()
	err := tag.BindFromForm(c)
	check(err)
	tag.Save(nil)
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

// TagsGet renders list of all tags
func TagsGet(c *gin.Context) {
	tags := models.GetAllTags()
	admintemplates.WriteTags(c.Writer.(io.Writer), c, tags)
}

// TagUploadImages handler receives tag :id, ?field and multipart form with one image
func TagUploadImages(c *gin.Context) {
	id := c.Param("id")
	fieldName := c.Query("field")
	num, _ := strconv.Atoi(id)
	tag := models.GetTag(num)
	if tag == nil {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	var field *string
	switch fieldName {
	case "logo":
		field = &tag.Logo
	}
	if field == nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	*field = uploadImages(c, "tags")[0]
	tag.Save(nil)
	c.String(http.StatusOK, "OK")
}

// TagRemoveImage unlinks file from disk by :src and removes it from tag ?field
func TagRemoveImage(c *gin.Context) {
	id := c.Param("id")
	fieldName := c.Query("field")
	num, _ := strconv.Atoi(id)
	tag := models.GetTag(num)
	if tag == nil {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	var field *string
	switch fieldName {
	case "logo":
		field = &tag.Logo
	}
	if field == nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}
	removeImage(c, []string{*field})
	*field = ""
	tag.Save([]string{fieldName})
	c.String(http.StatusOK, "OK")
}
