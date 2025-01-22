package utils

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"vproxie/constant"
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

//loading env values
func init(){
	log.Println("loading env values")
	filePaths, err := GetAllConfigFiles(constant.CONFIG_DIR_PATH)

	if err != nil {
		log.Fatal("error while reading env files")
		return
	}

	fmt.Println(filePaths)

	filteredConfigMaps := FilterConfigFiles(filePaths)

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

//loading global configs
// var requestHeadersToRemove []string
// var requestHeadersToModify map[string]string 
// func init(){
// 	requestHeadersToRemove = GlobalHeadersConfig.GetRequestHeadersToRemove()
// 	if requestHeadersToRemove == nil { requestHeadersToRemove = make([]string, 0) }
// 	requestHeadersToModify = GlobalHeadersConfig.GetRequestModifyHeadersMap()
// }

func DeleteHeaders(reqHeaders *http.Header) {
	requestHeadersToRemove := GlobalHeadersConfig.GetRequestHeadersToRemove() //change this to be global instead of getting data for each request
	for key := range *reqHeaders {
		if slices.Contains(requestHeadersToRemove, key){
			reqHeaders.Del(key)
		}
	}
}

func ModifyHeaders(reqHeaders *http.Header){
	requestHeadersToModify := GlobalHeadersConfig.GetRequestModifyHeadersMap() //change this to be global instead of getting data for each request
	for key, value := range requestHeadersToModify {
		reqHeaders.Set(key, value)
	}
}

func DeleteAndModifyHeaders(request *http.Header) {
	DeleteHeaders(request)
	ModifyHeaders(request)
}

func PrintHeaders(req *http.Request){
	for key, value := range req.Header {
		log.Print(key, value)
	}
}

func CreateCorelationHeader(headers *http.Header) {}

func SafetyCheck(){}
 
func RouteTo(req *http.Request) string {
	hostHeader := req.Header.Get(constant.HOST_HEADER)
	return hostHeader
}

