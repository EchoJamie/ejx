/*
Copyright © 2024 jamie HERE <EMAIL ADDRESS>
*/
package request

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Get(url string, reqBody io.Reader, header http.Header) *http.Response {
	return request(url, "GET", nil, header)
}

func Post(url string, reqBody io.Reader, header http.Header) *http.Response {
	return request(url, "POST", reqBody, header)
}

func request(url, method string, reqBody io.Reader, header http.Header) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, reqBody)

	req.Header = header
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求失败:", err)
		os.Exit(1)
	}
	return resp
}

func JsonHeader() http.Header {
	headers := http.Header{}
	headers.Set("Accept", "application/json")
	headers.Set("Content-Type", "application/json")
	return headers
}

/*
	CloseRespBody 关闭响应体
	@Example
		defer closeRespBody(resp.Body)
*/

func CloseRespBody(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		fmt.Println("关闭响应体失败:", err)
	}
}
