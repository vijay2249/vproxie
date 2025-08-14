package utils

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)


func GetAllConfigFiles(configDirPaths ...string) ([]string, error) {
	var fileNames []string
	for _, configDirPath := range configDirPaths {
		err := filepath.Walk(configDirPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				ErrorLogger.Println("Error while loading directory data")
				return err
			}

			fileInfo, err := os.Stat(path)
			if err != nil {
				ErrorLogger.Println("Error while getting file info")
				return err
			}
			if !fileInfo.IsDir() {
				fileNames = append(fileNames, path)
			}
			return nil
		})
		
		if err != nil {
			ErrorLogger.Println("Error while reading config files")
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

func CreateCorelationHeader(headers *http.Header) {}

func SafetyCheck(){}

