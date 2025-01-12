package utils

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/wedonttrack.org/vproxie/constant"
)

var ENV_DIR_RELATIVE_PATH string = "./.env/"

func GetAllConfigFiles(config_dir_paths ...string) ([]string, error) {
	var fileNames []string
	for _, config_dir_path := range config_dir_paths {
		err := filepath.Walk(config_dir_path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Fatal("Error while loading directory data")
				return err
			}

			fileInfo, err := os.Stat(path)
			if err != nil {
				log.Fatal("Error while getting file info")
				return err
			}
			if !fileInfo.IsDir() {
				fileNames = append(fileNames, path)
			}
			return nil
		})
		
		if err != nil {
			log.Fatal("Error while reading config files")
			return nil, err
		}
	}

	return fileNames, nil
}

func FilterConfigFiles(filePaths []string) (map[string][]string) {
	var configMapTypes  = make(map[string][]string) 
	for _, value := range filePaths {
		splitVals := strings.Split(value, ".")
		configMapTypes[splitVals[1]] = append(configMapTypes[splitVals[1]], value)
	}
	return configMapTypes
}

func init(){
	log.Println("loading env values")
	filePaths, err := GetAllConfigFiles(constant.CONFIG_DIR_PATH)

	if err != nil {
		log.Fatal("error while reading env files")
		return
	}

	fmt.Println(filePaths)

	filteredConfigMaps := FilterConfigFiles(filePaths)

	// vals, err := dotenv.Read(filteredConfigMaps[".env"]...)
	vals, err := LoadEnvConfigValues(filteredConfigMaps[".env"]...)

	if err != nil {
		log.Fatal("Unable to load .env files in .env folder")
		fmt.Println(err)
		return
	}

	log.Println("========== ENV VALUES ==============")
	log.Println(vals)
	log.Println("========== ENV VALUES ==============")
}

func ModifyHeaders(reqHeaders *http.Header){
	for key := range *reqHeaders {
		if slices.Contains(constant.HEADERS_TO_EXCLUDE, strings.ToLower(key)){
			reqHeaders.Del(key)
		}
	}
}

func PrintHeaders(req *http.Request){
	for key, value := range req.Header {
		fmt.Println(key, value)
	}
}

func CreateCorelationHeader(){}

func SafetyCheck(){}
 
func RouteTo(req *http.Request) string {
	hostHeader := req.Header.Get(constant.HOST_HEADER)
	return hostHeader
}

