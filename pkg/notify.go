package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

//import "os"

type slackBody struct{
	text   string
}

func SlackNotify(nodeName string) (resp Response, err error) {
	SLACK_URL :=  os.Getenv("SLACK_URL")
	client := &Client{
		Endpoint: SLACK_URL,
		httpCli: &http.Client{
			Timeout: 5,
		},
	}
	body := slackBody{
		nodeName,
	}
	slack_body, err := json.Marshal(body)
	fmt.Println(slack_body)
	req := Request{
		Method:"POST",
		Body: slack_body,
		URL:"",
	}
	resp, err = client.DoRequest(req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
