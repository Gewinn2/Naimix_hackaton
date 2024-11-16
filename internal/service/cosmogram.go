package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
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
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(3)
	comment := ""
	if value == 100 {
		firstVariant := "Ваши способы общения идеально совпадают. Вы понимаете друг друга с полуслова, и ваши разговоры всегда наполнены смыслом и взаимопониманием. Это создает крепкую основу для доверительных отношений."
		secondVariant := "Вы оба на одной волне, и ваши разговоры полны взаимопонимания. Это создает атмосферу доверия, в которой вы можете открыто делиться своими мыслями и чувствами, что укрепляет ваши отношения."
		thirdVariant := "Ваше общение — это настоящая симфония. Вы легко понимаете друг друга и делитесь мыслями, что создает прочную основу для ваших отношений."
		fourthVariant := "Вы оба обладаете уникальной способностью слушать и слышать друг друга, что позволяет вам находить общий язык даже в сложных ситуациях."
		commentsVariations := []string{firstVariant, secondVariant, thirdVariant, fourthVariant}
		comment = commentsVariations[randomIndex]
	} else if value == 90 {
		firstVariant := "Хотя ваши стили общения могут отличаться, это создает интересный динамичный обмен мнениями. Вы способны учиться друг у друга, даже если иногда возникают разногласия. Это может стать источником роста и понимания."
		secondVariant := "Ваши стили общения могут быть разными, но это не мешает вам находить общий язык. Эти различия могут обогащать ваши дискуссии, открывая новые перспективы и углубляя ваше понимание друг друга."
		thirdVariant := "Несмотря на различия в стилях общения, вы находите способы эффективно взаимодействовать. Это может обогатить ваши обсуждения и привести к новым идеям."
		fourthVariant := "Ваши разные подходы к общению могут иногда вызывать недопонимания, но вы оба готовы работать над решением этих проблем."
		commentsVariations := []string{firstVariant, secondVariant, thirdVariant, fourthVariant}
		comment = commentsVariations[randomIndex]
	} else if value == 50 {
		firstVariant := "Ваши подходы к общению могут быть немного конфликтными. Иногда возникают недопонимания, но при желании вы можете найти общий язык. Важно работать над тем, чтобы ваши различия не стали преградой для общения."
		secondVariant := "Ваши подходы к общению могут иногда вызывать трения, но это не означает, что вы не можете работать над улучшением. Важно быть терпеливыми и открытыми к компромиссам, чтобы наладить диалог."
		thirdVariant := "Ваши разговоры могут быть полны напряжения, но это также может стать возможностью для роста. Старайтесь обсуждать трудные моменты открыто и честно."
		fourthVariant := "Несмотря на конфликты в общении, вы оба способны находить компромиссы, что может помочь укрепить ваши отношения."
		commentsVariations := []string{firstVariant, secondVariant, thirdVariant, fourthVariant}
		comment = commentsVariations[randomIndex]
	} else if value == 35 {
		firstVariant := "Ваши стили общения могут быть довольно противоположными, что иногда приводит к напряженности. Необходимо прикладывать усилия, чтобы избежать конфликтов и найти способы конструктивного диалога."
		secondVariant := "Хотя ваши стили общения могут быть противоположными, это может стать источником интересных обсуждений. Постарайтесь использовать свои различия как возможность для роста и обучения друг у друга."
		thirdVariant := "Ваши стили общения могут быть несовместимыми, что иногда приводит к недопониманию. Попробуйте использовать активное слушание, чтобы лучше понимать друг друга."
		fourthVariant := "Хотя ваши подходы к общению отличаются, это может стать шансом для изучения новых способов выражения своих мыслей и чувств."
		commentsVariations := []string{firstVariant, secondVariant, thirdVariant, fourthVariant}
		comment = commentsVariations[randomIndex]
	} else {
		firstVariant := "К сожалению, ваши способы общения находятся на противоположных полюсах. Вы можете испытывать трудности в понимании друг друга, что может приводить к частым недопониманиям и конфликтам. Возможно, стоит обратить внимание на свои стили общения и попытаться найти точки соприкосновения, чтобы улучшить взаимопонимание."
		secondVariant := "Ваше общение похоже на постоянное столкновение. Вам может быть сложно понять друг друга, и это вызывает много напряжения. Возможно, стоит рассмотреть возможность работы с профессионалом, чтобы улучшить коммуникацию."
		thirdVariant := "Ваши способы общения кажутся несовместимыми. Возможно, вы часто не понимаете друг друга, что приводит к недопониманиям и конфликтам. Это может быть признаком того, что стоит поработать над улучшением коммуникации и активным слушанием."
		fourthVariant := "Кажется, что ваши стили общения абсолютно несовместимы. Это может приводить к постоянным конфликтам и чувству отчуждения. Обратите внимание на то, как вы можете изменить свои подходы к общению."
		commentsVariations := []string{firstVariant, secondVariant, thirdVariant, fourthVariant}
		comment = commentsVariations[randomIndex]
	}

	return comment
}

