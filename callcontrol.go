package voip3cx

import "fmt"

type GetCallControlResponse struct {
	Dn      string `json:"dn"`
	Type    string `json:"type"`
	Devices []struct {
		Dn        string `json:"dn"`
		DeviceId  string `json:"device_id"`
		UserAgent string `json:"user_agent"`
	} `json:"devices"`
	Participants []struct {
		Id               int    `json:"id"`
		Status           string `json:"status"`
		Dn               string `json:"dn"`
		PartyCallerName  string `json:"party_caller_name"`
		PartyDn          string `json:"party_dn"`
		PartyCallerId    string `json:"party_caller_id"`
		PartyDid         string `json:"party_did"`
		DeviceId         string `json:"device_id"`
		PartyDnType      string `json:"party_dn_type"`
		DirectControl    bool   `json:"direct_control"`
		OriginatedByDn   string `json:"originated_by_dn"`
		OriginatedByType string `json:"originated_by_type"`
		ReferredByDn     string `json:"referred_by_dn"`
		ReferredByType   string `json:"referred_by_type"`
		OnBehalfOfDn     string `json:"on_behalf_of_dn"`
		OnBehalfOfType   string `json:"on_behalf_of_type"`
		Callid           int    `json:"callid"`
		Legid            int    `json:"legid"`
	} `json:"participants"`
}

type GetCallControlDnNumberResponse struct {
	Dn      string `json:"dn"`
	Type    string `json:"type"`
	Devices []struct {
		Dn        string `json:"dn"`
		DeviceId  string `json:"device_id"`
		UserAgent string `json:"user_agent"`
	} `json:"devices"`
	Participants []struct {
		Id               int    `json:"id"`
		Status           string `json:"status"`
		Dn               string `json:"dn"`
		PartyCallerName  string `json:"party_caller_name"`
		PartyDn          string `json:"party_dn"`
		PartyCallerId    string `json:"party_caller_id"`
		PartyDid         string `json:"party_did"`
		DeviceId         string `json:"device_id"`
		PartyDnType      string `json:"party_dn_type"`
		DirectControl    bool   `json:"direct_control"`
		OriginatedByDn   string `json:"originated_by_dn"`
		OriginatedByType string `json:"originated_by_type"`
		ReferredByDn     string `json:"referred_by_dn"`
		ReferredByType   string `json:"referred_by_type"`
		OnBehalfOfDn     string `json:"on_behalf_of_dn"`
		OnBehalfOfType   string `json:"on_behalf_of_type"`
		Callid           int    `json:"callid"`
		Legid            int    `json:"legid"`
	} `json:"participants"`
}

type GetCallControlDevice struct {
	Dn        string `json:"dn"`
	DeviceId  string `json:"device_id"`
	UserAgent string `json:"user_agent"`
}

type GetCallControlParticipant struct {
	Id               int    `json:"id"`
	Status           string `json:"status"`
	Dn               string `json:"dn"`
	PartyCallerName  string `json:"party_caller_name"`
	PartyDn          string `json:"party_dn"`
	PartyCallerId    string `json:"party_caller_id"`
	PartyDid         string `json:"party_did"`
	DeviceId         string `json:"device_id"`
	PartyDnType      string `json:"party_dn_type"`
	DirectControl    bool   `json:"direct_control"`
	OriginatedByDn   string `json:"originated_by_dn"`
	OriginatedByType string `json:"originated_by_type"`
	ReferredByDn     string `json:"referred_by_dn"`
	ReferredByType   string `json:"referred_by_type"`
	OnBehalfOfDn     string `json:"on_behalf_of_dn"`
	OnBehalfOfType   string `json:"on_behalf_of_type"`
	Callid           int    `json:"callid"`
	Legid            int    `json:"legid"`
}

