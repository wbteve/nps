package version

const VERSION = "0.26.10-private.1"

// Compulsory minimum version, Minimum downward compatibility to this version
func GetVersion() string {
	return "0.26.0"
}
