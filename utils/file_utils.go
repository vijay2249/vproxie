package utils

import (
	"log"
	"os"
)

func ReadFileData(filePath string) (data []byte, err error){
	data, err = os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error while reading file: %v", filePath)
		return nil, err
	}
	return data, nil
}