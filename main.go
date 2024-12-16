package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wedonttrack.org/vproxie/utils"
	//"github.com/gin-gonic/gin"
)

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

func init(){
	//customTransport related init function
}

func handleRequests(w http.ResponseWriter, req *http.Request){
	targetURL := req.URL
	//proxy request 
	_ , err := http.NewRequest(req.Method, targetURL.String(), req.Body)
	if err != nil {
		http.Error(w, "error sending request,", http.StatusInternalServerError)
		return
	}

	//copy headers
	utils.ModifyHeaders(&req.Header)
	utils.PrintHeaders(req)

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
