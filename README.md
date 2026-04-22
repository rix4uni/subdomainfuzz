## subdomainfuzz

A lightweight tool for inserting custom payloads into URL subdomains, for domain mutation after that pass this output to ffuf. I got this tool idea from https://x.com/ArmanSameer95/status/1680811916053078019

## Prerequisites

For `-ffuf` mode, install ffuf first:
```
go install github.com/ffuf/ffuf/v2@latest
```

## Installation
```
go install github.com/rix4uni/subdomainfuzz@latest
```

## Download prebuilt binaries
```
wget https://github.com/rix4uni/subdomainfuzz/releases/download/v0.0.2/subdomainfuzz-linux-amd64-0.0.2.tgz
tar -xvzf subdomainfuzz-linux-amd64-0.0.2.tgz
rm -rf subdomainfuzz-linux-amd64-0.0.2.tgz
mv subdomainfuzz ~/go/bin/subdomainfuzz
```
Or download [binary release](https://github.com/rix4uni/subdomainfuzz/releases) for your platform.

## Compile from source
```
git clone --depth 1 github.com/rix4uni/subdomainfuzz.git
cd subdomainfuzz; go install
```

## Usage
```console
Usage of subdomainfuzz:
  -d string
        Specify the domain to exclude
  -ffuf
        Pipe output directly to ffuf
  -H string
        User-Agent header for ffuf (default "Mozilla/5.0 (Windows NT 10.0; Win64; x64)...")
  -payload string
        Specify the payload to insert (default "FUZZ")
  -silent
        silent mode.
  -version
        Print the version of the tool and exit.
  -w string
        Wordlist path for ffuf mode (auto-downloads default if not provided)
```

## Usage Example
```console
▶ cat alivesubs.txt
https://rix4uni.domain.com
http://pentestingdorks.domain.com:8080
https://linkinspector.admin.domain.com
```

### Output
```console
▶ cat alivesubs.txt | subdomainfuzz -silent -d domain.com -payload "FUZZ"
https://FUZZ.rix4uni.domain.com
https://rix4uni.FUZZ.domain.com
http://FUZZ.pentestingdorks.domain.com:8080
http://pentestingdorks.FUZZ.domain.com:8080
https://FUZZ.linkinspector.admin.domain.com
https://linkinspector.FUZZ.admin.domain.com
https://linkinspector.admin.FUZZ.domain.com
```

## ffuf Mode

Run ffuf directly without manually piping output:

### With custom wordlist
```console
▶ echo "https://outlet.dell.com" | subdomainfuzz -silent -d dell.com -ffuf -w wordlist.txt
```

### With auto-downloaded wordlist (2m-subdomains.txt)
```console
▶ echo "https://outlet.dell.com" | subdomainfuzz -silent -d dell.com -ffuf
Downloading default wordlist to ~/.config/subdomainfuzz/2m-subdomains.txt...
Wordlist downloaded successfully!
```

### With custom User-Agent
```console
▶ echo "https://outlet.dell.com" | subdomainfuzz -silent -d dell.com -ffuf -H "Custom-Agent/1.0"
```