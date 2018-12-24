package api

import (
	"fmt"
	"net/url"
	"path"
	"net/http"
	"io/ioutil"
	"strings"
)

const (
	APIBaseURL = "https://api.chatwork.com"
	Version    = "v2"
)

func RequestJSON(u *url.URL, method string, token string) ([]byte, error) {
	var req *http.Request
	var err error
	switch method {
		case "GET":
			url := u.String()
			req, err = http.NewRequest(method, url, nil)
		case "POST", "PUT", "DELETE":
			url := u.Scheme + "://" + u.Host + u.Path
			req, err = http.NewRequest(method, url, strings.NewReader(u.RawQuery))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	if err != nil {
		fmt.Println("Error at API request:%#v", err)
	}

	req.Header.Set("X-ChatWorkToken", token)

  client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error at API request:%#v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, nil
}

func buildAPIEndpoint(p string) (*url.URL, error) {
	u, err := url.Parse(APIBaseURL + "/" + Version)
	if err != nil {
		return nil, fmt.Errorf("Parse error: %#v", err)
	}
	u.Path = path.Join(u.Path, p)
	return u, err
}