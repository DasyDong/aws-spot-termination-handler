package pkg

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Request struct {
	URL    string
	Method string
	Body   []byte
}

type Response struct {
	Body       []byte
	StatusCode int
}

type Client struct {
	Endpoint      string
	httpCli       *http.Client
}

func (clt *Client) DoRequest(req Request) (resp Response, err error) {
	url := clt.Endpoint + req.URL
	fmt.Println(url)
	httpReq, err := http.NewRequest(req.Method, url, bytes.NewBuffer(req.Body))
	fmt.Println(httpReq.Body)
	if err != nil {
		return
	}
	httpReq.Header.Set("Content-Type", "application/json")
	//httpReq.Header.Set("Accept", "application/json")

	var httpResp *http.Response
	httpResp, err = clt.httpCli.Do(httpReq)
	fmt.Println(httpResp.Body)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	body, err := ioutil.ReadAll(httpResp.Body)
	fmt.Println(body)
	if err != nil {
		return Response{}, err
	}

	if httpResp.StatusCode == 204 {
		return Response{StatusCode: httpResp.StatusCode}, nil
	}

	if httpResp.StatusCode >= 300 {
		return Response{}, fmt.Errorf("%s %s http code: %d, details: %s", req.Method, req.URL, httpResp.StatusCode, body)
	}

	return Response{Body: body, StatusCode: httpResp.StatusCode}, nil
}