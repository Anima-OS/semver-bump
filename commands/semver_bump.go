package commands

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var projectVersion string
var versionFile string
var versionStorageType string = "file"
var versionStorageLocalDefaultVersion string

var SemverBumpCommand = &cobra.Command{
	Use:   "semver-bump",
	Short: "Semantic Versioning Bumper",
	Long:  `A semantic versioning file bumper built by giantswarm`,
	Run: func(cmd *cobra.Command, args []string) {
		s, err := getVersionStorage()

		if err != nil {
			log.Fatal(err)
		}

		currentVersion, err := s.ReadVersionFile(versionFile)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Current version is: %s", currentVersion)
	},
}

func Execute(v string) {
	projectVersion = v

	AddGlobalFlags()
	AddCommands()

	SemverBumpCommand.Execute()
}

func AddCommands() {
	SemverBumpCommand.AddCommand(bumpMajorCommand)
	SemverBumpCommand.AddCommand(bumpMinorCommand)
	SemverBumpCommand.AddCommand(bumpPatchCommand)
	SemverBumpCommand.AddCommand(initCommand)
	SemverBumpCommand.AddCommand(versionCommand)
}

func AddGlobalFlags() {
	SemverBumpCommand.PersistentFlags().StringVarP(&versionFile, "version-file", "f", "VERSION", "Version file to use")
	SemverBumpCommand.PersistentFlags().StringVarP(&versionStorageType, "storage-type", "s", "file", "Storage backend to use for version information")
	SemverBumpCommand.PersistentFlags().StringVarP(&versionStorageLocalDefaultVersion, "storage-local-default-version", "V", "0.0.1", "Default version to use when using the local storage backend")

	initCommand.Flags().StringVarP(&initialVersionString, "initial-version", "i", "0.1.0", "The initial version of the project")
}