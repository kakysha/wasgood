package main

import (
	"database/sql"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"wasgood/app"
	"wasgood/controllers"
	"wasgood/controllers/admin"
	"wasgood/controllers/auth"
)

func main() {
	// Logger + Recovery, public routes
	r := gin.Default()
	// Cookie Sessions
	store := sessions.NewCookieStore([]byte("возгудsecret"))
	r.Use(sessions.Sessions("session", store))
	r.Use(auth.LoginFromSession)
	// logged users only route group
	userroutes := r.Group("/", auth.UserCheck)
	// admin route group
	adminroutes := userroutes.Group("/admin", auth.AdminCheck)
	aclroutes := userroutes.Group("/admin", auth.ACLPresenceCheck)

	// Site routes
	r.GET("/", controllers.Index)
	r.GET("/auth/vk/login", auth.VKLogin)
	r.GET("/auth/vk/callback", auth.VKCallback)
	r.GET("/auth/fb/login", auth.FBLogin)
	r.GET("/auth/fb/callback", auth.FBCallback)
	userroutes.GET("/auth/logout", auth.Logout)
	r.GET("/rating/good", func(c *gin.Context) { controllers.Rating(c, "DESC") })
	r.GET("/rating/bad", func(c *gin.Context) { controllers.Rating(c, "ASC") })
	r.GET("/liquids/:id", controllers.Liquid)
	r.GET("/items/:id", controllers.Item)
	userroutes.POST("/items/:id/vote/:voice", controllers.Vote)
	userroutes.POST("/items/:id/review", controllers.PostReview)
	adminroutes.DELETE("/reviews/:id", controllers.ReviewDelete)
	adminroutes.POST("/reviews/:id/approve", controllers.ReviewApprove)
	r.GET("/brand/:slug", func(c *gin.Context) { controllers.Tag(c, "brand") })
	r.GET("/flavor/:slug", func(c *gin.Context) { controllers.Tag(c, "flavors") })
	r.GET("/search", controllers.Search)

	// Admin panel routes
	aclroutes.GET("/", func(c *gin.Context) { c.Redirect(http.StatusMovedPermanently, "/admin/liquids/1") })
	adminroutes.GET("/items/:id", admin.ItemGet) // id can be "new"
	adminroutes.POST("/items/:id", admin.ItemPost)
	aclroutes.POST("/items/:id/images", admin.ItemUploadImages)
	aclroutes.DELETE("/items/:id/images/*src", admin.ItemRemoveImage)
	aclroutes.POST("/items/:id/images/sort", admin.ItemSortImages)

	aclroutes.GET("/liquids/:id", admin.LiquidGet) // id can be "new"
	aclroutes.POST("/liquids/:id", admin.LiquidPost)

	adminroutes.GET("/reviews", admin.UnapprovedReviews)

	adminroutes.GET("/tags/:id", admin.TagGet)
	adminroutes.POST("/tags", admin.TagPost)
	adminroutes.GET("/tags", admin.TagsGet)
	adminroutes.POST("/tags/:id/images", admin.TagUploadImages)
	adminroutes.DELETE("/tags/:id/images/*src", admin.TagRemoveImage)

	var err error
	app.DB, err = sql.Open("postgres", "host=/var/run/postgresql user=wasgood dbname=wasgood sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = app.DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer app.DB.Close()

	r.Run("localhost:6354")
}
