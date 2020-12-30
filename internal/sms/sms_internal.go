package sms

import (
	"encoding/json"
	"github.com/GhvstCode/twillight/internal/app"
	"github.com/GhvstCode/twillight/internal/utils"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func InternalNewOutgoingMessage(APIClient app.Client, to string, from string, msgbody string, opts utils.SmsOpts) (*ResponseSms, error){

		requestUrl := APIClient.BaseUrl + "/Accounts/" + APIClient.AccountSid + "/Messages.json"
		method := "POST"

	vp := opts.ValidityPeriod
	if vp == "" {
		vp = "14400"
	}
	pf := strconv.FormatBool(opts.ProvideFeedback)
	Data := url.Values{}
	Data.Set("To",to)
	Data.Set("From",from)
	Data.Set("Body",msgbody)
	Data.Set("ValidityPeriod",vp)
	Data.Set("ProvideFeedback",pf)
	Data.Set("StatusCallback",opts.StatusCallback)
	DataReader := strings.NewReader(Data.Encode())

		//payload := strings.NewReader("To=%2B2347032541112&From=%2B16592045850&Body=FOR%20YOU%20BABY&ProvideFeedback=true&MediaUrl=https%3A%2F%2Fdemo.twilio.com%2Fowl.png")

		client := APIClient.Configuration.HTTPClient
		//Errors from the API request usually have a
		req, _ := http.NewRequest(method, requestUrl, DataReader)

		req.BasicAuth()

		req.Header.Add("Authorization", APIClient.BasicAuth)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		res, err := client.Do(req)

		if err != nil {
			return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
		}

		defer res.Body.Close()

	var e app.ErrorResponse
		var r ResponseSms
	if res.StatusCode  != http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
		}
		return nil, &e

	}

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

return &r, nil
}

func InternalNewOutgoingWhatsappMessage(APIClient app.Client, to string, from string, msgbody string, opts utils.SmsOpts) (*ResponseSms, error){

	requestUrl := APIClient.BaseUrl + "/Accounts/" + APIClient.AccountSid + "/Messages.json"
	method := "POST"

	vp := opts.ValidityPeriod
	if vp == "" {
		vp = "14400"
	}
	pf := strconv.FormatBool(opts.ProvideFeedback)
	Data := url.Values{}
	Data.Set("To","whatsapp:" + to)
	Data.Set("From","whatsapp:" + from)
	Data.Set("Body",msgbody)
	Data.Set("ValidityPeriod",vp)
	Data.Set("ProvideFeedback",pf)
	Data.Set("StatusCallback",opts.StatusCallback)
	DataReader := strings.NewReader(Data.Encode())

	//payload := strings.NewReader("To=%2B2347032541112&From=%2B16592045850&Body=FOR%20YOU%20BABY&ProvideFeedback=true&MediaUrl=https%3A%2F%2Fdemo.twilio.com%2Fowl.png")

	client := APIClient.Configuration.HTTPClient
	//Errors from the API request usually have a
	req, _ := http.NewRequest(method, requestUrl, DataReader)

	req.BasicAuth()

	req.Header.Add("Authorization", APIClient.BasicAuth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)

	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	defer res.Body.Close()

	var e app.ErrorResponse
	var r ResponseSms
	if res.StatusCode  != http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
		}
		return nil, &e

	}

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	return &r, nil
}

func InternalNewOutgoingMediaMessage(APIClient app.Client, to string, from string, msgbody string, mediaurl string, opts utils.SmsOpts) (*ResponseSms, error){

	requestUrl := APIClient.BaseUrl + "/Accounts/" + APIClient.AccountSid + "/Messages.json"
	method := "POST"

	vp := opts.ValidityPeriod
	if vp == "" {
		vp = "14400"
	}
	pf := strconv.FormatBool(opts.ProvideFeedback)
	Data := url.Values{}
	Data.Set("To",to)
	Data.Set("From",from)
	Data.Set("Body",msgbody)
	Data.Set("ValidityPeriod",vp)
	Data.Set("ProvideFeedback",pf)
	Data.Set("MediaUrl",mediaurl)
	Data.Set("StatusCallback",opts.StatusCallback)
	DataReader := strings.NewReader(Data.Encode())

	//payload := strings.NewReader("To=%2B2347032541112&From=%2B16592045850&Body=FOR%20YOU%20BABY&ProvideFeedback=true&MediaUrl=https%3A%2F%2Fdemo.twilio.com%2Fowl.png")

	client := APIClient.Configuration.HTTPClient
	//Errors from the API request usually have a
	req, _ := http.NewRequest(method, requestUrl, DataReader)

	req.BasicAuth()

	req.Header.Add("Authorization", APIClient.BasicAuth)


	res, err := client.Do(req)

	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	defer res.Body.Close()

	var e app.ErrorResponse
	var r ResponseSms
	if res.StatusCode  != http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
		}
		return nil, &e

	}

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	return &r, nil
}

func InternalRetrieveAllMessagesMedia(APIClient app.Client, messageSid string) (*ResponseAllMessageMedia, error){

	requestUrl := APIClient.BaseUrl + "/Accounts/" + APIClient.AccountSid + "/Messages" + messageSid+ "/Media.json"
	method := "GET"


	//payload := strings.NewReader("To=%2B2347032541112&From=%2B16592045850&Body=FOR%20YOU%20BABY&ProvideFeedback=true&MediaUrl=https%3A%2F%2Fdemo.twilio.com%2Fowl.png")

	client := APIClient.Configuration.HTTPClient
	//Errors from the API request usually have a
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.BasicAuth()

	req.Header.Add("Authorization", APIClient.BasicAuth)

	res, err := client.Do(req)

	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	defer res.Body.Close()

	var e app.ErrorResponse
	var r ResponseAllMessageMedia
	if res.StatusCode  != http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
		}
		return nil, &e

	}

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	return &r, nil
}

