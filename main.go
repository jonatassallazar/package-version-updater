package main

import (
	"flag"
	"package-version-updater/core"
)

func main() {
	pu := core.PackageUpdater{}

	flag.BoolVar(&pu.NestedScan, "nested", false, "Scan nested directories")
	flag.StringVar(&pu.CustomPackageName, "package", "", "Custom package name")
	flag.BoolVar(&pu.IsMajor, "major", false, "Update major version")
	flag.BoolVar(&pu.IsMinor, "minor", false, "Update minor version")
	flag.BoolVar(&pu.IsPatch, "patch", false, "Update patch version")
	rootDir := flag.String("dir", "./", "Root directory")

	flag.Parse()

	pu.ScanAllFiles(*rootDir)
}
