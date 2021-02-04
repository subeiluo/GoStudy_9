package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := "/Users/subeiluo/Downloads/ctext.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("open:%v failed, err:%v\n", filename, err)
		return
	}

	defer func() {
		file.Close()
	}()
	//等同于
	//defer file.Close()

	var content []byte
	var buf [4096]byte
	for {
		n, err := file.Read(buf[:]) //判断文件是否报错,且报错不是文件结束
		if err != nil && err != io.EOF {
			fmt.Printf("read:%s faile, err:%v\n", filename, err)
			return
		}
		if err == io.EOF { //如果文件结束,break
			break
		}
		validBuf := buf[0:n]
		//fmt.Printf("%s\n",string(validBuf))
		content = append(content, validBuf...) //将一个切片append到另一个切片中用... 展开
	}
	fmt.Printf("content:%s\n", content)

}
