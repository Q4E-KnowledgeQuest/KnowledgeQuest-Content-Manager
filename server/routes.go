/*
 * File: routes.go
 * File Created: Tuesday, 21st June 2022 11:43:32 pm
 * Author: Akhil Datla
 * © 2022 Akhil Datla
 */

package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
)

var e *echo.Echo

func Start(port int, log bool) {
	e = echo.New()
	e.HideBanner = true

	if log {
		e.Use(middleware.Logger())
	}
	e.Use(middleware.Recover())

	DefaultCORSConfig := middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}

	e.Use(middleware.CORSWithConfig(DefaultCORSConfig))

	initializeRoutes()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

func initializeRoutes() {
	e.POST("/courses/create", createCourse)
	e.GET("/courses/info", getCourse)
	e.GET("/courses/all", getAllCourses)
	e.POST("/courses/update", updateCourse)
	e.DELETE("/courses/delete", deleteCourse)
	e.POST("/licenses/create", generateLicense)
	e.POST("/licenses/register", registerLicense)
	e.DELETE("/licenses/revoke", revokeLicense)
	e.POST("/download", downloadCourses)

}
