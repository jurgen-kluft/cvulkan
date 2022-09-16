package cvulkan

import (
	"github.com/jurgen-kluft/ccode/denv"
)

// GetPackage returns the package object of 'cvulkan'
func GetPackage() *denv.Package {

	// The main package
	mainpkg := denv.NewPackage("cvulkan")

	// 'cvulkan' library
	mainlib := denv.SetupDefaultCppLibProject("cvulkan", "github.com\\jurgen-kluft\\cvulkan")

	mainpkg.AddMainLib(mainlib)

	if denv.OS == "windows" {
		//mainlib.AddDefine("_GLFW_WIN32;_GLFW_WGL;WIN32")
	} else if denv.OS == "darwin" {
		//mainlib.AddDefine("_GLFW_COCOA;MACOSX")
	} else if denv.OS == "linux" {
		//mainlib.AddDefine("_GLFW_X11;_GLFW_GFX;LINUX")
	}

	return mainpkg
}
