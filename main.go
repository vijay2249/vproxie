package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wedonttrack.org/vproxie/constant"
	"github.com/wedonttrack.org/vproxie/utils"
)

var filePaths []string;

// BANNER
func init(){
	fmt.Println(`
																																																										 
	||   / /                                                                              ||   / / //   ) )  /|    / / 
	||  / /          ___      __      ___             ( )  ___         / __               ||  / / //        //|   / /  
	|| / /   ____  //   ) ) //  ) ) //   ) ) \\ / /  / / //___) )     //   ) ) //   / /   || / / //        // |  / /   
	||/ /         //___/ / //      //   / /   \/ /  / / //           //   / / ((___/ /    ||/ / //        //  | / /    
	|  /         //       //      ((___/ /    / /\ / / ((____       ((___/ /      / /     |  / ((____/ / //   |/ /     
	`)
}

// load yaml config
func init(){
	var err error;
	filePaths, err = utils.GetAllConfigFiles(constant.CONFIG_DIR_PATH)
	if err != nil {
		log.Fatal("Error while getting config file paths")
		return
	}

	configMapTypes := utils.FilterConfigFiles(filePaths)
	utils.LoadEnvConfigValues(configMapTypes["env"]...)
	utils.LoadHeadersConfig(configMapTypes["yaml"]...)
	utils.PrintHeadersYamlConfig()

	utils.LoadHostsConfig(configMapTypes["yaml"]...)
	utils.PrintHostsForwardConfigYamlConfig()
}

func handleRequests(w http.ResponseWriter, req *http.Request){

	//copy headers
	utils.DeleteAndModifyHeaders(&req.Header)
	utils.PrintHeaders(req)

	//call backend service
	userHost := req.Header.Get("Referer")
	fmt.Println(userHost)
}

func main(){
	fmt.Println("Proxy server by Vijay - for custom projects")
	targets()
	startServer()
}

func startServer(){
	server := http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(handleRequests),
	}

	log.Println("Starting proxy server on :8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting proxy server: ", err)
	}
}


func targets(){
	fmt.Println(`
	1. .env file to store the hosts and based on that host - send to respective backend server
	2. Dummy backend server to get the responses
	3. Authentication on both proxy and backend server
	`)
}
