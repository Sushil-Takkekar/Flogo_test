package SurveyMonkey_GetResponse

import (
	"github.com/Sushil-Takkekar/Flogo_test/Activity_logic/surveyMonkeyCode"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/tidwall/gjson"
	"fmt"
	"io/ioutil"
	"net/http"
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
func (a *SurveyMonkeyGetResponseActivity) Eval(context activity.Context) (done bool, err error)  {

	// do eval
	fmt.Println("Starting the SurveyMonkey application...")
		//accessToken := "z8UFEI9i5ua1WWhI40S1xo8yLlFJFsOPMdwtsB83YYAJy.1fr.zPLQ9mfrh7a2qTZHqdCwwnMHHn9.U0OvXcyx5SjYLRjcMUsE-YE6mcZAB0fg4lP2zoDNg-sL8fxDoQ"
		//surveyName := "FLG_2_QA_Variety"

		accessToken := context.GetInput("Access_Token").(string)
		surveyName := context.GetInput("Survey_Name").(string)
		
		jsonstr := ""
		jsonSR := ""
		activityOutput := `{ "survey" : { "questions" : [] } }`
		result_return := ""
		err_return := ""

		request, _ := http.NewRequest("GET", "https://api.surveymonkey.com/v3/surveys?title="+surveyName, nil)
		request.Header.Set("Authorization", "bearer "+accessToken)
		request.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		res_surveyID, err_surveyID := client.Do(request)
		surveyID := ""
		if err_surveyID != nil {
			//set return
			err_return = "The HTTP request for getting SurveyID failed with error: "+err_surveyID.Error()
			//context.SetOutput("Response_Json", err_return)
			activityLog.Errorf(err_return)
			return false, err_return
		} else {
			res_surveyID, _ := ioutil.ReadAll(res_surveyID.Body)
			fmt.Println("res_surveyID= ", string(res_surveyID))

			invalidSurveyName := gjson.Get(string(res_surveyID), "data.#").String()
			errorCode := gjson.Get(string(res_surveyID), "error.http_status_code").String()
			if errorCode=="401" {
				//set return
				err_return = gjson.Get(string(res_surveyID), "error.message").String()
				//context.SetOutput("Response_Json", err_return)
				activityLog.Errorf(err_return)
				return false, err_return
			} else if errorCode=="404" {
				//set return
				err_return = gjson.Get(string(res_surveyID), "error.message").String()
				//context.SetOutput("Response_Json", err_return)
				activityLog.Errorf(err_return)
				return false, err_return
			}	else if invalidSurveyName=="0" {
				//set return
				err_return = "Invalid Survey name !!"
				//context.SetOutput("Response_Json", err_return)
				activityLog.Errorf(err_return)
				return false, err_return
			} else {
				surveyID = gjson.Get(string(res_surveyID), "data.0.id").String()
			}
			//fmt.Println("res_surveyID", surveyID)
		}

		link_surveyDetails := "https://api.surveymonkey.com/v3/surveys/"+surveyID+"/details"
		request, _ = http.NewRequest("GET", link_surveyDetails, nil)
		request.Header.Set("Authorization", "bearer "+accessToken)
		request.Header.Set("Content-Type", "application/json")
		client = &http.Client{}
		res_surveyDetails, err_surveyDetails := client.Do(request)
		if err_surveyDetails != nil {
			//set return
			err_return = "The HTTP request for getting SurveyID failed with error: "+err_surveyDetails.Error()
			//context.SetOutput("Response_Json", err_return)
			activityLog.Errorf(err_return)
			return false, err_return
		} else {
			surveyDetails, _ := ioutil.ReadAll(res_surveyDetails.Body)
			//fmt.Println(string(surveyDetails))
			errorCode := gjson.Get(string(surveyDetails), "error.http_status_code").String()
			if errorCode=="404" {
				//set return
				err_return = gjson.Get(string(surveyDetails), "error.message").String()
				//context.SetOutput("Response_Json", err_return)
				activityLog.Errorf(err_return)
				return false, err_return
			}
			//set surveyDetails parent element
			jsonstr = 	`{ "surveydetails" : `+string(surveyDetails)+`}`
		}

		link_surveyResponse := "https://api.surveymonkey.com/v3/surveys/"+surveyID+"/responses/bulk"
		request, _ = http.NewRequest("GET", link_surveyResponse, nil)
		request.Header.Set("Authorization", "bearer "+accessToken)
		request.Header.Set("Content-Type", "application/json")
		client = &http.Client{}
		res_surveyResponse, err_surveyResponse := client.Do(request)
		if err_surveyResponse != nil {
			//set return
			err_return = "The HTTP request for getting SurveyID failed with error: "+err_surveyResponse.Error()
			//context.SetOutput("Response_Json", err_return)
			activityLog.Errorf(err_return)
			return false, err_return
		} else {
			surveyResponse, _ := ioutil.ReadAll(res_surveyResponse.Body)
			//fmt.Println(string(surveyResponse))
			errorCode := gjson.Get(string(surveyResponse), "error.http_status_code").String()
			if errorCode=="404" {
				//set return
				err_return = gjson.Get(string(surveyResponse), "error.message").String()
				//context.SetOutput("Response_Json", err_return)
				activityLog.Errorf(err_return)
				return false, err_return
			}
			//set surveyresponses
			jsonSR = `{ "surveyresponses" : `+string(surveyResponse)+`}`
		}
		//fmt.Println(jsonstr)
		//fmt.Println(jsonSR)
/*-----------------------------------------------------------------------------------------------------------*/

		activityOutput = surveyMonkeyCode.SetSurveyDetails(jsonstr, jsonSR, activityOutput)
		context.SetOutput("Response_Json", activityOutput)
		
	return true, nil
}
