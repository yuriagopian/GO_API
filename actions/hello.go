package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// HelloHandler is a default handler to serve up
// a hello page.
func HelloHandler(c buffalo.Context) error {
	//return c.Render(http.StatusOK, r.JSON(map[string]string{"message": "Hello Full Cycle"}))
	return c.Render(http.StatusOK, r.String("Hello Full Cycle"))
}
