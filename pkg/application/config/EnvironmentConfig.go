package config

import (
	"YellowPepper-FundsTransfers/pkg/misc/environment"
	"YellowPepper-FundsTransfers/pkg/misc/properties"
	"bufio"
	"fmt"
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
	cmdSource := properties.NewDefaultPropertySource(CMD_PROPERTY_SOURCE_NAME, properties.NewDefaultProperties().FromArray(&cmdArgs).Build())

	osArgs := os.Environ()
	osSource := properties.NewDefaultPropertySource(OS_PROPERTY_SOURCE_NAME, properties.NewDefaultProperties().FromArray(&osArgs).Build())

	env := environment.NewDefaultEnvironment().WithPropertySources(cmdSource, osSource).Build()

	resourcesFolder := retrieveResourcesFolder(env)
	profile := env.GetValueOrDefault(PROFILE, PROFILE_DEFAULT_VALUE).AsString()

	mainFileArgs := readFile(filepath.Join(resourcesFolder, "application.properties"))
	mainFileSource := properties.NewDefaultPropertySource(MAIN_FILE_PROPERTY_SOURCE_NAME, properties.NewDefaultProperties().FromArray(mainFileArgs).Build())

	profileFileArgs := readFile(filepath.Join(resourcesFolder, "application-"+profile+".properties"))
	profileFileSource := properties.NewDefaultPropertySource(PROFILE_FILE_PROPERTY_SOURCE_NAME, properties.NewDefaultProperties().FromArray(profileFileArgs).Build())

	env.AppendPropertySources(mainFileSource, profileFileSource)

	return env
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
	sourceFolderName := env.GetValueOrDefault(SOURCE_FOLDER_NAME, SOURCE_FOLDER_NAME_DEFAULT_VALUE).AsString()
	sourceRootDirectory := filepath.Join(workingDirectory, sourceFolderName)
	if !validateIfFolderExists(sourceRootDirectory) {
		sourceRootDirectory = filepath.Join(workingDirectory)
	}
	log.Println(fmt.Sprintf("Configured Resources Folder: %s", sourceRootDirectory))
	return filepath.Join(sourceRootDirectory, ".resources")
}

func validateIfFolderExists(directory string) bool {
	_, err := os.Stat(directory)
	if err != nil {
		return false
	}
	return true
}
