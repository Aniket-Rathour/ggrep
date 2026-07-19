package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/fs"
	"path"
	"strconv"
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
	
}

func grepLines(dir fs.FS, pattern string, addon string , name string , impPath string) ([]string, int){
	
	if addon == "-r" {
		var final []string
		var count int
		entries , err :=fs.ReadDir(dir , impPath )
		if err!= nil {
			return []string{"there was a error"} , 0
		} 
		
		for _, files := range entries {
			if files.IsDir(){
				newPath := path.Join(impPath,files.Name())
				
				final1 , count2 := grepLines(dir , pattern , addon , name , newPath)
				final = append(final, final1...)
				count += count2

			}else{
			
				newPath :=path.Join(impPath,files.Name())
				content ,err  := fs.ReadFile(dir , newPath)
				if err!= nil {
				return []string{"there was a error"} , 0
				} 
				
				reader := bytes.NewReader(content)
				scanner := bufio.NewScanner(reader)
				for scanner.Scan() {
					line := scanner.Text()
					if strings.Contains(line, pattern){
						final = append(final , line )
						count++
					}
					
				}
			}

		}
		return final, count
	}	
	mainfile , err := fs.ReadFile(dir , name)
	if err != nil {
		return []string{"error in reading file "} , 0
	}

	reader := bytes.NewReader(mainfile)
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
					matches = append(matches , line)
					counter++
				}
				fmt.Printf("%d" , counter)

			case "-n":
				line:=scanner.Text()
				if strings.Contains(line, pattern){
					toadd :=strconv.Itoa(linecount)
					matches=append(matches, toadd)
					
					counter++
				}

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