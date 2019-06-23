package libs

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func HttpRequest(addr string, method string, client *HTTPClient) string {

	httpRequest, err := http.NewRequest(method, addr, nil)
	CheckError(err)

	result := client.SendRequest(httpRequest)
	return result

}

func (client *HTTPClient) SendRequest(request *http.Request) string {

	resp, err := client.Client.Do(request)
	CheckError(err)

	result, err := ReadAll(resp.Body)
	CheckError(err)
	fmt.Println("result:", string(result))

	return string(result)

}

func ReadAll(src io.Reader) ([]byte, error) {
	if src == nil {
		return nil, nil
	}

	return ioutil.ReadAll(src)
}
