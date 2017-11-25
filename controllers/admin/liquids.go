package admin

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"wasgood/controllers/auth"
	"wasgood/models"
	"wasgood/templates/admintemplates"
)

// LiquidGet renders liquid form for item with provided id or empty form to create new item
func LiquidGet(c *gin.Context) {
	item := models.NewLiquid()
	if getItemOfType(c, item) {
		if c.Param("id") != "new" && !auth.CheckACLAccess(c, item) {
			return
		}
		admintemplates.WriteLiquid(c.Writer.(io.Writer), c, item)
	}
}

// LiquidPost updates info about item
func LiquidPost(c *gin.Context) {
	item := models.NewLiquid()
	err := c.Bind(item)
	check(err)
	// check access to old item
	if item.ID != 0 {
		oldItem := models.NewLiquid()
		if ok := models.GetItem(item.ID, oldItem); ok {
			if !auth.CheckACLAccess(c, oldItem) {
				return
			}
		}
	}
	// check access to new item
	if !auth.CheckACLAccess(c, item) {
		return
	}
	models.SaveItem(item, nil)
	c.JSON(http.StatusOK, gin.H{"message": "OK", "item": item})
}
