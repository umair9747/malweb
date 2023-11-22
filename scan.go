package main

import (
	"log"
)

func scanTargets() {
	for _, target := range targets {
		log.Println("Scanning:", target)

		//PHASE 1 - 200 ONLY
		log.Println("[Phase 1] Making a request on target URL/domain...")
		targetResp := makeReq(target)
		if targetResp == "" {
			log.Println("Target returned no response body... [Skipping all other phases.]")
			continue
		} else {
			log.Println("[Phase 2] Extracting URLs from response body...")
			extURLs := extractURLs(targetResp)
			log.Println("Discovered", len(extURLs), "URLs from the response body.")
			log.Println("[Phase 3] Matching the URLs against URLhaus...")
			checkURLhaus(extURLs)
			if len(malURLs) > 0 {
				log.Println("Found", len(malURLs), "malicious URLs from the list!")
			}
		}
	}
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
