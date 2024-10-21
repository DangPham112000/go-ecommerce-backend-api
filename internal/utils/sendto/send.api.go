package sendto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type MailRequest struct {
	ToEmail     string `json:"toEmail"`
	Subject     string `json:"subject"`
	MessageBody string `json:"messageBody"`
	Attachment  string `json:"attachment"`
}

func SendEmailToJavaAPI(otp string, email string, purpose string) error {
	// URL API
	postURL := "http://localhost:8080/email/send-text"

	// DATA json
	mailRequest := MailRequest{
		ToEmail:     email,
		Subject:     "Verify OTP " + purpose,
		MessageBody: "OTP is " + otp,
		Attachment:  "path/to/email",
	}

	// Convert struct to json
	requestBody, err := json.Marshal(mailRequest)
	if err != nil {
		return err
	}

	// Create request
	req, err := http.NewRequest("POST", postURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	// Set header
	req.Header.Set("Content-Type", "application/json")

	// Execute request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	fmt.Sprintln("Response status: ", res.Status)
	return nil
}
