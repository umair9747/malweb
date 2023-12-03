package main

const (
	banner = `
	__  __       ___        __   _     
	|  \/  | __ _| \ \      / /__| |__  
	| |\/| |/ _' | |\ \ /\ / / _ \ '_ \ 
	| |  | | (_| | | \ V  V /  __/ |_) |
	|_|  |_|\__,_|_|  \_/\_/ \___|_.__/ 
										
			- Developed by Umair Nehri (https://twitter.com/0x9747)
	`
)

var urlhausURLs []string

var cliArgs []string
var targets []string

var malURLs []string
var malDepthURLs []string

var saveFlag string
var depthFlag bool

type Result struct {
	Hit       string `json:"Hit"`
	Source    string `json:"Source"`
	DepthMode bool   `json:"DepthMode"`
}

type TargetResult struct {
	Target  string   `json:"Target"`
	Results []Result `json:"Results"`
}
