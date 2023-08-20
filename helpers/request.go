package helpers

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func ParseJsonBody(body io.ReadCloser) (map[string]interface{}, error) {
	resBody, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	resStr := string(resBody)
	resbytes := []byte(resStr)
	var jsonRes map[string]interface{}
	err = json.Unmarshal(resbytes, &jsonRes)
	if err != nil {
		return nil, err
	}
	return jsonRes, nil
}
