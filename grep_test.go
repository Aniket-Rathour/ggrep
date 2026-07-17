package main

import (
	"testing"
	"testing/fstest"
)

func TestCommands( t *testing.T){
	fakedir := fstest.MapFS{
		"aniket.txt": {Data : []byte("hiii my name is aniket\n i love bikes \ni wouldlove if you can join me ")},
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
	}{
		{
			name : "test for -i case",
			pattern : "aniket",
			addon : "dafault",
			file : "aniket.txt",
			wantLines : []string{"hiii my name is aniket"},
			wantCount: 1,
		},
	}

	for _ , tt := range test {
		result , num := grepLines(fakedir ,tt.pattern , tt.addon ,tt.file )

		CheckError(t, result , num , tt.wantLines , tt.wantCount)
	}
}

func CheckError(t testing.TB ,lines []string ,count int , wantline []string , wantCount int){
	t.Helper()
	if count != wantCount{
		t.Errorf("wanted return numer %d , got %d",wantCount ,count)
	}

	if len(lines) != len(wantline){
		t.Errorf("there was a diff in the len of bohto got  %q  want %d",len(lines) , len(wantline) )
	}
	for i , line := range lines{
		if line!= wantline[i]{
			t.Errorf("mismatch of the lines \n what was wanted was = %q ,what we got is %q ", wantline[i] , line)
		}
	}

	

}