package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataStrings []string, parsers []DataParser) ([]string, error) {
	if len(dataStrings) != len(parsers) {
		return nil, fmt.Errorf("количество строк данных (%d) не совпадает с количеством парсеров (%d)", len(dataStrings), len(parsers))
	}
	var results []string
	for i := 0; i < len(dataStrings); i++ {
		err := parsers[i].Parse(dataStrings[i])
		if err != nil {
			log.Printf("Ошибка парсинга строки")
			continue
		}
		report, err := parsers[i].ActionInfo()
		if err != nil {
			log.Printf("Ошибка формирования отчёта")
			continue
		}
		results = append(results, report)
	}
	return results, nil
}
