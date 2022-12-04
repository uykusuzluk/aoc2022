package aocutils

import (
	"bufio"
	"log"
	"net/http"
	"os"
)

func InputLinesFromURL(url string, q chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("cannot send http get req url: " + url + " err: " + err.Error())
	}
	defer resp.Body.Close()

	rdr := bufio.NewScanner(resp.Body)
	rdr.Split(bufio.ScanLines)

	for rdr.Scan() {
		q <- rdr.Text()
	}
	close(q)
}

func InputLinesFromFile(filepath string, q chan<- string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalln("cannot open file at: " + filepath + " err: " + err.Error())
	}
	defer file.Close()

	rdr := bufio.NewScanner(file)
	rdr.Split(bufio.ScanLines)

	for rdr.Scan() {
		q <- rdr.Text()
	}
	close(q)
}
