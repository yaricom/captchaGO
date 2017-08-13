# Overview
The simple captcha cracker in GO language.

# Usage

$**go run main.go** dataDir start end operation

where:

* dataDir - is the directory with input files and ground truth files
* start - the start index of first file to process (from dataDir/input)
* end - the end index of last file to process (from dataDir/input)
* operation - the operation to perform [learn, test]

## Example

To start captcha's fonte details learning execute following:

$**go run main.go** ./data 0 25 learn

To test learned details copy generated output into coresponding function and execute:

$**go run main.go** ./data 0 25 test
