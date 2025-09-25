package dto

type Response struct {
	Code int               `json:"code"`
	Obj  map[string]string `json:"obj,omitempty"`
}
