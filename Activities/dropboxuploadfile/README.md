# Dropbox Upload File

This activity allows you to upload a file to the Dropbox server. It takes access token, file contents and destination path to where you want to upload the file as an input. In the response you will get the 'Success' string if the file is uploaded successfully. In this you can pass the file in two ways, either by giving path of that file or directly giving contents in a binary format.

## Installation

### Flogo CLI

```
flogo install github.com/Sushil-Takkekar/Flogo_test/Activities/dropboxuploadfile
```

### Schema

```
{
  "inputs":[
    {
      "name": "accessToken",
      "type": "string",
      "required": true
    },
    {
      "name": "sourceType",
      "type": "string",
      "allowed": [
        "File path",
        "Binary data"
      ],
      "value": "File path",
      "required": true
    },
    {
      "name": "dropboxDestPath",
      "type": "string",
      "required": true
    },
    {
      "name": "sourceFilePath",
      "type": "string"
    },
    {
      "name": "binaryContent",
      "type": "any"
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "string"
    }
  ]
}
```

### Activity Input


| Name | Required | Type | Description |
| ---- | -------- | ---- |------------ |
| accessToken | True | String | Access Token of your Dropbox account |
| sourceType  | True | String | Type in which the file is passed for upload. Path or File contents(binary) |
| dropboxDestPath | True | String | Destination path to where you want to upload the file |
| sourceFilePath | True | String | Path of the file |
| binaryContent | True | Any | File contents in a binary format |


### Activity Output


| Name | Type | Description |
| ---- | ---- | ----------- |
| result | String | Success message, if a file is uploaded successfully |

### Example :
This activity will give the response in a following way,

```
Success
```
