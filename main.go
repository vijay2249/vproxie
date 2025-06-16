package main

import (
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
	utils.InfoLogger.Println(`
																																																										 
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

	// load yaml config values
	utils.InfoLogger.Println("Loading yaml config details - start")
	utils.LoadYamlConfigValues(configMapTypes[constant.YAML]...)
	utils.InfoLogger.Println("Loading yaml config details - completed")

	utils.PrintHeadersYamlConfig()
	utils.PrintHostsForwardConfigYamlConfig()
	utils.PrintLoggingConfigs()
}

func handleRequests(w http.ResponseWriter, req *http.Request){


	// validate request connection status
	/*
	validate SSL certificate, check blocked IP requests, check connection type(secure or not)
	*/

	// ----- modify request details  START ------
	// modify request headers
	utils.DeleteAndModifyHeaders(&req.Header)
	utils.PrintHeaders(req)


	// ----- modify request details  END --------


	//call backend service
	userHost := req.Header.Get(constant.REFERER)
	utils.InfoLogger.Println(userHost)

	// ----- modify response details  START ------

	// get response from backend service
	// modify headers from response
	// modify additional metadata details from response
	// send response to user

	// ----- modify request details  END ------

}

func main(){
	utils.InfoLogger.Println("Proxy server by Vijay - for custom projects")
	startServer()
}

func startServer(){
	server := http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(handleRequests),
	}

	utils.InfoLogger.Println("Starting proxy server on :8080")
	err := server.ListenAndServe()
	if err != nil {
		utils.ErrorLogger.Fatal("Error starting proxy server: ", err)
	}
}


func init(){
	utils.InfoLogger.Println(`
	[x] .env file to store the hosts and based on that host - send to respective backend server
	[ ] Dummy backend server to get the responses
	[ ] Authentication on both proxy and backend server
	`)
}
