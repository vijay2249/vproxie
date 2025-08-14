package route_utils

import (
	"slices"
	"net/http"
	"strings"
	"io"

	"github.com/vijay2249/vproxie/constant"
	"github.com/vijay2249/vproxie/custom/types"
	"github.com/vijay2249/vproxie/utils"
)



var (

	ROUTING_CONFIGS *types.RoutingConfig
	ALL_DOMAINS []string


	requestHeadersToRemove   []string
	requestHeadersToModify   map[string]string
	responseHeadersToRemove  []string
	responseHeadersToModify  map[string]string
	headersConfigInitialized bool = false
)

func InitRouteUtils() {
	if !headersConfigInitialized {
		requestHeadersToRemove = utils.GlobalHeadersConfig.GetRequestHeadersToRemove()
		requestHeadersToModify = utils.GlobalHeadersConfig.GetRequestModifyHeadersMap()
		responseHeadersToRemove = utils.GlobalHeadersConfig.GetResponseHeadersToRemove()
		responseHeadersToModify = utils.GlobalHeadersConfig.GetResponseModifyHeadersMap()
		headersConfigInitialized = true
	}

	ROUTING_CONFIGS = utils.GlobalRoutingConfig
	ALL_DOMAINS = ROUTING_CONFIGS.GetAllDomains()
}

func GetRouteConfig() *types.RoutingConfig { 
	if headersConfigInitialized{
		return ROUTING_CONFIGS
	} else {
		return nil
	}
}

func FindServerURL(host, endpoint string) string {
	// find routing config based on subdomain, even if part of host is present in domain.
	
	// port := strings.Split(host, ":")[1]
	host = strings.Split(host, ":")[0]

	//find largest common suffix string from allDomains that is common with host variable
	bestMatchingDomainConfig := utils.FindBiggestMatchingSuffix(ALL_DOMAINS, host)

	if bestMatchingDomainConfig != "" {
		utils.InfoLogger.Printf("Best matching domain suffix foing: %s for host: %s\n", bestMatchingDomainConfig, host)

		
		if domainConfig, err := ROUTING_CONFIGS.GetDomainConfigDetails(bestMatchingDomainConfig); err == nil {
			subdomain := strings.TrimSuffix(host, bestMatchingDomainConfig)
			subdomain = strings.TrimSuffix(subdomain, ".") // remove any trailing "."
	
			utils.InfoLogger.Printf("Extracted subdomain: %s\b", subdomain)
			// get URL by subdomain
			if subdomain != "" {
				if serverURL, err := domainConfig.GetServiceURLBySubdomain(subdomain); err == nil {
					utils.InfoLogger.Printf("found server URL by subdomain routing: %s\n", serverURL)
					return serverURL
				}
			}

			// fallback is get server URL by endpoint
			// remove prefix "/" at endpoint
			if serverURL, err := domainConfig.GetServiceURLByEndPoint(strings.TrimPrefix(endpoint, "/")); err == nil {
				utils.InfoLogger.Printf("found server URL by endpoint routing: %s\n", serverURL)
				return serverURL
			}

			// fallback option -> return default
			utils.InfoLogger.Printf("routing to default service url: %s\n", domainConfig.GetDefaultServiceURL())
			return domainConfig.GetDefaultServiceURL()
		}
	
	}
	return constant.NO_URL
}

func GetResponseFromServer(url string, request *http.Request) http.Response{
	if url == constant.NO_URL {
		//return 404 - no backend service found - return this response to client
		return http.Response{StatusCode: http.StatusNotFound}
	}

	// get response from url with same payload
	requestBody, err := request.GetBody()
	if err != nil {
        utils.ErrorLogger.Printf("error getting request body: %s\n", err)
		//return 400 - bad request - return this response to client
        return http.Response{StatusCode: http.StatusBadRequest}
    }
	reqMethod := request.Method

	newRequest, err := http.NewRequest(reqMethod, url, requestBody)
	if err != nil {
        utils.ErrorLogger.Printf("error creating new request: %s\n", err)
        return http.Response{StatusCode: http.StatusInternalServerError}
    }

	for key, values := range request.Header{
		for _, value := range values {
			newRequest.Header.Add(key, value)
		}
	}

	client := &http.Client{}
	clientResponse, err := client.Do(newRequest)
	if err != nil {
        utils.ErrorLogger.Printf("error making request to server: %s\n", err)
        return http.Response{StatusCode: http.StatusInternalServerError}
    }
	return *clientResponse;
}

func GetFormattedResponseFromServer(request *http.Request) http.Response {

	var host, endpoint = request.Host, request.URL.Path
	utils.InfoLogger.Printf("host: %s, endpoint: %s\n", host, endpoint)

	// get server url from route config
	url := FindServerURL(host, endpoint)

	// call server url and get response
	response := GetResponseFromServer(url, request)

	// format response
	FormatHeaders(&response.Header, constant.RESPONSE)

	// return formatted response
	return response
}

func ParseResponse(){}

func RefactorResponse(){}

func RespondToClient(){}

func DeleteHeaders(headers *http.Header, headersFrom string) {
	var headersToRemove []string
	switch headersFrom {
	case constant.REQUEST:
		headersToRemove = requestHeadersToRemove
	case constant.RESPONSE:
		headersToRemove = responseHeadersToRemove
    default:
		utils.WarnLogger.Printf("no config present to remove headers, ignoring this functionality") 
	}
	
	for key := range *headers {
		if slices.Contains(headersToRemove, key){
			headers.Del(key)
		}
	}
}

func ModifyHeaders(headers *http.Header, headersFrom string){
	var headersToModify map[string]string;
	switch headersFrom {
		case constant.REQUEST:
			headersToModify = requestHeadersToModify
		case constant.RESPONSE:
			headersToModify = responseHeadersToModify
		default:
			utils.WarnLogger.Printf("no config present to modfy headers, ignoring this functionality")
	}

	for key, value := range headersToModify {
		headers.Set(key, value)
	}
}


func FormatHeaders(headers *http.Header, headersFrom string){
	DeleteHeaders(headers, headersFrom)
	ModifyHeaders(headers, headersFrom)
}

func PrintHeaders(req *http.Request){
	for key, value := range req.Header {
		utils.InfoLogger.Print(key, value)
	}
}

func RouteTo(req *http.Request) string {
	hostHeader := req.Header.Get(constant.HOST_HEADER)
	return hostHeader
}

func CastResponseToResponseWriter(serverResponse *http.Response, w *http.ResponseWriter){
	for key, values := range serverResponse.Header {
        for _, value := range values {
            (*w).Header().Add(key, value)
        }
    }
    (*w).WriteHeader(serverResponse.StatusCode)
    io.Copy(*w, serverResponse.Body)
}