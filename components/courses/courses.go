/*
 * File: courses.go
 * File Created: Tuesday, 21st June 2022 7:46:58 pm
 * Author: Akhil Datla
 * © 2022 Akhil Datla
 */

package courses

import (
	"main/components/dbmanager"
	"os"
	"os/exec"
	"path/filepath"

	cp "github.com/otiai10/copy"
	uuid "github.com/satori/go.uuid"
)

type Course struct {
	ID       string `storm:"id"`
	Name     string `storm:"unique"`
	Filepath string
}

func CreateCourse(name, filepath string) (string, error) {
	course := &Course{
		ID:       uuid.NewV4().String(),
		Name:     name,
		Filepath: filepath,
	}
	err := dbmanager.Save(course)
	return course.ID, err
}

func GetCourse(id string) (Course, error) {
	var course Course
	err := dbmanager.Query("ID", id, &course)
	return course, err
}

func GetAllCourses() ([]Course, error) {
	var courses []Course
	err := dbmanager.QueryAll(&courses)
	return courses, err
}

func UpdateCourse(id, name, filepath string) error {
	course := &Course{
		ID:       id,
		Name:     name,
		Filepath: filepath,
	}
	err := dbmanager.Update(course)
	return err
}

func DeleteCourse(id string) error {
	course := &Course{
		ID: id,
	}
	err := dbmanager.Delete(course)
	return err
}

func GenerateWebsite(courseIDs []string) string {
	filePaths := make([]string, 0)
	courseNames := make([]string, 0)

	for _, id := range courseIDs {
		course, _ := GetCourse(id)
		filePaths = append(filePaths, course.Filepath)
		courseNames = append(courseNames, course.Name)
	}

	tempHugoDir, _ := os.MkdirTemp("", "")

	cp.Copy(filepath.Join(filepath.Dir(""), "hugo"), tempHugoDir)

	for i, filePath := range filePaths {

		os.MkdirAll(filepath.Join(tempHugoDir, "content", courseNames[i]), 0755)

		cp.Copy(filePath, filepath.Join(tempHugoDir, "content", courseNames[i]))

	}

	indexFile, _ := os.OpenFile(filepath.Join(tempHugoDir, "content", "_index.md"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)

	indexFile.WriteString(indexMDContent)

	buildCmd := exec.Command("hugo")
	buildCmd.Dir = tempHugoDir
	buildCmd.Run()

	buildZipName := uuid.NewV4().String() + ".zip"

	zipCmd := exec.Command("zip", "-r", buildZipName, "public")
	zipCmd.Dir = tempHugoDir
	zipCmd.Run()

	cp.Copy(filepath.Join(tempHugoDir, buildZipName), filepath.Join(filepath.Dir(""), buildZipName))

	os.RemoveAll(tempHugoDir)

	return buildZipName

}

var indexMDContent = 
`
---
title: "Quest 4 Excellence Learning Platform"
---

# Quest 4 Excellence Learning Platform

This learning platform is designed to help you engage with the Quest 4 Excellence cirriculum.

`
