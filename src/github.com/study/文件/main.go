package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/**
os.Open()函数能够打开一个文件，返回一个*File和一个err
对得到得文件实例调用Close()方法能够关闭文件
*/

func ordinaryRead(file *os.File) {
	//读文件 指定读得长度
	tmp := make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if err != nil {
			fmt.Println("read from file field err:%V", err)
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
		if n < len(tmp) {
			return
		}
	}
}

/**
buf 一行一行读
 */
func bufRead(file *os.File) {
	for {
		reader := bufio.NewReader(file)
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("read line error")
		}
		fmt.Println(line)
	}

}
/**
根据文件名读取 会帮我们执行关闭文件
 */
func readFromFileByIoutil(path string) {
	ret, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("read file failed ,err %s \n", err)
	}
	fmt.Println(string(ret))
}

func main() {
	path := "D:\\GoWorkspace\\src\\github.com\\study\\文件\\main.go"
	fileObj, err := os.Open(path)
	if err != nil {
		fmt.Printf("open file faield", err)
		return
	}
	//关闭文件
	defer fileObj.Close()
	//ordinaryRead(fileObj)
	//bufRead(fileObj)
	readFromFileByIoutil(path)
}
