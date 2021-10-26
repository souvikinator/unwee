# unwee

A simple tool written in `go` to unshorten short urls.
Takes bunch of URLs and returns the unshortened URLs

## TODO:

- [ ] Code Cleanup 😓 
- [x] Proper flags for URL inputs (currently only takes input from STDIN)
- [x] Flag for output files to store the results
- [x] Support for output format(in JSON, currently supports TXT)
- [ ] `-ex` (exclude) flag to not include unreachable/invalid URLs in the output file

more freatures if possible?

## Installation

```bash
go get -u github.com/DarthCucumber/unwee
```

## Usage

```
 _  _ _ ___ __ _____ ___ 
| || | ' \ V  V / -_) -_)
 \_,_|_||_\_/\_/\___\___| v1.1.0
 
usage: unwee [options...]

options: 

[-u] (url) takes in single URL as input and gives it's unshortened form.
[-o] (output) takes in file name as input and saves result in the file
[-f] (file) takes in file name containing list of shortened URLs
[-h] (help) prints help menu

Example:

- Single URL Input from stdin:
	> echo "http://shorturl/xyz" | go run main.go
- Input single url
	> go run main.go -u http://shorturl/xyz
- URL List file Input from stdin
	> cat url_list.txt | go run main.go 
- Input from file
	> go run main.go -f url_list.txt
- Setting output file to save results (default is saved at "output/<data_time>.txt")
	> go run main.go -o outputfile.txaer
```

### note
There is a sample txt file in the repo with some shortened URLs in it for testing. So go ahead give it a try.
