package dropboxdownloadfile

import (
	"github.com/Sushil-Takkekar/Flogo_test/Activity_logic/dropboxcode"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

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
	accessToken := context.GetInput("accessToken").(string)
	DropboxAPIArg := `{"path": "` + context.GetInput("downloadSourcePath").(string) + `"}`

	result, err := dropboxcode.DownloadFile(accessToken, DropboxAPIArg)

	if err != nil {
		activityLog.Errorf(err.Error())
		return false, err
	}

	//activityLog.Debugf("Result: %s", result)
	context.SetOutput("fileContents", result)
	return true, nil
}
