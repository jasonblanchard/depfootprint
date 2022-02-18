package main

import (
	"jasonblanchard/depfootprint/pkg/npmjs"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/api/tree/:package", func(c echo.Context) error {
		npmjsFetcher := &npmjs.NpmJS{}
		pkg := c.Param("package")

		deps, err := npmjsFetcher.Tree(pkg)
		if err != nil {
			e.Logger.Error(err)
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, deps)
	})
	e.Logger.Fatal(e.Start(":1323"))
}