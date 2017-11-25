package admin

import (
	"github.com/gin-gonic/gin"
	"io"
	"wasgood/models"
	"wasgood/templates/admintemplates"
)

// UnapprovedReviews renders list of items that has unapproved reviews
func UnapprovedReviews(c *gin.Context) {
	list := models.GetUnapprovedReviews()
	admintemplates.WriteUnapprovedReviews(c.Writer.(io.Writer), c, list)
}
