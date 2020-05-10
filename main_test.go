package main

import (
	"testing"
)

/**
go test -v
*/

func TestDefaultCrtPath(t *testing.T) {
	var crtPath *string
	crtDefPath := "certs"
	crtPath = &crtDefPath
	_, err := getCrtPath(crtPath)
	if err != nil {
		t.Error(err)
	}
}

func TestDefaultExtPath(t *testing.T) {
	var extPath *string
	crtExtPath := "alt_names.ext"
	extPath = &crtExtPath
	_, err := getExtPath(extPath)
	if err != nil {
		t.Error(err)
	}
}

/**
Generate certs using default params
*/
func TestGenerateCerts(t *testing.T) {
	CERTSPATH = "certs"
	EXTPATH = "alt_names.ext"
	DAYS = 1024
	CRTCOUNTRY = "UA"
	CRTSTATE = "Kyiv"
	CRTCITY = "Kyiv"
	CRTORG = "Localhost-Certificates"
	CRTCOMMONNAME = "localhost.local"

	_, err := runCommands()
	if err != nil {
		t.Error(err)
	}
}
