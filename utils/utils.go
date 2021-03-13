package utils

import (
		"os"
		"fmt"
		"bufio"
		"io"
		"strings"
		"sync"
		"net/http"
)


//displays usage details
func DisplayHelp(){
banner := `                          
 _  _ _ ___ __ _____ ___ 
| || | ' \ V  V / -_) -_)
 \_,_|_||_\_/\_/\___\___| v1.1.0

usage: unwee [options...]

options: 
[-u] (url) takes in single URL as input and gives it's unshortened form.
[-o] (output) takes in file name as input and saves result in the file
[-f] (file) takes in file name containing list of shortened URLs
[-h] (help) prints help menu
[-ex] (exclude) takes in status codes(404,403,...) separated by commas and doesn't include them in results
`
fmt.Println(banner)
os.Exit(0)
}

//check error
func CheckErr(err error){
	if err!=nil{
		fmt.Println("[unwee] ",err)
		os.Exit(0)
	}
}

//return slice of data in the file
func GetFileData(filePath string)([]string){
	var output []string
	fl,err:=os.Open(filePath)
	if !os.IsExist(err) {
		fmt.Println("[ERR] ",err)
		os.Exit(0)
	}
 	defer fl.Close()
	
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, strings.Split(input, "\n")[0])
	}

	return output
}

//checks if the results are piped to another command
func isPiped() bool {
	info, _ := os.Stdin.Stat()
	if info.Mode()&os.ModeCharDevice == 0 {
		return true
	} 
		return false
}

//read input from stdin
func GetStdin() []string {
	var output []string
	if !isPiped() {
		return output 
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, strings.Split(input, "\n")[0])
	}

	return output
}

//return unshortened URL
func Start(url string, wg *sync.WaitGroup){
	defer wg.Done()
	res,err:=http.Head(url)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("[%s] %s => %s\n",res.Status,url,res.Request.URL.String())
}
