package dropboxcode

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	srcFilePath      = "File path"
	srcBinaryContent = "Binary data"
)

// Downloaderror is the struct for download error
type Downloaderror struct {
	ErrorSummary string
}

func UploadFile(accessToken string, sourceType string, DropboxAPIArg string, sourceFilePath string, binaryContent []byte) (result []byte, err error) {
	// Get source file contents
	var srcFile, res []byte
	var errReadFile error

	// Check source type
	switch sourceType {
	case srcFilePath:
		srcFile, errReadFile = ioutil.ReadFile(sourceFilePath)
		if errReadFile != nil {
			return res, errReadFile
		}

	case srcBinaryContent:
		srcFile = binaryContent
	}

	// Make api call for upload
	request, _ := http.NewRequest("POST", "https://content.dropboxapi.com/2/files/upload", bytes.NewBuffer(srcFile))
	request.Header.Set("Authorization", "Bearer "+accessToken)
	request.Header.Set("Content-Type", "application/octet-stream")
	request.Header.Set("Dropbox-API-Arg", DropboxAPIArg)
	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resUploadFile, errUploadFile := client.Do(request)
	if errUploadFile != nil {
		return res, errUploadFile
	}
	// Close http connection
	defer resUploadFile.Body.Close()
	resUploadFileData, _ := ioutil.ReadAll(resUploadFile.Body)
	if strings.Contains(string(resUploadFileData), "Error in call to API function") {
		return res, errors.New(string(resUploadFileData))
	}
	if strings.Contains(string(resUploadFileData), "Unknown API function") {
		return res, errors.New(string(resUploadFileData))
	}

	var downloaderror Downloaderror
	json.Unmarshal([]byte(string(resUploadFileData)), &downloaderror)
	if downloaderror.ErrorSummary != "" {
		return res, errors.New(downloaderror.ErrorSummary)
	}

	return resUploadFileData, nil
}
