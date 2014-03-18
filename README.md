# go-harvest

go-harvest is a Go client library for accessing the Harvest Time Tracking API

## Disclaimer

As of 3/18/14 this library is still under active development and missing things that would make it useful on it's own.

## Usage

```go
import "github.com/backforty/go-harvest/harvest"
```

Construct a new Harvest client, then use the client to access the Harvest API.

```go
apiClient := harvest.NewAPIClientWithBasicAuth("YOUR_USERNAME", "YOUR_PASSWORD", "YOUR_SUBDOMAIN")
clients := apiClient.GetClients()
```

## License
This library is distributed under the MIT-style license found in the [LICENSE](./LICENSE) file.
