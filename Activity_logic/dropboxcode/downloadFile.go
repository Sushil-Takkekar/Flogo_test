package dropboxcode

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// SetSurveyDetails maps the required data from survey response to output
func DownloadFile(accessToken string, DropboxAPIArg string) (result []byte, err error) {

	var res []byte
	downloadUrl := "https://content.dropboxapi.com/2/files/download"

	request, _ := http.NewRequest("POST", downloadUrl, nil)
	request.Header.Set("Authorization", "Bearer "+accessToken)
	request.Header.Set("Dropbox-API-Arg", DropboxAPIArg)
	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resDownloadFile, errDownloadFile := client.Do(request)
	if errDownloadFile != nil {
		return res, errDownloadFile
	}
	// Close http connection
	defer resDownloadFile.Body.Close()

	resDownloadFileData, _ := ioutil.ReadAll(resDownloadFile.Body)
	if strings.Contains(string(resDownloadFileData), "Error in call to API function") {
		return res, errors.New(string(resDownloadFileData))
	}
	if strings.Contains(string(resDownloadFileData), "Unknown API function") {
		return res, errors.New(string(resDownloadFileData))
	}

	var downloaderror Downloaderror
	json.Unmarshal([]byte(string(resDownloadFileData)), &downloaderror)
	if downloaderror.ErrorSummary != "" {
		return res, errors.New(downloaderror.ErrorSummary)
	}

	return resDownloadFileData, nil
}
