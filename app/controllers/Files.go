package controllers

import (
	"GoBase/app/pkg/helpers"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type TestFile struct {
	ReadFile  string
	WriteFile string
}

type SumTYpe struct {
	EnCount    int64
	NumCount   int64
	SpaceCount int64
	OtherCount int64
}

func TestFiles() {
	TestFile := TestFile{}
	TestFile.ReadFile = "/Users/songsong/Downloads/1111.txt"
	TestFile.WriteFile = "/Users/songsong/Downloads/2222.txt"
	// TestFile.TestFile1()
	// TestFile.TestFile2()
	// TestFile.TestFile3()
	// TestFile.TestFile4()
	fmt.Println(TestFile.TestFile6())
}

// TestFile6 统计文件中不同类型的个数
func (f TestFile) TestFile6() SumTYpe {
	SumTYpe := SumTYpe{}
	if f.TestFile5() != true {
		return SumTYpe
	}
	contentList, err := helpers.ReadFile(f.ReadFile)
	if err != nil {
		fmt.Println(err)
		return SumTYpe
	}

	for _, content := range contentList {
		for _, v := range content {
			switch {
			case v >= 'a' && v <= 'z':
				SumTYpe.EnCount++
			case v >= 'A' && v <= 'Z':
				SumTYpe.EnCount++
			case v >= '0' && v <= '9':
				SumTYpe.NumCount++
			case v == ' ' || v == '\t':
				SumTYpe.SpaceCount++
			default:
				SumTYpe.OtherCount++
			}
		}
	}

	// fmt.Printf("字母:%v, 数字:%v, 空格：%v, 其他:%v", SumTYpe.EnCount, SumTYpe.NumCount, SumTYpe.SpaceCount, SumTYpe.OtherCount)
	return SumTYpe
}

// TestFile5 判断文件啊是否存在
func (f TestFile) TestFile5() bool {
	_, err := os.Stat(f.ReadFile)

	var isExit bool
	if err == nil {
		isExit = true
	}

	if os.IsNotExist(err) {
		isExit = false
	}

	return isExit
}

// TestFile4 文件a.txt 导入到 b.txt
func (f TestFile) TestFile4() {
	content, err := ioutil.ReadFile(f.ReadFile)
	if err != nil {
		fmt.Println("读取失败")
		return
	}

	err = ioutil.WriteFile(f.WriteFile, content, 0666)
	if err != nil {
		fmt.Println("导入失败")
		return
	}
}

// TestFile3 写文件
func (f TestFile) TestFile3() {
	//const (
	//	O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
	//	O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
	//	O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
	//	O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
	//	O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
	//	O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
	//	O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
	//	O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
	//)
	for i := 0; i < 5; i++ {
		file, err := helpers.WriteFile(f.WriteFile, "aaa"+strconv.Itoa(i))
		if err != nil {
			fmt.Printf("写入失败%d", i)
		}
		if file != true {
			fmt.Printf("写入失败%d", i)
		}
	}

	fmt.Println("ok")
}

func (f TestFile) TestFile2() {
	data := helpers.ReadAll(f.ReadFile)
	fmt.Println(data)
}

// TestFile1 按行读取文件
func (f TestFile) TestFile1() string {
	contentList, err := helpers.ReadFile(f.ReadFile)
	if err != nil {
		fmt.Println(err)
		return "err"
	}
	for k, v := range contentList {
		fmt.Printf("k:%d,v:%v", k, v)
	}

	return ""
}
