package config

import (
	"flag"
	"package-version-updater/core"
	"strings"
)

type Config struct {
	LogLevel uint8
	RootDirs []string
}

func (c *Config) GetFlags(pu *core.PackageUpdater) {
	flag.BoolVar(&pu.NestedScan, "nested", false, "Scan nested directories")
	flag.StringVar(&pu.CustomPackageName, "package", "", "Custom package name")
	flag.BoolVar(&pu.IsMajor, "major", false, "Update major version")
	flag.BoolVar(&pu.IsMinor, "minor", false, "Update minor version")
	flag.BoolVar(&pu.IsPatch, "patch", false, "Update patch version")
	level := flag.Uint("level", 4, "Log level - Default: 4 (INFO)")
	rootDir := flag.String("dir", ".", "Root directory")

	flag.Parse()
	c.LogLevel = uint8(*level)

	if strings.Contains(*rootDir, ",") {
		c.RootDirs = strings.Split(*rootDir, ",")
	} else {
		c.RootDirs = append(c.RootDirs, *rootDir)
	}
}
