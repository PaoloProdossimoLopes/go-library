package enviroment

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	Enviroment Env
)

func Init() error {
	loadDotEnvError := godotenv.Load(".env")
	if loadDotEnvError != nil {
		return fmt.Errorf("Probleam to load '.env' file: %s", loadDotEnvError.Error())
	}

	hostPort, hostPortError := validateVariable("HOST_PORT")
	if hostPortError != nil {
		return hostPortError
	}

	Enviroment = Env{
		hostPort: hostPort,
	}
	return nil
}

func validateVariable(variable string) (string, error) {
	variableValue := os.Getenv(variable)
	const emptyText = ""

	if variableValue == emptyText {
		formattedError := fmt.Errorf("Enviroment variable '%s' is not set", variable)
		return emptyText, formattedError
	}

	return variableValue, nil
}
