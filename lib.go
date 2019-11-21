package boldchat

import (
	"fmt"
)

// Constants that defines self-awareness
// LIBNAME    is the name
// MAJVERSION is the major version number
// MINVERSION is the minor version number
// RELVERSION is the release version number
const (
	LIBNAME    = "BoldChat"
	MAJVERSION = 0
	MINVERSION = 0
	RELVERSION = 0
)

// Version returns a string that contains
// the name and version all nice and neat
func Version() string {
	return fmt.Sprintf("%s v%d.%d.%d", LIBNAME,
		MAJVERSION, MINVERSION, RELVERSION)
}
