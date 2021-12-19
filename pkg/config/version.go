package config

import (
	"fmt"
	"time"
)

// Variables that are populated during build process
var appVersion = "?" // contains git tag as version
var buildTime = "?"  // when the program was built
var commit = "?"     // commit hash, getting from git rev-parse

// Version Utility structure
type Version struct {
	App         string `json:"app"`          // Application.
	AppVersion  string `json:"app_version"`  // Application version.
	BuildTime   string `json:"build_time"`   // Time when application was build.
	Commit      string `json:"commit"`       // Git commit hash.
	TimeZone    string `json:"time_zone"`    // Local TimeZone
	CurrentTime string `json:"current_time"` // Time that application was executed
}

// NewVersion creates version instance with application information.
// [app] contains command string
// Adding current time as execution time
func NewVersion(app string) *Version {
	return &Version{
		App:         app,
		AppVersion:  appVersion,
		BuildTime:   buildTime,
		Commit:      commit,
		TimeZone:    time.Local.String(),
		CurrentTime: time.Now().Round(0).String(),
	}
}

// PrintAppVersionInfo prints application version to console
func PrintAppVersionInfo() {
	fmt.Printf("Version: %s\nBuildTime: %s\nCommit: %s\n", appVersion, buildTime, commit)
}
