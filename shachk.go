package main

import (
	"os"
	"log"
	"io"
	"crypto/sha256"
	"fmt"
	"encoding/hex"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("\nYou have to provide two arguments. \n1: the file you wish to check. " +
			"\n2. the sha256 sum or a http/https link to a file containing the sha256 sum")
	}

	var fileFlag string

	fileFlag = os.Args[1]

	hasher := sha256.New()

	file, err := os.Open(fileFlag)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//Copy file to hasher
	if _, err := io.Copy(hasher, file); err != nil {
		log.Fatal(err)
	}

	fileSum := hex.EncodeToString(hasher.Sum(nil))

	trueSum := os.Args[len(os.Args) -1]

	if strings.ContainsAny(trueSum, "http | https") {
		if resp, err := http.Get(trueSum); err != nil {
			log.Fatal(err)
		} else {
			defer resp.Body.Close()
			body, _  := ioutil.ReadAll(resp.Body)
			trueSum = strings.Split(string(body), " ")[0]
		}
	}

	if fileSum == trueSum {
		fmt.Println("Sums are equal")
	} else {
		fmt.Println("WARNING! Sums are not equal")
	}

}