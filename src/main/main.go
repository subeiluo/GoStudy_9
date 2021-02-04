package main

import (
	"fmt"
	"os"
	"time"
)
const (
	bit = 1
	Kb  = bit * 1024
	Mb  = Kb * 1024
	Gb  = Mb * 1024
)

func main()  {
	n :=newfile("/Users/subeiluo/Downloads/logtest.log")
	n.sizeSplitLog()
}

type FileName struct {
	filename string
}
func newfile(name string) *FileName {
	return &FileName{
		filename: name,
	}
}

func (c *FileName)sizeSplitLog() {
	sizer,_:=os.Stat(c.filename)
	siz :=sizer.Size()
	fmt.Println(siz)
	if siz - 50*Mb >= 50 {
		fmt.Println(siz)
		nowhour := time.Now()
		newfilename := fmt.Sprintf(
			"%s-%04d-%02d-%02d-%02d", c.filename,
			nowhour.Year(), nowhour.Month(), nowhour.Day(), nowhour.Hour())
		os.Rename(c.filename, newfilename)

	}
}