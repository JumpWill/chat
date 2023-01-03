package constant

type OperationEnums string

const (
	READ   OperationEnums = "READ"
	ADD    OperationEnums = "ADD"
	UPDATE OperationEnums = "UPDATE"
	DELETE OperationEnums = "DELETE"
)

type LogLevel int8

const (
	Debug LogLevel = iota - 1
	Info
	Warn
	Error
	DPanic
	Panic
	Fatal
)
