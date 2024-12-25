package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Структура для обработки данных о друзьях
type Friend struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsClosed  bool   `json:"is_closed"`
	PhotoURL  string `json:"photo_100"`
}

// Основная функция
func main() {
	// Ваш Access Token
	token := "Сюда ваш токен "

	// ID пользователя (по умолчанию "self" для текущего пользователя)
	userID := "self"

	// URL API метода friends.get
	apiURL := "https://api.vk.com/method/friends.get"

	// Параметры запроса
	params := url.Values{}
	params.Set("user_id", userID)
	params.Set("fields", "first_name,last_name,photo_100") // Запрашиваем доп. данные
	params.Set("access_token", token)
	params.Set("v", "5.131") // Версия API

	// Выполняем запрос
	resp, err := http.Get(apiURL + "?" + params.Encode())

	if err != nil {
		log.Fatalf("Ошибка при выполнении запроса: %v", err)
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка при чтении ответа: %v", err)
	}

	// Распарсим JSON
	var result struct {
		Response struct {
			Count int      `json:"count"`
			Items []Friend `json:"items"`
		} `json:"response"`
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalf("Ошибка при парсинге JSON: %v", err)
	}

	// Вывод данных о друзьях
	fmt.Printf("Найдено друзей: %d\n", result.Response.Count)
	for _, friend := range result.Response.Items {
		fmt.Printf("ID: %d, Имя: %s %s, Фото: %s\n", friend.ID, friend.FirstName, friend.LastName, friend.PhotoURL)
	}
}
