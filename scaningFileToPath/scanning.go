package scanning

import (
	"fmt"
	log "github.com/sirupsen/logrus"

	"os"
	"path/filepath"
)

func ListDirByWalk(path string) (infoOfAllFiles map[string]os.FileInfo, err error) {
	log.SetFormatter(&log.JSONFormatter{})
	standartFields := log.Fields{
		"func": "LstDirByWalk",
	}
	slog := log.WithFields(standartFields)

	if _, err1 := os.Stat(path); err1 != nil {
		if os.IsNotExist(err1) {
			slog.Panicf("Неверно указан каталог, работа программы завершена")
			panic("Неверно указан каталог")
		} else if path == "" {
			defer os.Exit(2)
			err = fmt.Errorf("директория не указана! %s", err)
			slog.Errorf("Директория не указана, работа программы завершена")

		}
	}

	infoOfAllFiles = make(map[string]os.FileInfo)

	err = filepath.Walk(path, func(wPath string, info os.FileInfo, err error) error {

		//  возвращаем название папки
		if info.IsDir() {
			fmt.Printf("[Директория  %s]\n", wPath)
			return nil
		}
		// Выводится название файла рзмер и путь
		if wPath != path {

			fmt.Printf("Имя файла %s, Размер %d, Путь %s\n", info.Name(), info.Size(), wPath)
			infoOfAllFiles[wPath] = info
		}

		return nil

	})
	if err != nil {
		err = fmt.Errorf("ошибка рекурсивного обхода выбранной директории %s", err)
		slog.Panicf("Ошибка обхода директории, работа программы завершена")
		return
	}

	for i, v := range infoOfAllFiles {
		fmt.Printf("Список всех обнаруженных файлов: Имя файла %s, Размер %d Kb, Путь %s\n", v.Name(), v.Size(), i)

	}
	slog.Info("Попытка обхода директории завершилась удачно")
	return infoOfAllFiles, nil
}

/*
scanning.go:5:2: `github.com/sirupsen/logrus` is in the blacklist: logging is allowed only by logutils.Log (depguard)
        log "github.com/sirupsen/logrus"
        ^
scanning.go:24:10: ST1005: error strings should not be capitalized (stylecheck)
                        err = fmt.Errorf("Директория не указана! %s", err)

*/
