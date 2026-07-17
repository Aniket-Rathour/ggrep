package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/fs"
	"strconv"

	//"fmt"
	"log"
	"os"
	"strings"
)

func main (){
	if len(os.Args) < 3 {
		log.Fatal("file name is missing")
	}
	file, err := os.Open(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
        if err := file.Close(); err != nil {
            log.Printf("error closing file: %v\n", err)
        }
    }()
	//Scanner := bufio.NewScanner(file);
	//pattern := "5"
	result , count := grepLines(os.DirFS("aniket"), os.Args[1] , os.Args[3] , "aniket.txt")
	fmt.Print(result ,count)
}

func grepLines(dir fs.FS, pattern string, addon string , name string) ([]string, int){
	file , err := fs.ReadFile(dir , name)
	if err != nil {
		return []string{"error in reading file "} , 0
	}

	reader := bytes.NewReader(file)
	scanner := bufio.NewScanner(reader)
	var matches []string

	counter := 0
	linecount := 1
	for scanner.Scan() {
		switch addon {
			case "-i":
				line:= scanner.Text()
				lineLower := strings.ToLower(line)
				pattern = strings.ToLower(pattern)
				if strings.Contains(lineLower ,pattern ){
					matches = append(matches , line)
					counter++
				}

			case "-c":
				line := scanner.Text()
				if strings.Contains(line, pattern){
					matches = append(matches , line )
					counter++
				}
				fmt.Printf("%d" , counter)

			case "-n":
				//println("hiii i am from -n ")
				line:=scanner.Text()
				if strings.Contains(line, pattern){
					toadd :=strconv.Itoa(linecount)
					matches=append(matches, toadd)
					//fmt.Printf("so we found the word in line = %q \n", matches)
					counter++
				}
			case "-r":



			default:
				//println("hii i am from defult")
				line := scanner.Text()
				if strings.Contains(line, pattern){
					matches = append(matches , line )
					counter++
					
				}
				
		}
		linecount++
		
	}
	return matches,counter
}