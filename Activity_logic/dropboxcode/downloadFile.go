package dropboxcode

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// // Downloaderror is the struct for download error
// type Downloaderror struct {
// 	ErrorSummary string
// }

// SetSurveyDetails maps the required data from survey response to output
func DownloadFile(accessToken string, DropboxAPIArg string) (result []byte, err error) {

	//accessToken := "XO7WTFIqKvUAAAAAAAABdP3i3khVOQ7TBNPP-Gm3rg9GbtUl3TEH90MG3cNZ0-i-"
	//DropboxAPIArg := `{"path": "/Home/t1/tmp.txt"}`

	var res []byte
	request, _ := http.NewRequest("POST", "https://content.dropboxapi.com/2/files/download", nil)
	request.Header.Set("Authorization", "Bearer "+accessToken)
	request.Header.Set("Dropbox-API-Arg", DropboxAPIArg)
	client := &http.Client{
		Timeout: time.Second * 30,
	}
	resDownloadFile, errDownloadFile := client.Do(request)
	if errDownloadFile != nil {
		//fmt.Printf(errDownloadFile.Error())
		//activityLog.Errorf(errDownloadFile.Error())
		return res, errDownloadFile
	}
	// Close http connection
	defer resDownloadFile.Body.Close()
	resDownloadFileData, _ := ioutil.ReadAll(resDownloadFile.Body)
	if strings.Contains(string(resDownloadFileData), "Error in call to API function") {
		//fmt.Println("Error= ", string(resDownloadFileData))
		//activityLog.Errorf(string(resDownloadFileData))
		return res, errors.New(string(resDownloadFileData))
	}
	if strings.Contains(string(resDownloadFileData), "Unknown API function") {
		//fmt.Println("Error= ", string(resDownloadFileData))
		//activityLog.Errorf(string(resDownloadFileData))
		return res, errors.New(string(resDownloadFileData))
	}

	var downloaderror Downloaderror
	json.Unmarshal([]byte(string(resDownloadFileData)), &downloaderror)
	if downloaderror.ErrorSummary != "" {
		//fmt.Println("error_summary=", downloaderror.ErrorSummary)
		//activityLog.Errorf(downloaderror.ErrorSummary)
		return res, errors.New(downloaderror.ErrorSummary)
	}
	//fmt.Println("resDownloadFileData= ", string(resDownloadFileData))
	//activityLog.Debugf("resDownloadFileData: %s", resDownloadFileData)
	//context.SetOutput("fileContents", resDownloadFileData)
	return resDownloadFileData, nil
}
