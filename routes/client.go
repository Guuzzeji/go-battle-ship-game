package routes

import "github.com/gin-gonic/gin"

// CreateClientGroup is a function that takes a gin router and
// adds all the main routes to the client. The routes added are:
// - GET /: Renders the index page.
// - GET /play/:id/p/:playerid: Renders the game page.
// - GET /setup/:id/p/:playerid: Renders the setup page.
// - GET /error: Renders the error page.
// - NoRoute: Handles 404 errors by rendering the error page.
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
