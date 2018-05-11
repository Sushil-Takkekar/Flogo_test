package dropboxuploadfile

import (
	"github.com/Sushil-Takkekar/Flogo_test/Activity_logic/dropboxcode"
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

	accessToken := context.GetInput("accessToken").(string)
	sourceType := context.GetInput("sourceType").(string)
	DropboxAPIArg := `{"path": "` + context.GetInput("dropboxDestPath").(string) + `","mode": "add","autorename": true,"mute": false}`
	sourceFilePath := context.GetInput("sourceFilePath").(string)
	binaryContent := context.GetInput("binaryContent").([]byte)

	result, err := dropboxcode.UploadFile(accessToken, sourceType, DropboxAPIArg, sourceFilePath, binaryContent)

	if err != nil {
		activityLog.Errorf(err.Error())
		return false, err
	}

	//fmt.Println("resUploadFileData= ", string(resUploadFileData))
	activityLog.Debugf("resUploadFileData: %s", string(resUploadFileData))
	context.SetOutput("result", "Success")
	return true, nil
}
