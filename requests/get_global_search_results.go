package requests

import "github.com/irefoxy/s21client/gql"

type GetGlobalSearchResults_Variables struct {
	SearchString string   `json:"searchString"`
	Items        []string `json:"items"`
	Page         struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	} `json:"page"`
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
		`query getGlobalSearchResults($searchString: String!, $items: [SearchItem]!, $page: PagingInput!) {
            globalSearch {
                searchByText(searchString: $searchString, items: $items, page: $page) {
                    profiles {
                        ...GlobalSearchProfilesSearchResult
                        __typename
                    }
                    projects {
                        ...GlobalSearchProjectsSearchResult
                        __typename
                    }
                    studentCourses {
                        ...GlobalSearchCoursesSearchResult
                        __typename
                    }
                    __typename
                }
                __typename
            }
        }
        
        fragment GlobalSearchProfilesSearchResult on ProfilesSearchResult {
            count
            profiles {
                login
                firstName
                lastName
                level
                avatarUrl
                school {
                    shortName
                    __typename
                }
                __typename
            }
			__typename
        }
        
        fragment GlobalSearchProjectsSearchResult on ProjectsSearchResult {
            count
            projects {
                studentTaskId
                status
                finalPercentage
                finalPoint
                project {
                    goalId
                    goalName
                    __typename
                }
                executionType
                __typename
            }
        }
        
        fragment GlobalSearchCoursesSearchResult on CoursesSearchResult {
            count
            courses {
                goalId
                name
                displayedCourseStatus
                executionType
                finalPercentage
                experience
                courseType
                localCourseId
                goalStatus
                __typename
            }
        }`,
		variables,
	)

	return GqlRequest[GetGlobalSearchResults_Data](ctx, request)
}
