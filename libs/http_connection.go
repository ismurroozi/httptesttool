package libs

import (
	"net/http"
)

type HTTPClient struct {
	Client *http.Client
}

type HTTPClientList struct {
	ClientList []HTTPClient
}

func (client *HTTPClient) InitClient() {
	client.Client = &http.Client{}
}

func (clients *HTTPClientList) InitClients(number int) {
	for i := 0; i < number; i++ {
		client := HTTPClient{}
		client.InitClient()
		clients.ClientList = append(clients.ClientList, client)
	}

}
