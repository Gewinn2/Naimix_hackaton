package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"server/internal/entities"
	"server/util"
	"strings"
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

func (c *CosmogramService) Get(cosmogramRequestBody entities.CosmogramRequestBody) (*entities.CosmogramResponseBody, error) {
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
	candidate.Nation = "RU"
	candidate.Timezone = "Europe/Moscow"
	candidate.ZodiacType = "Tropic"

	var staff []entities.Subject
	for i := range cosmogramRequestBody.Staff {
		year, month, day, err = util.BirthDateParser(cosmogramRequestBody.Staff[i].BirthDate)
		if err != nil {
			return nil, err
		}
		var employee entities.Subject
		employee.Name = fmt.Sprintf("%s %s %s", cosmogramRequestBody.Candidate.Surname, cosmogramRequestBody.Candidate.Name, cosmogramRequestBody.Candidate.ThirdName)
		employee.Year = year
		employee.Month = month
		employee.Day = day
		employee.Hour = 12
		employee.Minute = 11
		employee.Longitude = 37.6156
		employee.Latitude = 55.7522
		employee.City = "Moscow"
		employee.Nation = "RU"
		employee.Timezone = "Europe/Moscow"
		employee.ZodiacType = "Tropic"
		staff = append(staff, employee)
	}

	var compatibilities []entities.Compatibility
	var cosmogram entities.Subject2
	for i := range staff {
		requestBody := entities.CosmogramApiRequestBody{
			FirstSubject:  candidate,
			SecondSubject: staff[i],
		}
		// Преобразуем тело в JSON
		jsonData, err := json.Marshal(requestBody)
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

		var apiResponse entities.CosmogramApiResponseBody
		err = json.Unmarshal(body, &apiResponse)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Ошибка при парсинге ответа: %v", err))
		}

		aspects := ComparePlanets(apiResponse.Data.FirstSubject, apiResponse.Data.SecondSubject)
		scores := EvaluateCategory(aspects)
		var compatibility entities.Compatibility
		compatibility.CandidateFullName = candidate.Name
		compatibility.EmployeeFullName = staff[i].Name
		compatibility.Communication = scores["Communication"]
		compatibility.CommunicationComment = AddCommentToCommunication(scores["Communication"])
		compatibility.Emotions = scores["Emotions"]
		compatibility.EmotionsComment = AddCommentToEmotions(scores["Emotions"])
		compatibility.Work = scores["Work"]
		compatibility.WorkComment = AddCommentToWork(scores["Work"])
		compatibilities = append(compatibilities, compatibility)

		cosmogram = apiResponse.Data.FirstSubject

		resp.Body.Close()
	}

	var response entities.CosmogramResponseBody
	response.Compatibilities = compatibilities
	response.Cosmogram = cosmogram
	return &response, nil
}

func EvaluateCategory(aspects map[string][]string) map[string]int {
	scores := map[string]int{
		"Communication": 0,
		"Emotions":      0,
		"Work":          0,
	}

	// Общение (Меркурий)
	if contains(aspects, "Conjunction", "Mercury", "Mercury") {
		scores["Communication"] += 100
	} else if contains(aspects, "Opposition", "Mercury", "Mercury") {
		scores["Communication"] += 90
	} else if contains(aspects, "Square", "Mercury", "Mercury") {
		scores["Communication"] += 50
	} else if contains(aspects, "Opposition", "Mercury", "Mercury") {
		scores["Communication"] += 35
	}

	// Эмоции (Луна)
	if contains(aspects, "Conjunction", "Moon", "Moon") {
		scores["Emotions"] += 100
	} else if contains(aspects, "Trine", "Moon", "Moon") {
		scores["Emotions"] += 90
	} else if contains(aspects, "Square", "Moon", "Moon") {
		scores["Emotions"] += 50
	} else if contains(aspects, "Opposition", "Moon", "Moon") {
		scores["Emotions"] += 35
	}

	// Работа (Солнце и Сатурн)
	if contains(aspects, "Trine", "Sun", "Saturn") {
		scores["Work"] += 100
	} else if contains(aspects, "Sextile", "Sun", "Saturn") {
		scores["Work"] += 90
	} else if contains(aspects, "Square", "Sun", "Saturn") {
		scores["Work"] += 40
	}

	return scores
}

func contains(aspects map[string][]string, aspectType, planetA, planetB string) bool {
	for _, connection := range aspects[aspectType] {
		if strings.Contains(connection, planetA) && strings.Contains(connection, planetB) {
			return true
		}
	}
	return false
}

