package config

import (
	"testing"
)

func TestNewVersion(t *testing.T) {
	type args struct {
		app string
	}

	tests := []struct {
		name    string
		args    args
		want    *Version
		prepare func()
	}{
		{
			name: "Must create NewVersion using values given",
			args: args{
				app: "tests",
			},
			want: &Version{
				App:         "tests",
				AppVersion:  "v1_test",
				BuildTime:   "2021-07-07created_using_makefile",
				Commit:      "tag",
				TimeZone:    "Local", // Cannot change to different using TZ env variable
				CurrentTime: "",      // not validated
			},
			prepare: func() {
				appVersion = "v1_test"
				buildTime = "2021-07-07created_using_makefile"
				commit = "tag"
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepare()
			got := NewVersion(tt.args.app)
			// Only CurrentTime not checked because it is always different
			if got.App != tt.want.App || got.AppVersion != tt.want.AppVersion || got.BuildTime != tt.want.BuildTime || got.Commit != tt.want.Commit || got.TimeZone != tt.want.TimeZone {
				t.Errorf("NewVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintAppVersionInfo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Must Println version without error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintAppVersionInfo()
		})
	}
}
