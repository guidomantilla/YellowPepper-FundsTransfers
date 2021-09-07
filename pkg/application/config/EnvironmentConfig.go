package config

import (
	"YellowPepper-FundsTransfers/pkg/misc/environment"
	"YellowPepper-FundsTransfers/pkg/misc/files"
	"YellowPepper-FundsTransfers/pkg/misc/properties"
	"os"
	"path/filepath"
)

const (
	CMD_PROPERTY_SOURCE_NAME          = "CMD"
	OS_PROPERTY_SOURCE_NAME           = "OS"
	MAIN_FILE_PROPERTY_SOURCE_NAME    = "MAIN-FILE"
	PROFILE_FILE_PROPERTY_SOURCE_NAME = "PROFILE-FILE"
)
func LoadEnvironment() environment.Environment {

	cmdArgs := os.Args[1:]
	cmdSource := properties.NewDefaultPropertySource(CMD_PROPERTY_SOURCE_NAME, properties.NewDefaultProperties().FromArray(&cmdArgs).Build())

	osArgs := os.Environ()
	osSource := properties.NewDefaultPropertySource(OS_PROPERTY_SOURCE_NAME, properties.NewDefaultProperties().FromArray(&osArgs).Build())

	env := environment.NewDefaultEnvironment().WithPropertySources(cmdSource, osSource).Build()

	resourcesFolder := retrieveResourcesFolder(env)
	profile := env.GetValueOrDefault(PROFILE, PROFILE_DEFAULT_VALUE).AsString()

	mainFileArgs := files.ReadFile(filepath.Join(resourcesFolder, "application.properties"))
	mainFileSource := properties.NewDefaultPropertySource(MAIN_FILE_PROPERTY_SOURCE_NAME, properties.NewDefaultProperties().FromArray(mainFileArgs).Build())

	profileFileArgs := files.ReadFile(filepath.Join(resourcesFolder, "application-"+profile+".properties"))
	profileFileSource := properties.NewDefaultPropertySource(PROFILE_FILE_PROPERTY_SOURCE_NAME, properties.NewDefaultProperties().FromArray(profileFileArgs).Build())

	env.AppendPropertySources(mainFileSource, profileFileSource)

	return env
}

func retrieveResourcesFolder(env environment.Environment) string {
	workingDirectory, _ := os.Getwd()
	sourceFolderName := env.GetValueOrDefault(SOURCE_FOLDER_NAME, SOURCE_FOLDER_NAME_DEFAULT_VALUE).AsString()
	sourceRootDirectory := filepath.Join(workingDirectory, sourceFolderName)
	if !files.ValidateIfFolderExists(sourceRootDirectory) {
		sourceRootDirectory = filepath.Join(workingDirectory)
	}
	return filepath.Join(sourceRootDirectory, ".resources")
}