func AddCommentToEmotions(value int) string {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(3)
	comment := ""
	if value == 100 {
		firstVariant := "Ваши эмоциональные связи невероятно крепки. Вы чувствуете друг друга на интуитивном уровне и можете поддерживать друг друга в любых ситуациях. Это создает глубокую привязанность и понимание."
		secondVariant := "Вы делитесь глубокими эмоциональными связями, которые позволяют вам чувствовать себя комфортно и безопасно рядом друг с другом. Это создает прочный фундамент для ваших отношений, полных любви и поддержки."
		thirdVariant := "Вы оба чувствуете эмоциональную связь, которая позволяет вам поддерживать друг друга в любых ситуациях. Это создает атмосферу любви и заботы."
		fourthVariant := "Ваши эмоции переплетены так гармонично, что вы понимаете друг друга без слов. Это делает ваши отношения особенно крепкими."
		commentsVariations := []string{firstVariant, secondVariant, thirdVariant, fourthVariant}
		comment = commentsVariations[randomIndex]
	} else if value == 90 {
		firstVariant := "Ваши эмоциональные реакции гармонично сочетаются. Вы способны поддерживать друг друга и находить общий язык в сложных ситуациях. Это создает атмосферу доверия и спокойствия в ваших отношениях."
		secondVariant := "Ваши эмоциональные реакции гармонично переплетаются, создавая атмосферу взаимопонимания и поддержки. Вы способны легко делиться своими чувствами, что укрепляет вашу связь."
		thirdVariant := "Ваши эмоции прекрасно дополняют друг друга, создавая ощущение безопасности и стабильности. Вы оба умеете поддерживать друг друга в трудные времена."
		fourthVariant := "Вы способны открыто делиться своими чувствами, что укрепляет вашу связь и помогает вам расти вместе."
		commentsVariations := []string{firstVariant, secondVariant, thirdVariant, fourthVariant}
		comment = commentsVariations[randomIndex]
	} else if value == 50 {
		firstVariant := "Ваши эмоциональные отклики могут иногда конфликтовать. Это может приводить к недопониманию и напряженности, но при желании вы можете работать над улучшением взаимопонимания."
		secondVariant := "Ваши эмоции могут иногда вступать в конфликт, создавая напряжение. Однако, если вы будете открыты к обсуждению своих чувств, это может привести к более глубокому пониманию и близости."
		thirdVariant := "Ваши эмоциональные реакции могут иногда вступать в конфликт, что приводит к недопониманию. Однако, если вы будете открыты к обсуждению, это может укрепить вашу связь."
		fourthVariant := "Хотя ваши эмоции могут иногда не совпадать, важно помнить, что каждый из вас имеет право на свои чувства. Открытое общение поможет вам лучше понять друг друга."
		commentsVariations := []string{firstVariant, secondVariant, thirdVariant, fourthVariant}
		comment = commentsVariations[randomIndex]
	} else if value == 35 {
		firstVariant := "Ваши эмоциональные стили могут быть противоположными, что иногда создает трудности в понимании друг друга. Важно быть терпеливыми и открытыми для обсуждения своих чувств, чтобы избежать конфликтов."
		secondVariant := "Ваши эмоциональные стили могут быть совершенно разными, что иногда затрудняет взаимопонимание. Важно научиться выражать свои чувства и быть открытыми к обсуждениям, чтобы избежать недопонимания."
		thirdVariant := "Ваши эмоциональные стили могут быть совершенно разными, что иногда создает барьеры. Постарайтесь быть терпеливыми и открытыми к чувствам друг друга."
		fourthVariant := "Хотя ваши эмоции могут не совпадать, это также может стать возможностью для роста и понимания. Обсуждение ваших различий может привести к более глубокой связи."
		commentsVariations := []string{firstVariant, secondVariant, thirdVariant, fourthVariant}
		comment = commentsVariations[randomIndex]
	} else {
		firstVariant := "Ваши эмоциональные связи кажутся очень слабыми или даже отсутствуют. Вы можете не понимать друг друга на глубоком уровне, что приводит к чувству отдаленности и отсутствию поддержки в трудные моменты. Это может быть сигналом к тому, чтобы поработать над развитием эмоциональной близости и открытости в ваших отношениях."
		secondVariant := "Ваши эмоциональные связи кажутся отсутствующими, и вы можете чувствовать себя изолированными друг от друга. Это может быть сигналом к тому, чтобы исследовать, как вы можете лучше понять и поддерживать друг друга в эмоциональном плане."
		thirdVariant := "Ваши эмоциональные связи кажутся отсутствующими, и вы можете чувствовать себя изолированными друг от друга. Это может быть сигналом к тому, чтобы начать работать над развитием эмоциональной близости."
		fourthVariant := "Кажется, что вы не понимаете и не поддерживаете друг друга на эмоциональном уровне. Это может вызывать чувство одиночества. Возможно, стоит рассмотреть возможность открытого обсуждения своих чувств."
		commentsVariations := []string{firstVariant, secondVariant, thirdVariant, fourthVariant}
		comment = commentsVariations[randomIndex]
	}

	return comment
}

