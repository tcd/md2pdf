package md2pdf

import "fmt"

// VersionInfo about a build of md2pdf.
type VersionInfo struct {
	Major string
	Minor string
	Patch string
	Hash  string
	Date  string
}

func (v VersionInfo) String() string {
	if v.Major == "" {
		return "md2pdf dev"
	}
	return fmt.Sprintf("md2pdf v%s.%s.%s (%s %s)", v.Major, v.Minor, v.Patch, v.Hash, v.Date)
}

var (
	// VersionMajor is set during build
	VersionMajor string
	// VersionMinor is set during build
	VersionMinor string
	// VersionPatch is set during build
	VersionPatch string
	// VersionHash is set during build
	VersionHash string
	// VersionDate is set during build
	VersionDate string
)

// Version is for use in the version subcommand.
var Version = VersionInfo{
	Major: VersionMajor,
	Minor: VersionMinor,
	Patch: VersionPatch,
	Hash:  VersionHash,
	Date:  VersionDate,
}
