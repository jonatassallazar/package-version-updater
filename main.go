package main

import (
	"strconv"

	"package-version-updater/config"
	"package-version-updater/core"
	"package-version-updater/logger"
)

func main() {
	logger := logger.Logger{}
	pu := core.PackageUpdater{
		Logger: &logger,
	}
	c := config.Config{}

	c.GetFlags(&pu)
	logger.SetLoggerLevel(c.LogLevel)

	if !pu.IsMajor && !pu.IsMinor && !pu.IsPatch {
		logger.Panic("No version update type selected.")
	}

	logger.Debug("Nested scan: " + strconv.FormatBool((pu.NestedScan)))
	logger.Debug("Package custom name: " + pu.CustomPackageName)

	for _, dir := range c.RootDirs {
		logger.Debug("Root directory: " + dir)
		logger.Log("Starting the package version updater...")

		pu.ScanAllFiles(dir)
	}

	logger.Log("Package version updater finished.")
	logger.BreakLine()
}
