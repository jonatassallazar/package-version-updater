package core

import (
	"fmt"
	"os"
	"package-version-updater/logger"
	"regexp"
	"strconv"
)

type VersionUpdater struct {
	Logger        *logger.Logger
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

// Read the file and return the data in bytes
func (vu *VersionUpdater) ReadFile() {
	data, err := os.ReadFile(vu.FilePath)
	if err != nil {
		vu.Logger.Fatal(err.Error())
	}

	vu.Data = data
}

// Get the current version of the bytes data
func (vu *VersionUpdater) GetVersion() {
	vu.Expression = regexp.MustCompile(`"version": "(\d+).(\d+).(\d+)"`)

	versionBytes := vu.Expression.Find(vu.Data)

	versionStr := string(versionBytes)
	vu.VersionString = versionStr

	versionExpr := regexp.MustCompile(`(\d+)`)
	versions := versionExpr.FindAllString(versionStr, -1)

	vu.Versions = versions
}

// Extract the version numbers from the version string
func (vu *VersionUpdater) ExtractVersionsInt() {
	major, err := strconv.ParseInt(vu.Versions[0], 10, 64)
	if err != nil {
		vu.Logger.Fatal(err.Error())
	}

	minor, err := strconv.ParseInt(vu.Versions[1], 10, 64)
	if err != nil {
		vu.Logger.Fatal(err.Error())
	}

	patch, err := strconv.ParseInt(vu.Versions[2], 10, 64)
	if err != nil {
		vu.Logger.Fatal(err.Error())
	}

	vu.Major = major
	vu.Minor = minor
	vu.Patch = patch
}

func (vu *VersionUpdater) UpdateMajorVersion() {
	vu.Major++
	vu.Minor = 0
	vu.Patch = 0
}

func (vu *VersionUpdater) UpdateMinorVersion() {
	vu.Minor++
	vu.Patch = 0
}

func (vu *VersionUpdater) UpdatePatchVersion() {
	vu.Patch++
}

// Update the version according to the type of update
func (vu *VersionUpdater) UpdateVersion(p *PackageUpdater) {
	if p.IsMajor {
		vu.UpdateMajorVersion()
	} else if p.IsMinor {
		vu.UpdateMinorVersion()
	} else if p.IsPatch {
		vu.UpdatePatchVersion()
	}

	// Update the version string
	vu.VersionString = fmt.Sprintf(`"version": "%d.%d.%d"`, vu.Major, vu.Minor, vu.Patch)
}

// Update the package.json file bytes
func (vu *VersionUpdater) UpdatePackageBytes() {
	vu.UpdatedData = vu.Expression.ReplaceAllString(string(vu.Data), vu.VersionString)
}

// Write the package.json file
func (vu *VersionUpdater) WritePackageFile() {
	err := os.WriteFile(vu.FilePath, []byte(vu.UpdatedData), os.ModePerm)
	if err != nil {
		vu.Logger.Fatal(err.Error())
	}
}
