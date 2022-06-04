package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"newmod3/remove"
	scanning "newmod3/scaningFileToPath"
	"newmod3/sortingFile"
)

func main() {

	log.SetFormatter(&log.JSONFormatter{})
	standartFields := log.Fields{
		"func": "Sorting",
	}
	Slog := log.WithFields(standartFields)

	infoOfAllFiles, err := scanning.ListDirByWalk("/home/anton/projects/golang-3/hw3/duplicateWithTesting")

	if err != nil {
		_ = fmt.Errorf("произошла ошибка выполнения! %w", err)

	}

	InfoOfCopy, err := sorting.Sorting(infoOfAllFiles)
	if err != nil {
		_ = fmt.Errorf("произошла ошибка выполнения! %w", err)

		Slog.Error(err)
	}
	Slog.Info("Сортировка файлов была успешно завершена")

	copyList := sorting.DuplicateList(InfoOfCopy)

	fmt.Println("Если Вы хотите удалить дубликаты файлов, введите Y, для отмены введите N")
	var y string
	_, _ = fmt.Scan(&y)

	err = remove.Remove(copyList, y)
	if err != nil {
		fmt.Println(err)
	}

}
