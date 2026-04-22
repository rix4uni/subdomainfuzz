package main

import (
    "bufio"
    "flag"
    "fmt"
    "io"
    "net/http"
    "os"
    "os/exec"
    "path/filepath"
    "runtime"
    "strings"

    "github.com/rix4uni/subdomainfuzz/banner"
)

const (
    defaultWordlistURL  = "https://wordlists-cdn.assetnote.io/data/manual/2m-subdomains.txt"
    defaultWordlistName = "2m-subdomains.txt"
)

func main() {
    // Parse command-line flags
    domain := flag.String("d", "", "Specify the domain to exclude")
    payload := flag.String("payload", "FUZZ", "Specify the payload to insert")
    silent := flag.Bool("silent", false, "silent mode.")
    versionFlag := flag.Bool("version", false, "Print the version of the tool and exit.")
    ffufMode := flag.Bool("ffuf", false, "Pipe output directly to ffuf")
    wordlist := flag.String("w", "", "Wordlist path for ffuf mode (auto-downloads default if not provided)")
    userAgent := flag.String("H", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/147.0.0.0 Safari/537.36", "User-Agent header for ffuf")
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

    // If ffuf mode is enabled but no wordlist provided, use/download default wordlist
    if *ffufMode && *wordlist == "" {
        defaultWordlist := getDefaultWordlistPath()
        if _, err := os.Stat(defaultWordlist); os.IsNotExist(err) {
            if !*silent {
                fmt.Printf("Downloading default wordlist to %s...\n", defaultWordlist)
            }
            if err := downloadDefaultWordlist(defaultWordlist); err != nil {
                fmt.Printf("Error downloading wordlist: %v\n", err)
                os.Exit(1)
            }
            if !*silent {
                fmt.Println("Wordlist downloaded successfully!")
            }
        }
        *wordlist = defaultWordlist
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
                targetURL := protocol + injected + suffix

                if *ffufMode {
                    // Run ffuf directly
                    cmd := exec.Command("ffuf", "-c", "-u", targetURL, "-w", *wordlist, "-v", "-H", "User-Agent: "+*userAgent)
                    cmd.Stdout = os.Stdout
                    cmd.Stderr = os.Stderr
                    cmd.Run()
                } else {
                    fmt.Println(targetURL)
                }
            }
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error reading input:", err)
    }
}

// getDefaultWordlistPath returns the path to the default wordlist based on OS
func getDefaultWordlistPath() string {
    var configDir string
    if runtime.GOOS == "windows" {
        configDir = os.Getenv("APPDATA")
        if configDir == "" {
            configDir = os.Getenv("USERPROFILE")
        }
    } else {
        configDir = os.Getenv("HOME")
        if configDir != "" {
            configDir = filepath.Join(configDir, ".config")
        }
    }
    return filepath.Join(configDir, "subdomainfuzz", defaultWordlistName)
}

// downloadDefaultWordlist downloads the default wordlist to the specified path
func downloadDefaultWordlist(filePath string) error {
    // Create directory if it doesn't exist
    dir := filepath.Dir(filePath)
    if err := os.MkdirAll(dir, 0755); err != nil {
        return fmt.Errorf("failed to create directory: %v", err)
    }

    // Download the file
    resp, err := http.Get(defaultWordlistURL)
    if err != nil {
        return fmt.Errorf("failed to download wordlist: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("bad status: %s", resp.Status)
    }

    // Create the file
    out, err := os.Create(filePath)
    if err != nil {
        return fmt.Errorf("failed to create file: %v", err)
    }
    defer out.Close()

    // Write the body to file
    _, err = io.Copy(out, resp.Body)
    if err != nil {
        return fmt.Errorf("failed to write file: %v", err)
    }

    return nil
}
