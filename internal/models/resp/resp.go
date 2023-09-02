package resp

type APIResponse[T any] struct {
	Code    string `json:"code"`
	Message string `json:"msg"`
	Data    T      `json:"data"`
}
