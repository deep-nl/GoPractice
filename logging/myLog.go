package logging

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

/*
这里利用基础的log库实现基本的log功能
*/
var (
	Trace   *log.Logger //
	Info    *log.Logger //
	Warning *log.Logger //
	Error   *log.Logger //
)

func init() {
	file, err := os.OpenFile("errors.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln("无法打开log文件：", err)
	}

	Trace = log.New(ioutil.Discard,
		"Trace: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout,
		"Info: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(os.Stdout,
		"Warning: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(file, os.Stderr),
		"Error: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
