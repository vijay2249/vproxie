package utils

import (
	"os"
	"reflect"

	customTypes "github.com/vijay2249/vproxie/custom/types"
	yaml "gopkg.in/yaml.v3"
)

var (
	GlobalHeadersConfig *customTypes.HeadersConfig
  GlobalRoutingConfig *customTypes.RoutingConfig
  GlobalLoggingConfig *customTypes.LoggingConfig
)

var ConfigsToLoad = []interface{}{&GlobalHeadersConfig, &GlobalRoutingConfig, &GlobalLoggingConfig}

func PrintHeadersYamlConfig(){
	InfoLogger.Println("Printing headers yaml config")
	InfoLogger.Println(*GlobalHeadersConfig)
}

func PrintHostsForwardConfigYamlConfig(){
	InfoLogger.Println("Printing hosts yaml config")
	InfoLogger.Println(*GlobalRoutingConfig)
}

func PrintLoggingConfigs(){
	InfoLogger.Println("Logging configs")
	InfoLogger.Println(*GlobalLoggingConfig)
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

		err = UnmarshallConfig(data)
		if err != nil {
			ErrorLogger.Fatal("Error while loading yaml config, please re-validate")
			return err
		}

	}
	return nil
}

func UnmarshallConfig(data []byte) (err error){
	for _, config := range ConfigsToLoad {
		configDataType := reflect.TypeOf(config)
		InfoLogger.Printf("Loading yaml config for %v", configDataType)
		err = yaml.Unmarshal(data, config)
		if err != nil {
			ErrorLogger.Printf("Error while loading %v config", configDataType)
			return err
		}
	}
	return nil
}

func UnmarshallEachConfig(data []byte, config interface{}) (err error){
	configDataType := reflect.TypeOf(config)
	InfoLogger.Printf("Loading yaml config for %v", configDataType)
	err = yaml.Unmarshal(data, &config)
	InfoLogger.Println("before")
	InfoLogger.Println(&config)
	if err != nil {
		ErrorLogger.Printf("Error while loading %v config", configDataType)
		return err
	}
	InfoLogger.Println("after")
	InfoLogger.Println(&config)
	return nil
}