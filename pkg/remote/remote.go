package remote

import (
	"io/ioutil"
	"net/http"
)

type Client struct {
	Method string
	Url    string
}

func NewClient(method string, url string) *Client {
	return &Client{Method: method, Url: url}
}

func (r *Client) Get() ([]byte, error){
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, r.Url, nil)

	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type Request interface {
	Get()  ([]byte, error)
}