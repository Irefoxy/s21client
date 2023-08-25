package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_GetAsapWidgets struct {
}


type Response_Data_GetAsapWidgets struct {
	Response_Student_GetAsapWidgets Response_Student_GetAsapWidgets `json:"student"`
}

type Response_Student_GetAsapWidgets struct {
	Response_GetAsapWidgetList_GetAsapWidgets Response_GetAsapWidgetList_GetAsapWidgets `json:"getAsapWidgetList"`
	Typename          string            `json:"__typename"`
}

type Response_GetAsapWidgetList_GetAsapWidgets struct {
	Response_WidgetList_GetAsapWidgets []Response_WidgetList_GetAsapWidgets `json:"widgetList"`
	Typename   string       `json:"__typename"`
}

type Response_WidgetList_GetAsapWidgets struct {
	ShortImg       string      `json:"shortImg"`
	ShortTitle     string      `json:"shortTitle"`
	ShortURL       interface{} `json:"shortUrl"`
	StartDate      string      `json:"startDate"`
	FinishDate     string      `json:"finishDate"`
	ShowFinishDate bool        `json:"showFinishDate"`
	FullTitle      string      `json:"fullTitle"`
	Text           string      `json:"text"`
	FullImgURL     string      `json:"fullImgUrl"`
	ADTTypeID      int64       `json:"adtTypeId"`
	ADTWidgetID    int64       `json:"adtWidgetId"`
	Typename       string      `json:"__typename"`
}

func (ctx *RequestContext) GetAsapWidgets(variables Request_Variables_GetAsapWidgets) (Response_Data_GetAsapWidgets, error) {
	request := gql.NewQueryRequest[Request_Variables_GetAsapWidgets](
		"query getAsapWidgets {\n  student {\n    getAsapWidgetList {\n      widgetList {\n        ...AsapWidget\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment AsapWidget on AsapWidgetInfo {\n  shortImg\n  shortTitle\n  shortUrl\n  startDate\n  finishDate\n  showFinishDate\n  fullTitle\n  text\n  fullImgUrl\n  adtTypeId\n  adtWidgetId\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_GetAsapWidgets](ctx, request)
}