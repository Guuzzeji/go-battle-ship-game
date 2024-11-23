package routes

import "github.com/gin-gonic/gin"

func CreateClientGroup(router *gin.Engine) {
	router.LoadHTMLGlob("html/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.GET("/play/:id/p/:playerid", func(c *gin.Context) {
		c.HTML(200, "game.html", nil)
	})

	router.GET("/setup/:id/p/:playerid", func(c *gin.Context) {
		c.HTML(200, "setup.html", nil)
	})

	router.GET("/error", func(c *gin.Context) {
		c.HTML(400, "error.html", nil)
	})

	router.NoRoute(func(c *gin.Context) {
		c.HTML(404, "error.html", nil)
	})
}
