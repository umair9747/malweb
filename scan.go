package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func scanTargets() {
	var resultArray []TargetResult
	for _, target := range targets {
		var targetResults []Result
		var extURLs []string
		log.Println("Scanning:", target)

		//PHASE 1 - 200 ONLY
		log.Println("[Phase 1] Making a request on target URL/domain...")
		targetResp := makeReq(target)
		if targetResp == "" {
			log.Println("Target returned no response body... [Skipping all other phases.]")
		} else {
			log.Println("[Phase 2] Extracting URL(s) from response body...")
			extURLs = extractURLs(targetResp)
		}
		extURLs = append(extURLs, target)
		extURLs = uniqueList(extURLs)
		log.Println("Discovered", len(extURLs), "URL(s) related to the target.")

		log.Println("[Phase 3] Matching the URL(s) against URLhaus DB...")
		checkURLhaus(extURLs)
		if len(malURLs) > 0 {
			log.Println("Found", len(malURLs), "malicious URL(s) from the list!")
			for _, malURL := range malURLs {
				res := Result{Hit: malURL, Source: "URLhaus", DepthMode: false}
				targetResults = append(targetResults, res)
			}
		}

		if depthFlag {
			log.Println("[Phase 4 - Depth Mode] Matching the domain(s)/IP Address(es) against URLhaus DB...")
			checkURLhausDepthMode(extURLs)
			if len(malDepthURLs) > 0 {
				log.Println("Found", len(malDepthURLs), "malicious domain(s)/IP Address(es) from the list!")
				for _, malURL := range malDepthURLs {
					res := Result{Hit: malURL, Source: "URLhaus", DepthMode: true}
					targetResults = append(targetResults, res)
				}
			}
		}

		targetResult := TargetResult{
			Target:  target,
			Results: targetResults,
		}
		resultArray = append(resultArray, targetResult)
		extURLs = nil
		malURLs = nil
		fmt.Println("")
	}
	resultJSON, err := json.Marshal(resultArray)
	if err != nil {
		log.Println("Error marshalling to JSON:", err)
		return
	}

	if saveFlag != "" {
		err := ioutil.WriteFile(saveFlag, resultJSON, 0644)
		if err != nil {
			fmt.Println("[Error]", err)
		}

		log.Println("Output saved to", saveFlag)
	}

	// Print or use the JSON as needed
	fmt.Println("\n\n" + string(resultJSON))
}

func checkURLhaus(extURLs []string) {
	for _, extURL := range extURLs {
		for _, urlhausURL := range urlhausURLs {
			if extURL == urlhausURL {
				malURLs = append(malURLs, extURL)
			}
		}
	}
}

func checkURLhausDepthMode(extURLs []string) {
	var targets []string
	for _, extURL := range extURLs {
		re := regexp.MustCompile(`(?m)(\bhttps?|ftp):\/\/([^\/:\s]+)(?::(\d+))?\b`)
		matches := re.FindAllStringSubmatch(extURL, -1)
		if len(matches[0]) > 2 {
			targets = append(targets, matches[0][2])
		}
	}
	targets = uniqueList(targets)
	for _, extTarget := range targets {
		for _, urlhausURL := range urlhausURLs {
			if strings.Contains(urlhausURL, extTarget) {
				malDepthURLs = append(malDepthURLs, extTarget)
			}
		}
	}
	malDepthURLs = uniqueList(malDepthURLs)
}
