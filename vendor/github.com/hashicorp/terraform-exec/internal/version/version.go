package version

const version = "0.18.0"

// ModuleVersion returns the current version of the github.com/hashicorp/terraform-exec Go module.
// This is a function to allow for future possible enhancement using debug.BuildInfo.
func ModuleVersion() string {
	return version
}
