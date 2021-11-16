package types

type Response struct {
	StatusCode int         `json:"-"`
	Data       interface{} `json:"data"`
	Meta       interface{} `json:"meta,omitempty"`
}
