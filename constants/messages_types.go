package constants

type RequestType int32

const (
	UNKNOWN RequestType = iota
	INITIALIZE
	INITIALIZED
	TEXT_DOCUMENT_DID_CHANGE
	TEXT_DOCUMENT_DEFINITION
	TEXT_DOCUMENT_COMPLETION
	TEXT_DOCUMENT_CODE_ACTION
	TEXT_DOCUMENT_DID_OPEN
	TEXT_DOCUMENT_HOVER
	SHUTDOWN
	EXIT
)

var requestTypeMap = map[string]RequestType{
	"initialize":              INITIALIZE,
	"initialized":             INITIALIZED,
	"textDocument/didChange":  TEXT_DOCUMENT_DID_CHANGE,
	"textDocument/definition": TEXT_DOCUMENT_DEFINITION,
	"textDocument/completion": TEXT_DOCUMENT_COMPLETION,
	"textDocument/codeAction": TEXT_DOCUMENT_CODE_ACTION,
	"textDocument/didOpen":    TEXT_DOCUMENT_DID_OPEN,
	"textDocument/hover":      TEXT_DOCUMENT_HOVER,
	"shutdown":                SHUTDOWN,
	"exit":                    EXIT,
}

func GetRequestType(method string) RequestType {
	if val, ok := requestTypeMap[method]; ok {
		return val
	}
	return UNKNOWN
}
