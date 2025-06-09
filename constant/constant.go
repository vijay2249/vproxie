package constant

const (

	// file extension constants
	ENV string = "env"
	YAML string = "yaml"

	//Header constants
	HOST_HEADER string = "Host"
	REFERER string = "Referer"

	//config constants
	CONFIG_DIR_PATH = "./config/"
	TOTAL_CUSTOM_CONFIGS = 3 //!!!change this if adding more custom types!!!

	//LOGGING constants
	WARN string = "warn"
	INFO string = "info"
	DEBUG string = "debug"
	TRACE string = "trace"
	ERROR string = "error"
	LOG_TIME_FORMAT string = "DD-MM-YYY hh:mm:ss ±hh:mm" //<day>-<month>-<year> <hour>:<minutes>:<seconds> <timezone>
)