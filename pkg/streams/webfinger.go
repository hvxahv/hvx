package streams

import (
	"io/ioutil"
	"net/http"
)

type Client struct {
	Method string
	Url    string
}

func NewRequest(method string, url string) *Client {
	return &Client{Method: method, Url: url}
}

func (r *Client) Get() ([]byte, error){
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, r.Url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/activity+json; charset=utf-8")
	req.Header.Add("Accept", "application/ld+json")

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
	Get() ([]byte, error)
}

