package main

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

type InNetworkFile struct {
	Description string `json:"description"`
	Location    string `json:"location"`
}

func main() {
	inputFile := "anthem_Index_2024-07-01.json.gz"
	outputFile := "anthem_ny_ppo_urls.txt"

	startTime := time.Now()
	extractURLs(inputFile, outputFile)
	duration := time.Since(startTime)

	fmt.Printf("Total execution time: %v\n", duration)
}

func extractURLs(inputFile string, outputFile string) {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	gz, err := gzip.NewReader(file)
	if err != nil {
		fmt.Println("Error creating gzip reader:", err)
		return
	}
	defer gz.Close()

	urlSet := make(map[string]struct{})
	decoder := json.NewDecoder(gz)

	objectCount := 0
	startTime := time.Now()

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}

		if str, ok := token.(string); ok && str == "in_network_files" {
			count := parseInNetworkFiles(decoder, urlSet)
			objectCount += count

			if objectCount%10000 == 0 {
				fmt.Printf("Processed %d objects. Elapsed time: %v\n", objectCount, time.Since(startTime))
			}
		}
	}

	fmt.Printf("Total objects processed: %d. Total processing time: %v\n", objectCount, time.Since(startTime))
	writeURLsToFile(outputFile, urlSet)
}

func parseInNetworkFiles(decoder *json.Decoder, urlSet map[string]struct{}) int {
	_, err := decoder.Token() // Expect '['
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return 0
	}

	count := 0
	for decoder.More() {
		var file InNetworkFile
		err := decoder.Decode(&file)
		if err != nil {
			fmt.Println("Error decoding in_network_file:", err)
			return count
		}

		count++
		if isNYPPO(file.Location, file.Description) {
			urlSet[file.Location] = struct{}{}
		}
	}

	_, err = decoder.Token() // Expect ']'
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}

	return count
}

func isNYPPO(url string, description string) bool {
	var urlRegex = regexp.MustCompile(`https://([^.]+)\.mrf\.bcbs\.com/.*\.json\.gz`)
	matches := urlRegex.FindStringSubmatch(url)
	if matches == nil {
		return false
	}

	subdomain := matches[1]

	isNY := subdomain == "empirebcbs"
	isPPO := strings.Contains(strings.ToLower(description), "ppo")

	return isNY && isPPO
}

func writeURLsToFile(filename string, urlSet map[string]struct{}) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for url := range urlSet {
		writer.WriteString(url + "\n")
	}

	fmt.Printf("URLs extracted and saved successfully. Total unique URLs: %d\n", len(urlSet))
}
