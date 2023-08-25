package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_GetCampusCurrentUser struct {
}


type Response_Data_GetCampusCurrentUser struct {
	Response_User_GetCampusCurrentUser Response_User_GetCampusCurrentUser `json:"user"`
}

type Response_User_GetCampusCurrentUser struct {
	Response_GetCurrentUser_GetCampusCurrentUser Response_GetCurrentUser_GetCampusCurrentUser `json:"getCurrentUser"`
	Typename       string         `json:"__typename"`
}

type Response_GetCurrentUser_GetCampusCurrentUser struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) GetCampusCurrentUser(variables Request_Variables_GetCampusCurrentUser) (Response_Data_GetCampusCurrentUser, error) {
	request := gql.NewQueryRequest[Request_Variables_GetCampusCurrentUser](
		"query getCampusCurrentUser {\n  user {\n    getCurrentUser {\n      id\n      login\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_GetCampusCurrentUser](ctx, request)
}