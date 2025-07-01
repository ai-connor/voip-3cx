package main

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type Group struct {
	GroupId    int    `json:"GroupId"`
	Number     string `json:"Number"`
	MemberName string `json:"MemberName"`
	Name       string `json:"Name"`
	Type       string `json:"Type"`
	CanDelete  bool   `json:"CanDelete"`
	Id         int    `json:"Id"`
	Rights     struct {
		RoleName string `json:"RoleName"`
	} `json:"Rights"`
}

type User struct {
	FirstName    string  `json:"FirstName"`
	LastName     string  `json:"LastName"`
	EmailAddress string  `json:"EmailAddress"`
	Number       string  `json:"Number"`
	Id           int     `json:"Id"`
	Groups       []Group `json:"Groups"`
}

type CheckUserExistsResponse struct {
	OdataContext string `json:"@odata.context"`
	Value        []User `json:"value"`
}

type GetUsersResponse struct {
	OdataContext string `json:"@odata.context"`
	Value        []User `json:"value"`
}

type CreateUserRequest struct {
	Require2FA           bool   `json:"Require2FA"`
	FirstName            string `json:"FirstName"`
	LastName             string `json:"LastName"`
	EmailAddress         string `json:"EmailAddress"`
	Number               string `json:"Number"`
	Id                   int    `json:"Id"`
	Language             string `json:"Language"`
	SendEmailMissedCalls bool   `json:"SendEmailMissedCalls"`
	VMEmailOptions       string `json:"VMEmailOptions"`
	PromptSet            string `json:"PromptSet"`
	AccessPassword       string `json:"AccessPassword"`
}

type CreateUserResponse struct {
	OdataContext         string `json:"@odata.context"`
	Enable2FA            bool   `json:"Enable2FA"`
	Require2FA           bool   `json:"Require2FA"`
	FirstName            string `json:"FirstName"`
	LastName             string `json:"LastName"`
	EmailAddress         string `json:"EmailAddress"`
	Number               string `json:"Number"`
	Id                   int    `json:"Id"`
	Language             string `json:"Language"`
	SendEmailMissedCalls bool   `json:"SendEmailMissedCalls"`
	VMEmailOptions       string `json:"VMEmailOptions"`
	PromptSet            string `json:"PromptSet"`
}

type AssignUserDepartmentRoleRequest struct {
	Groups []struct {
		GroupId int `json:"GroupId"`
		Rights  struct {
			RoleName string `json:"RoleName"`
		} `json:"Rights"`
	} `json:"Groups"`
	Id int `json:"Id"`
}

type CreateUserFriendlyUrlRequest struct {
	CallUsEnableChat       bool   `json:"CallUsEnableChat"`
	ClickToCallId          string `json:"ClickToCallId"`
	Id                     int    `json:"Id"`
	WebMeetingFriendlyName string `json:"WebMeetingFriendlyName"`
}

type ValidateUserFriendlyUrlRequest struct {
	Model struct {
		FriendlyName string `json:"FriendlyName"`
		Pair         string `json:"Pair"`
	} `json:"model"`
}

type BatchDeleteUsersRequest struct {
	Ids []int `json:"Ids"`
}

type BatchDeleteUsersResponse struct {
	OdataContext string        `json:"@odata.context"`
	Value        []interface{} `json:"value"`
}

type IUsers interface {
	GetUsers(top int, skip int, filter string, orderBy string, selectFields []string, expandRelatedEntities string) (*GetUsersResponse, error)
	CheckUserSameEmailExists(email string) (*GetUsersResponse, error)
	CreateUser(request CreateUserRequest) (*CreateUserResponse, error)
	AssignRole(userId int, request AssignUserDepartmentRoleRequest) error
	CreateUserFriendlyUrl(userId int, request CreateUserFriendlyUrlRequest) error
	ValidateUserFriendlyUrl(request ValidateUserFriendlyUrlRequest) error
	BatchDeleteUser(request BatchDeleteUsersRequest) (*BatchDeleteUsersResponse, error)
}

func (t *TCXApiClient) Users() IUsers {
	return t.WithTokenRenewal()
}

func (t *TCXApiClient) GetUsers(top int, skip int, filter string, orderBy string, selectFields []string, expandRelatedEntities string) (*GetUsersResponse, error) {
	queryParams := url.Values{}
	queryParams.Add("$top", strconv.Itoa(top))
	queryParams.Add("$skip", strconv.Itoa(skip))
	queryParams.Add("$filter", filter)
	queryParams.Add("$orderBy", orderBy)
	queryParams.Add("$selectFields", strings.Join(selectFields, ", "))
	queryParams.Add("$expandRelatedEntities", expandRelatedEntities)
	r := t.authorizedGetRequest("xapi/v1/Users?" + queryParams.Encode())
	resp, err := t.HttpClient.Do(r)
	return parseResponse[GetUsersResponse](err, resp)
}

func (t *TCXApiClient) CheckUserSameEmailExists(email string) (*GetUsersResponse, error) {
	return t.GetUsers(1, 0, fmt.Sprintf("tolower(EmailAddress) eq '%s'", email), "Number", make([]string, 0), "")
}

func (t *TCXApiClient) CreateUser(request CreateUserRequest) (*CreateUserResponse, error) {
	r := t.authorizedPostRequest("xapi/v1/Users", request)
	resp, err := t.HttpClient.Do(r)
	return parseResponse[CreateUserResponse](err, resp)
}

func (t *TCXApiClient) AssignRole(userId int, request AssignUserDepartmentRoleRequest) error {
	r := t.authorizedPatchRequest(fmt.Sprintf("xapi/v1/Users(%v)", userId), request)
	resp, err := t.HttpClient.Do(r)
	return parseNoResponse(err, resp)
}

func (t *TCXApiClient) CreateUserFriendlyUrl(userId int, request CreateUserFriendlyUrlRequest) error {
	r := t.authorizedPatchRequest(fmt.Sprintf("xapi/v1/Users(%v)", userId), request)
	resp, err := t.HttpClient.Do(r)
	return parseNoResponse(err, resp)
}

func (t *TCXApiClient) ValidateUserFriendlyUrl(request ValidateUserFriendlyUrlRequest) error {
	r := t.authorizedPostRequest("xapi/v1/WebsiteLinks/Pbx.ValidateLink", request)
	resp, err := t.HttpClient.Do(r)
	return parseNoResponse(err, resp)
}

func (t *TCXApiClient) BatchDeleteUser(request BatchDeleteUsersRequest) (*BatchDeleteUsersResponse, error) {
	r := t.authorizedPostRequest("xapi/v1/Users/Pbx.BatchDelete", request)
	resp, err := t.HttpClient.Do(r)
	return parseResponse[BatchDeleteUsersResponse](err, resp)
}
