package dto

type Response struct {
	Code int                    `json:"code"`
	Obj  map[string]interface{} `json:"obj,omitempty"`
}
