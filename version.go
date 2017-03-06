package govcert

import "fmt"

// Version represents the current version of the govcert library
const Version = "0.2.0"

// VCertVersion defines the version of the VCert binary to embed within the library
const VCertVersion = "2.17.1.0"

// GetVersion returns the current version of the govcert library as well as the
func GetVersion() string {
	return fmt.Sprintf("govcert library version: %s\nVCert Binary version: %s", Version, VCertVersion)
}
