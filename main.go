package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

const (
	apiURL  = "https://urlscan.io/api/v1/search/?q=domain:%s"
	version = "1.0"
)

type Result struct {
	Results []struct {
		Page struct {
			Domain string `json:"domain"`
		} `json:"page"`
	} `json:"results"`
}

func showBanner() {
	banner := `
 ██╗   ██╗██████╗ ██╗     ███████╗ ██████╗ █████╗ ███╗   ██╗███╗   ██╗███████╗██████╗ 
 ██║   ██║██╔══██╗██║     ██╔════╝██╔════╝██╔══██╗████╗  ██║████╗  ██║██╔════╝██╔══██╗
 ██║   ██║██████╔╝██║     ███████╗██║     ███████║██╔██╗ ██║██╔██╗ ██║█████╗  ██████╔╝
 ██║   ██║██╔══██╗██║     ╚════██║██║     ██╔══██║██║╚██╗██║██║╚██╗██║██╔══╝  ██╔══██╗
 ╚██████╔╝██║  ██║███████╗███████║╚██████╗██║  ██║██║ ╚████║██║ ╚████║███████╗██║  ██║
  ╚═════╝ ╚═╝  ╚═╝╚══════╝╚══════╝ ╚═════╝╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝  ╚═══╝╚══════╝╚═╝  ╚═╝

           URLScan.io Powered Subdomain Discovery Tool v%s
                           by: Psikoz
                    github.com/Psikoz-coder/urlscanner
`
	fmt.Printf(banner, version)
	fmt.Printf("\n%s\n\n", strings.Repeat("=", 70))
}

func showUsage() {
	fmt.Println("Usage:")
	fmt.Println("  urlscanner -set <API_KEY>           Set URLScan.io API key")
	fmt.Println("  urlscanner -d <domain>              Discover subdomains")
	fmt.Println("  urlscanner -dL <file>               Discover subdomains from domain list")
	fmt.Println("  urlscanner -d <domain> -o <file>    Save results to file")
	fmt.Println("  urlscanner -d <domain> -silent      Silent mode (no banner)")
	fmt.Println("\nExamples:")
	fmt.Println("  urlscanner -set abc123def456")
	fmt.Println("  urlscanner -d example.com")
	fmt.Println("  urlscanner -dL domains.txt -o results.txt")
	fmt.Println("  urlscanner -d example.com -o subdomains.txt -silent")
}

func getHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return usr.HomeDir
}

func saveAPIKey(apiKey string) error {
	dir := filepath.Join(getHomeDir(), ".urlscanner")
	os.MkdirAll(dir, 0755)
	return ioutil.WriteFile(filepath.Join(dir, "config.json"), []byte(`{"apikey":"`+apiKey+`"}`), 0644)
}

func loadAPIKey() (string, error) {
	data, err := ioutil.ReadFile(filepath.Join(getHomeDir(), ".urlscanner", "config.json"))
	if err != nil {
		return "", err
	}
	var conf map[string]string
	err = json.Unmarshal(data, &conf)
	if err != nil {
		return "", err
	}
	return conf["apikey"], nil
}

func fetchSubdomains(domain, apiKey string, silent bool) ([]string, error) {
	if !silent {
		fmt.Printf("[*] Scanning subdomains for: %s\n", domain)
	}
	
	url := fmt.Sprintf(apiURL, domain)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("API-Key", apiKey)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result Result
	json.NewDecoder(resp.Body).Decode(&result)

	subdomains := make(map[string]struct{})
	for _, entry := range result.Results {
		if strings.HasSuffix(entry.Page.Domain, domain) {
			subdomains[entry.Page.Domain] = struct{}{}
		}
	}

	unique := make([]string, 0, len(subdomains))
	for sub := range subdomains {
		unique = append(unique, sub)
	}

	if !silent {
		fmt.Printf("[+] Found %d unique subdomains\n", len(unique))
	}

	return unique, nil
}

func readDomainsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var domains []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domain := strings.TrimSpace(scanner.Text())
		if domain != "" {
			domains = append(domains, domain)
		}
	}

	return domains, scanner.Err()
}

func writeResultsToFile(filename string, results []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, result := range results {
		fmt.Fprintln(writer, result)
	}
	return writer.Flush()
}

func main() {
	var silent bool
	var outputFile string
	var domainList string
	var domain string
	
	// Parse arguments
	args := os.Args[1:]
	if len(args) == 0 {
		showBanner()
		showUsage()
		return
	}

	// Check for silent mode first
	for _, arg := range args {
		if arg == "-silent" {
			silent = true
			break
		}
	}

	if !silent {
		showBanner()
	}

	// Parse other arguments
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-set":
			if i+1 >= len(args) {
				fmt.Println("[!] Error: API key not specified")
				return
			}
			if err := saveAPIKey(args[i+1]); err != nil {
				fmt.Println("[!] Error: Could not save API key:", err)
				return
			}
			fmt.Println("[+] API key saved successfully")
			return

		case "-d":
			if i+1 >= len(args) {
				fmt.Println("[!] Error: Domain not specified")
				return
			}
			domain = args[i+1]
			i++

		case "-dL":
			if i+1 >= len(args) {
				fmt.Println("[!] Error: Domain list file not specified")
				return
			}
			domainList = args[i+1]
			i++

		case "-o":
			if i+1 >= len(args) {
				fmt.Println("[!] Error: Output file not specified")
				return
			}
			outputFile = args[i+1]
			i++

		case "-silent":
			// Already handled
			continue

		default:
			// Skip unknown arguments
		}
	}

	// Load API key
	apiKey, err := loadAPIKey()
	if err != nil {
		fmt.Println("[!] Error: Could not load API key. Please set it first using -set")
		return
	}

	var allSubdomains []string

	// Process single domain
	if domain != "" {
		subdomains, err := fetchSubdomains(domain, apiKey, silent)
		if err != nil {
			fmt.Println("[!] Error: Request failed:", err)
			return
		}
		allSubdomains = append(allSubdomains, subdomains...)
	}

	// Process domain list
	if domainList != "" {
		domains, err := readDomainsFromFile(domainList)
		if err != nil {
			fmt.Println("[!] Error: Could not read domain list:", err)
			return
		}

		if !silent {
			fmt.Printf("[*] Processing %d domains from file\n", len(domains))
		}

		for _, d := range domains {
			subdomains, err := fetchSubdomains(d, apiKey, silent)
			if err != nil {
				if !silent {
					fmt.Printf("[!] Error processing %s: %v\n", d, err)
				}
				continue
			}
			allSubdomains = append(allSubdomains, subdomains...)
		}
	}

	// Output results
	if len(allSubdomains) == 0 {
		if !silent {
			fmt.Println("[!] No subdomains found")
		}
		return
	}

	// Remove duplicates when processing multiple domains
	uniqueSubdomains := make(map[string]struct{})
	for _, sub := range allSubdomains {
		uniqueSubdomains[sub] = struct{}{}
	}

	finalResults := make([]string, 0, len(uniqueSubdomains))
	for sub := range uniqueSubdomains {
		finalResults = append(finalResults, sub)
	}

	// Write to file if specified
	if outputFile != "" {
		if err := writeResultsToFile(outputFile, finalResults); err != nil {
			fmt.Println("[!] Error: Could not write to file:", err)
			return
		}
		if !silent {
			fmt.Printf("[+] Results saved to: %s\n", outputFile)
		}
	} else {
		// Print to stdout
		for _, sub := range finalResults {
			fmt.Println(sub)
		}
	}

	if !silent {
		fmt.Printf("\n[+] Total unique subdomains found: %d\n", len(finalResults))
	}
}
