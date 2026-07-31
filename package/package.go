package cvulkan

import (
	"github.com/jurgen-kluft/gide/denv"
)

// GetPackage returns the package object of 'cvulkan'
func GetPackage() *denv.Package {

	// The main package
	mainpkg := denv.NewPackage("github.com\\jurgen-kluft", "cvulkan")

	// 'cvulkan' library
	mainlib := denv.SetupCppLibProject(mainpkg, "cvulkan")

	mainpkg.AddMainLib(mainlib)

	return mainpkg
}
