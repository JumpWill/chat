package constant

type OperationEnums string

const (
	READ   OperationEnums = "READ"
	ADD    OperationEnums = "ADD"
	UPDATE OperationEnums = "UPDATE"
	DELETE OperationEnums = "DELETE"
)

type LogLevel int8

type ResponseCode int

// Log Level
const (
	Debug LogLevel = iota - 1
	Info
	Warn
	Error
	DPanic
	Panic
	Fatal
)

// Response Code
const (
	Success ResponseCode = 200

	Not_Found = 404
	UnAuth
	UnKnown = 500
)
