package main

import (
	"log"

	"microman/applogic"
)

// main is the entry point for the microguy application.
func main() {
    // Creating a new api instance
    api := applogic.DefaultAPIClient()

    // Starting the server
    log.Println("Initializing " + applogic.GetTitle(api) + "...")
    applogic.RunDefaultClient(api)

}
