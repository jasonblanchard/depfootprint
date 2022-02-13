package main

import (
	"jasonblanchard/depfootprint/pkg/npmjs"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/api/tree/:package", func(c echo.Context) error {
		npmjsFetcher := &npmjs.NpmJS{}

		deps, err := npmjsFetcher.Tree("remix")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, deps)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
