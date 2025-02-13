package main

import (
	"flag"
	"package-version-updater/core"
	"package-version-updater/logger"
	"strconv"
)

func main() {
	logger := logger.Logger{}

	logger.Log("Starting the package version updater...")

	pu := core.PackageUpdater{}

	flag.BoolVar(&pu.NestedScan, "nested", false, "Scan nested directories")
	flag.StringVar(&pu.CustomPackageName, "package", "", "Custom package name")
	flag.BoolVar(&pu.IsMajor, "major", false, "Update major version")
	flag.BoolVar(&pu.IsMinor, "minor", false, "Update minor version")
	flag.BoolVar(&pu.IsPatch, "patch", false, "Update patch version")
	rootDir := flag.String("dir", "./", "Root directory")

	flag.Parse()

	logger.Verbose("Nested scan: " + strconv.FormatBool((pu.NestedScan)))
	logger.Log("Scanning all files in the directory...")
	logger.Info("Root directory: " + *rootDir)

	pu.ScanAllFiles(*rootDir)

	logger.Log("Package version updater finished.")
	logger.BreakLine()
}
