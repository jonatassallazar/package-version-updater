package core

import (
	"os"
	"package-version-updater/logger"
	"strings"
)

type PackageUpdater struct {
	NestedScan        bool
	CustomPackageName string
	IsMajor           bool
	IsMinor           bool
	IsPatch           bool
	Logger            *logger.Logger
}

func (p *PackageUpdater) ReadDir(dir string) []os.DirEntry {
	p.Logger.Debug("Reading directory: " + dir)

	// read os directory
	files, err := os.ReadDir(dir)
	if err != nil {
		p.Logger.Fatal(err.Error())
	}

	return files
}

func (p *PackageUpdater) ScanAllFiles(dir string) {
	// read all files in the directory
	files := p.ReadDir(dir)

	for _, file := range files {
		fileName := file.Name()
		path := dir + fileName + "/"

		// update the package.json file
		if fileName == "package.json" || (p.CustomPackageName != "" && fileName == p.CustomPackageName) {
			p.Logger.Log("Updating package version: " + dir + fileName)

			p.UpdatePackageVersion(dir + fileName)

			p.Logger.Log("Package version updated.")
		}

		// if the file is a directory, scan it (support for nested directories)
		if file.IsDir() && !strings.Contains(path, "node_modules") && p.NestedScan {
			p.Logger.Debug("Scanning nested directory: " + path)

			p.ScanAllFiles(path)
		}
	}
}

// Update the package version
func (p *PackageUpdater) UpdatePackageVersion(filePath string) {
	vu := VersionUpdater{
		FilePath: filePath,
	}

	vu.ReadFile()
	vu.ExtractVersionsInt()
	vu.UpdateVersion(p)
	vu.UpdatePackageBytes()
	vu.WritePackageFile()

}
