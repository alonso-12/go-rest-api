package joker

// ErrorCode defines supported error codes.
type ErrorCode uint

const (
	CodeUnknown ErrorCode = iota
	CodeInvalidArgument
	CodeInvalidToken
	CodeNotFound
	CodeNoContent
	CodeJsonMarshal
	CodeJsonUnmarshal
	CodeSqlQueryRow
	CodeSqlSelect
)
