package models 
type Endpoints struct {

	// Number of resources within the current page.
	Count int64 `json:"count,omitempty"`

	// Details of the queried resources.
	// Example: \"documents\": {\n        \"/codestream/api/executions/b80254a7-fcff-4918-ad88-501d08096337\": {\n            \"project\": \"test-project\",\n            \"id\": \"b80254a7-fcff-4918-ad88-501d08096337\",\n            \"name\": \"test\",\n            \"updatedAt\": \"2019-09-23 13:48:54.483\",\n            \"tags\": [],\n            \"_link\": \"/codestream/api/executions/b80254a7-fcff-4918-ad88-501d08096337\",\n            \"_updateTimeInMicros\": 1569226734483000,\n            \"_createTimeInMicros\": 1569226720988000,\n            \"index\": 1,\n            \"notifications\": [],\n            \"comments\": \"\",\n            \"starred\": {},\n            \"input\": {},\n            \"output\": {},\n            \"stageOrder\": [],\n            \"stages\": {},\n            \"status\": \"COMPLETED\",\n            \"statusMessage\": \"Execution Completed.\",\n            \"_durationInMicros\": 13495000,\n            \"_requestTimeInMicros\": 1569226720988000,\n            \"_executedBy\": \"exampleuser\",\n            \"_pipelineLink\": \"/codestream/api/pipelines/b49898f9-d42d-4f19-8bda-e77a373c41b9\",\n            \"_nested\": false,\n            \"_rollback\": false,\n            \"_inputMeta\": {},\n            \"_outputMeta\": {},\n            \"workspaceResults\": []\n        }\n    }
	Documents map[string]Endpoint `json:"documents,omitempty"`

	// Partial URLs representing the links to the queried resources.
	// Example: /codestream/api/executions/b80254a7-fcff-4918-ad88-501d08096337
	Links []string `json:"links"`

	// Number of resources across all pages.
	TotalCount int64 `json:"totalCount,omitempty"`
}