func InternalRetrieveAllMessages(APIClient app.Client) (*ResponseGetAllMessages, error){

	requestUrl := APIClient.BaseUrl + "/Accounts/" + APIClient.AccountSid + "/Messages.json"
	method := "GET"


	//payload := strings.NewReader("To=%2B2347032541112&From=%2B16592045850&Body=FOR%20YOU%20BABY&ProvideFeedback=true&MediaUrl=https%3A%2F%2Fdemo.twilio.com%2Fowl.png")

	client := APIClient.Configuration.HTTPClient
	//Errors from the API request usually have a
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.BasicAuth()

	req.Header.Add("Authorization", APIClient.BasicAuth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)

	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	defer res.Body.Close()

	var e app.ErrorResponse
	var r ResponseGetAllMessages
	if res.StatusCode  != http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
		}
		return nil, &e

	}

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	return &r, nil
}

func InternalRetrieveAMessage(APIClient app.Client, MessageSid string) (*ResponseSms, error){

	requestUrl := APIClient.BaseUrl + "/Accounts/" + APIClient.AccountSid + "/Messages"+ MessageSid +".json"
	method := "GET"


	//payload := strings.NewReader("To=%2B2347032541112&From=%2B16592045850&Body=FOR%20YOU%20BABY&ProvideFeedback=true&MediaUrl=https%3A%2F%2Fdemo.twilio.com%2Fowl.png")

	client := APIClient.Configuration.HTTPClient
	//Errors from the API request usually have a
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.BasicAuth()

	req.Header.Add("Authorization", APIClient.BasicAuth)

	res, err := client.Do(req)

	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	defer res.Body.Close()

	var e app.ErrorResponse
	var r ResponseSms
	if res.StatusCode  != http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
		}
		return nil, &e

	}

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	return &r, nil
}

func InternalUpdateMessage(APIClient app.Client, MessageSid, body string) (*ResponseSms, error){

	requestUrl := APIClient.BaseUrl + "/Accounts/" + APIClient.AccountSid + "/Messages/"+ MessageSid +".json"
	method := "POST"


	//payload := strings.NewReader("To=%2B2347032541112&From=%2B16592045850&Body=FOR%20YOU%20BABY&ProvideFeedback=true&MediaUrl=https%3A%2F%2Fdemo.twilio.com%2Fowl.png")

	Data := url.Values{}
	Data.Set("Body",body)
	DataReader := strings.NewReader(Data.Encode())

	client := APIClient.Configuration.HTTPClient
	//Errors from the API request usually have a
	req, _ := http.NewRequest(method, requestUrl, DataReader)

	req.BasicAuth()

	req.Header.Add("Authorization", APIClient.BasicAuth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)

	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	defer res.Body.Close()

	var e app.ErrorResponse
	var r ResponseSms
	if res.StatusCode  != http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
		}
		return nil, &e

	}

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	return &r, nil
}

func InternalSendMessageFeedback(APIClient app.Client, MessageSid, Outcome string) (*ResponseSendMessageFeedback, error){

	requestUrl := APIClient.BaseUrl + "/Accounts/" + APIClient.AccountSid + "/Messages/"+ MessageSid +"/Feedback.json"
	method := "POST"


	//payload := strings.NewReader("To=%2B2347032541112&From=%2B16592045850&Body=FOR%20YOU%20BABY&ProvideFeedback=true&MediaUrl=https%3A%2F%2Fdemo.twilio.com%2Fowl.png")

	Data := url.Values{}
	Data.Set("Outcome",Outcome)
	DataReader := strings.NewReader(Data.Encode())

	client := APIClient.Configuration.HTTPClient
	//Errors from the API request usually have a
	req, _ := http.NewRequest(method, requestUrl, DataReader)

	req.BasicAuth()

	req.Header.Add("Authorization", APIClient.BasicAuth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)

	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	defer res.Body.Close()

	var e app.ErrorResponse
	var r ResponseSendMessageFeedback
	if res.StatusCode  != http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
		}
		return nil, &e

	}

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	return &r, nil
}

func InternalDeleteMessage(APIClient app.Client, MessageSid string) (*ResponseSms, error){

	requestUrl := APIClient.BaseUrl + "/Accounts/" + APIClient.AccountSid + "/Messages/"+ MessageSid +".json"
	method := "GET"


	//payload := strings.NewReader("To=%2B2347032541112&From=%2B16592045850&Body=FOR%20YOU%20BABY&ProvideFeedback=true&MediaUrl=https%3A%2F%2Fdemo.twilio.com%2Fowl.png")


	client := APIClient.Configuration.HTTPClient
	//Errors from the API request usually have a
	req, _ := http.NewRequest(method, requestUrl, nil)

	req.BasicAuth()

	req.Header.Add("Authorization", APIClient.BasicAuth)
	res, err := client.Do(req)

	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	defer res.Body.Close()

	var e app.ErrorResponse
	var r ResponseSms
	if res.StatusCode  != http.StatusOK {
		err := json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
		}
		return nil, &e

	}

	err = json.NewDecoder(res.Body).Decode(&r)
	if err != nil {
		return nil, &app.ErrorResponse{Code: 0, Message: err.Error()}
	}

	return &r, nil
}

//Type A! which should not be imported.