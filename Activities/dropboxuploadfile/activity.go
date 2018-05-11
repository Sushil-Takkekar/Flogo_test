package dropboxuploadfile

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

const (
	srcFilePath      = "File path"
	srcBinaryContent = "Binary data"
)

type Downloaderror struct {
	Error_summary string
}

// ActivityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-flogo-Dropbox_UploadFile")

// DropboxUploadFileActivity is a stub for your Activity implementation
type DropboxUploadFileActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &DropboxUploadFileActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *DropboxUploadFileActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *DropboxUploadFileActivity) Eval(context activity.Context) (done bool, err error) {
	// Initialize parameters
	// accessToken := "XO7WTFIqKvUAAAAAAAABdP3i3khVOQ7TBNPP-Gm3rg9GbtUl3TEH90MG3cNZ0-i-"
	// DropboxAPIArg := `{"path": "/Home/TestUpload/myconfig.zip","mode": "add","autorename": true,"mute": false}`
	// sourceFile := "D:/BW6/BW6_Export/FilePoller.zip"

	accessToken := context.GetInput("accessToken").(string)
	sourceType := context.GetInput("sourceType").(string)
	DropboxAPIArg := `{"path": "` + context.GetInput("dropboxDestPath").(string) + `","mode": "add","autorename": true,"mute": false}`
	sourceFilePath := context.GetInput("sourceFilePath").(string)
	binaryContent := context.GetInput("binaryContent").([]byte)

	// Get source file contents
	var srcFile []byte
	var err_readFile error
	switch sourceType {
	case srcFilePath:
		//srcFile, err_readFile := os.Open("D:/tmp.txt")
		srcFile, err_readFile = ioutil.ReadFile(sourceFilePath)
		if err_readFile != nil {
			activityLog.Errorf(err_readFile.Error())
			return false, err_readFile
		}

	case srcBinaryContent:
		srcFile = binaryContent
	}

	request, _ := http.NewRequest("POST", "https://content.dropboxapi.com/2/files/upload", bytes.NewBuffer(srcFile))
	request.Header.Set("Authorization", "Bearer "+accessToken)
	request.Header.Set("Content-Type", "application/octet-stream")
	request.Header.Set("Dropbox-API-Arg", DropboxAPIArg)
	client := &http.Client{}
	res_uploadFile, err_uploadFile := client.Do(request)
	if err_uploadFile != nil {
		//fmt.Printf(err_uploadFile.Error())
		activityLog.Errorf(err_uploadFile.Error())
		return false, err_uploadFile
	}
	res_uploadFile_data, _ := ioutil.ReadAll(res_uploadFile.Body)
	if strings.Contains(string(res_uploadFile_data), "Error in call to API function") {
		//fmt.Println("Error= ", string(res_uploadFile_data))
		activityLog.Errorf(string(res_uploadFile_data))
		return false, errors.New(string(res_uploadFile_data))
	}
	if strings.Contains(string(res_uploadFile_data), "Unknown API function") {
		//fmt.Println("Error= ", string(res_uploadFile_data))
		activityLog.Errorf(string(res_uploadFile_data))
		return false, errors.New(string(res_uploadFile_data))
	}

	var downloaderror Downloaderror
	json.Unmarshal([]byte(string(res_uploadFile_data)), &downloaderror)
	if downloaderror.Error_summary != "" {
		//fmt.Println("error_summary=", downloaderror.Error_summary)
		activityLog.Errorf(downloaderror.Error_summary)
		return false, errors.New(downloaderror.Error_summary)
	}

	//fmt.Println("res_uploadFile_data= ", string(res_uploadFile_data))
	activityLog.Debugf("res_uploadFile_data: %s", string(res_uploadFile_data))
	context.SetOutput("result", "Success")

	return true, nil
}
