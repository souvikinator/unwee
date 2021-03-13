package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
    "strings"
    "github.com/fatih/color"
)

var wg sync.WaitGroup

func main() {

	banner := `                          
 _  _ _ ___ __ _____ ___ 
| || | ' \ V  V / -_) -_)
 \_,_|_||_\_/\_/\___\___| v1.0

usage: unwee [options...]

options: 
    [-l] (labeled) prints input url with corresponding result url
    [-h] (help) prints help menu
    [-m] (mute) mutes any error

example:

[1] passing single url:-
    echo "short_url" | unwee -l

[2] passing multiple urls:-
    cat urls.txt | unwee -l

note: In file each URL must be on new line

for detailed usage head over to : 

https://github.com/DarthCucumber/Tools/tree/master/unwee/README.md

`
	//urls var
	var inp_data []string
	//getting command line args
	help := flag.Bool("h", false, "prints help menu")
	labeled := flag.Bool("l", false, "prints input url and it's corresponding url")
	mute := flag.Bool("m", false, "prints input url and it's corresponding url")
	flag.Parse()

	if *help {
		fmt.Println(banner)
		os.Exit(0)
	}

	if is_piped() {
		//each url on new line
		inp_data = stdinInp() //taking input from stdin
		if len(inp_data) > 0 {
			for _, url := range inp_data {
                wg.Add(1)
                go getURL(url,labeled,mute)
			}
            wg.Wait()
		}
	} else {
		fmt.Println(banner)
		os.Exit(0)
	}

}

//utility function
func is_piped() bool {
	info, _ := os.Stdin.Stat()
	if info.Mode()&os.ModeCharDevice == 0 {
		return true
	} else {
		return false
	}
}

func stdinInp() []string {
	reader := bufio.NewReader(os.Stdin)
	var output []string
	for {
		input, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, strings.Split(input,"\n")[0])
	}

	return output
}

//make http request
func getURL(short_url string, labeled *bool, mute *bool) {
	defer wg.Done()
    //setting output colors
    s:=color.New(color.FgCyan)
    e:=color.New(color.FgRed)

    //making header request
	res, err := http.Head(short_url)

	if err != nil {
		fmt.Println("err code:", err)
	}

	if res.StatusCode == 200 {
        if *labeled{
		    s.Println(short_url, "=>",res.Request.URL)
        }else{
		    s.Println(res.Request.URL)
        }
    } else {
        if !*mute{
            e.Println(res.StatusCode,short_url)
        }
	}
}
