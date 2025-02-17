package main

import (
	"flag"
	"package-version-updater/core"
	"package-version-updater/logger"
	"strconv"
)

func main() {
	logger := logger.Logger{}

	pu := core.PackageUpdater{
		Logger: &logger,
	}

	flag.BoolVar(&pu.NestedScan, "nested", false, "Scan nested directories")
	flag.StringVar(&pu.CustomPackageName, "package", "", "Custom package name")
	flag.BoolVar(&pu.IsMajor, "major", false, "Update major version")
	flag.BoolVar(&pu.IsMinor, "minor", false, "Update minor version")
	flag.BoolVar(&pu.IsPatch, "patch", false, "Update patch version")
	level := flag.Uint("level", 4, "Log level - Default: 4 (INFO)")
	rootDir := flag.String("dir", "./", "Root directory")

	flag.Parse()
	logger.Level = uint8(*level)

	if !pu.IsMajor && !pu.IsMinor && !pu.IsPatch {
		logger.Panic("No version update type selected.")
	}

	logger.Debug("Nested scan: " + strconv.FormatBool((pu.NestedScan)))
	logger.Debug("Root directory: " + *rootDir)
	logger.Debug("Package custom name: " + pu.CustomPackageName)

	logger.Log("Starting the package version updater...")

	pu.ScanAllFiles(*rootDir)

	logger.Log("Package version updater finished.")
	logger.BreakLine()
}
