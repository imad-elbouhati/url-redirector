# **URL Redirector**

This is a simple URL redirector application written in Go. It allows you to define custom path-to-URL mappings and redirect incoming requests accordingly.

## How it works
The application uses a YAML file to define the path-to-URL mappings. When a request comes in, the application checks if the requested path matches any of the mappings. If a match is found, the client is redirected to the corresponding URL. If no match is found, a fallback handler is called.

## Prerequisites
Before running the application, make sure you have the following dependencies installed:

* Go programming language
* `gopkg.in/yaml.v2` package

## Installation
To install the application, follow these steps:

1. Clone the repository: `https://github.com/imad-elbouhati/url-redirector.git`
2. Change to the project directory: `cd url-redirector`
3. Install the necessary dependencies: `go get gopkg.in/yaml.v2`

## Usage

To use the URL redirector, follow these steps:

1. Create a YAML file (yaml-data.yaml) and define the path-to-URL mappings as shown in the example below:

```
- path: /path1
  url: https://example.com/destination1
- path: /path2
  url: https://example.com/destination2
```

2. Run the application: `go run main/main.go`

3. Open your web browser and navigate to http://localhost:8080/path1 or http://localhost:8080/path2. You should be redirected to the corresponding URLs specified in the YAML file.
