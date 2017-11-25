package admin

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
	"wasgood/controllers/auth"
	"wasgood/models"
	"wasgood/templates/admintemplates"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getItemOfType(c *gin.Context, item models.Model) bool {
	id := c.Param("id")
	if id != "new" {
		num, _ := strconv.Atoi(id)
		if ok := models.GetItem(num, item); !ok {
			c.String(http.StatusNotFound, "Not Found")
			return false
		}
	}
	return true
}

// ItemGet renders form for item with id or empty form to create new item
func ItemGet(c *gin.Context) {
	item := models.NewItem()
	if getItemOfType(c, item) {
		admintemplates.WriteItem(c.Writer.(io.Writer), c, item)
	}
}

// ItemPost updates info about item
func ItemPost(c *gin.Context) {
	item := models.NewItem()
	err := c.Bind(item)
	check(err)
	models.SaveItem(item, nil)
	c.JSON(http.StatusOK, gin.H{"message": "OK", "item": item})
}

// ItemUploadImages handler receives item :id and multipart form with multiple photos
func ItemUploadImages(c *gin.Context) {
	id := c.Param("id")
	item := models.NewItem()
	num, _ := strconv.Atoi(id)
	if ok := models.GetItem(num, item); !ok {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	if !auth.CheckACLAccess(c, item) {
		return
	}
	item.Images = append(item.Images, uploadImages(c, "items")...)
	models.SaveItem(item, nil)
	c.String(http.StatusOK, "OK")
}

// ItemRemoveImage unlinks file from disk by :src and removes it from Item.Images array
func ItemRemoveImage(c *gin.Context) {
	id := c.Param("id")
	item := models.NewItem()
	num, _ := strconv.Atoi(id)
	if ok := models.GetItem(num, item); !ok {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	if !auth.CheckACLAccess(c, item) {
		return
	}
	item.Images = removeImage(c, item.Images)
	models.SaveItem(item, []string{"images"})
	c.String(http.StatusOK, "OK")
}

// ItemSortImages stores received array of sorted src attributes in Item.Images
func ItemSortImages(c *gin.Context) {
	id := c.Param("id")
	item := models.NewItem()
	num, _ := strconv.Atoi(id)
	if ok := models.GetItem(num, item); !ok {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	if !auth.CheckACLAccess(c, item) {
		return
	}
	err := c.Bind(item)
	check(err)
	models.SaveItem(item, nil)
	c.String(http.StatusOK, "OK")
}
