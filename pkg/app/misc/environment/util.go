package environment

import (
	"YellowPepper-FundsTransfers/pkg/app/misc/files"
	"YellowPepper-FundsTransfers/pkg/app/misc/properties"
	"os"
	"path/filepath"
)

const (
	CMD_PROPERTY_SOURCE_NAME          = "CMD"
	MAIN_FILE_PROPERTY_SOURCE_NAME    = "MAIN-FILE"
	PROFILE_FILE_PROPERTY_SOURCE_NAME = "PROFILE-FILE"
	OS_PROPERTY_SOURCE_NAME           = "OS"
)

func LoadEnvironment(cmdArgs *[]string, profile string, profileDefaultValue string, sourceFolderName string, sourceFolderNameDefaultValue string) Environment {

	cmdSource := properties.NewDefaultPropertySource(CMD_PROPERTY_SOURCE_NAME, properties.NewDefaultProperties().FromArray(cmdArgs).Build())

	env := NewDefaultEnvironment().WithPropertySources(cmdSource).Build()

	resourcesFolder := retrieveResourcesFolder(env, sourceFolderName, sourceFolderNameDefaultValue)
	profile = env.GetValueOrDefault(profile, profileDefaultValue).AsString()

	mainFileArgs := files.ReadFile(filepath.Join(resourcesFolder, "application.properties"))
	mainFileSource := properties.NewDefaultPropertySource(MAIN_FILE_PROPERTY_SOURCE_NAME, properties.NewDefaultProperties().FromArray(mainFileArgs).Build())

	profileFileArgs := files.ReadFile(filepath.Join(resourcesFolder, "application-"+profile+".properties"))
	profileFileSource := properties.NewDefaultPropertySource(PROFILE_FILE_PROPERTY_SOURCE_NAME, properties.NewDefaultProperties().FromArray(profileFileArgs).Build())

	osArgs := os.Environ()
	osSource := properties.NewDefaultPropertySource(OS_PROPERTY_SOURCE_NAME, properties.NewDefaultProperties().FromArray(&osArgs).Build())

	env.AppendPropertySources(osSource, mainFileSource, profileFileSource)

	return env
}

func retrieveResourcesFolder(env Environment, sourceFolderName string, sourceFolderNameDefaultValue string) string {
	workingDirectory, _ := os.Getwd()
	sourceFolderName = env.GetValueOrDefault(sourceFolderName, sourceFolderNameDefaultValue).AsString()
	sourceRootDirectory := filepath.Join(workingDirectory, sourceFolderName)
	if !files.ValidateIfFolderExists(sourceRootDirectory) {
		sourceRootDirectory = filepath.Join(workingDirectory)
	}
	return filepath.Join(sourceRootDirectory, ".resources")
}
