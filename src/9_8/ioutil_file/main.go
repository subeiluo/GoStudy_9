package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filename := "/Users/subeiluo/Downloads/mmp.log"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("read file %s failed, err:%v\n", filename, err)

	}
	fmt.Printf("content:%s\n", string(content))
}
