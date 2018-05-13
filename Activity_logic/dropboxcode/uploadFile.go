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
	//var res []byte
	var errReadFile error
	switch sourceType {
	case srcFilePath:
		//srcFile, errReadFile := os.Open("D:/tmp.txt")
		srcFile, errReadFile = ioutil.ReadFile(sourceFilePath)
		if errReadFile != nil {
			//activityLog.Errorf(errReadFile.Error())
			return res, errReadFile
		}

	case srcBinaryContent:
		srcFile = binaryContent
	}

	request, _ := http.NewRequest("POST", "https://content.dropboxapi.com/2/files/upload", bytes.NewBuffer(srcFile))
	request.Header.Set("Authorization", "Bearer "+accessToken)
	request.Header.Set("Content-Type", "application/octet-stream")
	request.Header.Set("Dropbox-API-Arg", DropboxAPIArg)
	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resUploadFile, errUploadFile := client.Do(request)
	if errUploadFile != nil {
		//fmt.Printf(errUploadFile.Error())
		//activityLog.Errorf(errUploadFile.Error())
		return res, errUploadFile
	}
	// Close http connection
	defer resUploadFile.Body.Close()
	resUploadFileData, _ := ioutil.ReadAll(resUploadFile.Body)
	if strings.Contains(string(resUploadFileData), "Error in call to API function") {
		//fmt.Println("Error= ", string(resUploadFileData))
		//activityLog.Errorf(string(resUploadFileData))
		return res, errors.New(string(resUploadFileData))
	}
	if strings.Contains(string(resUploadFileData), "Unknown API function") {
		//fmt.Println("Error= ", string(resUploadFileData))
		//activityLog.Errorf(string(resUploadFileData))
		return res, errors.New(string(resUploadFileData))
	}

	var downloaderror Downloaderror
	json.Unmarshal([]byte(string(resUploadFileData)), &downloaderror)
	if downloaderror.ErrorSummary != "" {
		//fmt.Println("error_summary=", downloaderror.ErrorSummary)
		//activityLog.Errorf(downloaderror.ErrorSummary)
		return res, errors.New(downloaderror.ErrorSummary)
	}

	return resUploadFileData, nil
}
