package main

import (
	"testing"
	"testing/fstest"
)

func TestCommands( t *testing.T){
	fakedir := fstest.MapFS{
		"aniket.txt": {Data : []byte("hiii my name is aniket\n i love bikes\ni am capatial Aniket\ni wouldlove if you can join me \naniket")},
		"aniket2.txt": {Data : []byte("hiii my name is aniket ")},
		"dir1/aniket.txt": {Data : []byte("hiii my name is aniket ")},
		"dir2/aniket.txt": {Data : []byte("hiii my name is aniket ")},
	}

	test := []struct{
		name string
		pattern string
		addon	string
		file	string
		wantLines	[]string
		wantCount	int
		path	string
	}{
		{
			name : "test for -i case",
			pattern : "aniket",
			addon : "dafault",
			file : "aniket.txt",
			wantLines : []string{"hiii my name is aniket" , "aniket"},
			wantCount: 2,
			path: "aniket.txt",
		},
		{
			name : "test for -n case",
			pattern : "aniket",
			addon : "-n",
			file : "aniket.txt",
			wantLines : []string{"1", "5"},
			wantCount: 2,
			path: "aniket.txt",
		},
		{
			name : "test for -n case",
			pattern : "aniket",
			addon : "-i",
			file : "aniket.txt",
			wantLines : []string{"hiii my name is aniket","i am capatial Aniket" , "aniket"},
			wantCount: 3,
			path: "aniket.txt",
		},
		{
			name : "test for -n case",
			pattern : "aniket",
			addon : "-r",
			file : "",
			wantLines : []string{"hiii my name is aniket" , "aniket", "hiii my name is aniket " , "hiii my name is aniket " , "hiii my name is aniket "},
			wantCount: 5,
			path: ".",
		},
		{
			name : "test for -n case",
			pattern : "name",
			addon : "-r",
			file : "",
			wantLines : []string{"hiii my name is aniket", "hiii my name is aniket " , "hiii my name is aniket " , "hiii my name is aniket "},
			wantCount: 4,
			path: ".",
		},
		{
			name : "test for -n case",
			pattern : "bikes",
			addon : "-r",
			file : "",
			wantLines : []string{" i love bikes"},
			wantCount: 1,
			path: ".",
		},
		

	}

	for _ , tt := range test {
		result , num := grepLines(fakedir ,tt.pattern , tt.addon ,tt.file, tt.path )
		//println("\n\n\n the result len is ", len(result) , "\n numebr is " , num)
		CheckError(t, result , num , tt.wantLines , tt.wantCount)
	}
}

func CheckError(t testing.TB ,lines []string ,count int , wantline []string , wantCount int){
	t.Helper()
	if count != wantCount{
		t.Errorf("wanted return numer %d , got %d",wantCount ,count)
	}

	if len(lines) != len(wantline){
		t.Errorf("there was a diff in the len of bohto got  %q  want %d",len(lines), len(wantline) )
	}
	for i , line := range lines{
		if line!= wantline[i]{
			t.Errorf("mismatch of the lines %d \n what was wanted was = %q ,what we got is %q ",i, wantline[i] , line)
		}
	}

	

}