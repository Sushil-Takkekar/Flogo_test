# SurveyMonkey GetResponse

This activity allows you to get the responses for your survey.

### Activity Input


| Name | Required | Type | Description |
| ---- | -------- | ---- |------------ |
| Access_Token | Yes | String | Access Token of your surveymonkey App |
| Survey_Name  | Yes | String | Name of the survey |


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
                                "choice_id": "1969874111",
                                "row_id": "",
                                "text": "",
                                "title": "Male"
                            }
                        ],
                        "id": "6807750266"
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
