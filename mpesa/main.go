package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	appKey    = "zMG4r5EQwi9uRW0EqhmGTCV0F8KdWhPJ"
	appSecret = "lA2DKCLcvJmyhUIC"
	passkey   = "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919"
	baseUrl   = "https://sandbox.safaricom.co.ke"
	shortcode = "174379"
	partyA    = "254729308456"
	partyB    = "254708161146"
	// partyB   = "254729308456"
	// partyA   = "254708161146"
	callback = "https://callback.com/path"
)

// Mpesa is an application that will be making a transaction
type Mpesa struct {
	consumerKey    string
	consumerSecret string
	baseURL        string
	client         *http.Client
}

// MpesaOpts stores all the configuration keys we need to set up a Mpesa app,
type MpesaOpts struct {
	ConsumerKey    string
	ConsumerSecret string
	BaseURL        string
}

// MpesaAccessTokenResponse is the response sent back by Safaricom when we make a request to generate a token
type MpesaAccessTokenResponse struct {
	AccessToken  interface{} `json:"access_token"`
	ExpiresIn    interface{} `json:"expires_in"`
	RequestID    interface{} `json:"requestId"`
	ErrorCode    interface{} `json:"errorCode"`
	ErrorMessage interface{} `json:"errorMessage"`
}

// type MpesaAccessTokenResponse struct {
// 	AccessToken  string `json:"access_token"`
// 	ExpiresIn    string `json:"expires_in"`
// 	RequestID    string `json:"requestId"`
// 	ErrorCode    string `json:"errorCode"`
// 	ErrorMessage string `json:"errorMessage"`
// }

// STKPushRequestBody is the body with the parameters to be used to initiate an STK push request
type STKPushRequestBody struct {
	BusinessShortCode string `json:"BusinessShortCode"`
	Password          string `json:"Password"`
	Timestamp         string `json:"Timestamp"`
	TransactionType   string `json:"TransactionType"`
	Amount            string `json:"Amount"`
	PartyA            string `json:"PartyA"`
	PartyB            string `json:"PartyB"`
	PhoneNumber       string `json:"PhoneNumber"`
	CallBackURL       string `json:"CallBackURL"`
	AccountReference  string `json:"AccountReference"`
	TransactionDesc   string `json:"TransactionDesc"`
}

// STKPushRequestResponse is the response sent back after initiating an STK push request.
type STKPushRequestResponse struct {
	MerchantRequestID   string `json:"MerchantRequestID"`
	CheckoutRequestID   string `json:"CheckoutRequestID"`
	ResponseCode        string `json:"ResponseCode"`
	ResponseDescription string `json:"ResponseDescription"`
	CustomerMessage     string `json:"CustomerMessage"`
	RequestID           string `json:"requestId"`
	ErrorCode           string `json:"errorCode"`
	ErrorMessage        string `json:"errorMessage"`
}

// NewMpesa sets up and returns an instance of Mpesa
func NewMpesa(m *MpesaOpts) *Mpesa {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	return &Mpesa{
		consumerKey:    m.ConsumerKey,
		consumerSecret: m.ConsumerSecret,
		baseURL:        m.BaseURL,
		client:         client,
	}
}
func main() {
	mpesa := NewMpesa(&MpesaOpts{
		ConsumerKey:    appKey,
		ConsumerSecret: appSecret,
		BaseURL:        baseUrl,
	})

	// fmt.Println("+++++++++++++++++++++++=step 1")
	// The expected format is YYYYMMDDHHmmss
	timestamp := time.Now().Format("20060102150405")
	passwordToEncode := fmt.Sprintf("%s%s%s", shortcode, passkey, timestamp)
	password := base64.StdEncoding.EncodeToString([]byte(passwordToEncode))
	// fmt.Println("------------password", password)
	// fmt.Println("------------timestamp", timestamp)

	accessTokenResponse, err := mpesa.generateAccessToken()
	if err != nil {
		// fmt.Println("+++++++++++++++++++++++=step 1 errr", err)
		log.Fatalln(err)
	}

	// fmt.Println("+++++++++++++++++++++++=step 2")
	fmt.Printf("%+v\n", accessTokenResponse)

	// base64 encoding of the shortcode + passkey + timestamp
	// passwordToEncode := fmt.Sprintf("%s%s%s", shortcode, passkey, timestamp)
	// password := base64.StdEncoding.EncodeToString([]byte(passwordToEncode))
	// fmt.Println("------------password", password)

	response, err := mpesa.initiateSTKPushRequest(&STKPushRequestBody{
		BusinessShortCode: shortcode,
		Password:          password,
		Timestamp:         timestamp,
		TransactionType:   "CustomerPayBillOnline",
		Amount:            "1",    // Amount to be charged when checking out
		PartyA:            partyA, // 2547XXXXXXXX
		PartyB:            shortcode,
		PhoneNumber:       partyB,   // 2547XXXXXXXX
		CallBackURL:       callback, // https://
		AccountReference:  "TEST",
		TransactionDesc:   "Payment via STK push.",
	})

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", response)

}

// generateAccessToken sends a http request to generate new access token
func (m *Mpesa) generateAccessToken() (*MpesaAccessTokenResponse, error) {
	url := fmt.Sprintf("%s/oauth/v1/generate?grant_type=client_credentials", m.baseURL)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// fmt.Println("+++++++++++++++++++++++=step 1 d", req)
	req.SetBasicAuth(m.consumerKey, m.consumerSecret)
	req.Header.Set("Content-Type", "application/json")

	resp, err := m.makeRequest(req)
	if err != nil {
		return nil, err
	}

	accessTokenResponse := &MpesaAccessTokenResponse{}

	// fmt.Println("+++++++++++++++++++++++=step 1 d", string(resp))

	if err := json.Unmarshal(resp, &accessTokenResponse); err != nil {
		// fmt.Println("+++++++++++++++++++++++=step 1 d", err)
		return nil, err
	}

	return accessTokenResponse, nil
}

// initiateSTKPushRequest makes a http request performing an STK push request
func (m *Mpesa) initiateSTKPushRequest(body *STKPushRequestBody) (*STKPushRequestResponse, error) {
	url := fmt.Sprintf("%s/mpesa/stkpush/v1/processrequest", m.baseURL)

	requestBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	accessTokenResponse, err := m.generateAccessToken()
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessTokenResponse.AccessToken))

	resp, err := m.makeRequest(req)
	if err != nil {
		return nil, err
	}

	stkPushResponse := new(STKPushRequestResponse)
	if err := json.Unmarshal(resp, &stkPushResponse); err != nil {
		return nil, err
	}

	return stkPushResponse, nil
}

// makeRequest performs all the http requests for the specific app
func (m *Mpesa) makeRequest(req *http.Request) ([]byte, error) {
	resp, err := m.client.Do(req)
	if err != nil {
		return nil, err
	}

	// fmt.Println("+++++++++++++++++++++++=step 1 make request a")
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	// fmt.Println("+++++++++++++++++++++++=step 1 make request b", resp)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// fmt.Println("+++++++++++++++++++++++=step 1 make request c", body)
	return body, nil
}
