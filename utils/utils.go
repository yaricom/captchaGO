package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"strconv"
)

func ReadStringsData(path, suffix string, start, end int) []string {
	res := make([]string, end-start+1)
	for i := start; i <= end; i++ {
		file := fmt.Sprintf("%s/%s%02d.txt", path, suffix, i)
		res[i] = FileToString(file)
	}
	return res
}

func FileToString(file string) string {
	sBytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	return strings.TrimSpace(string(sBytes))
}

func CreateImgFingerprints(inputStr string) []string {
	fp := make([][]string, 5)
	for i := 0; i < 5; i++ {
		fp[i] = make([]string, 10)
	}
	rd := strings.NewReader(inputStr)
	var cols, rows int
	fmt.Fscanf(rd, "%d %d", &rows, &cols)
	pixels := cols * rows
	//fmt.Printf("cols: %d, rows: %d, pixels: %d\n", cols, rows, pixels)
	for i := 0; i < pixels; i++ {
		row := i / cols
		col := i - row * cols
		var r, g, b int
		_, err := fmt.Fscanf(rd, "%d,%d,%d", &r, &g, &b)
		if err != nil {
			// repeat scan when new line character found to get next triplet
			fmt.Fscanf(rd, "%d,%d,%d", &r, &g, &b)
		}
		if row >= 21 {
			// no need to iterate further
			break
		} else if row >= 11 {
			m := (r + g + b) / 3
			//fmt.Printf("r: %d, c: %d, m: %d\n", row, col, m)
			if col >= 5 && col < 49 {
				// find index
				c_ind := (col - 5) / 9
				if m < 50 {
					fp[c_ind][row - 11] += "1"
				} else {
					fp[c_ind][row - 11] += "0"
				}
			}
		}

	}

	// convert fingerprints
	res := make([]string, 5)
	for i := 0; i < 5; i++ {
		for _, v := range fp[i] {
			s, err := strconv.ParseInt(v, 2, 32)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			res[i] += fmt.Sprintf("%d-", s)
		}
	}

	return res
}

func MergeImgFingerprints(outStr string, fp []string) map[string]string {
	res := make(map[string]string)
	for i := 0; i < 5; i++ {
		str := string(outStr[i])
		res[fp[i]] = str
	}
	return res
}
