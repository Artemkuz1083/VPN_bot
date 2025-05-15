package getreqforpanel

//Здесь будем писать GET запросы к панели https://www.postman.com/hsanaei/3x-ui/documentation/q1l5l0u/3x-ui

import (
	database "bot/DataBase"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

// Аутентификация в панели
func Authenticate() []*http.Cookie {
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
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
	}
	defer res.Body.Close()

	cookies := res.Cookies()

	fmt.Println("Response Status:", res.Status)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	fmt.Println("Response Body:", string(body))

	return cookies
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

	clientSettings := map[string]interface{}{
		"id":       4,
		"settings": string(mustJSONMarshal(settings)),
	}

	fmt.Println(string(mustJSONMarshal(settings)))
	fmt.Println(json.Marshal(clientSettings))

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
func AddNewUser(email string, limitIP int, totalGB int, expiryTime int, enable bool, telegram_id int, payment bool) {
	url := "http://212.113.116.19:1860/5GdMclkztE8Cn3g/panel/api/inbounds/addClient"
	method := "POST"

	// Генерация UUID для клиента
	clientID := uuid.New().String()

	cookies := Authenticate()

	var cookieName string
	var cookieValue string

	for _, cookie := range cookies {
		cookieName = cookie.Name
		cookieValue = cookie.Value
		fmt.Printf("Name: %s, Value: %s\n", cookieName, cookieValue)
	}

	cookieReq := &http.Cookie{
		Name:  cookieName,
		Value: cookieValue,
	}

	// Создаем структуру clients
	clients := map[string]interface{}{
		"id":         clientID,
		"flow":       "",
		"email":      email,
		"limitIp":    limitIP,
		"totalGB":    totalGB,
		"expiryTime": expiryTime,
		"enable":     enable,
		"tgId":       "",
		"subId":      fmt.Sprint(telegram_id),
		"reset":      0,
	}

	// Сериализуем clients в JSON-строку
	settings := map[string]interface{}{
		"clients": []interface{}{clients},
	}
	settingsJSON, _ := json.Marshal(settings)

	// Формируем основной запрос
	requestBody := map[string]interface{}{
		"id":       4,
		"settings": string(settingsJSON),
	}

	// Двойная сериализация
	payload, _ := json.Marshal(requestBody)

	fmt.Println("Final JSON:", string(payload))

	client := &http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.AddCookie(cookieReq)

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

	database.AddNewUser(clientID, telegram_id, email, limitIP, totalGB, expiryTime, enable, payment, time.Now(), time.Now().AddDate(0, 0, 30))
}

// Генерация нового UUID
func NewUuid() string {
	newID := uuid.New()
	return newID.String()
}

func GetInbound() {
	url := "http://212.113.116.19:1860/5GdMclkztE8Cn3g/panel/api/inbounds/getClientTraffics/s729v2km"
	method := "GET"

	cookies := Authenticate()

	var cookieName string
	var cookieValue string

	for _, cookie := range cookies {
		cookieName = cookie.Name
		cookieValue = cookie.Value
	}

	cookieReq := &http.Cookie{
		Name:  cookieName,
		Value: cookieValue,
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	req.AddCookie(cookieReq)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
