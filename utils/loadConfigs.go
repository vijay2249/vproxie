package utils

import (
	"fmt"
	"os"
	"reflect"

	dotenv "github.com/joho/godotenv"
	customTypes "github.com/vijay2249/vproxie/custom/types"
	yaml "gopkg.in/yaml.v3"
)

var (
	GlobalHeadersConfig *customTypes.HeadersConfig
  GlobalHostsForwardConfig *customTypes.ForwardRequestToConfig
  GlobalLoggingConfig *customTypes.LoggingConfig
)

var ConfigsToLoad = []interface{}{GlobalHeadersConfig , GlobalHostsForwardConfig, GlobalLoggingConfig}

func PrintHeadersYamlConfig(){
	fmt.Println("Printing headers yaml config")
	fmt.Println(*GlobalHeadersConfig)
}

func PrintHostsForwardConfigYamlConfig(){
	fmt.Println("Printing hosts yaml config")
	fmt.Println(*GlobalHostsForwardConfig)
}

func PrintLoggingConfigs(){
	fmt.Println("Logging configs")
	fmt.Println(*GlobalLoggingConfig)
}

func LoadEnvConfigValues(filePaths ...string) (map[string]string, error){
	InfoLogger.Printf("env config files: %v", filePaths)
	var err error
	if len(filePaths) == 0 {
		WarnLogger.Println("config files are empty, ignoring this loading env configs")
		return make(map[string]string, 0), nil
	}
	vals, err := dotenv.Read(filePaths...)
	InfoLogger.Println("Completed loading env config details")
	if err != nil {
		ErrorLogger.Println("Unable to load env config files")
		return nil, err
	}
	return vals, nil
}

func LoadYamlConfigValues(filePaths ...string) error {
	InfoLogger.Printf("reading yaml config from files: %v", filePaths)
	for _, file := range filePaths {
		data, err := os.ReadFile(file)
		if err != nil {
			ErrorLogger.Printf("Error while reading %v file to laod yaml config", file)
			return err
		}
		InfoLogger.Printf("unloading yaml config from %v", file)

		err = UnmarshallYamlConfig(data)
		if err != nil {
			ErrorLogger.Printf("Error while loading config from %v file", file)
			return err
		}
	}
	return nil
}


func UnmarshallYamlConfig(data []byte) (err error) {
	for _, config := range ConfigsToLoad {
		configDataType := reflect.TypeOf(config)
		fmt.Println("Before")
		fmt.Println(config)
		InfoLogger.Printf("Loading yaml config for %v", configDataType)
		err = yaml.Unmarshal(data, &config)
		if err != nil {
			ErrorLogger.Printf("Error while loading %v config", configDataType)
			return err
		}
		fmt.Println("After")
		fmt.Println(config)
	}

	fmt.Println("Headers config")
	fmt.Println(GlobalHeadersConfig)
	return nil
}