package actioninfo

import (
	"log"
)

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, line := range dataset {
		if err := dp.Parse(line); err != nil {
			log.Printf("ошибка при разборе строки")
			continue
		}
		infoStr, err := dp.ActionInfo()
		if err != nil {
			log.Printf("ошибка при формировании информации об активности")
			continue
		}
		log.Println(infoStr)
	}
}
