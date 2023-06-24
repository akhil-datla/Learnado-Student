/*
 * File: server.go
 * File Created: Tuesday, 13th June 2023 3:19:45 pm
 * Last Modified: Friday, 23rd June 2023 1:03:05 am
 * Author: Akhil Datla
 * Copyright © Akhil Datla 2023
 */

package server

import (
	"fmt"
	"main/components/courses"
	"net/http"
	"path/filepath"

	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
)

var e *echo.Echo

// Start initializes and starts the server on the specified port.
func Start(port int) {
	e = echo.New()
	e.HideBanner = true

	e.Use(middleware.Recover())

	// Configure CORS middleware
	DefaultCORSConfig := middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}

	e.Use(middleware.CORSWithConfig(DefaultCORSConfig))

	initializeRoutes()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

// initializeRoutes sets up the server routes and handlers.
func initializeRoutes() {
	// Custom file server handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path == "/" {
			path = "/index.html"
		}

		// Open the requested file from the in-memory file system
		f, err := courses.AppFs.Open(path)
		if err != nil {
			http.Error(w, "File not found", 404)
			return
		}
		defer f.Close()

		fi, err := f.Stat()
		if err != nil {
			http.Error(w, "File not found", 404)
			return
		}

		// Check if the file is a directory
		if fi.IsDir() {
			f.Close() // Close the directory file
			path = filepath.Join(path, "index.html")

			// Open and serve the index file directly without checking if it's a directory again
			f, err = courses.AppFs.Open(path)
			if err != nil {
				http.Error(w, "File not found", 404)
				return
			}
			defer f.Close()

			fi, err = f.Stat()
			if err != nil {
				http.Error(w, "File not found", 404)
				return
			}
		}

		// Serve the file content
		http.ServeContent(w, r, path, fi.ModTime(), f)
	})

	// Register the custom handler with Echo framework
	e.GET("/*", echo.WrapHandler(handler))

	// Register the license registration endpoint
	e.GET("/register/:key", func(c echo.Context) error {
		key := c.Param("key")

		// Register the license key
		output, err := courses.RegisterLicense(key)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// Download the courses
		err = courses.DownloadCourses()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		// Load the courses
		err = courses.LoadCourses()
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		return c.String(http.StatusOK, output)
	})
}
