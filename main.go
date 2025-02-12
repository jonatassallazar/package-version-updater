package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	ScanAllFiles("./", "")
}

func ReadDir(dir string) []os.DirEntry {
	// read os directory
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	return files
}

func ScanAllFiles(path, prefix string) {
	// read all files in the directory
	files := ReadDir(path)

	for _, file := range files {
		fileName := file.Name()
		dirPath := path + fileName + "/"

		// update the package.json file
		if fileName == "package.json" {
			UpdatePackageVersion(path + fileName)
		}

		// if the file is a directory, scan it (support for nested directories)
		if file.IsDir() && !strings.Contains(dirPath, "node_modules") {
			ScanAllFiles(dirPath, "  ")
		}
	}
}

func UpdatePackageVersion(filePath string) {
	var versionStr string
	var IsMajor bool
	var IsMinor bool
	var isPatch = true

	// Read the package.json file
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Get the current version
	exp := regexp.MustCompile(`"version": "(\d+).(\d+).(\d+)"`)

	versionBytes := exp.Find(data)
	versionStr = string(versionBytes)

	versionExpr := regexp.MustCompile(`(\d+)`)
	versions := versionExpr.FindAllString(versionStr, -1)

	major, err := strconv.ParseInt(versions[0], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	minor, err := strconv.ParseInt(versions[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	patch, err := strconv.ParseInt(versions[2], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	// Update the version according to the type of update
	if IsMajor {
		major++
	}

	if IsMinor {
		minor++
	}

	if isPatch {
		patch++
	}

	// Update the version string
	versionStr = fmt.Sprintf(`"version": "%d.%d.%d"`, major, minor, patch)

	// Update package.json text with the new version
	newPackageUpdated := exp.ReplaceAllString(string(data), versionStr)

	// Write the package.json file
	err = os.WriteFile(filePath, []byte(newPackageUpdated), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}
