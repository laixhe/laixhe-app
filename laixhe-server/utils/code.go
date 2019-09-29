package utils

const (
	SUCCESS = iota
	ERROR_UNKNOWN
	ERROR_AUTH
	ERROR_SERVER
	ERROR_DB
	ERROR_NOT_SUPPORT_API
	ERROR_INVALID_JSON
	ERROR_PARAMETER_EMPTY
	ERROR_PARAMETER
	ERROR_RESOURCES
	ERROR_RESOURCES_REPEAT
)

var codeText = map[int]string{
	SUCCESS:                "Success",
	ERROR_UNKNOWN:          "Unknown Error",
	ERROR_AUTH:             "AUTH Error",
	ERROR_SERVER:           "Server Error",
	ERROR_DB:               "DB Error",
	ERROR_NOT_SUPPORT_API:  "Not Support Api",
	ERROR_INVALID_JSON:     "Invalid json",
	ERROR_PARAMETER_EMPTY:  "Parameter Empty",
	ERROR_PARAMETER:        "Parameter Error",
	ERROR_RESOURCES:        "Resources Error",
	ERROR_RESOURCES_REPEAT: "Resources Repeat",
}

func CodeText(code int) string {
	return codeText[code]
}
