package file_utils

import (
	"log"
	"os"
	"path/filepath"
)

func ReadFileData(filePath string) (data []byte, err error){
	data, err = os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error while reading file: %v", filePath)
		return nil, err
	}
	return data, nil
}

func GetAllFileNamesInFolderAndSubFolder(config_dir_paths ...string) ([]string, error) {
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