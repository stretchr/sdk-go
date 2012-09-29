package main

import (
	"../stretchr"
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

/*
	To build this tool, run the 'build.sh' file, or use 'go bulid'.

	We recommend specifying the "-o stretchr" flag so the executable file generated
	is easy to use.
*/

const (
	PROJ_NAME_ENV string = "StretchrProjectName"
	PUB_KEY_ENV   string = "StretchrPublicKey"
	PRIV_KEY_ENV  string = "StretchrPrivateKey"
	HOST_ENV      string = "StretchrHost"
)

var (
	NoErrors          error = nil
	MissingProject    error = errors.New(fmt.Sprintf("Unknown project, set '%s' envorinment variable or specify -project flag.", PROJ_NAME_ENV))
	MissingPublicKey  error = errors.New(fmt.Sprintf("Unknown public key, set '%s' envorinment variable", PUB_KEY_ENV))
	MissingPrivateKey error = errors.New(fmt.Sprintf("Unknown private key, set '%s' envorinment variable", PRIV_KEY_ENV))
)

/*
	Flags
*/
var help = flag.Bool("help", false, "Whether to print out this help or not.")
var verbose = flag.Bool("verbose", false, "Whether to be verbose or not.")
var method = flag.String("method", "GET", "The HTTP method to use in this request (GET, POST, PUT or DELETE).")
var path = flag.String("path", "", "The path of the request to make.")
var host = flag.String("host", "stretchr.com", "The host of the service to connect to.")
var projectFlag = flag.String("project", "", fmt.Sprintf("The name of the Stretchr project to connect to.  Alternativly set the %s environment variable.", PROJ_NAME_ENV))
var publicKeyFlag = flag.String("key", "", fmt.Sprintf("The public key used to access the services.  Alternativly set the %s environment variable.", PUB_KEY_ENV))
var privateKeyFlag = flag.String("privatekey", "", fmt.Sprintf("The private key used when signing requests for security.  Alternativly set the %s environment variable.", PRIV_KEY_ENV))
var query = flag.String("query", "", "The query part of the request to modify behaviour.")
var json = flag.String("json", "", "The JSON body to send in the request.")

/*
	Stretchr variable holds the reference to a stretchr.Session object
	from which all Stretchr actions are performed.
*/
var Stretchr *stretchr.Session

func main() {

	flag.Parse()

	if *help || len(os.Args) < 2 {
		print("Usage for %s:", os.Args[0])
		flag.PrintDefaults()
		return
	}

	if err := prepareSession(); err != nil {
		printError(err)
		return
	}

	if *verbose {
		fmt.Println("Stretchr command line tool - v0.2 - by Mat Ryer")
		fmt.Println("Verbose mode on")
		fmt.Println("")
	}

	/*
		Print session details
	*/
	if *verbose {
		print("Session")
		print("--------------------------------------")
		print("    Protocol: %s", Stretchr.Protocol)
		print("     Version: %s", Stretchr.Version)
		print("        Host: %s", Stretchr.Host)
		print("Project name: %s", Stretchr.Project)
		print("  Public key: %s", Stretchr.PublicKey)
		print(" Private key: %s", trunc(Stretchr.PrivateKey, 20))
		print("")
	}

	/*
		Prepare the request
	*/
	fullUrl := fmt.Sprintf("%s?%s", Stretchr.Url(*path), *query)
	u, urlParseErr := url.ParseRequestURI(fullUrl)
	if urlParseErr != nil {
		print("Failed to parse URL: %s", urlParseErr)
		os.Exit(1)
	}
	fullUrl = u.String()

	// add the public key
	var fullUrlWithKey string = fullUrl

	if strings.Contains(fullUrl, "?") {
		fullUrlWithKey = fmt.Sprintf("%s&~key=%s", fullUrl, Stretchr.PublicKey)
	} else {
		fullUrlWithKey = fmt.Sprintf("%s?~key=%s", fullUrl, Stretchr.PublicKey)
	}

	signedUrl, signingErr := stretchr.GetSignedURL(*method, fullUrlWithKey, *json, Stretchr.PrivateKey)

	if signingErr != nil {
		print("Failed to sign request: %s", signingErr)
		os.Exit(1)
	}

	if *verbose {
		print("Request")
		print("--------------------------------------")
		print("    Method: %s", *method)
		print("      Path: %s", *path)
		print("     Query: %s", *query)
		print("  Full URL: %s", fullUrl)
		print(" Keyed URL: %s", fullUrlWithKey)
		print("Signed URL: %s", signedUrl)
		print("")

		print("Request Body")
		print("--------------------------------------")
		print(*json)
		print("")
	}

	// make the request
	sro, response, requestErr := stretchr.ActiveRequester.MakeRequest(*method, fullUrl, *json, Stretchr.PublicKey, Stretchr.PrivateKey)

	if requestErr != nil {
		print("Error when making the request: %s", requestErr)
		os.Exit(1)
	}

	if *verbose {

		print("")
		print("")

		print("Response")
		print("--------------------------------------")
		print("HTTP Status code: %d", sro.StatusCode)
		print("")

		print("Response Headers")
		print("--------------------------------------")
		for k, vs := range response.Header {
			for _, v := range vs {
				print("%s: %s", k, v)
			}
		}

		print("")

		print("Response Body")
		print("--------------------------------------")
		print(sro.ResponseBody)

		print("")
		print("")
		print("")

	}

	// print the request
	print("%s %s", *method, signedUrl)
	print("")
	print("Status: %d", sro.StatusCode)
	print("%s", sro.ResponseBody)

	// print a blank line
	fmt.Println("")

}

func printError(err error) {
	fmt.Println(fmt.Sprintf("%s", err))
}

func prepareSession() error {

	project := def(def(*projectFlag, os.Getenv(PROJ_NAME_ENV)), "test")
	publicKey := def(*publicKeyFlag, os.Getenv(PUB_KEY_ENV))
	privateKey := def(*privateKeyFlag, os.Getenv(PRIV_KEY_ENV))

	if !assertNotBlank(project) {
		return MissingProject
	}
	if !assertNotBlank(publicKey) {
		return MissingPublicKey
	}
	if !assertNotBlank(privateKey) {
		return MissingPrivateKey
	}

	host := def(*host, def(os.Getenv(HOST_ENV), "stretchr.com"))

	Stretchr = stretchr.InProject(project).WithKeys(publicKey, privateKey)
	Stretchr.Host = host

	return NoErrors

}
func print(f string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(f, args...))
}
func trunc(f string, l int) string {
	if len(f) > l {
		return fmt.Sprintf("%s...", f[0:l])
	}
	return f
}
func def(s, def string) string {
	if len(s) == 0 {
		return def
	}
	return s
}
func assertNotBlank(s string) bool {
	return len(s) != 0
}
