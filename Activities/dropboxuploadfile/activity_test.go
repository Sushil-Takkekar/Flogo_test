package dropboxuploadfile

import (
	"fmt"
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

	//TestCase No-1 (File path)
	//setup attrs
	var tmp []byte
	tc.SetInput("accessToken", "XO7WTFIqKvUAAAAAAAABdP3i3khVOQ7TBNPP-Gm3rg9GbtUl3TEH90MG3cNZ0-i-")
	tc.SetInput("sourceType", "File path")
	tc.SetInput("dropboxDestPath", "/Home/TestUpload/myconfig.zip")
	tc.SetInput("sourceFilePath", "D:/BW6/BW6_Export/FilePoller.zip")
	tc.SetInput("binaryContent", tmp)
	act.Eval(tc)
	//check result attr
	result := tc.GetOutput("result")
	assert.Equal(t, "Success", result)

	//TestCase No-2 (Binary data)
	//setup attrs
	tc.SetInput("accessToken", "XO7WTFIqKvUAAAAAAAABdP3i3khVOQ7TBNPP-Gm3rg9GbtUl3TEH90MG3cNZ0-i-")
	tc.SetInput("sourceType", "Binary data")
	tc.SetInput("dropboxDestPath", "/Home/TestUpload/sample.zip")
	tc.SetInput("sourceFilePath", "")
	// Read binary data as a file
	binaryData, err_binaryData := ioutil.ReadFile("D:/Flogo/sample.zip")
	if err_binaryData != nil {
		fmt.Println(err_binaryData.Error())
	}
	tc.SetInput("binaryContent", binaryData)
	act.Eval(tc)
	//check result attr
	result2 := tc.GetOutput("result")
	assert.Equal(t, "Success", result2)
}
