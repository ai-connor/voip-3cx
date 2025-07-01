package main

import (
	"net/url"
)

type CheckDepartmentExistsResponse struct {
	OdataContext string `json:"@odata.context"`
	Value        []struct {
		Name       string `json:"Name"`
		IsDefault  bool   `json:"IsDefault"`
		HasMembers bool   `json:"HasMembers"`
		Number     string `json:"Number"`
		Id         int    `json:"Id"`
	} `json:"value"`
}

type CreateDepartmentRequest struct {
	AllowCallService bool   `json:"AllowCallService"`
	Id               int    `json:"Id"`
	Language         string `json:"Language"`
	Name             string `json:"Name"`
	PromptSet        string `json:"PromptSet"`
	Props            struct {
		LiveChatMaxCount         int    `json:"LiveChatMaxCount"`
		PersonalContactsMaxCount int    `json:"PersonalContactsMaxCount"`
		PromptsMaxCount          int    `json:"PromptsMaxCount"`
		SystemNumberFrom         string `json:"SystemNumberFrom"`
		SystemNumberTo           string `json:"SystemNumberTo"`
		TrunkNumberFrom          string `json:"TrunkNumberFrom"`
		TrunkNumberTo            string `json:"TrunkNumberTo"`
		UserNumberFrom           string `json:"UserNumberFrom"`
		UserNumberTo             string `json:"UserNumberTo"`
	} `json:"Props"`
	TimeZoneId          string `json:"TimeZoneId"`
	DisableCustomPrompt bool   `json:"DisableCustomPrompt"`
}

type CreateDepartmentResponse struct {
	OdataContext string `json:"@odata.context"`
	Name         string `json:"Name"`
	Id           int    `json:"Id"`
	Language     string `json:"Language"`
	Props        struct {
		LiveChatMaxCount         int `json:"LiveChatMaxCount"`
		PersonalContactsMaxCount int `json:"PersonalContactsMaxCount"`
		PromptsMaxCount          int `json:"PromptsMaxCount"`
	} `json:"Props"`
	TimeZoneId string `json:"TimeZoneId"`
}

type CheckLiveChatUrlExistsResponse struct {
	OdataContext string `json:"@odata.context"`
	Value        []struct {
		Id           int    `json:"Id"`
		Link         string `json:"Link"`
		ChatEnabled  bool   `json:"ChatEnabled"`
		CallsEnabled bool   `json:"CallsEnabled"`
		General      struct {
			Greeting       string `json:"Greeting"`
			Authentication string `json:"Authentication"`
		} `json:"General"`
		Advanced struct {
			EnableDirectCall     bool   `json:"EnableDirectCall"`
			CommunicationOptions string `json:"CommunicationOptions"`
		} `json:"Advanced"`
	} `json:"value"`
}

type CreateLiveChatUrlRequest struct {
	Advanced struct {
		CallTitle            string `json:"CallTitle"`
		CommunicationOptions string `json:"CommunicationOptions"`
		EnableDirectCall     bool   `json:"EnableDirectCall"`
		IgnoreQueueOwnership bool   `json:"IgnoreQueueOwnership"`
	} `json:"Advanced"`
	CallsEnabled  bool `json:"CallsEnabled"`
	ChatEnabled   bool `json:"ChatEnabled"`
	DefaultRecord bool `json:"DefaultRecord"`
	DN            struct {
		Id     int    `json:"Id"`
		Name   string `json:"Name"`
		Number string `json:"Number"`
		Type   string `json:"Type"`
	} `json:"DN"`
	General struct {
		AllowSoundNotifications bool   `json:"AllowSoundNotifications"`
		Authentication          string `json:"Authentication"`
		DisableOfflineMessages  bool   `json:"DisableOfflineMessages"`
		Greeting                string `json:"Greeting"`
	} `json:"General"`
	Group   string `json:"Group"`
	Link    string `json:"Link"`
	Name    string `json:"Name"`
	Styling struct {
		Animation string `json:"Animation"`
		Minimized bool   `json:"Minimized"`
	} `json:"Styling"`
	Translations struct {
		GreetingMessage     string `json:"GreetingMessage"`
		StartChatButtonText string `json:"StartChatButtonText"`
		UnavailableMessage  string `json:"UnavailableMessage"`
	} `json:"Translations"`
	Website []string `json:"Website"`
}

type CreateLiveChatUrlResponse struct {
	OdataContext  string `json:"@odata.context"`
	Id            int    `json:"Id"`
	Group         string `json:"Group"`
	Link          string `json:"Link"`
	ChatEnabled   bool   `json:"ChatEnabled"`
	CallsEnabled  bool   `json:"CallsEnabled"`
	DefaultRecord bool   `json:"DefaultRecord"`
}

