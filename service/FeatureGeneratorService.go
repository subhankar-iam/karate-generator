package service

import (
	"featureGen/helper"
	"featureGen/model"
	"path/filepath"
	"strings"
)

func GenerateFeatureFiles(jsonString string) (string, error) {
	content, err := model.GenerateContent(jsonString)
	if err != nil {
		return "", err
	}
	return content, nil
}

func SaveFile(filePath string, fileContent string, baseDir string) error {
	filePath = strings.Replace(filePath, `classpath:`, ``, -1)
	baseFilePath := filepath.Dir(filePath)
	fileName := filepath.Base(filePath)
	if !strings.Contains(fileName, ".feature") {
		fileName += ".feature"
	}
	err := helper.SaveInDrive(baseFilePath, fileName, fileContent, baseDir)
	if err != nil {
		return err
	}
	return nil
}
