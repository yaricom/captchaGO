package main

import (
	"os"
	"fmt"
	"strconv"
	"sort"
	"github.com/yaricom/captchaGO/utils"
)

func main() {
	if len(os.Args) < 4 {
		printHelp()
		os.Exit(0)
	}
	inputPath := os.Args[1] + "/input"
	outPath := os.Args[1] + "/output"
	start, _ := strconv.Atoi(os.Args[2])
	end, _ := strconv.Atoi(os.Args[3])

	outStrs := utils.ReadStringsData(outPath, "output", start, end)
	inStrs := utils.ReadStringsData(inputPath, "input", start, end)

	if len(outStrs) != len(inStrs) {
		fmt.Printf("The output strings count: %d not equal to input strings count: %d",
			len(outStrs), len(inStrs))
	}

	if os.Args[4] == "learn" {
		learnSymbolsMap(outStrs, inStrs)
	} else if os.Args[4] == "test" {
		testSymbolsMap(outStrs, inStrs)
	}
}

func learnSymbolsMap(outStrs[]string, inStrs[]string) {
	symbMap := make(map[string]string)
	for i := 0; i < len(inStrs); i++ {
		fp := utils.CreateImgFingerprints(inStrs[i])
		outMap := utils.MergeImgFingerprints(outStrs[i], fp)
		for k, v := range outMap {
			if val, ok := symbMap[k]; ok {
				if val != v {
					// symbol registered but with different fingerprint
					fmt.Printf("Fingerprint: %s registered for different symbols: %s <> %s\n",
						k, v, val)
					os.Exit(1)
				}
			} else {
				symbMap[k] = v
			}
		}
	}
	// sort keys
	sortedKeys := make([]string, len(symbMap))
	i := 0
	for k := range symbMap {
		sortedKeys[i] = k
		i += 1
	}
	sort.Strings(sortedKeys)

	// output collected symbols map
	fmt.Println("symbMap := map[string]string {")
	for _, k := range sortedKeys {
		fmt.Printf("\"%s\" : \"%s\",\n", k, symbMap[k])
	}
	fmt.Println("}")
}

func testSymbolsMap(outStrs[]string, inStrs[]string) {
	symbMap := map[string]string{
		"12-28-60-108-204-396-510-12-12-12-" : "4",
		"120-204-388-384-440-460-390-390-204-120-" : "6",
		"120-204-390-204-120-204-390-390-204-120-" : "8",
		"120-204-390-390-206-118-6-134-204-120-" : "9",
		"120-204-390-390-390-390-390-390-204-120-" : "O",
		"120-204-390-390-390-390-438-414-204-122-" : "Q",
		"120-204-390-6-12-24-48-96-192-510-" : "2",
		"124-198-384-384-384-398-390-390-198-124-" : "G",
		"124-198-386-384-384-384-384-386-198-124-" : "C",
		"248-396-6-12-56-12-6-6-396-248-" : "3",
		"252-390-384-384-252-6-6-6-390-252-" : "S",
		"252-48-48-48-48-48-48-48-48-252-" : "I",
		"384-384-384-384-384-384-384-384-384-508-" : "L",
		"390-390-204-120-48-48-120-204-390-390-" : "X",
		"390-390-204-120-48-48-48-48-48-48-" : "Y",
		"390-390-390-204-204-204-120-120-48-48-" : "V",
		"390-390-390-390-390-390-390-390-204-120-" : "U",
		"390-390-390-390-438-438-438-510-462-390-" : "W",
		"390-390-390-390-510-390-390-390-390-390-" : "H",
		"390-396-408-432-480-480-432-408-396-390-" : "K",
		"390-454-486-486-438-438-414-398-398-390-" : "N",
		"390-462-510-438-438-438-390-390-390-390-" : "M",
		"48-112-240-48-48-48-48-48-48-252-" : "1",
		"48-120-204-390-390-390-390-204-120-48-" : "0",
		"48-120-204-390-390-390-510-390-390-390-" : "A",
		"504-396-390-390-390-390-390-390-396-504-" : "D",
		"504-396-390-396-504-396-390-390-396-504-" : "B",
		"508-12-12-24-48-96-192-384-384-508-" : "Z",
		"508-384-384-384-504-384-384-384-384-508-" : "E",
		"508-384-384-440-460-6-6-390-204-120-" : "5",
		"508-390-390-390-508-384-384-384-384-384-" : "P",
		"508-390-390-390-508-496-408-396-390-390-" : "R",
		"510-384-384-384-504-384-384-384-384-384-" : "F",
		"510-48-48-48-48-48-48-48-48-48-" : "T",
		"510-6-6-12-24-48-96-192-384-384-" : "7",
		"60-12-12-12-12-12-12-140-216-112-" : "J",
	}
	for i := 0; i < len(inStrs); i++ {
		fps := utils.CreateImgFingerprints(inStrs[i])
		captcha := ""
		for _, fp := range fps {
			fmt.Printf("%s : %s\n", fp, symbMap[fp])
			captcha += symbMap[fp]
		}
		fmt.Println(captcha)
		if captcha != outStrs[i] {
			fmt.Printf("Parsed captcha: [%s] is not equal to GT: [%s]",
			captcha, outStrs[i])
			os.Exit(1)
		}
	}

	fmt.Println("Captcha test passed OK")
}

func printHelp() {
	fmt.Println("Arguments:")
	fmt.Println("dataDir - the path to the directory with data files")
	fmt.Println("start - the start suffix")
	fmt.Println("end - the end suffix")
	fmt.Println("operation - the operation to perform [learn, test]")
}
