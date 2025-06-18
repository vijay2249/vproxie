package route_utils

import (
	"errors"
	"net/http"
	// "os"

	"github.com/vijay2249/vproxie/utils"
)

var ROUTING_CONFIG = utils.GlobalRoutingConfig

func GetBackendServiceURL(domain, subdomain, endpoint string) (url string, err error) {
	//return default root service
	if subdomain == "" && endpoint == "" { 
		return domain, nil
	}
	//get service based on endpoint
	if subdomain == "" {
		url, err = ROUTING_CONFIG.GetBackendServiceURLBySubdomain(domain, endpoint)
		if err != nil { return "", err }
		return url, nil
	} 
	// get service based on endpoint
	if endpoint == "" {
		url, err = ROUTING_CONFIG.GetBackendServiceURLBySubdomain(domain, subdomain)
		if err != nil { return "", err }
		return url, nil
	}

	utils.WarnLogger.Printf("No such subdomain exists to route request to. Please check configurations")
	return "", errors.ErrUnsupported
}

func GetSubdomain(host string) {
	// a.b.example.com
	// example.com
	// https://a.b.example.com
	// https://example.com


}

func GetResponseFromHost(req *http.Request){
	var host, endpoint string = req.Host, req.URL.Path
	utils.DebugLogger.Printf("req host: %v\treq endpoint: %v", host, endpoint)

	
	//get backend url
}

func ParseResponse(){}

func RefactorResponse(){}

func RespondToClient(){}

