package api

import (
	"fmt"
	"net/url"
	"path"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

const (
	APIBaseURL = "https://api.chatwork.com"
	Version    = "v2"
)

func RequestJSON(url string, token string) (result, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	} 
	req.Header.Set("X-ChatWorkToken", token)
	fmt.Println(token)

    client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error at API request:%#v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var result []Task
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func buildAPIEndpoint(p string) (*url.URL, error) {
	u, err := url.Parse(APIBaseURL + "/" + Version)
	if err != nil {
		return nil, fmt.Errorf("Parse error: %#v", err)
	}
	u.Path = path.Join(u.Path, p)
	return u, err
}