package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetMySuggestedActivities struct {
	Variables_Page_GetMySuggestedActivities Variables_Page_GetMySuggestedActivities `json:"page"`
}

type Variables_Page_GetMySuggestedActivities struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}


type Data_GetMySuggestedActivities struct {
	Data_School21_GetMySuggestedActivities Data_School21_GetMySuggestedActivities `json:"school21"`
}

type Data_School21_GetMySuggestedActivities struct {
	GetMySuggestedActivities []interface{} `json:"getMySuggestedActivities"`
	Typename                 string        `json:"__typename"`
}


func (ctx *RequestContext) GetMySuggestedActivities(variables Variables_GetMySuggestedActivities) (Data_GetMySuggestedActivities, error) {
	request := gql.NewQueryRequest[Variables_GetMySuggestedActivities](
		"query getMySuggestedActivities($page: PagingInput!, $statuses: [ParticipantEventStatus]) {\n  school21 {\n    getMySuggestedActivities(page: $page, statuses: $statuses) {\n      id\n      start\n      end\n      eventType\n      description\n      eventCode\n      activity {\n        organizers {\n          id\n          login\n          __typename\n        }\n        eventCreator {\n          id\n          login\n          __typename\n        }\n        comments {\n          type\n          createTs\n          comment\n          __typename\n        }\n        averageFeedbackRating\n        isVisible\n        activityType\n        status\n        activityEventId\n        eventId\n        name\n        description\n        location\n        currentStudentsCount\n        maxStudentCount\n        isRegistered\n        isInWaitList\n        isWaitListActive\n        stopRegisterDate\n        beginDate\n        endDate\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetMySuggestedActivities](ctx, request)
}