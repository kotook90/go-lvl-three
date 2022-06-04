package sorting

import (
	"fmt"
	"os"
	"testing"
)

func TestSorting(t *testing.T) {

	result := 2 // два файла, индекс длина индекса реального результата должна быть 2

	mapa := map[string]os.FileInfo{}
	f, err := os.Open("test2.file")
	if err != nil {
		panic(err)
	}
	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}

	f2, err := os.Open("test.file")
	if err != nil {
		panic(err)
	}
	f3, err := f2.Stat()
	if err != nil {
		panic(err)
	}

	mapa["a"] = fi
	mapa["b"] = f3

	realResult, _ := Sorting(mapa)
	for i, v := range realResult {
		fmt.Printf("NAME %s, SIZE %d, PATH %s, INDEX %d\n", v.Name, v.Size, v.Path, i)

	}
	if len(realResult) != result {
		fmt.Println("Сортиовка не работает")
	} else {
		fmt.Println("ok")
	}

}
