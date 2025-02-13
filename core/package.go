package core

import (
	"log"
	"os"
	"regexp"
	"strings"
)

type PackageUpdater struct {
	NestedScan        bool
	CustomPackageName string
	IsMajor           bool
	IsMinor           bool
	IsPatch           bool
}

type VersionUpdater struct {
	FilePath      string
	Data          []byte
	UpdatedData   string
	Expression    *regexp.Regexp
	VersionString string
	Versions      []string
	Major         int64
	Minor         int64
	Patch         int64
}

func (p *PackageUpdater) ReadDir(dir string) []os.DirEntry {
	// read os directory
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
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
			p.UpdatePackageVersion(dir + fileName)
		}

		// if the file is a directory, scan it (support for nested directories)
		if file.IsDir() && !strings.Contains(path, "node_modules") && p.NestedScan {
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
