package sorting

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

type CopyFiles struct {
	Name string
	Size int64
	Path string
}

var LetFile []CopyFiles

func Sorting(infoOfAllFiles map[string]os.FileInfo) (infoOfCopy []CopyFiles, err error) {

	//тут я перезаписываю все имена, размеры и пути найденных файлов из мапы в слайс. Ключ мапы это путь к файлу(string)

	for k, v := range infoOfAllFiles {
		LetFile = append(LetFile, CopyFiles{v.Name(), v.Size(), k}) //letFile это слайс, который содержит имя, размер и путь всех найденных файлов

	}
	infoOfCopy = LetFile
	//fmt.Println(InfoOfCopy)

	return infoOfCopy, nil
}

// следующий лог
func DuplicateList(infoOfCopy []CopyFiles) (duplicateAmount map[CopyFiles]int) {

	log.SetFormatter(&log.JSONFormatter{})
	standartFields := log.Fields{
		"func": "Sorting",
	}
	Dlog := log.WithFields(standartFields)

	//var copyList2 []CopyFiles
	duplicateAmount = make(map[CopyFiles]int)
	for i := 0; i < len(infoOfCopy); i++ {
		for j := i + 1; j < len(infoOfCopy); j++ {
			if infoOfCopy[i].Name == infoOfCopy[j].Name && infoOfCopy[i].Size == infoOfCopy[j].Size {

				duplicateAmount[infoOfCopy[j]]++
			}
		}

	}

	if len(duplicateAmount) == 0 {
		fmt.Println("Дубликаты не найдены")
		Dlog.Info("Дубликаты не найдены")
		os.Exit(1)
	}
	log.Printf("Найдены следующие дубликаты")
	Dlog.Info("Дубликаты  найдены")
	for i, v := range duplicateAmount {
		fmt.Printf("Имя %s, Размер. %dKb - Дубликат номер: %d, Расположение дубликата: %s\n", i.Name, i.Size, v, i.Path)
	}

	return duplicateAmount
}
