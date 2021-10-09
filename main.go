package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/DarthCucumber/unwee/utils"
)

type config struct {
	outputFile  string
	urlList     []string
	resultList  []string
	excludeList []int
}

func main() {
	urlFlag := flag.String("u", "", "(url) takes in single URL as input and gives it's unshortened form")
	outputFileFlag := flag.String("o", "", "(output) takes in file name as input and saves result in the file")
	urlFileFlag := flag.String("f", "", "(file) takes in file name containing list of shortened URLs")
	helpFlag := flag.Bool("h", false, "(help) prints help menu")
	excludeFlag := flag.String("ex", "", "(exclude) takes in comma seperated response code(403,404,...etc) as input and excludes the result corresponding to those response code")
	outputFileExtFlag := flag.Bool("j", false, "(output file extension) true if you want it to save in json file else the result will be saved in a text file")

	flag.Parse()

	sessionConfig := config{}

	if *helpFlag {
		utils.DisplayHelp()
	}
	//get url input from -u flag
	if len(*urlFlag) > 0 {
		sessionConfig.urlList = append(sessionConfig.urlList, *urlFlag)
	}

	if len(*excludeFlag) > 0 {
		e := strings.Split(*excludeFlag, ",")
		for _, code := range e {
			p, err := strconv.Atoi(code)
			utils.CheckErr(err)
			sessionConfig.excludeList = append(sessionConfig.excludeList, p)
		}
	}
	//get url input from file
	if len(*urlFileFlag) > 0 {
		urlList := utils.GetFileData(*urlFileFlag)
		if len(urlList) == 0 {
			fmt.Println("[ERR] ", "Empty URL list provided")
			os.Exit(0)
		}
		sessionConfig.urlList = append(sessionConfig.urlList, urlList...)
	}
	//get url input from stdin
	stdinUrl := utils.GetStdin()
	sessionConfig.urlList = append(sessionConfig.urlList, stdinUrl...)

	urls := sessionConfig.urlList
	if len(urls) == 0 {
		fmt.Println("[ERR] No URL provided")
		os.Exit(0)
	}

	//TODO: add support to export in JSON,CSV and other
	//after all the process is done

	// Type of file
	fileExt := ".txt" // default value
	/* Check if the FileExtFlag is true,
	If it is then the fileExt value will be json else it will be default .txt */
	if *outputFileExtFlag {
		fileExt = ".json"
	}
	if len(*outputFileFlag) == 0 {
		*outputFileFlag = "./output/" + time.Now().Format("01-02-2006_15:04:05") + fileExt
	} else {
		// If Json FileName is given as input
		fileNameArr := strings.Split(*outputFileFlag, ".")
		if fileNameArr[1] == "json" {
			fileExt = ".json"
		}
	}

	sessionConfig.outputFile = *outputFileFlag

	sessionConfig.display()

	fmt.Printf("\n[*] Starting...\n")

	//do the main stuff
	var wg sync.WaitGroup
	for _, u := range urls {
		wg.Add(1)
		go utils.Start(u, &wg, &sessionConfig.resultList, &sessionConfig.excludeList)
	}
	wg.Wait()

	fmt.Printf("\n[*] Done")
	/*
		if the fileExt is .txt, then call WriteToTextFile
		else check whether the fileExt is .json, if it's true then call WriteToJsonFile
	*/
	fmt.Printf("\n[*] Saving results to text file: %s\n", *outputFileFlag)
	if fileExt == ".txt" {
		utils.WriteToTextFile(sessionConfig.resultList, *outputFileFlag)
	} else if fileExt == ".json" {
		utils.WriteToJsonFile(sessionConfig.resultList, *outputFileFlag)
	}

	fmt.Printf("[*] Saved\n")
}

//prints config before starting
func (c config) display() {
	var cf1, cf2 string
	if len(c.outputFile) > 0 {
		cf1 = fmt.Sprintf("Output File: %s", c.outputFile)
	}
	cf2 = fmt.Sprintf("Number of URLs: %d", len(c.urlList))
	fmt.Printf(`
 _  _ _ ___ __ _____ ___ 
| || | ' \ V  V / -_) -_)
 \_,_|_||_\_/\_/\___\___| v1.1.0

---------------------------------------
%s
%s
---------------------------------------`, cf2, cf1)
}
