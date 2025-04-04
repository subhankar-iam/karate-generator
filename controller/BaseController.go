package controller

import (
	"encoding/json"
	"featureGen/model"
	"featureGen/service"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var baseDir string

func init() {
	godotenv.Load(".env")
	baseDir = os.Getenv("BASE_FILE_PATH")
}

func Orchestrate(req map[string]interface{}) {

	fileName := req["file_name"].(string)
	delete(req, "file_name")
	req_json, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}
	content, err := model.GenerateContent(string(req_json))
	if err != nil {
		panic(err)
	}
	if err := service.SaveFile(fileName, content, baseDir); err != nil {
		fmt.Println(err)
	}
	service.CommitAndPush(baseDir)
}
