package utils

import (
	"fmt"
	"log"
	"os"

	dotenv "github.com/joho/godotenv"
	customTypes "github.com/vijay2249/vproxie/custom/types"
	yaml "gopkg.in/yaml.v3"
)

var GlobalHeadersConfig *customTypes.HeadersConfig
var GlobalHostsForwardConfig *customTypes.ForwardRequestToConfig

func LoadHeadersConfig(filePaths ...string) (err error) {
	for _, file := range filePaths {
		data, err := os.ReadFile(file)
		if err != nil {
			log.Fatal("Error while loading yaml file")
			return err
		}
		fmt.Println("unloading yaml configs")
		err = yaml.Unmarshal(data, &GlobalHeadersConfig)
		if err != nil {
			log.Fatal("Error while unmarshall the yaml config")
			return err
		}
		// if GlobalHeadersConfig == nil {
		// 	log.Printf("Required config data type is not found in %v file", file)
		// }
	}
	return nil
}

func LoadHostsConfig(filePaths ...string) (err error) {
	for _, file := range filePaths {
		data, err := os.ReadFile(file)
		if err != nil {
			log.Fatal("Error while loading yaml file")
			return err
		}
		fmt.Println("unloading yaml configs")
		err = yaml.Unmarshal(data, &GlobalHostsForwardConfig)
		if err != nil {
			log.Fatal("Error while unmarshall the yaml config")
			return err
		}
	}
	return nil
}

func PrintHeadersYamlConfig(){
	fmt.Println("Printing yaml config")
	fmt.Println(*GlobalHeadersConfig)
}

func PrintHostsForwardConfigYamlConfig(){
	fmt.Println("Printing yaml config")
	fmt.Println(*GlobalHostsForwardConfig)
}


func LoadEnvConfigValues(filePaths ...string) (map[string]string, error){
	fmt.Println(filePaths)
	if len(filePaths) == 0 {
		return make(map[string]string, 0), nil
	}
	vals, err := dotenv.Read(filePaths...)
	
	if err != nil {
		log.Fatal("Unable to load env config files")
		return nil, err
	}
	return vals, nil
}

func ValidateConfigFiles(){
	
}