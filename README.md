# Overview
The simple captcha cracker in GO language. It applies super fast, rather naive approach to recognize captions in captchas and works only for most simple captchas generated using not distorted font, but it can tolerate small distortion as well. Presented algorithm will try to build fingerprint for each symbol during training and later at test stage generated symbols fingeprints will be used to recognize captchas. It is really surpising how many systems in the world are protected with simple captchas which can be easy cracked with described method!

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
