package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

type KanjiInfo struct {
	Kanji               string   `json:"kanji"`
	KunReadings         []string `json:"kun_readings"`
	OnReadings          []string `json:"on_readings"`
	NameReadings        []string `json:"name_readings"`
	Meanings            []string `json:"meanings"`
	StrokeCount         int      `json:"stroke_count"`
	Grade               int      `json:"grade"`
	JLPT                int      `json:"jlpt"`
	HeisigEn            string   `json:"heisig_en"`
	FreqMainichiShinbun int      `json:"freq_mainichi_shinbun"`
}

func getKanjiInfo(kanji string) (*KanjiInfo, error) {
	resp, err := http.Get(fmt.Sprintf("https://kanjiapi.dev/v1/kanji/%s", kanji))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var kanjiInfo KanjiInfo
	err = json.Unmarshal(body, &kanjiInfo)
	if err != nil {
		return nil, err
	}

	return &kanjiInfo, nil
}

func main() {
	// Инициализация генератора случайных чисел
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Отправка GET-запроса на конечную точку kanjiapi.dev для получения всех кандзи
	resp, err := http.Get("https://kanjiapi.dev/v1/kanji/grade-1")
	if err != nil {
		fmt.Println("Ошибка при получении данных:", err)
		return
	}
	defer resp.Body.Close()

	// Чтение ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	// Распаковка JSON-ответа в срез строк
	var kanjiList []string
	err = json.Unmarshal(body, &kanjiList)
	if err != nil {
		fmt.Println("Ошибка при парсинге JSON:", err)
		return
	}

	// Проверка, что список кандзи не пуст
	if len(kanjiList) == 0 {
		fmt.Println("Кандзи не найдены.")
		return
	}

	// Выбор случайного кандзи из списка
	randomIndex := r.Intn(len(kanjiList))
	randomKanji := kanjiList[randomIndex]

	// Получение информации о случайном кандзи
	kanjiInfo, err := getKanjiInfo(randomKanji)
	if err != nil {
		fmt.Println("Ошибка при получении информации о кандзи:", err)
		return
	}

	// Вывод информации о случайном кандзи
	fmt.Printf("Случайное кандзи: %s\n", randomKanji)
	fmt.Printf("Чтения Kun: %v\n", kanjiInfo.KunReadings)
	fmt.Printf("Чтения On: %v\n", kanjiInfo.OnReadings)
	fmt.Printf("Чтения, используемые в именах: %v\n", kanjiInfo.NameReadings)
	fmt.Printf("Значения на английском: %v\n", kanjiInfo.Meanings)
	fmt.Printf("Количество перемещений: %d\n", kanjiInfo.StrokeCount)
	fmt.Printf("Официальный уровень: %d\n", kanjiInfo.Grade)
	fmt.Printf("Уровень JLPT: %d\n", kanjiInfo.JLPT)
	fmt.Printf("Ключевое слово Heisig: %s\n", kanjiInfo.HeisigEn)
	fmt.Printf("Частота в Mainichi Shinbun: %d\n", kanjiInfo.FreqMainichiShinbun)
}
