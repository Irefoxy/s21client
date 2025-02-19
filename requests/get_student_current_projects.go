package requests

import "github.com/irefoxy/s21client/gql"

type GetStudentCurrentProjects_Variables struct {
	UserID string `json:"userId"`
}

type GetStudentCurrentProjects_Data struct {
	Student GetStudentCurrentProjects_Data_Student `json:"student"`
}

type GetStudentCurrentProjects_Data_Student struct {
	GetStudentCurrentProjects []GetStudentCurrentProjects_Data_GetStudentCurrentProject `json:"getStudentCurrentProjects"`
	Typename                  string                                                    `json:"__typename"`
}

type GetStudentCurrentProjects_Data_GetStudentCurrentProject struct {
	GoalID                         *int64      `json:"goalId"`
	Name                           string      `json:"name"`
	Description                    string      `json:"description"`
	Experience                     int64       `json:"experience"`
	DateTime                       interface{} `json:"dateTime"`
	FinalPercentage                interface{} `json:"finalPercentage"`
	Laboriousness                  int64       `json:"laboriousness"`
	ExecutionType                  *string     `json:"executionType"`
	GoalStatus                     *string     `json:"goalStatus"`
	CourseType                     *string     `json:"courseType"`
	DisplayedCourseStatus          *string     `json:"displayedCourseStatus"`
	AmountAnswers                  *int64      `json:"amountAnswers"`
	AmountMembers                  *int64      `json:"amountMembers"`
	AmountJoinedMembers            *int64      `json:"amountJoinedMembers"`
	AmountReviewedAnswers          *int64      `json:"amountReviewedAnswers"`
	AmountCodeReviewMembers        *int64      `json:"amountCodeReviewMembers"`
	AmountCurrentCodeReviewMembers *int64      `json:"amountCurrentCodeReviewMembers"`
	GroupName                      string      `json:"groupName"`
	LocalCourseID                  *int64      `json:"localCourseId"`
	Typename                       string      `json:"__typename"`
}

func (ctx *RequestContext) GetStudentCurrentProjects(variables GetStudentCurrentProjects_Variables) (GetStudentCurrentProjects_Data, error) {
	request := gql.NewQueryRequest[GetStudentCurrentProjects_Variables](
		"query getStudentCurrentProjects($userId: ID!) {\n  student {\n    getStudentCurrentProjects(userId: $userId) {\n      ...StudentProjectItem\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment StudentProjectItem on StudentItem {\n  goalId\n  name\n  description\n  experience\n  dateTime\n  finalPercentage\n  laboriousness\n  executionType\n  goalStatus\n  courseType\n  displayedCourseStatus\n  amountAnswers\n  amountMembers\n  amountJoinedMembers\n  amountReviewedAnswers\n  amountCodeReviewMembers\n  amountCurrentCodeReviewMembers\n  groupName\n  localCourseId\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[GetStudentCurrentProjects_Data](ctx, request)
}
