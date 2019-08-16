package engine

type Request struct {
	Url string
	ParserFunc func([]byte) ParseRequest
}
type ParseRequest struct {
	Requests []Request
	Items []interface{}
}

func NilParser([]byte) ParseRequest {
	return ParseRequest{}
}
