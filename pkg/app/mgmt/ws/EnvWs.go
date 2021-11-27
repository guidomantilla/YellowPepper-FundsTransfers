package ws

import (
	"YellowPepper-FundsTransfers/pkg/app/misc/environment"
	"net/http"

	"github.com/gin-gonic/gin"
)

/* TYPES DEFINITION */

type EnvWsResponse struct {
	EnvironmentVariables []map[string]interface{} `json:"environment-variables"`
}

type EnvWs interface {
	Get(context *gin.Context)
}

type DefaultEnvWs struct {
	environment environment.Environment
}

/* TYPES CONSTRUCTOR */

func NewDefaultEnvWs(environment environment.Environment) *DefaultEnvWs {
	return &DefaultEnvWs{
		environment: environment,
	}
}

/* DefaultEnvWs METHODS */

func (ws *DefaultEnvWs) Get(context *gin.Context) {

	response := EnvWsResponse{
		EnvironmentVariables: make([]map[string]interface{}, len(ws.environment.GetPropertySources())),
	}
	for index, env := range ws.environment.GetPropertySources() {
		response.EnvironmentVariables[index] = env.AsMap()
	}
	context.JSON(http.StatusOK, response)
}
