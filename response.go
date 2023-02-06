package rest

import "encoding/json"

type Response interface {
	Marshal() ([]byte, error)
}

type response struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
	Result  any    `json:"result"`
}

func NewResponse(ok bool, msg string, result any) Response {
	return response{Ok: ok, Message: msg, Result: result}
}

func (r response) Marshal() ([]byte, error) {
	bytes, err := json.Marshal(r)
	if err != nil {
		return []byte{}, err
	}
	return bytes, nil
}
