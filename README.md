# unwee

A simple tool written in `go` to unshorten short urls.
Takes bunch of URLs and returns the unshortened URLs

## Installation

```bash
go get -u github.com/DarthCucumber/Tools/unwee
```

## Usage

```
usage: unwee [options...]

options:

[-l] (labelled) prints input url with corresponding result url

[-h] (help) prints help menu

[-m] (mute) mutes any error

```

## Example

### Passing single URL
```bash
echo "http://vsurl.ta/aha8q12" | unwee -l
```
Result:
```bash
http://vsurl.ta/rndpth => http://veryshoturl.ta/random/path
```

### Passing multiple URLs
let us have a text file `test_urls.txt` with list of short urls in it. Lets 					
`unwee` them ; )
```bash
cat test_url.txt | unwee -l
```
and you know what the result will be.

in addition user can pass options with `unwee` as per their needs.

### `-l` or labelled   

Allows user to print result in following format 
`[short url] => [corresponding long url]`

but sometimes one may need only the results so **excluding this flag** will print only results

### `-m` or mute   

Allows user to mute any URL related errors (404,403...)
Without this flag any error gets printed in the screen.

### note
There is a sample txt file in the repo with some shortened URLs in it for testing. So go ahead give it a try.
