package httplib

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func PostJSON(url string, v interface{}) (response []byte, err error) {
	bs, err := json.Marshal(v)
	if err != nil {
		return response, err
	}

	bf := bytes.NewBuffer(bs)

	resp, err := http.Post(url, "application/json", bf)
	if err != nil {
		return response, err
	}

	if resp.StatusCode != 200 {
		return []byte{}, errors.New("status code <> 200")
	}

	if resp.Body == nil {
		return []byte{}, nil
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
