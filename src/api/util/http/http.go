/**
* @Author: zy
* @Date: 2020/04/04 15:00
 */
package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

const (
	ContentType = "application/json"
)

func NewHttpTool() httpTool {
	return httpTool{}
}

type httpTool struct {
}

func (tool *httpTool) Get(url string, queryParam map[string]string) (int, []byte, error) {
	return tool.handle(url, http.MethodGet, []byte{}, queryParam)
}

func (tool *httpTool) Post(url string, bodyParam []byte) (int, []byte, error) {
	return tool.handle(url, http.MethodPost, bodyParam, map[string]string{})
}

func (tool *httpTool) Put(url string, bodyParam []byte) (int, []byte, error) {
	return tool.handle(url, http.MethodPut, bodyParam, map[string]string{})
}

func (tool *httpTool) Delete(url string, bodyParam []byte) (int, []byte, error) {
	return tool.handle(url, http.MethodDelete, bodyParam, map[string]string{})
}

func (tool *httpTool) handle(url string, method string, bodyParam []byte, queryParam map[string]string) (int, []byte, error) {
	resp := make([]byte, 0)
	body := bytes.NewBuffer([]byte(bodyParam))
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return http.StatusInternalServerError, resp, err
	}
	request.Header.Add("Content-Type", ContentType)
	if method == http.MethodGet {
		query := request.URL.Query()
		for key, value := range queryParam {
			query.Add(key, value)
		}
		request.URL.RawQuery = query.Encode()
	}
	httpClient := &http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		return http.StatusInternalServerError, resp, err
	}
	resp, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return http.StatusInternalServerError, resp, err
	}
	return response.StatusCode, resp, nil
}
