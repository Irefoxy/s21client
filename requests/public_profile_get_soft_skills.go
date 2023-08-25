package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_PublicProfileGetSoftSkills struct {
	StudentID string `json:"studentId"`
}


type Response_Data_PublicProfileGetSoftSkills struct {
	Response_School21_PublicProfileGetSoftSkills Response_School21_PublicProfileGetSoftSkills `json:"school21"`
}

type Response_School21_PublicProfileGetSoftSkills struct {
	Response_GetSoftSkillsByStudentID_PublicProfileGetSoftSkills     []Response_GetSoftSkillsByStudentID_PublicProfileGetSoftSkills   `json:"getSoftSkillsByStudentId"`
	Response_GetSoftSkillLimitByStudentID_PublicProfileGetSoftSkills Response_GetSoftSkillLimitByStudentID_PublicProfileGetSoftSkills `json:"getSoftSkillLimitByStudentId"`
	Typename                     string                       `json:"__typename"`
}

type Response_GetSoftSkillLimitByStudentID_PublicProfileGetSoftSkills struct {
	PowerLimit int64  `json:"powerLimit"`
	Typename   string `json:"__typename"`
}

type Response_GetSoftSkillsByStudentID_PublicProfileGetSoftSkills struct {
	ID                     string `json:"id"`
	Type                   string `json:"type"`
	Code                   string `json:"code"`
	TotalPower             int64  `json:"totalPower"`
	HueSaturationLightness string `json:"hueSaturationLightness"`
	Typename               string `json:"__typename"`
}

func (ctx *RequestContext) PublicProfileGetSoftSkills(variables Request_Variables_PublicProfileGetSoftSkills) (Response_Data_PublicProfileGetSoftSkills, error) {
	request := gql.NewQueryRequest[Request_Variables_PublicProfileGetSoftSkills](
		"query publicProfileGetSoftSkills($studentId: UUID!) {\n  school21 {\n    getSoftSkillsByStudentId(studentId: $studentId) {\n      id\n      type\n      code\n      totalPower\n      hueSaturationLightness\n      __typename\n    }\n    getSoftSkillLimitByStudentId(studentId: $studentId) {\n      powerLimit\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_PublicProfileGetSoftSkills](ctx, request)
}