package dropboxdownloadfile

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

type Downloaderror struct {
	Error_summary string
}

// ActivityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-flogo-Dropbox_DownloadFile")

// DropboxDownloadFileActivity is a stub for your Activity implementation
type DropboxDownloadFileActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &DropboxDownloadFileActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *DropboxDownloadFileActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *DropboxDownloadFileActivity) Eval(context activity.Context) (done bool, err error) {
	// Initialize parameters
	//accessToken := "XO7WTFIqKvUAAAAAAAABdP3i3khVOQ7TBNPP-Gm3rg9GbtUl3TEH90MG3cNZ0-i-"
	//DropboxAPIArg := `{"path": "/Home/t1/tmp.txt"}`

	accessToken := context.GetInput("accessToken").(string)
	DropboxAPIArg := `{"path": "` + context.GetInput("downloadSourcePath").(string) + `"}`

	request, _ := http.NewRequest("POST", "https://content.dropboxapi.com/2/files/download", nil)
	request.Header.Set("Authorization", "Bearer "+accessToken)
	request.Header.Set("Dropbox-API-Arg", DropboxAPIArg)
	client := &http.Client{}
	res_downloadFile, err_downloadFile := client.Do(request)
	if err_downloadFile != nil {
		//fmt.Printf(err_downloadFile.Error())
		activityLog.Errorf(err_downloadFile.Error())
		return false, err_downloadFile
	}
	res_downloadFile_data, _ := ioutil.ReadAll(res_downloadFile.Body)
	if strings.Contains(string(res_downloadFile_data), "Error in call to API function") {
		//fmt.Println("Error= ", string(res_downloadFile_data))
		activityLog.Errorf(string(res_downloadFile_data))
		return false, errors.New(string(res_downloadFile_data))
	}
	if strings.Contains(string(res_downloadFile_data), "Unknown API function") {
		//fmt.Println("Error= ", string(res_downloadFile_data))
		activityLog.Errorf(string(res_downloadFile_data))
		return false, errors.New(string(res_downloadFile_data))
	}

	var downloaderror Downloaderror
	json.Unmarshal([]byte(string(res_downloadFile_data)), &downloaderror)
	if downloaderror.Error_summary != "" {
		//fmt.Println("error_summary=", downloaderror.Error_summary)
		activityLog.Errorf(downloaderror.Error_summary)
		return false, errors.New(downloaderror.Error_summary)
	}

	//fmt.Println("res_downloadFile_data= ", string(res_downloadFile_data))
	activityLog.Debugf("res_downloadFile_data: %s", res_downloadFile_data)
	context.SetOutput("fileContents", res_downloadFile_data)
	return true, nil
}
