package requests

import "github.com/irefoxy/s21client/gql"

type GetGlobalSearchResults_Variables struct {
	SearchString string                           `json:"searchString"`
	Items        []string                         `json:"items"`
	Page         GetGlobalSearchResults_Data_Page `json:"page"`
}

type GetGlobalSearchResults_Data_Page struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type Profile struct {
	Login     string `json:"login"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Level     int64  `json:"level"`
	AvatarURL string `json:"avatarUrl"`
	School    School `json:"school"`
	Typename  string `json:"__typename"`
}

type School struct {
	ShortName string `json:"shortName"`
	Typename  string `json:"__typename"`
}

type ProfilesSearchResult struct {
	Count    int64     `json:"count"`
	Profiles []Profile `json:"profiles"`
	Typename string    `json:"__typename"`
}

type GlobalSearchResult struct {
	Profiles       ProfilesSearchResult `json:"profiles"`
	Projects       interface{}          `json:"projects"`
	StudentCourses interface{}          `json:"studentCourses"`
	Typename       string               `json:"__typename"`
}

type GetGlobalSearchResults_Data struct {
	GlobalSearch GlobalSearchResult `json:"globalSearch"`
	Typename     string             `json:"__typename"`
}

func (ctx *RequestContext) GetGlobalSearchResults(variables GetGlobalSearchResults_Variables) (GetGlobalSearchResults_Data, error) {
	request := gql.NewQueryRequest[GetGlobalSearchResults_Variables](
		"query getGlobalSearchResults($searchString: String!, $items: [SearchItem]!, $page: PagingInput!) {\n  globalSearch {\n    searchByText(searchString: $searchString, items: $items, page: $page) {\n      profiles {\n        ...GlobalSearchProfilesSearchResult\n        __typename\n      }\n      projects {\n        ...GlobalSearchProjectsSearchResult\n        __typename\n      }\n      studentCourses {\n        ...GlobalSearchCoursesSearchResult\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment GlobalSearchProfilesSearchResult on ProfilesSearchResult {\n  count\n  profiles {\n    login\n    firstName\n    lastName\n    level\n    avatarUrl\n    school {\n      shortName\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment GlobalSearchProjectsSearchResult on ProjectsSearchResult {\n  count\n  projects {\n    studentTaskId\n    status\n    finalPercentage\n    finalPoint\n    project {\n      goalId\n      goalName\n      __typename\n    }\n    executionType\n    __typename\n  }\n  __typename\n}\n\nfragment GlobalSearchCoursesSearchResult on CoursesSearchResult {\n  count\n  courses {\n    goalId\n    name\n    displayedCourseStatus\n    executionType\n    finalPercentage\n    experience\n    courseType\n    localCourseId\n    goalStatus\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[GetGlobalSearchResults_Data](ctx, request)
}
