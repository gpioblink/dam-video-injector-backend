package apicall

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type SearchSongInput struct {
	AppVer string `json:"appVer"`
	CategoryCd string `json:"categoryCd"`
	DeviceId string `json:"deviceId"`
	DeviceNm string `json:"deviceNm"`
	OsVer string `json:"osVer"`
	Page string `json:"page"`
	SongMatchType string `json:"songMatchType"`
	SerialNo string `json:"serialNo"`
	SongName string `json:"songName"`
}

type ConnectInput struct {
	AppVer string `json:"appVer"`
	DeviceId string `json:"deviceId"`
	DeviceNm string `json:"deviceNm"`
	OsVer string `json:"osVer"`
	SerialNo string `json:"serialNo"`
	QRcode string `json:"QRcode"`
}

type RemoconInput struct {
	AppVer string `json:"appVer"`
	DeviceId string `json:"deviceId"`
	DeviceNm string `json:"deviceNm"`
	OsVer string `json:"osVer"`
	SerialNo string `json:"serialNo"`
	QRcode string `json:"QRcode"`
	RemoconCode string `json:"remoconCode"`
}

type SendInput struct {
	DeviceNm string `json:"deviceNm"`
	OsVer string `json:"osVer"`
	DeviceId string `json:"deviceId"`
	AppVer string `json:"appVer"`
	QRcode string `json:"QRcode"`
	SerialNo string `json:"serialNo"`
	ReqNo string `json:"reqNo"`
	MyKey int `json:"myKey"`
	SendDate string `json:"sendDate"`
	ContentsKind string `json:"contentsKind"`
	Interrupt int `json:"interrupt"`
}

func SendRequest(qrCode, reqNo string, interrupt int) {
	client := &http.Client{}

	data, _ := json.Marshal(SendInput{
		AppVer:   "2.7.4",
		DeviceId: "f6057f8333624957a6a5bd4d0c49178b997854cf6a6ad0c09219a0864dfb1216",
		DeviceNm: "Android Nexus 6",
		OsVer:    "9",
		SerialNo: "AM103094",
		QRcode:   qrCode,
		ReqNo: reqNo,
		MyKey: 0,
		SendDate: time.Now().Format("2006/01/02 15:04:05"),
		Interrupt: interrupt,
	})

	resp, _ := client.Post(
		"https://denmoku.clubdam.com/dkdenmoku/DkDamSendServlet",
		"application/json",
		bytes.NewBuffer(data),
	)
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}

func ConnectDevice(qrCode string) {
	client := &http.Client{}

	data, _ := json.Marshal(ConnectInput{
		AppVer:   "2.7.4",
		DeviceId: "f6057f8333624957a6a5bd4d0c49178b997854cf6a6ad0c09219a0864dfb1216",
		DeviceNm: "Android Nexus 6",
		OsVer:    "9",
		SerialNo: "AM103094",
		QRcode:   qrCode,
	})

	resp, _ := client.Post(
		"https://denmoku.clubdam.com/dkdenmoku/DkDamConnectServlet",
		"application/json",
		bytes.NewBuffer(data),
	)
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}

func SendRemocon(qrCode string, remoconCode string) {
	client := &http.Client{}

	data, _ := json.Marshal(RemoconInput{
		AppVer:   "2.7.4",
		DeviceId: "f6057f8333624957a6a5bd4d0c49178b997854cf6a6ad0c09219a0864dfb1216",
		DeviceNm: "Android Nexus 6",
		OsVer:    "9",
		SerialNo: "AM103094",
		QRcode: qrCode,
		RemoconCode: remoconCode,
	})

	resp, _ := client.Post(
		"https://denmoku.clubdam.com/dkdenmoku/DkDamRemoconSendServlet",
		"application/json",
		bytes.NewBuffer(data),
	)
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}

func SearchSongs(songName string) {
	client := &http.Client{}

	data, _ := json.Marshal(SearchSongInput{
		AppVer:        "2.1.0",
		CategoryCd:    "020000",
		DeviceId:      "abcdef123456789",
		DeviceNm:      "cURL",
		OsVer:         "4.4.4",
		Page:          "1",
		SongMatchType: "0",
		SerialNo:      "AB316238",
		SongName:      songName,
	})

	resp, _ := client.Post(
		"https://denmoku.clubdam.com/dkdenmoku/DkDamSearchServlet",
		"application/json",
		bytes.NewBuffer(data),
	)
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}