func ComparePlanets(subjectA, subjectB entities.Subject2) map[string][]string {
	aspects := make(map[string][]string)

	// Список планет для анализа
	planets := []struct {
		name    string
		planetA entities.Planet
		planetB entities.Planet
	}{
		{"Sun", subjectA.Sun, subjectB.Sun},
		{"Moon", subjectA.Moon, subjectB.Moon},
		{"Mercury", subjectA.Mercury, subjectB.Mercury},
		{"Venus", subjectA.Venus, subjectB.Venus},
		{"Mars", subjectA.Mars, subjectB.Mars},
		{"Jupiter", subjectA.Jupiter, subjectB.Jupiter},
		{"Saturn", subjectA.Saturn, subjectB.Saturn},
		{"Uranus", subjectA.Uranus, subjectB.Uranus},
		{"Neptune", subjectA.Neptune, subjectB.Neptune},
		{"Pluto", subjectA.Pluto, subjectB.Pluto},
		{"Chiron", subjectA.Chiron, subjectB.Chiron},
		{"MeanLilith", subjectA.MeanLilith, subjectB.MeanLilith},
		{"MeanNode", subjectA.MeanNode, subjectB.MeanNode},
		{"MeanSouthNode", subjectA.MeanSouthNode, subjectB.MeanSouthNode},
	}

	// Сравнение планет между двумя субъектами
	for _, pair := range planets {
		aspect := CalculateAspect(pair.planetA, pair.planetB)
		if aspect != "No Major Aspect" {
			aspects[aspect] = append(aspects[aspect], fmt.Sprintf("%s (%s) ↔ %s (%s)", pair.name, subjectA.Name, pair.name, subjectB.Name))
		}
	}

	return aspects
}

func CalculateAspect(planetA, planetB entities.Planet) string {
	diff := math.Abs(planetA.AbsPos - planetB.AbsPos)
	if diff > 180 {
		diff = 360 - diff
	}

	switch {
	case math.Abs(diff) <= 5:
		return "Conjunction" // Соединение
	case math.Abs(diff-60) <= 5:
		return "Sextile" // Секстиль
	case math.Abs(diff-90) <= 5:
		return "Square" // Квадрат
	case math.Abs(diff-120) <= 5:
		return "Trine" // Тригон
	case math.Abs(diff-180) <= 5:
		return "Opposition" // Оппозиция
	default:
		return "No Major Aspect"
	}
}

func AddCommentToCommunication(value int) string {
	comment := ""
	if value == 100 {
		comment = "Ваши способы общения идеально совпадают. Вы понимаете друг друга с полуслова, и ваши разговоры всегда наполнены смыслом и взаимопониманием. Это создает крепкую основу для доверительных отношений."
	} else if value == 90 {
		comment = "Хотя ваши стили общения могут отличаться, это создает интересный динамичный обмен мнениями. Вы способны учиться друг у друга, даже если иногда возникают разногласия. Это может стать источником роста и понимания."
	} else if value == 50 {
		comment = "Ваши подходы к общению могут быть немного конфликтными. Иногда возникают недопонимания, но при желании вы можете найти общий язык. Важно работать над тем, чтобы ваши различия не стали преградой для общения."
	} else if value == 35 {
		comment = "Ваши стили общения могут быть довольно противоположными, что иногда приводит к напряженности. Необходимо прикладывать усилия, чтобы избежать конфликтов и найти способы конструктивного диалога."
	} else {
		comment = "К сожалению, ваши способы общения находятся на противоположных полюсах. Вы можете испытывать трудности в понимании друг друга, что может приводить к частым недопониманиям и конфликтам. Возможно, стоит обратить внимание на свои стили общения и попытаться найти точки соприкосновения, чтобы улучшить взаимопонимание."
	}

	return comment
}

func AddCommentToEmotions(value int) string {
	comment := ""
	if value == 100 {
		comment = "Ваши эмоциональные связи невероятно крепки. Вы чувствуете друг друга на интуитивном уровне и можете поддерживать друг друга в любых ситуациях. Это создает глубокую привязанность и понимание."
	} else if value == 90 {
		comment = "Ваши эмоциональные реакции гармонично сочетаются. Вы способны поддерживать друг друга и находить общий язык в сложных ситуациях. Это создает атмосферу доверия и спокойствия в ваших отношениях."
	} else if value == 50 {
		comment = "Ваши эмоциональные отклики могут иногда конфликтовать. Это может приводить к недопониманию и напряженности, но при желании вы можете работать над улучшением взаимопонимания."
	} else if value == 35 {
		comment = "Ваши эмоциональные стили могут быть противоположными, что иногда создает трудности в понимании друг друга. Важно быть терпеливыми и открытыми для обсуждения своих чувств, чтобы избежать конфликтов."
	} else {
		comment = "Ваши эмоциональные связи кажутся очень слабыми или даже отсутствуют. Вы можете не понимать друг друга на глубоком уровне, что приводит к чувству отдаленности и отсутствию поддержки в трудные моменты. Это может быть сигналом к тому, чтобы поработать над развитием эмоциональной близости и открытости в ваших отношениях."
	}

	return comment
}

func AddCommentToWork(value int) string {
	comment := ""
	if value == 100 {
		comment = "Ваши профессиональные стремления и цели идеально совпадают. Вы способны эффективно работать вместе, поддерживая друг друга на пути к успеху. Это создает отличную команду, которая может достигать высоких результатов."
	} else if value == 90 {
		comment = "Вы оба понимаете важность работы и стремитесь к достижению общих целей. Ваше сотрудничество приносит плоды, и вы можете легко находить компромиссы в рабочих вопросах."
	} else if value == 40 {
		comment = "Ваши профессиональные взгляды могут иногда конфликтовать, что создает определенные трудности в совместной работе. Однако это может быть и возможностью для роста, если вы будете открыты для обсуждения и поиска решений"
	} else {
		comment = "Ваши профессиональные взгляды и подходы к работе полностью не совпадают. Это может создавать значительные препятствия в совместной деятельности и мешать достижению общих целей. Возможно, стоит обсудить свои ожидания и найти способы для улучшения сотрудничества, чтобы избежать конфликтов и недопонимания."
	}

	return comment
}
