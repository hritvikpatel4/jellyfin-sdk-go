<p align="center">
<h1 align="center">Jellyfin Client SDK - Go</h1>
<p align="center">Auto-generated Go client package for <a href="https://jellyfin.org/">Jellyfin</a> using <a href="https://github.com/OpenAPITools/openapi-generator">OpenAPI Generator</a></p>
<p align="center">See the <a href="/api/README.md">API Documentation</a></p>
</p>
<p align="center">
<p align="center"><a href="https://github.com/hritvikpatel4/jellyfin-sdk-go/actions/workflows/update.yml?query=branch%3Amain"><img src="https://github.com/hritvikpatel4/jellyfin-sdk-go/actions/workflows/update.yml/badge.svg?branch=main" alt="Build Status"></a> <a href="https://goreportcard.com/report/github.com/hritvikpatel4/jellyfin-sdk-go"><img src="https://goreportcard.com/badge/hritvikpatel4/jellyfin-sdk-go" alt="Go Report Card"></a> <a href="https://github.com/hritvikpatel4/jellyfin-sdk-go/releases/latest"><img src="https://img.shields.io/badge/version-10.10.3-blue.svg" alt="Release Version"></a> <a href="https://pkg.go.dev/github.com/hritvikpatel4/jellyfin-sdk-go"><img src="https://pkg.go.dev/badge/github.com/hritvikpatel4/jellyfin-sdk-go" alt="GoDoc"></a> <a href="LICENSE"><img src="https://img.shields.io/github/license/hritvikpatel4/jellyfin-sdk-go.svg" alt="License"></a></p>
</p>

## Example

Generate an API token in the Jellyfin WebUI at Administration -> Overview -> Advanced -> API Token.

Minimal example to get started:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	jellyfin "github.com/hritvikpatel4/jellyfin-sdk-go/api"
)

func main() {
	config := &jellyfin.Configuration{
		Servers:       jellyfin.ServerConfigurations{{URL: "http://<ADDRESS>:<PORT>"}},
		DefaultHeader: map[string]string{"Authorization": `MediaBrowser Token="<API_TOKEN>"`},
	}
	client := jellyfin.NewAPIClient(config)

	result, resp, err := client.ActivityLogAPI.GetLogEntries(context.Background()).Execute()
	if err != nil {
		log.Printf("Error when calling `GetLogEntries``: %v\n", err)
		log.Printf("Full HTTP response: %v\n", resp)
		os.Exit(1)
	}

	for _, i := range result.Items {
		fmt.Println(*i.Name)
	}
}
```

## Development

The `configuration.go` file was adjusted (renamed the struct `ServerConfiguration` to `APIServerConfiguration`) and ignored from further generations to avoid conflicts with the same struct from the `model_server_configuration.go` file. In case the `configuration.go` needs to be regenerated, the struct needs to be renamed again.