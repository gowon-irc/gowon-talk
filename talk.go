package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	apiURL       = "https://gpt2-default-dstdu4u23a-uc.a.run.app/"
	dataTemplate = `{"prefix":"%s","length":"100","temperature":"0.7","top_k":"40"}`
)

type responseJSON struct {
	Text string `json:"text"`
}

func post(s string) (out []byte, err error) {
	body := strings.NewReader(fmt.Sprintf(dataTemplate, s))

	req, err := http.NewRequest("POST", apiURL, body)
	if err != nil {
		return out, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return out, err
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return out, err
	}

	return contents, nil
}

func talk(args string) (msg string, err error) {
	j := &responseJSON{}
	req, err := post(args)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(req, j)
	if err != nil {
		return "", err
	}

	return j.Text, nil
}
