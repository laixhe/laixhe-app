package utils

const (
	SUCCESS                 = 0
	ERROR_UNKNOWN           = 1
	ERROR_AUTH              = 2
	ERROR_SERVER            = 3
	ERROR_WEBSOCKET         = 4
	ERROR_DB                = 5
	ERROR_NOT_SUPPORT_API   = 6
	ERROR_INVALID_JSON      = 7
	ERROR_PARAMETER_EMPTY   = 8
	ERROR_PARAMETER         = 9
	ERROR_RESOURCES         = 10
	ERROR_RESOURCES_REPEAT  = 11
	ERROR_NOT_LOGGED_IN     = 12
	ERROR_OPERATION_FAILURE = 13
	ERROR_ROUTING_NOT_EXIST = 14
)

var codeText = map[int]string{
	SUCCESS:                 "Success",
	ERROR_UNKNOWN:           "Unknown Error",
	ERROR_AUTH:              "AUTH Error",
	ERROR_SERVER:            "Server Error",
	ERROR_WEBSOCKET:         "WebSocket Error",
	ERROR_DB:                "DB Error",
	ERROR_NOT_SUPPORT_API:   "Not Support Api",
	ERROR_INVALID_JSON:      "Invalid json",
	ERROR_PARAMETER_EMPTY:   "Parameter Empty",
	ERROR_PARAMETER:         "Parameter Error",
	ERROR_RESOURCES:         "Resources Error",
	ERROR_RESOURCES_REPEAT:  "Resources Repeat",
	ERROR_NOT_LOGGED_IN:     "Not Logged In",
	ERROR_OPERATION_FAILURE: "Operation Failure",
	ERROR_ROUTING_NOT_EXIST: "Routing Not Exist",
}

func CodeText(code int) string {
	return codeText[code]
}
