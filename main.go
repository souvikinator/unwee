package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	//"time"

	"github.com/DarthCucumber/unwee/utils"
)

type config struct {
	outputFile string
	inputFile  string
	urlList    []string
}

func main() {
	urlFlag := flag.String("u", "", "(url) takes in single URL as input and gives it's unshortened form")
	//outputFileFlag := flag.String("o", "", "(output) takes in file name as input and saves result in the file")
	urlFileFlag := flag.String("f", "", "(file) takes in file name containing list of shortened URLs")
	helpFlag := flag.Bool("h", false, "(help) prints help menu")
	//excludeFlag := flag.String("ex", "", "(exclude) takes in comma seperated response code(403,404,...etc) as input and excludes the result corresponding to those response code")
	flag.Parse()

	sessionConfig := config{}

	if *helpFlag {
		utils.DisplayHelp()
	}
	//get url input from -u flag
	if len(*urlFlag) > 0 {
		sessionConfig.urlList = append(sessionConfig.urlList, *urlFlag)
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
	if len(urls)==0{
		fmt.Println("[ERR] No URL provided")
		os.Exit(0)
	}
	
	//TODO: output file
	////after all the process is done
	//if len(*outputFileFlag) == 0 {
		////generate one with name: date_time.txt
		//*outputFileFlag = "./output/"+time.Now().Format("01-02-2006 15:04:05") + ".txt"
	//}
	////if file doesn't exists then create one
	//outFile, err := os.Create(*outputFileFlag)
	//utils.CheckErr(err)
	//outFile.Close()
	//sessionConfig.outputFile = *outputFileFlag

	fmt.Printf("[unwee] output file : %s\n", sessionConfig.outputFile)
	fmt.Printf("[unwee] Input file : %s\n", sessionConfig.inputFile)

	//do the main stuff
	var wg sync.WaitGroup
	for _, u := range urls {
		wg.Add(1)
		go utils.Start(u,&wg)
	}
	wg.Wait()

	fmt.Println("[unwee] Done")

}
