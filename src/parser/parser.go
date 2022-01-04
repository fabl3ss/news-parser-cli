package parser

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func formatJSON(data []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", "    ")
	if err == nil {
		return out.Bytes(), err
	}
	return data, nil
}

func ParseWithParams(lang string, coutry string, category string, kwords []string) ([]byte, error) {
	req_link := "https://newsdata.io/api/1/news?apikey=pub_3102ecc2c3dff7232f448fe1c2aca7e8f903&language=" + lang

	if coutry != "" {
		req_link += "&country=" + coutry
	}

	if category != "" {
		req_link += "&category=" + category
	}

	if len(kwords) != 0 {
		req_link += "&q="
		for _, v := range kwords {
			req_link += "%20AND%20" + v
		}
	}

	resp, err := http.Get(req_link)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteJson, _ := ioutil.ReadAll(resp.Body)

	prettyJSON, err := formatJSON(byteJson)
	return prettyJSON, err
}
