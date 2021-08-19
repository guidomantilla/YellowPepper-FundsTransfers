package config

import (
	"YellowPepper-FundsTransfers/misc/collection"
	"os"
	"strings"
)

var applicationProperties *collection.Properties
func loadApplicationProperties(args []string) {
	applicationProperties = collection.NewProperties()
	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		applicationProperties.Add(pair[0], pair[1])
	}
}

func LoadApplicationProperties(args []string) *collection.Properties {
	if applicationProperties == nil {
		loadApplicationProperties(args)
	}
	return applicationProperties
}
