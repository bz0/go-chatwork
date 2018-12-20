package api

import (
	"fmt"
	"net/url"
	"path"
	"net/http"
	"io/ioutil"
)

const (
	APIBaseURL = "https://api.chatwork.com"
	Version    = "v2"
)

func RequestJSON(url string, token string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
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