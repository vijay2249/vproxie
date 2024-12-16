package utils

import (
	//"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"

	dotenv "github.com/joho/godotenv"
	constant "github.com/wedonttrack.org/vproxie/constant"
)

var ENV_DIR_RELATIVE_PATH string = "..\\.env\\"

func getEnvDirPath() (string, error) {
	path, err := os.Getwd()

	if err != nil {
		log.Fatal("Unable to get the current working directory location")
		return "", err
	}

	path += "\\" + ENV_DIR_RELATIVE_PATH

	//check if folder exists or not
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("required folder/folders (%v) are not present", path)
		return "", err
	}
	
	return path, nil
}

func getAllEnvFiles() ([]string, error) {
	var fileNames []string

	envFolderPath, err := getEnvDirPath()
	if err != nil  || envFolderPath == "" {
		log.Fatal("Some error while getting env files")
		return nil, err
	}

	err = filepath.Walk(envFolderPath, func(path string, info os.FileInfo, err error) error{
		if err != nil {
			log.Panicf("unable to read file: %s", path)
			//fmt.Println(err)
			return err
		}
		fileNames = append(fileNames, path)
		return nil
	})

	if err != nil {
		log.Fatal("Error while reading env files")
		return nil, err
	}

	return fileNames, err
}

func init(){
	log.Println("loading env values")
	fileNames, err := getAllEnvFiles()

	if err != nil {
		log.Fatal("error while reading env files")
		return
	}

	vals, err := dotenv.Read(fileNames...)

	if err != nil {
		log.Fatal("Unable to load .env files in .env folder")
		return
	}

	log.Println("========== ENV VALUES ==============")
	log.Println(vals)
	log.Println("========== ENV VALUES ==============")
}

func ModifyHeaders(reqHeaders *http.Header) {
	for key := range *reqHeaders {
		if slices.Contains(constant.HEADERS_TO_EXCLUDE, strings.ToLower(key)){
			reqHeaders.Del(key)
		}
	}
}

func PrintHeaders(req *http.Request){
	for key, value := range req.Header {
		log.Print(key, value)
		//fmt.Println(key, value)
	}
}

func CreateCorelationHeader(headers *http.Header) {}

func SafetyCheck(){}

func RouteTo(req *http.Request) string {
	hostHeader := req.Header.Get(constant.HOST_HEADER)
	return hostHeader
}