type ConfigureDepartmentCallRoutingRequest struct {
	Id         int `json:"Id"`
	BreakRoute struct {
		IsPromptEnabled bool `json:"IsPromptEnabled"`
		Route           struct {
			External string `json:"External"`
			Number   string `json:"Number"`
			To       string `json:"To"`
		} `json:"Route"`
	} `json:"BreakRoute"`
	OfficeRoute struct {
		IsPromptEnabled bool `json:"IsPromptEnabled"`
		Route           struct {
			External string `json:"External"`
			Number   string `json:"Number"`
			To       string `json:"To"`
		} `json:"Route"`
	} `json:"OfficeRoute"`
	OutOfOfficeRoute struct {
		IsPromptEnabled bool `json:"IsPromptEnabled"`
		Route           struct {
			External string `json:"External"`
			Number   string `json:"Number"`
			To       string `json:"To"`
		} `json:"Route"`
	} `json:"OutOfOfficeRoute"`
	HolidaysRoute struct {
		IsPromptEnabled bool `json:"IsPromptEnabled"`
		Route           struct {
			External string `json:"External"`
			Number   string `json:"Number"`
			To       string `json:"To"`
		} `json:"Route"`
	} `json:"HolidaysRoute"`
}

type UpdateDepartmentRequest struct {
	Id    int    `json:"Id"`
	Name  string `json:"Name"`
	Props struct {
		LiveChatMaxCount         int    `json:"LiveChatMaxCount"`
		PersonalContactsMaxCount int    `json:"PersonalContactsMaxCount"`
		PromptsMaxCount          int    `json:"PromptsMaxCount"`
		SbcMaxCount              int    `json:"SbcMaxCount"`
		SystemNumberFrom         string `json:"SystemNumberFrom"`
		SystemNumberTo           string `json:"SystemNumberTo"`
		TrunkNumberFrom          string `json:"TrunkNumberFrom"`
		TrunkNumberTo            string `json:"TrunkNumberTo"`
		UserNumberFrom           string `json:"UserNumberFrom"`
		UserNumberTo             string `json:"UserNumberTo"`
	} `json:"Props"`
}

type DeleteDepartmentRequest struct {
	Id int `json:"id"`
}

func (t *TCXApiClient) CheckDepartmentExists(filter string) (*CheckDepartmentExistsResponse, error) {
	queryParams := url.Values{}
	queryParams.Add("$filter", filter)
	r := t.authorizedGetRequest("xapi/v1/Groups?" + queryParams.Encode())
	resp, err := t.HttpClient.Do(r)
	return parseResponse[CheckDepartmentExistsResponse](err, resp)
}

func (t *TCXApiClient) CreateDepartment(body CreateDepartmentRequest) (*CreateDepartmentResponse, error) {
	r := t.authorizedPostRequest("xapi/v1/Groups", body)
	resp, err := t.HttpClient.Do(r)
	return parseResponse[CreateDepartmentResponse](err, resp)
}

func (t *TCXApiClient) CheckLiveChatUrlExists(filter string) (*CheckLiveChatUrlExistsResponse, error) {
	queryParams := url.Values{}
	queryParams.Add("$filter", filter)
	r := t.authorizedGetRequest("xapi/v1/WebsiteLinks?" + queryParams.Encode())
	resp, err := t.HttpClient.Do(r)
	return parseResponse[CheckLiveChatUrlExistsResponse](err, resp)
}

func (t *TCXApiClient) CreateLiveChat(body CreateLiveChatUrlRequest) (*CreateLiveChatUrlResponse, error) {
	r := t.authorizedPostRequest("xapi/v1/WebsiteLinks", body)
	resp, err := t.HttpClient.Do(r)
	return parseResponse[CreateLiveChatUrlResponse](err, resp)
}

func (t *TCXApiClient) ConfigureDepartmentCallRouting(body ConfigureDepartmentCallRoutingRequest) error {
	r := t.authorizedPatchRequest("xapi/v1/Groups", body)
	resp, err := t.HttpClient.Do(r)
	return parseNoResponse(err, resp)
}

func (t *TCXApiClient) DeleteDepartment(departmentId int) error {
	r := t.authorizedPostRequest("xapi/v1/Pbx.DeleteCompanyById", &DeleteDepartmentRequest{
		Id: departmentId,
	})
	resp, err := t.HttpClient.Do(r)
	return parseNoResponse(err, resp)
}

func (t *TCXApiClient) UpdateDepartment(body UpdateDepartmentRequest) error {
	r := t.authorizedPatchRequest("xapi/v1/Groups", body)
	resp, err := t.HttpClient.Do(r)
	return parseNoResponse(err, resp)
}

type IDepartments interface {
	CheckDepartmentExists(filter string) (*CheckDepartmentExistsResponse, error)
	CreateDepartment(body CreateDepartmentRequest) (*CreateDepartmentResponse, error)
	CheckLiveChatUrlExists(filter string) (*CheckLiveChatUrlExistsResponse, error)
	CreateLiveChat(body CreateLiveChatUrlRequest) (*CreateLiveChatUrlResponse, error)
	ConfigureDepartmentCallRouting(body ConfigureDepartmentCallRoutingRequest) error
	DeleteDepartment(departmentId int) error
	UpdateDepartment(body UpdateDepartmentRequest) error
}

func (t *TCXApiClient) Departments() IDepartments {
	return t.WithTokenRenewal()
}
