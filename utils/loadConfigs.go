package utils

import (
	"fmt"
	"log"

	dotenv "github.com/joho/godotenv"
	yaml "gopkg.in/yaml.v3"
)

var globalYaml map[string]interface{}

func LoadYamlConfig(data []byte) (err error) {
	err = yaml.Unmarshal(data, &globalYaml)
	if err != nil {
		log.Fatalf("Unable to load yaml config")
		return err
	}
	return nil
}

func PrintYamlConfig(){
	for key, value := range globalYaml {
		log.Println(key, ": ", value)
	}
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