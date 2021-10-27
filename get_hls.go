package main
import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)
/*this application is used to get rtsp stream data for me*/
func main() {
	var Host = ""
	var Tim = 1
	for _, a := range os.Args[1:] {
		if string(a) == "--help" {
			fmt.Println("用法为-g 输入goal")
			fmt.Println("用法为-t 输入threat num")
		}
		if string(a) == "/?" {
			fmt.Println("用法为-g 输入goal")
			fmt.Println("用法为-t 输入threat num")
		}
		m, err := regexp.MatchString("^-[g,t]", string(a))
		reg := regexp.MustCompile(`^-[g,t]`)
		if m {
			host := reg.FindAllString(string(a), -2)[0]
			if host == "-g" {
				pp := []rune(a)
				Host = string(pp[2:len(pp)])
			}
			if host == "-t" {
				pp := []rune(a)
				tim, err := strconv.Atoi(string(pp[2:len(pp)]))
				Tim = tim
				if err != nil {
					fmt.Printf("输入数字哦")
					return
				}
			}
		}
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	if Host == "" {
		fmt.Println("please input /? for help")
		fmt.Println(string(Tim))
		return
	} else {
		for i := 1; i < Tim; i++ {
			fmt.Printf("拉起第%d个文件流\n", i)
			go getHttpHls(Host)
		}
		fmt.Println("拉起主线程文件流")
		getHttpHls(Host)
	}
}
func getHttpHls(h string) {
	req, _ := http.NewRequest("GET", h, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	//here we open a loop to read stream
	buffer := make([]byte, 8192)
	for {
		time.Sleep(500)
		_, _ = resp.Body.Read(buffer)
	}
}