package surveymonkeygetresponse

import (
	"github.com/Sushil-Takkekar/Flogo_test/Activity_logic/surveymonkeycode"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// ActivityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-flogo-SurveyMonkey_GetResponse")

// SurveyMonkeyGetResponseActivity is a stub for your Activity implementation
type SurveyMonkeyGetResponseActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &SurveyMonkeyGetResponseActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *SurveyMonkeyGetResponseActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *SurveyMonkeyGetResponseActivity) Eval(context activity.Context) (done bool, err error) {

	// do eval
	accessToken := context.GetInput("Access_Token").(string)
	surveyName := context.GetInput("Survey_Name").(string)

	/*-----------------------------------------------------------------------------------------------------------*/

	result, err := surveymonkeycode.SetSurveyDetails(accessToken, surveyName)
	if err != nil {
		activityLog.Errorf(err.Error())
		return false, err
	}

	activityLog.Debugf("Result: %s", result)
	context.SetOutput("Response_Json", result)

	return true, nil
}
