package notification

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/NaySoftware/go-fcm"

	"bytes"
	"io"
	"net/http"
	"os"
)

func SendNotification() (*fcm.FcmResponseStatus, error) {

	const (
		serverKey = "AAAAlnYz40U:APA91bGxAn5QaUBUFWOGXIFS3RaARSkuUC1gaxQOJefYKF1-D-f9788ehl8zC9oc3_4YOisugx6YwRprfFdNgxqfsWN_F7PP_CXOo9PHznj3i-wZiXsBwSpKGYpO2C3s5qWLZdGwsXYK"
		topic     = "/topics/someTopic"
	)

	var NP fcm.NotificationPayload
	NP.Title = "hello"
	NP.Body = "world"

	data := map[string]string{
		"msg": "Hello World1",
		"sum": "Happy Day",
	}

	ids := []string{
		"e9b2EB_-SeGy34-UEK6J3V:APA91bGi-9ttujaz1qgUo-5YAJAksfm2bZMtVYwsWZmGRbB1soRl6_EJUa_a7qvxMNgjdxQHpzh9kNZwnDLpLZBBKR5P9RhZSdQOX5obHJ_709_TRNrEDY_DEQo86AKz-1Tkpm9lJRgo",
	}

	// xds := []string{
	// 	"e9b2EB_-SeGy34-UEK6J3V:APA91bGi-9ttujaz1qgUo-5YAJAksfm2bZMtVYwsWZmGRbB1soRl6_EJUa_a7qvxMNgjdxQHpzh9kNZwnDLpLZBBKR5P9RhZSdQOX5obHJ_709_TRNrEDY_DEQo86AKz-1Tkpm9lJRgo",
	// }

	c := fcm.NewFcmClient(serverKey)
	c.NewFcmRegIdsMsg(ids, data)
	// c.AppendDevices(xds)
	c.SetNotificationPayload(&NP)
	status, err := c.Send()
	if err == nil {
		status.PrintResults()
	} else {
		fmt.Println(err)
	}

	return status, nil
}

func SendNotificationFromHTTPRequestTest() (string, error) {
	const (
		serverKey     = "AAAAlnYz40U:APA91bGxAn5QaUBUFWOGXIFS3RaARSkuUC1gaxQOJefYKF1-D-f9788ehl8zC9oc3_4YOisugx6YwRprfFdNgxqfsWN_F7PP_CXOo9PHznj3i-wZiXsBwSpKGYpO2C3s5qWLZdGwsXYK"
		topic         = "/topics/someTopic"
		urlFCM        = "https://fcm.googleapis.com/fcm/send"
		authorization = "key = " + serverKey
		content_type  = "application/json"
	)

	data := []byte(`{
		"to" : "e9b2EB_-SeGy34-UEK6J3V:APA91bGi-9ttujaz1qgUo-5YAJAksfm2bZMtVYwsWZmGRbB1soRl6_EJUa_a7qvxMNgjdxQHpzh9kNZwnDLpLZBBKR5P9RhZSdQOX5obHJ_709_TRNrEDY_DEQo86AKz-1Tkpm9lJRgo",
		"notification" : {
			"body" : "Hello guys",
			"title": "Title of Your Notification"
		}
	   }`)

	client := &http.Client{Timeout: 5 * time.Hour}
	req, err := http.NewRequest("POST", urlFCM, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", content_type)
	req.Header.Add("Authorization", authorization)

	resp, err := client.Do(req)
	if os.IsTimeout(err) {
		return "", err
	}

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return string(body), err
	}

	return string(data), nil
}

func SendNotificationFromHTTPRequest(userToken, titleNotif, bodyNotif, trans_no string) (string, error) {
	const (
		serverKey     = "AAAAlnYz40U:APA91bGxAn5QaUBUFWOGXIFS3RaARSkuUC1gaxQOJefYKF1-D-f9788ehl8zC9oc3_4YOisugx6YwRprfFdNgxqfsWN_F7PP_CXOo9PHznj3i-wZiXsBwSpKGYpO2C3s5qWLZdGwsXYK"
		topic         = "/topics/someTopic"
		urlFCM        = "https://fcm.googleapis.com/fcm/send"
		authorization = "key = " + serverKey
		content_type  = "application/json"
	)

	type Notif struct {
		To           string      `json:"to"`
		Data         interface{} `json:"data"`
		Notification interface{} `json:"notification"`
	}

	type Notification struct {
		Body  string `json:"body"`
		Title string `json:"title"`
	}

	type Data struct {
		ReffNo string `json:"reff_no"`
	}

	message := &Notification{
		Title: titleNotif,
		Body:  bodyNotif,
	}

	dataMessage := &Data{
		ReffNo: trans_no,
	}

	bodyMessage := &Notif{
		To:           userToken,
		Data:         dataMessage,
		Notification: message,
	}

	dataJson, _ := json.Marshal(bodyMessage)

	// data := []byte(dataJson)

	client := &http.Client{}
	req, err := http.NewRequest("POST", urlFCM, bytes.NewBuffer(dataJson))
	if err != nil {
		return "error new request", err
	}
	req.Header.Add("Content-Type", content_type)
	req.Header.Add("Authorization", authorization)

	resp, err := client.Do(req)
	if os.IsTimeout(err) {
		return "error time out", err
	}

	if err != nil {
		return "error client Do", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return string(body), err
	}

	return string(userToken), nil
}
