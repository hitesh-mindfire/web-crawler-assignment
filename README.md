DeepScanBot(Web Crawler) Documentation

# Overview

This project is a customizable web crawler written in Go.

# Features

Customizable crawl depth: Set the maximum depth to crawl web pages.
Timeout management: Set a timeout for each HTTP request.
Proxy support: Specify a proxy server for the HTTP requests.
Output options: Choose between plain text or JSON output.
Page size limit: Skip pages exceeding a certain size.
Disable redirects: Option to disable HTTP redirects.
TLS verification: Option to disable TLS verification for HTTPS requests.
Unique URL tracking: Ensures URLs are crawled only once if enabled.
Show URL source: Display where each URL was found (e.g., in <a> tags, <script> tags).

# Usage

To run the web crawler, use the following command:

go mod download

go run main.go -url <starting_url> [options]
or
go build

# Flags

-url <string>: Required. The starting URL for the crawler.
-depth <int>: Maximum depth to crawl. Default is 2.
-timeout <int>: Timeout for each HTTP request in seconds. Default is 2.
-proxy <string>: Proxy URL for HTTP requests. Example: http://127.0.0.1:8080.
-json: Output results in JSON format. Default is false.
-size <int>: Limit page size in KB. Default is -1 (no limit).
-dr: Disable following HTTP redirects. Default is false.
-s: Show the source of the URL based on where it was found. Default is false.
-insecure: Disable TLS verification. Default is false.
-u: Ensure unique URLs are crawled. Default is false.
-h: Show help message.
