package HelloWorld

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//TestCase No-1
	//setup attrs
	tc.SetInput("Your_Name", "Sushil")
	tc.SetInput("Your_Age", "22")
	act.Eval(tc)
	//check result attr
	result1 := tc.GetOutput("result")
	assert.Equal(t, result1, "Name= Sushil Age= 22")

	//TestCase No-2
	//setup attrs
	tc.SetInput("Your_Name", "hello")
	tc.SetInput("Your_Age", "23")
	act.Eval(tc)
	//check result attr
	result2 := tc.GetOutput("result")
	assert.Equal(t, result2, "Name= hello Age= 23")

}
