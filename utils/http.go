package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// HTTPGet -- get 请求
func HTTPGet(uri string) ([]byte, error) {
	response, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error: uri=%v, statusCode=%v", uri, response.StatusCode)
	}

	return ioutil.ReadAll(response.Body)
}
