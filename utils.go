package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

func loadURLhaus() {
	currentDate := time.Now().Format("2006-01-02")
	fileName := currentDate + "-urlhausurls.txt"

	_, err := os.Stat(fileName)

	// CHECK IF FILE EXISTS
	if os.IsNotExist(err) { // No
		log.Println("Looks like the URLhaus file doesn't exist. Fetching the latest data, hold tight...")
		fetchAndSaveURLhausData(fileName)
	} else { // Yes
		fileInfo, _ := os.Stat(fileName)
		modTime := fileInfo.ModTime().Format("2006-01-02")

		// Check if the file was modified on the present date
		if currentDate == modTime { // Yes
			log.Println("URLhaus file is already up-to-date. Reading...")
		} else { // No
			log.Println("Looks like the URLhaus file is old. Fetching the latest data, hold tight...")
			fetchAndSaveURLhausData(fileName)
		}
	}
	log.Println("Loading the URLHaus data into the memory...")
	loadURLhausdata(fileName)
}

func loadURLhausdata(fileName string) {
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	scanner := bufio.NewScanner(strings.NewReader(string(fileContent)))

	for scanner.Scan() {
		urlhausURLs = append(urlhausURLs, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanning file content:", err)
	}
}

func fetchAndSaveURLhausData(fileName string) {
	response, err := http.Get("https://urlhaus.abuse.ch/downloads/text/")
	if err != nil {
		log.Fatal("Error fetching URLhaus data:", err)
	}

	defer response.Body.Close()

	scanner := bufio.NewScanner(response.Body)
	var filteredLines []string

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "#") {
			filteredLines = append(filteredLines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading response body:", err)
	}

	data := strings.Join(filteredLines, "\n")
	err = ioutil.WriteFile(fileName, []byte(data), 0644)
	if err != nil {
		log.Fatal("Error saving data to file:", err)
	}

	log.Println("URLhaus data fetched and saved to", fileName)
}

func takeInput() {
	// Define the flags
	flag.BoolVar(&depthFlag, "depth", false, "Specifies whether to use the depth mode. [Scan for root domains and IP addresses.]")
	flag.StringVar(&saveFlag, "save", "output.json", "Specifies the output file name. If kept empty, default output file name is output.json")

	flag.Parse()

	cliArgs = flag.Args()
	if len(cliArgs) > 0 {
		for _, cliArg := range cliArgs {
			if strings.HasPrefix(cliArg, "https") {
				targets = append(targets, cliArg)
			} else {
				targets = append(targets, "http://"+cliArg)
				targets = append(targets, "https://"+cliArg)
			}
		}
	} else {
		log.Println("Looks like no arguments were specified! Exiting...")
		os.Exit(0)
	}
}

func makeReq(urlToFetch string) string {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // Ignore SSL certificate validation
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// Allow redirects
			log.Println("Redirecting to:", req.URL)
			return nil
		},
		Timeout: 30 * time.Second, // Set timeout to 30 seconds
	}

	// Make a GET request
	response, err := client.Get(urlToFetch)
	if err != nil {
		log.Println("Error:", err)
		return ""
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Println("Error reading response body:", err)
			return ""
		} else {
			return string(body)
		}
	} else {
		log.Println("Target returned a non-200 status code...")
		return ""
	}
}

func extractURLs(input string) []string {
	urlRegex := regexp.MustCompile(`(?i)\b(https?|ftp):\/\/(?:[-a-zA-Z0-9@:%._\+~#=]+|\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})(?::\d+)?(?:\/[-a-zA-Z0-9@:%_\+.~#?&//=]*)?\b|\bwww\.(?:[-a-zA-Z0-9@:%._\+~#=]+|\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})(?::\d+)?(?:\/[-a-zA-Z0-9@:%_\+.~#?&//=]*)?\b`)
	matches := urlRegex.FindAllString(input, -1)

	return matches
}

func uniqueList(input []string) []string {
	uniqueMap := make(map[string]bool)
	uniqueList := make([]string, 0)

	for _, item := range input {
		if !uniqueMap[item] {
			uniqueMap[item] = true
			uniqueList = append(uniqueList, item)
		}
	}

	return uniqueList
}
