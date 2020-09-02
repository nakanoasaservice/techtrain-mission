/*
 * TechTrain MISSION Game API
 *
 * TechTrain MISSION ゲームAPI入門仕様  まずはこのAPI仕様に沿って機能を実装してみましょう。    この画面の各APIの「Try it out」->「Execute」を利用することで  ローカル環境で起動中のAPIにAPIリクエストをすることができます。
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()

	router.Use(newCorsMiddleware())
	router.Use(newHandleErrorMiddleware())

	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

func newCorsMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowHeaders = append(config.AllowHeaders, "x-token")
	config.AllowAllOrigins = true
	return cors.New(config)
}

func newHandleErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors.Errors() {
			log.Print(err)
		}
	}
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		http.MethodGet,
		"/",
		Index,
	},

	{
		"CharacterListGet",
		http.MethodGet,
		"/character/list",
		CharacterListGet,
	},

	{
		"GachaDrawPost",
		http.MethodPost,
		"/gacha/draw",
		GachaDrawPost,
	},

	{
		"UserCreatePost",
		http.MethodPost,
		"/user/create",
		UserCreatePost,
	},

	{
		"UserGetGet",
		http.MethodGet,
		"/user/get",
		UserGetGet,
	},

	{
		"UserUpdatePut",
		http.MethodPut,
		"/user/update",
		UserUpdatePut,
	},
}
