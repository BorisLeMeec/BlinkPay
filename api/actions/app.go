package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gobuffalo/buffalo/middleware/ssl"
	"github.com/gobuffalo/envy"
	"github.com/unrolled/secure"

	"github.com/gobuffalo/x/sessions"
	"github.com/rs/cors"
	//"gopkg.in/yaml.v2"
	//"github.com/gobuffalo/buffalo/binding"
	//"net/http"
	//"io/ioutil"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")
var app *buffalo.App

//func init() {
//	binding.Register("text/yaml", func(req *http.Request, model interface{}) error {
//		b, err := ioutil.ReadAll(req.Body)
//		if err != nil {
//			return err
//		}
//		return yaml.Unmarshal(b, model)
//	})
//}

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.New(buffalo.Options{
			Env:          ENV,
			SessionStore: sessions.Null{},
			PreWares: []buffalo.PreWare{
				cors.Default().Handler,
			},
			SessionName: "_api_session",
			LooseSlash: true,
		})
		// Automatically redirect to SSL
		app.Use(ssl.ForceSSL(secure.Options{
			SSLRedirect:     ENV == "production",
			SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
		}))

		// Set the request content type to JSON
		app.Use(middleware.SetContentType("application/json"))

		if ENV == "development" {
			app.Use(middleware.ParameterLogger)
		}

		app.GET("/", HomeHandler)
		g := app.Group("/api/v1")

		// new UserResource
		ur := &UserResource{}

		g.GET("/test", TestURL)
		g.GET("/users", ur.List)
		g.POST("/users", ur.Create)
		g.GET("/users/{id}", ur.Show)
		g.POST("/users/check", ur.Check)

	}

	return app
}

func TestURL(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{"message": "Test successful!"}))
}