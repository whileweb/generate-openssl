package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var CERTSPATH string
var EXTPATH string
var DAYS uint
var CRTCOUNTRY string
var CRTSTATE string
var CRTCITY string
var CRTORG string
var CRTCOMMONNAME string

func main() {
	crtPath := flag.String("crtpath", "certs", "a string")
	extPath := flag.String("extpath", "alt_names.ext", "a string")
	days := flag.Uint("days", 1024, "a uint16")
	crtCountry := flag.String("C", "UA", "a string")
	crtState := flag.String("ST", "Kyiv", "a string")
	crtCity := flag.String("L", "Kyiv", "a string")
	crtOrg := flag.String("O", "Localhost-Certificates", "a string")
	crtCommonName := flag.String("CN", "localhost.local", "a string")

	flag.Parse()
	var err error

	CERTSPATH, err = getCrtPath(crtPath)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error:%v\n\n", err)
		os.Exit(1)
	}

	EXTPATH, err = getExtPath(extPath)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error:%v\n\n", err)
		os.Exit(1)
	}

	DAYS = *days
	CRTCOUNTRY = *crtCountry
	CRTSTATE = *crtState
	CRTCITY = *crtCity
	CRTORG = *crtOrg
	CRTCOMMONNAME = *crtCommonName

	_, err = runCommands()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error:%v\n\n", err)
		os.Exit(1)
	}
	fmt.Printf("OpenSSL certificates have been generated successfully [%s]\n", CERTSPATH)
}

/**
Return absolute path to target directory
Will be used `certs` as default directory in current folder
*/
func getCrtPath(crtPath *string) (string, error) {
	if *crtPath == "certs" {
		// will be created into current directory
		err := os.MkdirAll(*crtPath, os.ModePerm)
		return "certs", err
	} else {
		_, err := os.Stat(*crtPath)
		os.IsNotExist(err)
		return *crtPath, err
	}
}

/**
Return absolute path to target directory
Will be used `alt_names.ext` as default filename in current directory
*/
func getExtPath(extPath *string) (string, error) {
	if *extPath == "" {
		// get alt_names.ext from current directory
		_, err := os.Stat("alt_names.ext")
		os.IsNotExist(err)
		return "alt_names.ext", err
	} else {
		_, err := os.Stat(*extPath)
		os.IsNotExist(err)
		return *extPath, err
	}
}

func runCommands() (bool, error) {
	_, err := exec.Command("openssl",
		"req",
		"-x509",
		"-nodes",
		"-new",
		"-sha256",
		"-days",
		strconv.Itoa(int(DAYS)),
		"-newkey",
		"rsa:2048",
		"-keyout",
		strings.Replace("$CERTSPATH/RootCA.key", "$CERTSPATH", CERTSPATH, 1),
		"-out",
		strings.Replace("$CERTSPATH/RootCA.pem", "$CERTSPATH", CERTSPATH, 1),
		"-subj",
		"/C=UA/CN=Localhost-Root-CA",
	).Output()

	if err != nil {
		return false, err
	}

	_, err = exec.Command("openssl",
		"x509",
		"-outform",
		"pem",
		"-in",
		strings.Replace("$CERTSPATH/RootCA.pem", "$CERTSPATH", CERTSPATH, 1),
		"-out",
		strings.Replace("$CERTSPATH/RootCA.crt", "$CERTSPATH", CERTSPATH, 1),
	).Output()

	if err != nil {
		return false, err
	}

	_, err = exec.Command("openssl",
		"req",
		"-new",
		"-nodes",
		"-newkey",
		"rsa:2048",
		"-keyout",
		strings.Replace("$CERTSPATH/localhost.key", "$CERTSPATH", CERTSPATH, 1),
		"-out",
		strings.Replace("$CERTSPATH/localhost.csr", "$CERTSPATH", CERTSPATH, 1),
		"-subj",
		fmt.Sprintf("/C=%s/ST=%s/L=%s/O=%s/CN=%s", CRTCOUNTRY, CRTSTATE, CRTCITY, CRTORG, CRTCOMMONNAME),
	).Output()

	if err != nil {
		return false, err
	}

	_, err = exec.Command("openssl",
		"x509",
		"-req",
		"-sha256",
		"-days",
		strconv.Itoa(int(DAYS)),
		"-in",
		strings.Replace("$CERTSPATH/localhost.csr", "$CERTSPATH", CERTSPATH, 1),
		"-CA",
		strings.Replace("$CERTSPATH/RootCA.pem", "$CERTSPATH", CERTSPATH, 1),
		"-CAkey",
		strings.Replace("$CERTSPATH/RootCA.key", "$CERTSPATH", CERTSPATH, 1),
		"-CAcreateserial",
		"-extfile",
		EXTPATH,
		"-out",
		strings.Replace("$CERTSPATH/localhost.crt", "$CERTSPATH", CERTSPATH, 1),
	).Output()

	if err != nil {
		return false, err
	}

	return true, nil
}
