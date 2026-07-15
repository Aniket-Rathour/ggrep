package main

import (
	"bufio"
	"fmt"
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
	defer file.Close()
	Scanner := bufio.NewScanner(file);
	//pattern := "5"
	result , count := grepLines(Scanner , os.Args[1] , os.Args[3])
	fmt.Print(result ,count)
}

func grepLines(scanner *bufio.Scanner , pattern string, addon string) ([]string, int){
	var matches []string
	counter := 0
	linecount := 0
	for scanner.Scan() {
		switch addon {
			case "-i":
				line:= scanner.Text()
				line = strings.ToLower(line)
				pattern = strings.ToLower(pattern)
				if strings.Contains(line ,pattern ){
					matches = append(matches , line)
				}
			case "-c":
				line := scanner.Text()
				if strings.Contains(line, pattern){
					matches = append(matches , line )
					counter++
				}
				fmt.Printf("%d" , counter)

			case "-n":
				line:=scanner.Text()
				if strings.Contains(line, pattern){
					matches=append(matches, pattern)
					fmt.Printf("so we found the word in line = %d \n", linecount)
				}
			default:
				line := scanner.Text()
				if strings.Contains(line, pattern){
					matches = append(matches , line )
					
				}
				
		}
		linecount++
		
	}
	return matches,counter
}