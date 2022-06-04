package remove

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	sorting "newmod3/sortingFile"
	"os"
	"strconv"
	"strings"
)

func Remove(duplicateAmount map[sorting.CopyFiles]int, yes string) (err error) {

	log.SetFormatter(&log.JSONFormatter{})
	standartFields := log.Fields{
		"func": "Remove",
	}
	Rlog := log.WithFields(standartFields)

	yes = strings.ToUpper(yes)

	for i := range duplicateAmount {

		switch yes {
		case "Y":

			Rlog.Info("Пользователь согласился удалить дубликаты (Y)")
			err = os.Remove(i.Path)
			if err != nil {
				err = fmt.Errorf("ошибка при удалении дубликатов %s", err)
				Rlog.Errorf("ошибка при удалении дубликатов %s", err)
				return err
			}
			fmt.Println("Дубликаты файлов успешно удалены")
			Rlog.Info("Дубликаты успешно удалены")

		case "N":

			Rlog.Info("Пользователь отказался удалять дубликаты (N)")
			fmt.Println("Отмена")
			return

		case yes:

			if _, err = strconv.Atoi(yes); err == nil {
				fmt.Printf("%q Вы ввели число, операция прервана\n", yes)
				Rlog.Warnf("Пользователь ввел недопустимый символ, операция прервана %q", yes)
				return
			}
			if yes != "N" {
				fmt.Printf("Вы ввели неверную букву: '%s', операция прервана\n", yes)
				Rlog.Warnf("Пользователь ввел недопустимый символ %s, операция прервана", yes)
				return
			} else if yes != "Y" {
				fmt.Printf("Вы ввели неверную букву: '%s', операция прервана\n", yes)
				Rlog.Warnf("Пользователь ввел недопустимый символ, операция прервана %s", yes)
				return
			}

		}

	}
	return nil
}

/*
remove.go:5:2: `github.com/sirupsen/logrus` is in the blacklist: logging is allowed only by logutils.Log (depguard)
        log "github.com/sirupsen/logrus"
        ^
remove.go:26:4: unnecessaryBlock: case statement doesn't require a block statement (gocritic)
                        {
                        ^
remove.go:38:4: unnecessaryBlock: case statement doesn't require a block statement (gocritic)
                        {
                        ^
remove.go:44:4: unnecessaryBlock: case statement doesn't require a block statement (gocritic)
                        {

*/
