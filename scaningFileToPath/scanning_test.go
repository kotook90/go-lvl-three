package scanning_test

import (
	"fmt"
	scanning "newmod3/scaningFileToPath"
	"os"
	"reflect"
	"testing"
)

func TestListDirByWalk(t *testing.T) {
	path := "/home/anton/projects/golang-3/hw3/duplicateWithTesting"

	inf, _ := scanning.ListDirByWalk(path)

	if inf == nil {
		fmt.Println("Пустая директория")
	} else {
		fmt.Println("mapa наполнена")
	}
	for range inf {
		fmt.Println(inf)
	}

}

func TestListDirByWalk1(t *testing.T) {
	type args struct {
		path string
	}

	tests := []struct {
		name               string
		args               args
		wantInfoOfAllFiles map[string]os.FileInfo
		wantErr            bool
	}{
		{name: "file1", args: args{path: "/home/anton/projects/golang-3/hw3/duplicateWithTesting/file1"}, wantInfoOfAllFiles: map[string]os.FileInfo{}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInfoOfAllFiles, err := scanning.ListDirByWalk(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListDirByWalk() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfoOfAllFiles, tt.wantInfoOfAllFiles) {
				t.Errorf("ListDirByWalk() gotInfoOfAllFiles = %v, want %v", gotInfoOfAllFiles, tt.wantInfoOfAllFiles)
			}
		})
	}
}
