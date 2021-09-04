package config

import (
	"YellowPepper-FundsTransfers/pkg/misc/environment"
	"YellowPepper-FundsTransfers/pkg/misc/properties"
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	CMD_PROPERTY_SOURCE_NAME          = "CMD"
	OS_PROPERTY_SOURCE_NAME           = "OS"
	MAIN_FILE_PROPERTY_SOURCE_NAME    = "MAIN-FILE"
	PROFILE_FILE_PROPERTY_SOURCE_NAME = "PROFILE-FILE"
)

func LoadEnvironment() environment.Environment {

	cmdArgs := os.Args[1:]
	cmdSource := properties.NewDefaultPropertySource(CMD_PROPERTY_SOURCE_NAME, properties.NewPropertiesFromArray(&cmdArgs))

	osArgs := os.Environ()
	osSource := properties.NewDefaultPropertySource(OS_PROPERTY_SOURCE_NAME, properties.NewPropertiesFromArray(&osArgs))

	env := environment.NewDefaultEnvironment().WithPropertySources(cmdSource, osSource).Build()

	mainFileArgs := loadMainFile(env)
	mainFileSource := properties.NewDefaultPropertySource(MAIN_FILE_PROPERTY_SOURCE_NAME, properties.NewPropertiesFromArray(mainFileArgs))

	profileFileArgs := loadProfileFile(env)
	profileFileSource := properties.NewDefaultPropertySource(PROFILE_FILE_PROPERTY_SOURCE_NAME, properties.NewPropertiesFromArray(profileFileArgs))

	env.AppendPropertySources(mainFileSource, profileFileSource)

	return env
}

func loadMainFile(env environment.Environment) *[]string {
	resourcesDirectory := retrieveResourcesFolder(env)
	return readFile(filepath.Join(resourcesDirectory, "application.properties"))
}

func loadProfileFile(env environment.Environment) *[]string {
	resourcesDirectory := retrieveResourcesFolder(env)
	profile := env.GetValueOrDefault(PROFILE, PROFILE_DEFAULT_VALUE).AsString()
	return readFile(filepath.Join(resourcesDirectory, "application-"+profile+".properties"))
}

func readFile(fullFileName string) *[]string {

	lines := make([]string, 0)

	file, err := os.Open(fullFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &lines
}

func retrieveResourcesFolder(env environment.Environment) string {
	workingDirectory, _ := os.Getwd()
	sourceRootDirectory := filepath.Join(workingDirectory, env.GetValueOrDefault(SOURCE_FOLDER_NAME, SOURCE_FOLDER_NAME_DEFAULT_VALUE).AsString())
	if !validateIfFolderExists(sourceRootDirectory) {
		sourceRootDirectory = filepath.Join(workingDirectory)
	}
	return filepath.Join(sourceRootDirectory, ".resources")
}

func validateIfFolderExists(directory string) bool {
	_, err := os.Stat(directory)
	if err != nil {
		return false
	}
	return true
}
