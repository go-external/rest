package rest

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func Unmarshal(body io.ReadCloser, dest any) error {
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		return nil
	}
	defer body.Close()

	if err := json.Unmarshal(bytes, &dest); err != nil {
		return err
	}
	return nil
}
