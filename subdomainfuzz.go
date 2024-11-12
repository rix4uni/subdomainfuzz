package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "strings"

    "github.com/rix4uni/subdomainfuzz/banner"
)

func main() {
    // Parse command-line flags
    domain := flag.String("d", "", "Specify the domain to exclude")
    payload := flag.String("payload", "", "Specify the payload to insert")
    silent := flag.Bool("silent", false, "silent mode.")
    versionFlag := flag.Bool("version", false, "Print the version of the tool and exit.")
    flag.Parse()

    if *versionFlag {
        banner.PrintBanner()
        banner.PrintVersion()
        return
    }

    if !*silent {
        banner.PrintBanner()
    }

    if *domain == "" || *payload == "" {
        fmt.Println("Usage: subdomainfuzz -d <domain> -payload <payload>")
        os.Exit(1)
    }

    scanner := bufio.NewScanner(os.Stdin)

    // Read each line (URL) from input
    for scanner.Scan() {
        url := scanner.Text()

        // Check if the URL contains the domain to exclude
        if strings.Contains(url, *domain) {
            // Find the protocol (http:// or https://) and the rest of the URL
            var protocol, address string
            if strings.HasPrefix(url, "http://") {
                protocol = "http://"
                address = strings.TrimPrefix(url, "http://")
            } else if strings.HasPrefix(url, "https://") {
                protocol = "https://"
                address = strings.TrimPrefix(url, "https://")
            } else {
                // If there's no protocol, treat the whole URL as the address
                address = url
            }

            // Find the index of the domain in the address
            domainIndex := strings.Index(address, *domain)
            prefix := address[:domainIndex]
            suffix := address[domainIndex:]

            // Insert the payload at each insertion point before the domain
            segments := strings.Split(prefix, ".")
            for i := range segments {
                injected := strings.Join(append(segments[:i], append([]string{*payload}, segments[i:]...)...), ".")
                fmt.Println(protocol + injected + suffix)
            }
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error reading input:", err)
    }
}
