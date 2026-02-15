package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/i5hwar-ka1m39h/order_lock/api/db/generated"
)

func (conf *Config) handleEDCreate(w http.ResponseWriter, r *http.Request, email_delivery generated.EmailDelivery) {
	type body struct {
		Recipient  string `json:"recipient"`
		Sender     string `json:"sender"`
		Cc         string `json:"cc"`
		Bcc        string `json:"bcc"`
		Subject    string `json:"subject"`
		Body       string `json:"body"`
		
	}
	decoder := json.NewDecoder(r.Body)

	email_body := body{}
	err := decoder.Decode(&email_body)
	if err != nil{
		log.Println("error while decoding the email body", err)
		errorResponse(w, 400, fmt.Sprintln("error parsing email body", err))
		return
	}

	email_resp, err := conf.DB.CreateEmailDelivery(r.Context(), generated.CreateEmailDeliveryParams{
		Recipient: email_body.Recipient,
		Sender: email_body.Sender,
		Cc: email_delivery.Cc,
		Bcc: email_delivery.Bcc,
		Subject: email_body.Subject,
		Body: email_body.Body,

	})

	if err != nil {
		log.Println("error occured while saving in db", err)
		errorResponse(w, 500, fmt.Sprintln("error occured while saving in db", err))
		return
	}

	jsonResponseWriter(w, 201, email_resp)

}
