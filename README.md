## subdomainfuzz

A lightweight tool for inserting custom payloads into URL subdomains, for domain mutation after that pass this output to ffuf. I got this tool idea from https://x.com/ArmanSameer95/status/1680811916053078019

## Installation
```
go install github.com/rix4uni/subdomainfuzz@latest
```

## Download prebuilt binaries
```
wget https://github.com/rix4uni/subdomainfuzz/releases/download/v0.0.1/subdomainfuzz-linux-amd64-0.0.1.tgz
tar -xvzf subdomainfuzz-linux-amd64-0.0.1.tgz
rm -rf subdomainfuzz-linux-amd64-0.0.1.tgz
mv subdomainfuzz ~/go/bin/subdomainfuzz
```
Or download [binary release](https://github.com/rix4uni/subdomainfuzz/releases) for your platform.

## Compile from source
```
git clone --depth 1 github.com/rix4uni/subdomainfuzz.git
cd subdomainfuzz; go install
```

## Usage
```
Usage of subdomainfuzz:
  -d string
        Specify the domain to exclude
  -payload string
        Specify the payload to insert
  -silent
        silent mode.
  -version
        Print the version of the tool and exit.
```

## Usage Example
```
▶ cat alivesubs.txt
https://rix4uni.domain.com
http://pentestingdorks.domain.com:8080
https://linkinspector.admin.domain.com
```

### Output
```
▶ cat alivesubs.txt | subdomainfuzz -silent -d domain.com -payload "FUZZ"
https://FUZZ.rix4uni.domain.com
https://rix4uni.FUZZ.domain.com
http://FUZZ.pentestingdorks.domain.com:8080
http://pentestingdorks.FUZZ.domain.com:8080
https://FUZZ.linkinspector.admin.domain.com
https://linkinspector.FUZZ.admin.domain.com
https://linkinspector.admin.FUZZ.domain.com
```