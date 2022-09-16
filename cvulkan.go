package main

import (
	"github.com/jurgen-kluft/ccode"
	"github.com/jurgen-kluft/cvulkan/package"
)

func main() {
	ccode.Init()
	ccode.Generate(cvulkan.GetPackage())
}
