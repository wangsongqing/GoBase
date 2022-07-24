package helpers

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// BubbleSort 排序
func BubbleSort(arr *[]int) []int {
	temp := 0
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}

	return *arr
}

func BinaryFind(arr []int, findVal int) int {

	keys := -1
	for k, v := range arr {
		if v == findVal {
			keys = k
			break
		}
	}

	return keys
}

func ReadFile(fileName string) ([]string, error) {
	var fileData []string

	files, err := os.Open(fileName)
	if err != nil {
		return fileData, err
	}

	defer func(files *os.File) {
		err := files.Close()
		if err != nil {
			return
		}
	}(files)

	render := bufio.NewReader(files)

	for {
		str, err := render.ReadString('\n')
		_str := fmt.Sprintf("%s", str)
		fileData = append(fileData, _str)
		if err == io.EOF { // io.EOF 表示文件末行
			break
		}
	}

	return fileData, nil
}

func ReadAll(fileName string) string {
	contentList, err := ioutil.ReadFile(fileName)

	if err != nil {
		return ""
	}

	return string(contentList)
}

// WriteFile 追加方式写入文件,如果文件不存在则新建文件
func WriteFile(fileName string, content string) (bool, error) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return false, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	write := bufio.NewWriter(file)
	writeString, err := write.WriteString(content + "\r\n")
	if err != nil {
		return false, err
	}

	if writeString < 0 {
		return false, err
	}

	err = write.Flush()
	if err != nil {
		return false, err
	}

	return true, nil
}
