package model

import (
	"bytes"
	"encoding/json"
	"featureGen/contants"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Response struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

var API_KEY string

func init() {
	err := godotenv.Load()
	if err != nil {
		return
	}
	API_KEY = os.Getenv("GEMINI_API_KEY")

}

func sendGeminiRequest(prompt string, url string) (string, error) {
	request_body := bytes.NewBuffer([]byte(prompt))
	res, err := http.Post(url, "application/json", request_body)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	var response_json string
	if len(response.Candidates) > 0 && len(response.Candidates[0].Content.Parts[0].Text) > 0 {
		response_json = strings.Replace(response.Candidates[0].Content.Parts[0].Text, "```gherkin", "", -1)
		response_json = strings.Replace(response_json, "```", "", -1)
	}
	fmt.Println(response_json)
	return response_json, nil

}

func GenerateContent(json_data string) (string, error) {
	gemini_prompt := fmt.Sprintf(contants.User_Prompt, strings.ReplaceAll(json_data, `"`, `''`))
	gemini_url := fmt.Sprintf(contants.Gemini_url, API_KEY)
	gemini_request := fmt.Sprintf(contants.Gemini_Request, gemini_prompt)

	if resp, err := sendGeminiRequest(gemini_request, gemini_url); err != nil {
		return "", err
	} else {
		return resp, nil
	}

}
