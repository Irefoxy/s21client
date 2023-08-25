package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_GetGitlabSettingsToken struct {
	TaskID string `json:"taskId"`
}


type Response_Data_GetGitlabSettingsToken struct {
	Student interface{} `json:"student"`
}

func (ctx *RequestContext) GetGitlabSettingsToken(variables Request_Variables_GetGitlabSettingsToken) (Response_Data_GetGitlabSettingsToken, error) {
	request := gql.NewQueryRequest[Request_Variables_GetGitlabSettingsToken](
		"query getGitlabSettingsToken($taskId: ID!) {\n  student {\n    getLinkToPrivateStudentGitlabProjectByTaskId(taskId: $taskId) {\n      runnersToken\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_GetGitlabSettingsToken](ctx, request)
}