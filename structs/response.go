package structs

type Response struct {
	Error   bool        `json:"error"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}