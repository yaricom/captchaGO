package main

import(
	"os"
	//"bufio"
	"fmt"
	"io/ioutil"
	"strconv"

)

func readStringsData(path, suffix string, start, end int) [] string {
	res := make([]string, end - start + 1)
	for i := start; i <= end; i++ {
		file := fmt.Sprintf("%s/%s%02d.txt", path, suffix, i)
		res[i] = fileToString(file)
	}
	return res
}

func fileToString(file string) string  {
	sBytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	return string(sBytes)
}


func main() {
	if len(os.Args) < 4 {
		printHelp()
		os.Exit(0)
	}
	inputPath := os.Args[1] + "/input"
	outPath := os.Args[1] + "/output"
	start, _ := strconv.Atoi(os.Args[2])
	end, _ := strconv.Atoi(os.Args[3])

	outStrs := readStringsData(outPath, "output", start, end)
	inStrs := readStringsData(inputPath, "input", start, end)

	if len(outStrs) != len(inStrs) {
		fmt.Printf("The output strings count: %d not equal to input strings count: %d",
			len(outStrs), len(inStrs))
	}
}

func printHelp()  {
	fmt.Println("Arguments:")
	fmt.Println("dataDir - the path to the directory with data files")
	fmt.Println("start - the start suffix")
	fmt.Println("end - the end suffix")
}
