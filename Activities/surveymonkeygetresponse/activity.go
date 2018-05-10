package surveymonkeygetresponse

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Sushil-Takkekar/Flogo_test/Activity_logic/surveymonkeycode"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/tidwall/gjson"
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
	fmt.Println("Starting the SurveyMonkey application...")
	//accessToken := "z8UFEI9i5ua1WWhI40S1xo8yLlFJFsOPMdwtsB83YYAJy.1fr.zPLQ9mfrh7a2qTZHqdCwwnMHHn9.U0OvXcyx5SjYLRjcMUsE-YE6mcZAB0fg4lP2zoDNg-sL8fxDoQ"
	//surveyName := "FLG_2_QA_Variety"

	accessToken := context.GetInput("Access_Token").(string)
	surveyName := context.GetInput("Survey_Name").(string)

	jsonstr := ""
	jsonSR := ""
	activityOutput := `{ "survey" : { "questions" : [] } }`
	errReturn := ""
	errorCode := ""

	request, _ := http.NewRequest("GET", "https://api.surveymonkey.com/v3/surveys?title="+surveyName, nil)
	request.Header.Set("Authorization", "bearer "+accessToken)
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resSurveyIDData, errSurveyID := client.Do(request)
	surveyID := ""
	if errSurveyID != nil {
		//set return
		errReturn = "The HTTP request for getting SurveyID failed with error: " + errSurveyID.Error()
		activityLog.Errorf(errReturn)
		return false, errors.New(errReturn)
	}
	resSurveyID, _ := ioutil.ReadAll(resSurveyIDData.Body)
	//fmt.Println("resSurveyID= ", string(resSurveyID))

	invalidSurveyName := gjson.Get(string(resSurveyID), "data.#").String()
	errorCode = gjson.Get(string(resSurveyID), "error.http_status_code").String()
	if errorCode == "401" {
		//set return
		errReturn = gjson.Get(string(resSurveyID), "error.message").String()
		activityLog.Errorf(errReturn)
		return false, errors.New(errReturn)
	} else if errorCode == "404" {
		//set return
		errReturn = gjson.Get(string(resSurveyID), "error.message").String()
		activityLog.Errorf(errReturn)
		return false, errors.New(errReturn)
	} else if invalidSurveyName == "0" {
		//set return
		errReturn = "Invalid Survey name !!"
		activityLog.Errorf(errReturn)
		return false, errors.New(errReturn)
	} else {
		surveyID = gjson.Get(string(resSurveyID), "data.0.id").String()
	}
	//fmt.Println("resSurveyID", surveyID)

	linkSurveyDetails := "https://api.surveymonkey.com/v3/surveys/" + surveyID + "/details"
	request, _ = http.NewRequest("GET", linkSurveyDetails, nil)
	request.Header.Set("Authorization", "bearer "+accessToken)
	request.Header.Set("Content-Type", "application/json")
	client = &http.Client{}
	resSurveyDetails, errSurveyDetails := client.Do(request)
	if errSurveyDetails != nil {
		//set return
		errReturn = "The HTTP request for getting SurveyID failed with error: " + errSurveyDetails.Error()
		activityLog.Errorf(errReturn)
		return false, errors.New(errReturn)
	}
	surveyDetails, _ := ioutil.ReadAll(resSurveyDetails.Body)
	//fmt.Println(string(surveyDetails))
	errorCode = gjson.Get(string(surveyDetails), "error.http_status_code").String()
	if errorCode == "404" {
		//set return
		errReturn = gjson.Get(string(surveyDetails), "error.message").String()
		activityLog.Errorf(errReturn)
		return false, errors.New(errReturn)
	}
	//set surveyDetails parent element
	jsonstr = `{ "surveydetails" : ` + string(surveyDetails) + `}`

	linkSurveyResponse := "https://api.surveymonkey.com/v3/surveys/" + surveyID + "/responses/bulk"
	request, _ = http.NewRequest("GET", linkSurveyResponse, nil)
	request.Header.Set("Authorization", "bearer "+accessToken)
	request.Header.Set("Content-Type", "application/json")
	client = &http.Client{}
	resSurveyResponse, errSurveyResponse := client.Do(request)
	if errSurveyResponse != nil {
		//set return
		errReturn = "The HTTP request for getting SurveyID failed with error: " + errSurveyResponse.Error()
		activityLog.Errorf(errReturn)
		return false, errors.New(errReturn)
	}
	surveyResponse, _ := ioutil.ReadAll(resSurveyResponse.Body)
	//fmt.Println(string(surveyResponse))
	errorCode = gjson.Get(string(surveyResponse), "error.http_status_code").String()
	if errorCode == "404" {
		//set return
		errReturn = gjson.Get(string(surveyResponse), "error.message").String()
		activityLog.Errorf(errReturn)
		return false, errors.New(errReturn)
	}
	//set surveyresponses
	jsonSR = `{ "surveyresponses" : ` + string(surveyResponse) + `}`

	/*-----------------------------------------------------------------------------------------------------------*/

	activityOutput = surveymonkeycode.SetSurveyDetails(jsonstr, jsonSR, activityOutput)
	activityLog.Debugf("Response from survey: %s", activityOutput)
	context.SetOutput("Response_Json", activityOutput)

	return true, nil
}
