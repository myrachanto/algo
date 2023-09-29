package main

import (
	"fmt"
	"log"
	"time"

	"github.com/AndroidStudyOpenSource/mpesa-api-go"
)

const (
	appKey           = "zfeseGPBVO4tIbVtgqGPPSKrISm1Zk2A"
	appSecret        = "zAAuwEauigARNdkC"
	authorizationurl = "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"
)

type Mpesa struct{}

func New() *Mpesa {
	return &Mpesa{}
}

func main() {
	mpesa := New()
	mpesa.MPESAExpress()
}
func (*Mpesa) MPESAExpress() {

	svc, err := mpesa.New(appKey, appSecret, mpesa.SANDBOX)
	if err != nil {
		panic(err)
	}

	res, err := svc.Simulation(mpesa.Express{
		BusinessShortCode: "174379",
		Password:          "MTc0Mzc5YmZiMjc5ZjlhYTliZGJjZjE1OGU5N2RkNzFhNDY3Y2QyZTBjODkzMDU5YjEwZjc4ZTZiNzJhZGExZWQyYzkxOTIwMjMwMzAyMTEzNDM1",
		Timestamp:         fmt.Sprintf("%v", time.Now()),
		TransactionType:   "CustomerPayBillOnline",
		Amount:            501,
		PartyA:            "254729308456",
		PartyB:            "174379",
		PhoneNumber:       "254723174434",
		CallBackURL:       "https://callback.com/path",
		AccountReference:  "chantosweb",
		TransactionDesc:   "Payment of a bill",
	})

	if err != nil {
		log.Println(err)
	}
	log.Println(res)
}
func (*Mpesa) C2B() {
	svc, err := mpesa.New(appKey, appSecret, mpesa.SANDBOX)
	if err != nil {
		panic(err)
	}

	res, err := svc.C2BSimulation(mpesa.C2B{
		ShortCode:     "",
		CommandID:     "",
		Amount:        502,
		Msisdn:        "",
		BillRefNumber: "",
	})

	if err != nil {
		log.Println(err)
	}
	log.Println(res)
}
func (*Mpesa) B2C() {
	svc, err := mpesa.New(appKey, appSecret, mpesa.SANDBOX)
	if err != nil {
		panic(err)
	}

	res, err := svc.B2CRequest(mpesa.B2C{
		InitiatorName:      "",
		SecurityCredential: "",
		CommandID:          "",
		Amount:             503,
		PartyA:             "",
		PartyB:             "",
		Remarks:            "",
		QueueTimeOutURL:    "",
		ResultURL:          "",
		Occassion:          "",
	})

	if err != nil {
		log.Println(err)
	}
	log.Println(res)
}
func (*Mpesa) B2B() {
	svc, err := mpesa.New(appKey, appSecret, mpesa.SANDBOX)
	if err != nil {
		panic(err)
	}

	res, err := svc.B2BRequest(mpesa.B2B{
		Initiator:              "",
		SecurityCredential:     "",
		CommandID:              "",
		SenderIdentifierType:   "",
		RecieverIdentifierType: "",
		Remarks:                "",
		Amount:                 504,
		PartyA:                 "",
		PartyB:                 "",
		AccountReference:       "",
		QueueTimeOutURL:        "",
		ResultURL:              "",
	})

	if err != nil {
		log.Println(err)
	}
	log.Println(res)
}
func (*Mpesa) AccountBalance() {
	svc, err := mpesa.New(appKey, appSecret, mpesa.SANDBOX)
	if err != nil {
		panic(err)
	}

	res, err := svc.BalanceInquiry(mpesa.BalanceInquiry{
		Initiator:          "",
		SecurityCredential: "",
		CommandID:          "",
		PartyA:             "",
		IdentifierType:     "",
		Remarks:            "",
		QueueTimeOutURL:    "",
		ResultURL:          "",
	})

	if err != nil {
		log.Println(err)
	}
	log.Println(res)
}
func (*Mpesa) Reversal() {

	svc, err := mpesa.New(appKey, appSecret, mpesa.SANDBOX)
	if err != nil {
		panic(err)
	}

	res, err := svc.Reversal(mpesa.Reversal{
		Initiator:              "",
		SecurityCredential:     "",
		CommandID:              "",
		TransactionID:          "",
		Amount:                 505,
		ReceiverParty:          "",
		ReceiverIdentifierType: "",
		QueueTimeOutURL:        "",
		ResultURL:              "",
		Remarks:                "",
		Occassion:              "",
	})

	if err != nil {
		log.Println(err)
	}
	log.Println(res)

}
