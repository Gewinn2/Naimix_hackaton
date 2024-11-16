package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"server/internal/entities"
	"server/util"
	"time"
)

type CosmogramService struct {
	timeout time.Duration
}

func NewCosmogramService() *CosmogramService {
	return &CosmogramService{
		timeout: time.Duration(10) * time.Second,
	}
}

func (c *CosmogramService) Get(cosmogramRequestBody entities.CosmogramRequestBody) ([]entities.CosmogramApiResponseBody, error) {
	year, month, day, err := util.BirthDateParser(cosmogramRequestBody.Candidate.BirthDate)
	if err != nil {
		return nil, err
	}
	var candidate entities.Subject
	candidate.Name = fmt.Sprintf("%s %s %s", cosmogramRequestBody.Candidate.Surname, cosmogramRequestBody.Candidate.Name, cosmogramRequestBody.Candidate.ThirdName)
	candidate.Year = year
	candidate.Month = month
	candidate.Day = day
	candidate.Hour = 12
	candidate.Minute = 11
	candidate.Longitude = 37.6156
	candidate.Latitude = 55.7522
	candidate.City = "Moscow"
	candidate.Nation = "RUS"
	candidate.Timezone = "Europe/Moscow"
	candidate.ZodiacType = "Tropic"

	var staff []entities.Subject
	for i := range cosmogramRequestBody.Staff {
		year, month, day, err = util.BirthDateParser(cosmogramRequestBody.Staff[i].BirthDate)
		if err != nil {
			return nil, err
		}
		var personal entities.Subject
		personal.Name = fmt.Sprintf("%s %s %s", cosmogramRequestBody.Candidate.Surname, cosmogramRequestBody.Candidate.Name, cosmogramRequestBody.Candidate.ThirdName)
		personal.Year = year
		personal.Month = month
		personal.Day = day
		personal.Hour = 12
		personal.Minute = 11
		personal.Longitude = 37.6156
		personal.Latitude = 55.7522
		personal.City = "Moscow"
		personal.Nation = "RUS"
		personal.Timezone = "Europe/Moscow"
		personal.ZodiacType = "Tropic"
		staff = append(staff, personal)
	}

	var responses []entities.CosmogramApiResponseBody
	for i := range staff {
		// Преобразуем тело в JSON
		jsonData, err := json.Marshal(staff[i])
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Ошибка при маршаллинге данных: %v", err))
		}

		// Создаем новый POST-запрос
		url := "https://astrologer.p.rapidapi.com/api/v4/relationship-score"
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Ошибка при создании запроса: %v", err))
		}

		// Устанавливаем необходимые заголовки
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("x-rapidapi-ua", "RapidAPI-Playground")
		req.Header.Set("x-rapidapi-key", "2889ef6b0fmsh4d6c7d819751b98p1685f1jsn332f4e932558")
		req.Header.Set("x-rapidapi-host", "astrologer.p.rapidapi.com")

		// Отправляем запрос
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Ошибка при отправке запроса: %v", err))
		}

		// Читаем и выводим ответ
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Ошибка при чтении ответа: %v", err))
		}

		fmt.Printf("Ответ от сервера: %s\n", body)

		var response entities.CosmogramApiResponseBody
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Ошибка при парсинге ответа: %v", err))
		}
		responses = append(responses, response)
		resp.Body.Close()
	}
	return responses, nil
}
