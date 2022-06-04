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

func Sorting(infoOfAllFiles map[string]os.FileInfo) (infoOfCopy []CopyFiles, err error) {

	var LetFile []CopyFiles
	//тут я перезаписываю все имена, размеры и пути найденных файлов из мапы в слайс. Ключ мапы это путь к файлу(string)
	for k, v := range infoOfAllFiles {
		//letFile это слайс, который содержит имя, размер и путь всех найденных файлов
		LetFile = append(LetFile, CopyFiles{v.Name(), v.Size(), k})

	}
	infoOfCopy = LetFile

	return infoOfCopy, nil
}

func DuplicateList(infoOfCopy []CopyFiles) (duplicateAmount map[CopyFiles]int) {

	log.SetFormatter(&log.JSONFormatter{})
	standartFields := log.Fields{
		"func": "Sorting",
	}
	Dlog := log.WithFields(standartFields)

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
	Dlog.Info("Дубликаты  найдены")
	fmt.Printf("Найдены следующие дубликаты: \n")
	for i, v := range duplicateAmount {
		fmt.Printf("Имя %s, Размер. %dKb - Дубликат номер: %d, Расположение дубликата: %s\n", i.Name, i.Size, v, i.Path)
	}

	return duplicateAmount
}

/*
sorting.go:5:2: `github.com/sirupsen/logrus` is in the blacklist: logging is allowed only by logutils.Log (depguard)
        log "github.com/sirupsen/logrus"
        ^
sorting.go:15:54: captLocal: `InfoOfCopy' should not be capitalized (gocritic)
func Sorting(infoOfAllFiles map[string]os.FileInfo) (InfoOfCopy []CopyFiles, err error) {
                                                     ^
sorting.go:31:20: captLocal: `InfoOfCopy' should not be capitalized (gocritic)
func DuplicateList(InfoOfCopy []CopyFiles) (duplicateAmount map[CopyFiles]int) {
                   ^
sorting.go:19:2: commentFormatting: put a space between `//` and comment text (gocritic)
        //тут я перезаписываю все имена, размеры и пути найденных файлов из мапы в слайс. Ключ мапы это путь к файлу(string)
        ^
sorting.go:22:63: commentFormatting: put a space between `//` and comment text (gocritic)
                LetFile = append(LetFile, CopyFiles{v.Name(), v.Size(), k}) //letFile это слайс, который содержит имя, размер и путь всех найденных файлов
                                                                            ^
sorting.go:26:2: commentFormatting: put a space between `//` and comment text (gocritic)
        //fmt.Println(InfoOfCopy)
        ^

*/
