<h1 align="center">MalWeb</h1>
<p align="center"><b>Scan for malicious URLs or domains present on a web page.
</b></p>
<p align="center">
<a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/license-MIT-_red.svg"></a>
<a href="https://github.com/umair9747/malweb/issues"><img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat"></a>
<a href="https://github.com/umair9747/malweb/releases"><img src="https://img.shields.io/github/release/umair9747/malweb"></a>
<a href="https://twitter.com/0x9747"><img src="https://img.shields.io/twitter/follow/0x9747.svg?logo=twitter"></a>
</p>

<p align="center">
  <a href="#description">Description</a> •
  <a href="#setup">Setup</a> •
  <a href="#usage">Usage</a> •
  <a href="#contact">Contact</a> </p>


## Description
MalWeb is a go-lang based tool to quickly check if your target URLs or their response bodies contain malicious URLs, Domains, and IP Addresses by scanning it against the <a href="https://urlhaus.abuse.ch/">URLhaus</a> database dumps. <br>
This can be quite helpful in cases such as looking for compromised or malicious CDNs within your infrastructure, looking for URLs that are associated with malicious campaigns and many more!

![Tool Sample Output](<toolOP.png>)


## Setup

Since the tool is written in Golang, it is require for you to have it <a href="https://go.dev/doc/install">installed</a> on their systems to build the binaries.

After making sure that you have successfully installed Golang on your system, you can run the below command and get set going.

 ```
 go build
 ```

This will get all the required dependencies/packages and build a binary based on your environment.

## Usage

Using MalWeb is fairly easy and straightforward.

<h4>Basic Scan:</h4>

```
./malweb <target>

Example: 
./malweb https://pastebin.com/raw/something
```

Using the above approach you can quickly see if the URL that you have specified or any URLs within its response body are present in the malicious entries DB of URLhaus.

<h4>Depth Mode:</h4>

```
./malweb -depth <target>

Example: 
./malweb -depth https://pastebin.com/raw/something
```

Using depth mode you can go beyond scanning just URLs and see if the root domains or IP addresses have any presence within the malicious entries DB.
Example would be to check for all malicious entries of facebook.com when the input URL is https://test.facebook.com/something .

<h4>Save Flag:</h4>

```
./malweb -save <target>
./malweb -save test.json <target>
Example:
./malweb -save https://pastebin.com/raw/something
```

This flag can be used for saving the JSON output that MalWeb generates in an output file. You can provide the filename as an argument, but if kept empty, it will save to output.json by default.

The tool generates the output in the following format:
```
type Result struct {
	Hit       string `json:"Hit"`
	Source    string `json:"Source"`
	DepthMode bool   `json:"DepthMode"`
}

type TargetResult struct {
	Target  string   `json:"Target"`
	Results []Result `json:"Results"`
}
```

An example would be:

```
[{"Target":"https://pastebin.com/X7W7gcgz","Results":[{"Hit":"http://185.150.26.225/Kukri.mpsl","Source":"URLhaus","DepthMode":false},{"Hit":"http://185.150.26.225/Kukri.arm7","Source":"URLhaus","DepthMode":false},{"Hit":"https://sahmnx.mynetav.org/saham.apk","Source":"URLhaus","DepthMode":false},{"Hit":"https://sahlmnh.vizvaz.com/saham.apk","Source":"URLhaus","DepthMode":false},{"Hit":"https://sadldh.mrface.com/saham.apk","Source":"URLhaus","DepthMode":false},{"Hit":"https://adl-gh.fartit.com/saham.apk","Source":"URLhaus","DepthMode":false},{"Hit":"https://saldhg.my03.com/saham.apk","Source":"URLhaus","DepthMode":false}]}]
```

## Contact

You can reach me out through the following channels, <br>
 • LinkedIn: https://www.linkedin.com/in/umair-nehri-49699317a/<br> 
 • Twitter: https://twitter.com/0x9747<br>
 • Email: umairnehri9747@gmail.com