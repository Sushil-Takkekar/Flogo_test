# SurveyMonkey_GetResponse

This activity allows you to get the responses for your survey in a json format. It takes survey name and access token as an input and gives all the responses for that survey in a compact json format, which can be used to represent the stats graphically.

## Installation

### Flogo CLI

```
flogo install github.com/Sushil-Takkekar/Flogo_test/Activities/SurveyMonkey_GetResponse
```

### Third-party libraries used 
- #### GJSON : 
	GJSON is a Go package that provides a fast and simple way to get values from a json document. It has features such as one line retrieval, dot notation paths, iteration, and parsing json lines.
- #### SJSON : 
	SJSON is a Go package that provides a very fast and simple way to set a value in a json document. The purpose for this library is to provide efficient json updating in the SurveyMonkey_GetResponse activity.

### Schema

```
{
 "inputs":[
    {
      "name": "Access_Token",
      "required": true,
      "type": "string"
    },
	{
      "name": "Survey_Name",
      "required": true,
      "type": "string"
    }
  ],
  "outputs": [
    {
      "name": "Response_Json",
      "type": "string"
    }
  ]
}
```

### Activity Input


| Name | Required | Type | Description |
| ---- | -------- | ---- |------------ |
| Access_Token | True | String | Access Token of your surveymonkey App |
| Survey_Name  | True | String | Name of the survey |


### Activity Output


| Name | Type | Description |
| ---- | ---- | ----------- |
| Response_Json | String | Survey response output in json format |

### Example :
This activity will give the response in a following way,

```
{
    "survey": {
        "questions": [
              {
                "answers": {
                    "choices": [
                        {
                            "description": "",
                            "id": "1969874111",
                            "is_na": "",
                            "position": "1",
                            "text": "Male",
                            "visible": "true",
                            "weight": ""
                        },
                        {
                            "description": "",
                            "id": "1969875809",
                            "is_na": "",
                            "position": "2",
                            "text": "Female",
                            "visible": "true",
                            "weight": ""
                        },
                        {
                            "description": "",
                            "id": "1969875810",
                            "is_na": "",
                            "position": "3",
                            "text": "Other",
                            "visible": "true",
                            "weight": ""
                        }
                    ],
                    "other": {
                        "apply_all_rows": "",
                        "error_text": "",
                        "id": "",
                        "is_answer_choice": "",
                        "num_chars": "",
                        "num_lines": "",
                        "position": "",
                        "text": "",
                        "visible": ""
                    },
                    "rows": [
                        {
                            "id": ""
                        }
                    ]
                },
                "family": "single_choice",
                "id": "289100901",
                "position": "2",
                "responses": [
                    {
                        "answers": [
                            {
                                "choice_id": "1969874111",
                                "row_id": "",
                                "text": "",
                                "title": "Male"
                            }
                        ],
                        "id": "6807749079"
                    },
                    {
                        "answers": [
                            {
                                "choice_id": "1969875809",
                                "row_id": "",
                                "text": "",
                                "title": "Female"
                            }
                        ],
                        "id": "6807764827"
                    }
                ],
                "subtype": "vertical",
                "title": "Gender_multiplechoice",
                "type": "all",
                "visible": "true"
            }
        ]
    }
}
```

### Note :
You can use this output json to represent your stats in a graphical format.
