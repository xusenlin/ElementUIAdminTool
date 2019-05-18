package helper

import (
	"fmt"
	"io/ioutil"
	"os"
)

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func Mkdir(dir string) bool {
	err := os.Mkdir(dir, os.ModePerm)
	if err != nil{
		fmt.Println("生成"+ dir + "目录失败！")
		return false
	}
	return true
}

func WriteFile(fileName string,fileByte []byte) bool {

	err := ioutil.WriteFile(fileName,[]byte(fileByte),os.ModePerm)
	if err != nil{
		fmt.Println("生成"+ fileName + "文件失败！")
		return false
	}

	return true
}