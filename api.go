package voip3cx

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type TCXTokenResponse struct {
	TokenType    string      `json:"token_type"`
	ExpiresIn    int         `json:"expires_in"`
	AccessToken  string      `json:"access_token"`
	RefreshToken interface{} `json:"refresh_token"`
}

type TCXApiClient struct {
	PbxFQDN        string
	clientId       string
	clientSecret   string
	tokenStartTime time.Time
	Token          *TCXTokenResponse
	HttpClient     *http.Client
}

func New3CXClient(clientUrl, clientId, clientSecret string) TCXApiClient {
	return TCXApiClient{
		HttpClient:     &http.Client{},
		PbxFQDN:        clientUrl,
		clientId:       clientId,
		clientSecret:   clientSecret,
		tokenStartTime: time.Now(),
		Token:          nil,
	}
}

func (t *TCXApiClient) WithTokenRenewal() *TCXApiClient {
	// Check if token expired, and renew it.
	if t.Token == nil || time.Now().Sub(t.tokenStartTime).Minutes() >= float64(t.Token.ExpiresIn) {
		t.getAuthToken()
	}
	return t
}

func (t *TCXApiClient) authorizedRequest(route string) *http.Request {
	reqPath, _ := url.JoinPath(t.PbxFQDN, route)
	reqUrl, _ := url.Parse(reqPath)
	r := &http.Request{
		URL:    reqUrl,
		Header: make(http.Header),
	}

	r.Header.Set("Authorization", "Bearer "+t.Token.AccessToken)
	return r
}

func (t *TCXApiClient) authorizedGetRequest(route string) *http.Request {
	r := t.authorizedRequest(route)
	r.Method = "GET"
	return r
}

func (t *TCXApiClient) authorizedPatchRequest(route string, body any) *http.Request {
	r := t.authorizedPostRequest(route, body)
	r.Method = "PATCH"
	return r
}

func (t *TCXApiClient) authorizedPostRequest(route string, body any) *http.Request {
	r := t.authorizedRequest(route)
	r.Method = "POST"
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return nil // Handle JSON encoding errors
	}

	r.Body = io.NopCloser(&buf)
	return r
}

func (t *TCXApiClient) getAuthToken() {
	authUrl, _ := url.JoinPath(t.PbxFQDN, "connect/token")

	bodyValues := url.Values{}
	bodyValues.Set("client_id", "server_principal_id")
	bodyValues.Set("client_secret", t.clientSecret)
	bodyValues.Set("grant_type", "client_credentials")
	req, err := http.NewRequest("POST", authUrl, bytes.NewBufferString(bodyValues.Encode()))
	auth := base64.StdEncoding.EncodeToString([]byte(t.clientId + ":" + t.clientSecret))
	if err != nil {
		fmt.Println("Error creating request:", err)
	}
	req.Header.Add("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	res, err := parseResponse[TCXTokenResponse](err, resp)
	t.Token = res
}

func parseResponse[T any](err error, resp *http.Response) (*T, error) {
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, errors.New("error making request")
	}
	if resp.StatusCode >= 400 {
		fmt.Println("Error Returned from request:", resp)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, errors.New("error reading response body")
	}

	var tokenResponse T
	if err = json.Unmarshal(body, &tokenResponse); err != nil {
		log.Println("Failed to parse request body")
		return nil, errors.New("failed to parse request body")
	}
	return &tokenResponse, nil
}

func parseNoResponse(err error, resp *http.Response) error {
	_, err2 := parseResponse[any](err, resp)
	return err2
}