func AddCommentToWork(value int) string {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(3)
	comment := ""
	if value == 100 {
		firstVariant := "Ваши профессиональные стремления и цели идеально совпадают. Вы способны эффективно работать вместе, поддерживая друг друга на пути к успеху. Это создает отличную команду, которая может достигать высоких результатов."
		secondVariant := "Ваши профессиональные цели и амбиции идеально совпадают, что позволяет вам работать как единая команда. Вы вдохновляете друг друга, и вместе можете достигать выдающихся результатов."
		thirdVariant := "Вы оба стремитесь к одним и тем же целям и поддерживаете друг друга в профессиональных начинаниях. Это создает мощную команду, способную на многое."
		fourthVariant := "Ваше сотрудничество приносит отличные результаты, и вы оба чувствуете, что можете полагаться друг на друга в любых ситуациях."
		commentsVariations := []string{firstVariant, secondVariant, thirdVariant, fourthVariant}
		comment = commentsVariations[randomIndex]
	} else if value == 90 {
		firstVariant := "Вы оба понимаете важность работы и стремитесь к достижению общих целей. Ваше сотрудничество приносит плоды, и вы можете легко находить компромиссы в рабочих вопросах."
		secondVariant := "Вы оба понимаете важность совместной работы и стремитесь к общим целям. Это создает отличную основу для сотрудничества, где каждый может внести свой вклад."
		thirdVariant := "Ваши профессиональные подходы прекрасно дополняют друг друга, что позволяет вам эффективно работать в команде. Вы оба понимаете, как использовать сильные стороны друг друга."
		fourthVariant := "Вы способны легко находить общий язык в рабочих вопросах, что делает вашу совместную деятельность продуктивной и приятной."
		commentsVariations := []string{firstVariant, secondVariant, thirdVariant, fourthVariant}
		comment = commentsVariations[randomIndex]
	} else if value == 40 {
		firstVariant := "Ваши профессиональные взгляды могут иногда конфликтовать, что создает определенные трудности в совместной работе. Однако это может быть и возможностью для роста, если вы будете открыты для обсуждения и поиска решений"
		secondVariant := "Ваши профессиональные взгляды могут иногда конфликтовать, создавая трудности в совместной работе. Однако, если вы будете открыты к обсуждению и компромиссам, это может привести к улучшению ваших отношений."
		thirdVariant := "Ваши рабочие стили могут иногда конфликтовать, что приводит к напряжению. Однако, если вы будете открыты к компромиссам, это может помочь вам достичь лучших результатов."
		fourthVariant := "Несмотря на различия в подходах к работе, вы оба можете извлечь уроки из конфликтов и использовать их для улучшения своей совместной деятельности."
		commentsVariations := []string{firstVariant, secondVariant, thirdVariant, fourthVariant}
		comment = commentsVariations[randomIndex]
	} else {
		firstVariant := "Ваши профессиональные взгляды и подходы к работе полностью не совпадают. Это может создавать значительные препятствия в совместной деятельности и мешать достижению общих целей. Возможно, стоит обсудить свои ожидания и найти способы для улучшения сотрудничества, чтобы избежать конфликтов и недопонимания."
		secondVariant := "Ваши профессиональные подходы кажутся абсолютно несовместимыми. Это может создавать значительные препятствия в работе вместе и мешать достижению общих целей. Возможно, стоит пересмотреть свои ожидания и обсудить, как вы можете лучше сотрудничать."
		thirdVariant := "Ваше сотрудничество похоже на постоянные конфликты, и вам может быть сложно работать вместе. Возможно, стоит рассмотреть возможность работы с коучем для улучшения командной динамики."
		fourthVariant := "Кажется, что вы не можете найти общий язык в рабочих вопросах, что приводит к постоянным недопониманиям. Обратите внимание на то, как вы можете изменить свои подходы к совместной работе."
		commentsVariations := []string{firstVariant, secondVariant, thirdVariant, fourthVariant}
		comment = commentsVariations[randomIndex]
	}

	return comment
}
