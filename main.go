package main

import (
	"fmt"
	"os"

	"github.com/aldinokemal/go-whatsapp-web-multidevice/cmd"
)

// @title			Go WhatsApp Web Multi Device API
// @version		2.0
// @description	This is a Go implementation of WhatsApp Web Multi Device API
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	MIT
// @license.url	https://opensource.org/licenses/MIT
// @host			localhost:3000
// @BasePath		/
func main() {
	if err := cmd.Execute(); err != nil {
		// Print the error with a clearer prefix and exit with a non-zero status code
		// Note: using exit code 2 to distinguish application errors from OS-level errors (exit code 1)
		fmt.Fprintf(os.Stderr, "fatal error: %v\n", err)
		os.Exit(2)
	}
}
