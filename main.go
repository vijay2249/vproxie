package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vijay2249/vproxie/constant"
	"github.com/vijay2249/vproxie/utils"
)

var (
	filePaths []string;
	configMapTypes map[string][]string
)

// BANNER
func init(){
	fmt.Println(`
																																																										 
	||   / /                                                                              ||   / / //   ) )  /|    / / 
	||  / /          ___      __      ___             ( )  ___         / __               ||  / / //        //|   / /  
	|| / /   ____  //   ) ) //  ) ) //   ) ) \\ / /  / / //___) )     //   ) ) //   / /   || / / //        // |  / /   
	||/ /         //___/ / //      //   / /   \/ /  / / //           //   / / ((___/ /    ||/ / //        //  | / /    
	|  /         //       //      ((___/ /    / /\ / / ((____       ((___/ /      / /     |  / ((____/ / //   |/ /     
	
																																											- VijayCN
	`)
}

func init(){
	// get config giles
	var err error
	filePaths, err = utils.GetAllConfigFiles(constant.CONFIG_DIR_PATH)
	if err != nil {
		utils.ErrorLogger.Println("unable to get config file paths")
		utils.ErrorLogger.Println(err)
		return
	}
	utils.InfoLogger.Printf("config filePaths: %v", filePaths)

	configMapTypes = utils.FilterConfigFiles(filePaths)

	// validate if all config files are present or not, if not <idk>

	// load env config values
	utils.InfoLogger.Println("Loaing env config files - start")
	// envConfigValues
	_, err = utils.LoadEnvConfigValues(configMapTypes[constant.ENV]...)
	if err != nil {
		utils.ErrorLogger.Println("Unable to load env config values, please check error")
		utils.ErrorLogger.Println(err)
		utils.ErrorLogger.Println("Loaing env config files - failure")
		return
	}

	utils.InfoLogger.Println("Loading env config files - completed")

	// load yaml config values
	utils.InfoLogger.Println("Loading yaml config details - start")
	utils.LoadYamlConfigValues(configMapTypes[constant.YAML]...)
	utils.InfoLogger.Println("Loading yaml config details - completed")

	utils.PrintHeadersYamlConfig()
	utils.PrintHostsForwardConfigYamlConfig()
	utils.PrintLoggingConfigs()
}

func handleRequests(w http.ResponseWriter, req *http.Request){

	//copy headers
	utils.DeleteAndModifyHeaders(&req.Header)
	utils.PrintHeaders(req)

	//call backend service
	userHost := req.Header.Get(constant.REFERER)
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
	[x] .env file to store the hosts and based on that host - send to respective backend server
	[ ] Dummy backend server to get the responses
	[ ] Authentication on both proxy and backend server
	`)
}
