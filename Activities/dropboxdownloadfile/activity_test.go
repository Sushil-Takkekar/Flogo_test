package dropboxdownloadfile

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

	/*---------------------------------------------------------------------------------------------------------*/
	//TestCase No-1
	//setup attrs
	tc.SetInput("accessToken", "XO7WTFIqKvUAAAAAAAABdP3i3khVOQ7TBNPP-Gm3rg9GbtUl3TEH90MG3cNZ0-i-")
	tc.SetInput("downloadSourcePath", "/test/hello.txt")
	act.Eval(tc)
	//check result attr
	fileContents := tc.GetOutput("fileContents")
	// Read expected data from file
	expFile, errExpFile := ioutil.ReadFile("D:/Flogo/hello.txt")
	if errExpFile != nil {
		assert.Equal(t, "", fileContents)
	} else {
		assert.Equal(t, expFile, fileContents)
	}

	/*---------------------------------------------------------------------------------------------------------*/
	//TestCase No-2
	//setup attrs
	tc.SetInput("accessToken", "XO7WTFIqKvUAAAAAAAABdP3i3khVOQ7TBNPP-Gm3rg9GbtUl3TEH90MG3cNZ0-i-")
	tc.SetInput("downloadSourcePath", "/test/sample.zip")
	act.Eval(tc)
	//check result attr
	fileContents2 := tc.GetOutput("fileContents")
	// Read expected data from file
	expFile1, errExpFile1 := ioutil.ReadFile("D:/Flogo/sample.zip")
	if errExpFile1 != nil {
		assert.Equal(t, "", fileContents2)
	} else {
		assert.Equal(t, expFile1, fileContents2)
	}
}
