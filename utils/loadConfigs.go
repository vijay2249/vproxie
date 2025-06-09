package utils

import (
	"os"
	"reflect"

	"github.com/vijay2249/vproxie/constant"
	customTypes "github.com/vijay2249/vproxie/custom/types"
	yaml "gopkg.in/yaml.v3"
)

var (
	GlobalHeadersConfig *customTypes.HeadersConfig
  GlobalHostsForwardConfig *customTypes.ForwardRequestToConfig
  GlobalLoggingConfig *customTypes.LoggingConfig
)

func PrintHeadersYamlConfig(){
	InfoLogger.Println("Printing headers yaml config")
	InfoLogger.Println(*GlobalHeadersConfig)
}

func PrintHostsForwardConfigYamlConfig(){
	InfoLogger.Println("Printing hosts yaml config")
	InfoLogger.Println(*GlobalHostsForwardConfig)
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


		channels := make([]chan error, constant.TOTAL_CUSTOM_CONFIGS)
		go UnmarshallEachConfig(data, GlobalHeadersConfig, channels[0])
		go UnmarshallEachConfig(data, GlobalHostsForwardConfig, channels[1])
		go UnmarshallEachConfig(data, GlobalLoggingConfig, channels[2])

		for _, chans := range channels {
			InfoLogger.Println("inside channels for loop")
			err := <- chans
			InfoLogger.Println(err)
			InfoLogger.Println("inside channels for loop")
			if err != nil {
				ErrorLogger.Fatalf("Error while loading yaml configs, please revalidate")
			}
			close(chans)
		}
	}
	return nil
}

func UnmarshallEachConfig(data []byte, config interface{}, channel chan error) {
	configDataType := reflect.TypeOf(config)
	InfoLogger.Printf("Loading yaml config for %v", configDataType)
	err := yaml.Unmarshal(data, &config)
	InfoLogger.Println("before")
	InfoLogger.Println(&config)
	if err != nil {
		ErrorLogger.Printf("Error while loading %v config", configDataType)
		channel <- err
	}
	InfoLogger.Println("after")
	InfoLogger.Println(&config)
	channel <- nil
}