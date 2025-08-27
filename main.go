package main

import (
	"os"
	"errors"
	"net/http"

	"github.com/vijay2249/vproxie/constant"
	"github.com/vijay2249/vproxie/utils"
	route_utils "github.com/vijay2249/vproxie/utils/route"
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

// project targets
func init(){
	utils.InfoLogger.Println(`
	[x] .yaml file to store the hosts and based on that host - send to respective backend server
	[x] add logging mechanism

	// TODO
	[ ] make port and server address configurable
	[ ] remove PII related metadata from files
	[ ] files uploaded should be validated from virustotal first then should be send to backend servers
	[ ] Dummy backend server to get the responses
	[ ] Authentication on both proxy and backend server
	`)
}

func init(){
	// get config files
	var err error
	filePaths, err = utils.GetAllConfigFiles(constant.CONFIG_DIR_PATH)
	if err != nil {
		utils.ErrorLogger.Println("unable to get config file paths")
		utils.ErrorLogger.Fatal(err)
		return
	}
	utils.InfoLogger.Printf("config filePaths: %v", filePaths)

	configMapTypes = utils.FilterConfigFiles(filePaths)

	// load yaml config values
	utils.InfoLogger.Println("Loading yaml config details - start")
	utils.LoadYamlConfigValues(configMapTypes[constant.YAML]...)
	utils.InfoLogger.Println("Loading yaml config details - completed")

	// print configurations to check if load function is working or not
	utils.PrintHeadersYamlConfig()
	utils.PrintHostsForwardConfigYamlConfig()
	utils.PrintLoggingConfigs()

	// run all necessary functions to populate necessary details in utilities
	route_utils.InitRouteUtils()
}

func handleRequests(w http.ResponseWriter, req *http.Request){

	defer req.Body.Close()
	
	// validate connection details
	// check for HTTPS connection and check for valid SSL certificate 
	// check for blocked IP connections and blocked fingerprints
	
	//format request details --- START
	// delete metadata from request
	// copy headers  --- START
	// utils.PrintHeaders(req)
	route_utils.FormatHeaders(&req.Header, constant.REQUEST)
	// delete metadata from request payload
	

	// get backend service URL  --- START
	userHost := req.Header.Get(constant.REFERER)
	utils.InfoLogger.Println(userHost)
	var host, endpoint = req.Host, req.URL.Path
	utils.InfoLogger.Printf("host: %s, endpoint: %s\n", host, endpoint)
	backendServerURL := route_utils.FindServerURL(host, endpoint)

	// get response from server
	serverResponse := route_utils.GetResponseFromServer(backendServerURL, req)

	// format response from server
	route_utils.FormatHeaders(&serverResponse.Header, constant.RESPONSE)

	// return response to user
	route_utils.CastResponseToResponseWriter(&serverResponse, &w)
}

func main(){
	utils.InfoLogger.Println("Proxy server by Vijay - for custom projects")
	startServer()
}

func startServer(){
	server := http.Server{
		Addr: ":8080", //TODO - make it configurable
		Handler: http.HandlerFunc(handleRequests),
	}

	utils.InfoLogger.Println("Starting proxy server on :8080")
	err := server.ListenAndServe()
	// server.ListenAndServeTLS()
	if errors.Is(err, http.ErrServerClosed){
		utils.InfoLogger.Println("Server closed gracefully")
	}else if err != nil {
		utils.ErrorLogger.Fatal("Error starting proxy server: ", err)
		os.Exit(1)
	}
}
