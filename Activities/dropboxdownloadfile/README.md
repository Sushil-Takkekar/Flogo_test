# Dropbox Download File

This activity allows you to download a file from the Dropbox server. It takes access token and source path from where you want to download the file as an input. In the response you will get file contents downloaded in a binary format.

## Installation

### Flogo CLI

```
flogo install github.com/Sushil-Takkekar/Flogo_test/Activities/dropboxdownloadfile
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
      "name": "downloadSourcePath",
      "type": "string",
      "required": true
    }
  ],
  "outputs": [
    {
      "name": "fileContents",
      "type": "any"
    }
  ]
}
```

### Activity Input


| Name | Required | Type | Description |
| ---- | -------- | ---- |------------ |
| accessToken | True | String | Access Token of your Dropbox account |
| downloadSourcePath  | True | String | Path from where you want to download |


### Activity Output


| Name | Type | Description |
| ---- | ---- | ----------- |
| fileContents | Any | Contents of downloaded file |

### Example :
This activity will give the response in a following way,

```
Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.
```

### Note :
Do not use this to upload a file larger than 150 MB.