type MakeCallRequest struct {
	Destination  string `json:"destination"`
	Timeout      int    `json:"timeout"`
	Attacheddata struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"attacheddata"`
}

type MakeCallResponse struct {
	Finalstatus string `json:"finalstatus"`
	Reason      string `json:"reason"`
	Result      struct {
		Id               int    `json:"id"`
		Status           string `json:"status"`
		Dn               string `json:"dn"`
		PartyCallerName  string `json:"party_caller_name"`
		PartyDn          string `json:"party_dn"`
		PartyCallerId    string `json:"party_caller_id"`
		PartyDid         string `json:"party_did"`
		DeviceId         string `json:"device_id"`
		PartyDnType      string `json:"party_dn_type"`
		DirectControl    bool   `json:"direct_control"`
		OriginatedByDn   string `json:"originated_by_dn"`
		OriginatedByType string `json:"originated_by_type"`
		ReferredByDn     string `json:"referred_by_dn"`
		ReferredByType   string `json:"referred_by_type"`
		OnBehalfOfDn     string `json:"on_behalf_of_dn"`
		OnBehalfOfType   string `json:"on_behalf_of_type"`
		Callid           int    `json:"callid"`
		Legid            int    `json:"legid"`
	} `json:"result"`
	Reasontext string `json:"reasontext"`
}

type PostParticipantActionRequest struct {
	Reason       string `json:"reason"`
	Destination  string `json:"destination"`
	Timeout      int    `json:"timeout"`
	Attacheddata struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"attacheddata"`
}

func (t *TCXApiClient) GetCallControl() (*[]GetCallControlDevice, error) {
	r := t.authorizedGetRequest("callcontrol")
	resp, err := t.HttpClient.Do(r)
	return parseResponse[[]GetCallControlDevice](err, resp)
}

func (t *TCXApiClient) GetCallControlWebSocket() error {
	r := t.authorizedGetRequest("callcontrol/ws")
	resp, err := t.HttpClient.Do(r)
	return parseNoResponse(err, resp)
}

func (t *TCXApiClient) GetConnections(dnNumber string) (*GetCallControlDnNumberResponse, error) {
	r := t.authorizedGetRequest(fmt.Sprintf("callcontrol/%s", dnNumber))
	resp, err := t.HttpClient.Do(r)
	return parseResponse[GetCallControlDnNumberResponse](err, resp)
}

func (t *TCXApiClient) GetDevices(dnNumber string) (*[]GetCallControlDevice, error) {
	r := t.authorizedGetRequest(fmt.Sprintf("callcontrol/%s/devices", dnNumber))
	resp, err := t.HttpClient.Do(r)
	return parseResponse[[]GetCallControlDevice](err, resp)
}

func (t *TCXApiClient) GetDeviceDetails(dnNumber, deviceId string) (*GetCallControlDevice, error) {
	r := t.authorizedGetRequest(fmt.Sprintf("callcontrol/%s/devices/%s", dnNumber, deviceId))
	resp, err := t.HttpClient.Do(r)
	return parseResponse[GetCallControlDevice](err, resp)
}

func (t *TCXApiClient) GetParticipants(dnNumber string) (*[]GetCallControlParticipant, error) {
	r := t.authorizedGetRequest(fmt.Sprintf("callcontrol/%s/participants", dnNumber))
	resp, err := t.HttpClient.Do(r)
	return parseResponse[[]GetCallControlParticipant](err, resp)
}

func (t *TCXApiClient) GetParticipantDetails(dnNumber string, participantId int) (*GetCallControlParticipant, error) {
	r := t.authorizedGetRequest(fmt.Sprintf("callcontrol/%s/participants/%d", dnNumber, participantId))
	resp, err := t.HttpClient.Do(r)
	return parseResponse[GetCallControlParticipant](err, resp)
}

func (t *TCXApiClient) PostMakeCall(dnNumber string, request MakeCallRequest) (*MakeCallResponse, error) {
	r := t.authorizedPostRequest(fmt.Sprintf("callcontrol/%s/makecall", dnNumber), request)
	resp, err := t.HttpClient.Do(r)
	return parseResponse[MakeCallResponse](err, resp)
}

func (t *TCXApiClient) PostMakeCallByDevice(dnNumber, deviceId string, request MakeCallRequest) (*MakeCallResponse, error) {
	r := t.authorizedPostRequest(fmt.Sprintf("callcontrol/%s/devices/%s/makecall", dnNumber, deviceId), request)
	resp, err := t.HttpClient.Do(r)
	return parseResponse[MakeCallResponse](err, resp)
}

// TODO - Add the stream endpoints, and support ?channels?

func (t *TCXApiClient) PostCallAction(dnNumber string, participantId int, action string, request PostParticipantActionRequest) (*MakeCallResponse, error) {
	r := t.authorizedPostRequest(fmt.Sprintf("callcontrol/%s/participants/%d/%s", dnNumber, participantId, action), request)
	resp, err := t.HttpClient.Do(r)
	return parseResponse[MakeCallResponse](err, resp)
}
