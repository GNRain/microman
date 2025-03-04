package applogic

import (
	"os"

	"github.com/gorilla/mux"
)

// Settings struct holds settings for Api
type Settings struct {
	// Title of the api instance
	Title string
	// Hostname of the api server (default: localhost)
	Hostname string
	// Port that the server will listen on
	Port string
	// Version of the Api
	Version string
	// Secret is the secret key of the app
}

// Template is a struct that holds the default settings of the Api
type Template struct {
	// Default title of Api instance
	Title string
	// Default hostname of Api server (default: localhost)
	Hostname string
	// Default port that server will listen on
	Port string
	// Default version of Api
	Version string
}

// DefaultTemplate is the default template of the Api settings
func DefaultTemplate() Template {
	return Template{
		Title:    "Micro Dude Api",
		Hostname: "localhost",
		Port:     "6969",
		Version:  "0.420.69",
	}
}

// CheckTitle is a function that checks if the title is set from .env
func CheckTitle(api *Api, template string) {
	// Check the variables from env and set them to the settings
	// if they are not set, set them to the default values
	// get title from env
	api.Settings.Title = os.Getenv("TITLE")
	if api.Settings.Title == "" {
		api.Settings.Title = template
	}
}

// CheckHostname is a function that checks if the hostname is set from .env
func CheckHostname(api *Api, template string) {
	// Check the variables from env and set them to the settings
	// if they are not set, set them to the default values
	// get hostname from env
	api.Settings.Hostname = os.Getenv("HOSTNAME")
	if api.Settings.Hostname == "" {
		api.Settings.Hostname = template
	}
}

// CheckPort is a function that checks if the port is set from .env
func CheckPort(api *Api, template string) {
	api.Settings.Port = os.Getenv("PORT")
	if api.Settings.Port == "" {
		api.Settings.Port = template
	}
}

// CheckVersion is a function that checks if the version is set from .env
func CheckVersion(api *Api, template string) {
	api.Settings.Version = os.Getenv("VERSION")
	if api.Settings.Version == "" {
		api.Settings.Version = template
	}
}

// CheckSettings is a function that checks if the settings are set from .env
func CheckSettings(api *Api) {
	CheckTitle(api, DefaultTemplate().Title)
	CheckHostname(api, DefaultTemplate().Hostname)
	CheckPort(api, DefaultTemplate().Port)
	CheckVersion(api, DefaultTemplate().Version)
}

// SetPort is a function that sets the port of the Api
func SetPort(api *Api, port string) {
	api.Settings.Port = port
}

// GetPort is a function that returns the port of the Api
func GetPort(api *Api) string {
	return api.Settings.Port
}

// ShowPort is a function that returns the router of the Api
func ShowPort(api *Api) {
	println(api.Settings.Port)
}

// SetTitle is a function that sets the title of the Api
func SetTitle(api *Api, title string) {
	api.Settings.Title = title
}

// GetTitle is a function that returns the title of the Api
func GetTitle(api *Api) string {
	return api.Settings.Title
}

// ShowTitle is a function that prints the title of the Api
func ShowTitle(api *Api) {
	println(api.Settings.Title)
}

// SetHostname is a function that sets the hostname of the Api
func SetHostname(api *Api, hostname string) {
	api.Settings.Hostname = hostname
}

// GetHostname is a function that returns the hostname of the Api
func GetHostname(api *Api) string {
	return api.Settings.Hostname
}

// ShowHostname is a function that prints the hostname of the Api
func ShowHostname(api *Api) {
	println(api.Settings.Hostname)
}

// SetVersion is a function that sets the version of the Api
func SetVersion(api *Api, version string) {
	api.Settings.Version = version
}

// GetVersion is a function that returns the version of the Api
func GetVersion(api *Api) string {
	return api.Settings.Version
}

// ShowVersion is a function that prints the version of Api
func ShowVersion(api *Api) {
	println(api.Settings.Version)
}

// GetSettings is a function that returns the settings of the Api
func GetSettings(api *Api) Settings {
	return api.Settings
}

// GetFullPath is a function that returns the full path of the Api
func GetFullPath(api *Api) string {
	return GetHostname(api) + ":" + GetPort(api)
}

// GetRnP is a function that returns port and router of the Api for http serving
func GetRnP(api *Api) (string, *mux.Router) {
	return ":" + GetPort(api), GetRouter(api)
}
