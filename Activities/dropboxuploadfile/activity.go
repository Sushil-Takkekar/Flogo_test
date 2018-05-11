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

// Downloaderror is the struct for download error
type Downloaderror struct {
	ErrorSummary string
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
	var errReadFile error
	switch sourceType {
	case srcFilePath:
		//srcFile, errReadFile := os.Open("D:/tmp.txt")
		srcFile, errReadFile = ioutil.ReadFile(sourceFilePath)
		if errReadFile != nil {
			activityLog.Errorf(errReadFile.Error())
			return false, errReadFile
		}

	case srcBinaryContent:
		srcFile = binaryContent
	}

	request, _ := http.NewRequest("POST", "https://content.dropboxapi.com/2/files/upload", bytes.NewBuffer(srcFile))
	request.Header.Set("Authorization", "Bearer "+accessToken)
	request.Header.Set("Content-Type", "application/octet-stream")
	request.Header.Set("Dropbox-API-Arg", DropboxAPIArg)
	client := &http.Client{}
	resUploadFile, errUploadFile := client.Do(request)
	if errUploadFile != nil {
		//fmt.Printf(errUploadFile.Error())
		activityLog.Errorf(errUploadFile.Error())
		return false, errUploadFile
	}
	resUploadFileData, _ := ioutil.ReadAll(resUploadFile.Body)
	if strings.Contains(string(resUploadFileData), "Error in call to API function") {
		//fmt.Println("Error= ", string(resUploadFileData))
		activityLog.Errorf(string(resUploadFileData))
		return false, errors.New(string(resUploadFileData))
	}
	if strings.Contains(string(resUploadFileData), "Unknown API function") {
		//fmt.Println("Error= ", string(resUploadFileData))
		activityLog.Errorf(string(resUploadFileData))
		return false, errors.New(string(resUploadFileData))
	}

	var downloaderror Downloaderror
	json.Unmarshal([]byte(string(resUploadFileData)), &downloaderror)
	if downloaderror.ErrorSummary != "" {
		//fmt.Println("error_summary=", downloaderror.ErrorSummary)
		activityLog.Errorf(downloaderror.ErrorSummary)
		return false, errors.New(downloaderror.ErrorSummary)
	}

	//fmt.Println("resUploadFileData= ", string(resUploadFileData))
	activityLog.Debugf("resUploadFileData: %s", string(resUploadFileData))
	context.SetOutput("result", "Success")

	return true, nil
}
