/*
 * File: licensing.go
 * File Created: Tuesday, 21st June 2022 11:29:56 pm
 * Author: Akhil Datla
 * © 2022 Akhil Datla
 */

package licensing

import (
	"errors"
	"main/components/courses"
	"main/components/dbmanager"

	uuid "github.com/satori/go.uuid"
)

type License struct {
	ID       string `storm:"id"`
	CourseID string `storm:"index"`
}

type Entitlement struct {
	ID         string `storm:"id"`
	CourseID   string `storm:"index"`
	HardwareID string `storm:"index"`
}

func GenerateLicense(courseID string) (string, error) {
	license := &License{
		ID:       uuid.NewV4().String(),
		CourseID: courseID,
	}
	err := dbmanager.Save(license)
	return license.ID, err
}

func RegisterLicense(licenseID, hardwareID string) error {
	var license License
	err := dbmanager.Query("ID", licenseID, &license)
	if err != nil {
		return errors.New("invalid license key")
	}

	entitlement := &Entitlement{
		ID:         uuid.NewV4().String(),
		CourseID:   license.CourseID,
		HardwareID: hardwareID,
	}

	err = dbmanager.Save(entitlement)

	dbmanager.Delete(&license)

	return err
}

func RevokeLicense(courseID, hardwareID string) error {
	entitlement := &Entitlement{
		CourseID:   courseID,
		HardwareID: hardwareID,
	}
	err := dbmanager.Delete(entitlement)
	return err
}

func DownloadCourses(hardwareID string) (string, error) {
	var entitlements []Entitlement
	err := dbmanager.GroupQuery("HardwareID", hardwareID, &entitlements)
	if err != nil {
		return "", err
	}

	courseIDs := make([]string, 0)
	for _, entitlement := range entitlements {
		courseIDs = append(courseIDs, entitlement.CourseID)
	}

	return courses.GenerateWebsite(courseIDs), nil
}
