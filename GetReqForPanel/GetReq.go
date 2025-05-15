package getreqforpanel

//Здесь будем писать GET запросы к панели https://www.postman.com/hsanaei/3x-ui/documentation/q1l5l0u/3x-ui

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type ClientSettings struct {
	ID       int    `json:"id"`
	Settings string `json:"settings"`
}

// Аутентификация в панели
func Authenticate() {
	urlStr := "http://212.113.116.19:1860/5GdMclkztE8Cn3g/login"
	method := "POST"

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	username := os.Getenv("USERNAME_PANEL")
	password := os.Getenv("PASSWORD_PANEL")

	data := url.Values{}
	data.Set("username", username)
	data.Set("password", password)

	client := &http.Client{}
	req, err := http.NewRequest(method, urlStr, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Response Status:", res.Status)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response Body Length:", len(body))

	fmt.Printf("Response Body (hex): %x\n", body)

	fmt.Println("Response Body:", string(body))
}

// Обноваление данных о клиенте
func UpdateClient() {
	uuid := "pass"
	url := "http://212.113.116.19:1860/panel/api/inbounds/updateClient/" + uuid
	method := "POST"

	settings := map[string]interface{}{
		"clients": []map[string]interface{}{
			{
				"id":         "95e4e7bb-7796-47e7-e8a7-f4055194f776",
				"alterId":    0,
				"email":      "test123",
				"limitIp":    2,
				"totalGB":    42949672960,
				"expiryTime": 1682864675944,
				"enable":     true,
				"tgId":       "",
				"subId":      "",
			},
		},
	}

	clientSettings := ClientSettings{
		ID:       3,
		Settings: string(mustJSONMarshal(settings)),
	}

	payload, err := json.Marshal(clientSettings)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println(string(body))
}

func mustJSONMarshal(v interface{}) []byte {
	data, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return data
}

// Добавление нового клиента
func AddNewUser(email string, limitIP int, totalGB int, expiryTime int, enable bool, telegram_id int) {
	url := "http://localhost:2053/panel/api/inbounds/addClient"
	method := "POST"

	id := NewUuid()

	settings := map[string]interface{}{
		"clients": []map[string]interface{}{
			{
				"id":         id,
				"alterId":    0,
				"email":      email,
				"limitIp":    limitIP,
				"totalGB":    totalGB,
				"expiryTime": expiryTime,
				"enable":     enable,
				"tgId":       "",
				"subId":      "",
			},
		},
	}

	clientSettings := ClientSettings{
		ID:       3,
		Settings: string(mustJSONMarshal(settings)),
	}

	payload, err := json.Marshal(clientSettings)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response Status:", res.Status)
	fmt.Println("Response Body:", string(body))
}

// Генерация нового UUID
func NewUuid() string {
	newID := uuid.New()
	return newID.String()
}
