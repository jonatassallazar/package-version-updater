# Package Version Updater

This program allows you to update the package versions of multiple JavaScript/TypeScript projects simultaneously. It scans the specified directories for `package.json` files and updates the version numbers according to the provided criteria. This is particularly useful for maintaining consistency across multiple projects or repositories.

Features:
- Batch update package versions
- Customizable version increment (major, minor, patch)
- Generates a summary log report of updated packages

Usage:
1. Specify the root directory containing the projects.
2. Choose the version increment type (major, minor, patch).
3. Run the updater to apply the changes.

Example:
```bash
# Run the updater with minor version increment
./package-version-updater.exe -minor --dir ./project1 -nested 
```

## Installation

Instructions on how to clone the project and compile locally.

```bash
# Clone the repository
git clone https://github.com/jonatassallazar/package-version-updater

# Navigate to the project directory
cd package-version-updater

# Install dependencies
go mod tidy

# Compile
go build

# Or run directly
go run main.go
```

## Usage

Run from command line the compiled software.

```bash
# Run the program
./package-version-updater.exe
```

## Flags

- __`-nested`__: Scan nested directories for `package.json` files. Default is `false`.
- __`-package`__: Specify a custom package name to update. Default is "package.json".
- __`-major`__: Update the major version of the packages. Default is `false`. *Priority order 1
- __`-minor`__: Update the minor version of the packages. Default is `false`. *Priority order 2
- __`-patch`__: Update the patch version of the packages. Default is `false`. *Priority order 3
- __`-level`__: Set the log level for the output. Default is `4` (INFO).
- - __5__ - DEBUG
- - __4__ - INFO
- - __3__ - WARN
- - __2__ - ERROR
- - __1__ - FATAL
- - __0__ - PANIC
- __`-dir`__: Specify the root directory or multiple directories to scan for `package.json` files. Separate multiple paths with commas. Default is `.`.  
  __Example__: `-dir ./project1,./project2,../project3` or `-dir ./root_folder -nested`

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.