package constant


var HEADERS_TO_EXCLUDE = []string{
	"user-agent",
	"sec-ch-ua-platform",
	"sec-ch-ua",
	"sec-fetch-user",
	"sec-ch-ua-mobile",
}

var HOST_HEADER string = "Host"

var CONFIG_DIR_PATH = "./config/"