/*
 * File: handlers.go
 * File Created: Tuesday, 21st June 2022 11:43:35 pm
 * Author: Akhil Datla
 * © 2022 Akhil Datla
 */

package server

import (
	"encoding/json"
	"main/components/courses"
	"main/components/licensing"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func createCourse(c echo.Context) error {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error parsing request body")
	}
	name := json_map["name"].(string)
	filepath := json_map["filepath"].(string)
	id, err := courses.CreateCourse(name, filepath)

	if err != nil {
		return c.String(http.StatusInternalServerError, "Error creating course")
	}

	return c.JSON(http.StatusOK, map[string]string{"courseID": id})
}

func getCourse(c echo.Context) error {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return err
	}
	id := json_map["id"].(string)
	course, err := courses.GetCourse(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error getting course")
	}
	return c.JSON(http.StatusOK, course)
}

func getAllCourses(c echo.Context) error {
	courses, err := courses.GetAllCourses()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error getting courses")
	}
	return c.JSON(http.StatusOK, courses)
}

func updateCourse(c echo.Context) error {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error parsing request body")
	}
	id := json_map["id"].(string)
	name := json_map["name"].(string)
	filepath := json_map["filepath"].(string)
	err = courses.UpdateCourse(id, name, filepath)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error updating course")
	}
	return c.String(http.StatusOK, "Course updated")
}

func deleteCourse(c echo.Context) error {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error parsing request body")
	}
	id := json_map["id"].(string)
	err = courses.DeleteCourse(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error deleting course")
	}
	return c.String(http.StatusOK, "Course deleted")
}

func generateLicense(c echo.Context) error {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error parsing request body")
	}
	courseID := json_map["courseID"].(string)
	license, err := licensing.GenerateLicense(courseID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error generating license")
	}
	return c.JSON(http.StatusOK, map[string]string{"licenseKey": license})
}

func registerLicense(c echo.Context) error {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error parsing request body")
	}
	licenseKey := json_map["licenseKey"].(string)
	hardwareID := json_map["hardwareID"].(string)
	err = licensing.RegisterLicense(licenseKey, hardwareID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error registering license")
	}
	return c.String(http.StatusOK, "License registered")
}

func revokeLicense(c echo.Context) error {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error parsing request body")
	}
	hardwareID := json_map["hardwareID"].(string)
	courseID := json_map["courseID"].(string)
	err = licensing.RevokeLicense(hardwareID, courseID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error revoking license")
	}
	return c.String(http.StatusOK, "License revoked")
}

func downloadCourses(c echo.Context) error {
	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error parsing request body")
	}
	hardwareID := json_map["hardwareID"].(string)
	file, err := licensing.DownloadCourses(hardwareID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error downloading courses")
	}

	defer os.Remove(file)

	return c.File(file)
}
