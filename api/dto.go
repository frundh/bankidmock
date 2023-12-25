package api

import "time"

type Requirement struct {
	PinCode             bool     `json:"pinCode"`
	MRTD                bool     `json:"mrtd"`
	CardReader          string   `json:"cardReader"`
	CertificatePolicies []string `json:"certificatePolicies"`
	PersonalNumber      string   `json:"personalNumber"`
}

type AuthRequestDTO struct {
	EndUserIP             string      `json:"endUserIp"`
	Requirement           Requirement `json:"requirement"`
	UserVisibleData       string      `json:"userVisibleData"`
	UserNonVisibleData    string      `json:"userNonVisibleData"`
	UserVisibleDataFormat string      `json:"userVisibleDataFormat"`
}

type AuthResponseDTO struct {
	OrderRef       string `json:"orderRef"`
	AutoStartToken string `json:"autoStartToken"`
	QRStartToken   string `json:"qrStartToken"`
	QRStartSecret  string `json:"qrStartSecret"`
}

type CollectRequestDTO struct {
	OrderRef string `json:"orderRef"`
}

type CollectStatus string

var (
	CollectStatusComplete CollectStatus = "complete"
	CollectStatusPending  CollectStatus = "pending"
	CollectStatusFailed   CollectStatus = "failed"
)

type HintCode string

var (
	HintCodeOutstandingTransaction HintCode = "outstandingTransaction"
	HintCodeNoClient               HintCode = "noClient"
	HintCodeStarted                HintCode = "started"
	HintCodeUserMrtd               HintCode = "userMrtd"
	HintCodeUserCallConfirm        HintCode = "userCallConfirm"
	HintCodeUserSign               HintCode = "userSign"
)

type User struct {
	PersonalNumber string `json:"personalNumber"`
	Name           string `json:"name"`
	GivenName      string `json:"givenName"`
	Surname        string `json:"surname"`
}

type Device struct {
	IPAddress        string `json:"ipAddress"`
	UniqueHardwareID string `json:"uhi"`
}

type CompletionData struct {
	User            User      `json:"user"`
	Device          Device    `json:"device"`
	BankIDIssueDate time.Time `json:"bankIdIssueDate"`
	// StepUp
	Signature    string `json:"signature"`
	OCSPResponse string `json:"ocspResponse"`
}

type CollectResponseDTO struct {
	OrderRef string        `json:"orderRef"`
	Status   CollectStatus `json:"status"`
	HintCode HintCode      `json:"hintCode"`
}

type CancelRequestDTO struct {
	OrderRef string `json:"orderRef"`
}
