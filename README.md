# Cloutility API Client

Cloutility API Client is a GoLang project aimed at providing a client library for the Auwau Cloutility API. It also includes a Command Line Interface (CLI) client based on the Cobra library.

## Installation

To install Cloutility API Client, you can use the `go get` command:

```
go get github.com/safespring/cloutility-api-client
```

## Usage

### Using the Library

To use the Cloutility API Client library in your GoLang project, you can import the library using:

```go
go get "github.com/safespring/cloutility-api-client"
```

To query the API you need an authenticated client which is instansiated using:

```go
	client, err := cloutapi.Init(
		context.Background(),
		client_id,
		client_origin,
		username,
		password,
		baseurl,
	)
    if err != nil {
        // handle error
    }
```

Where `client_id` is your Cloutility API key and `client_origin` is the url specified when creating the API key. `username` and `password` is the same credentials you would use to login to the cloutility portal. `baseurl` is the 

You can then use the client to access the various endpoints of the API. For example, to get a list of all the consumers / consumption-units:

```go
	consumers, err := client.GetConsumers(business-unit-id)
	if err != nil {
	// handle error
	}

	for _, consumer := range consumers {
        fmt.Println(consumer.ID, consumer.Name)
	}
```

### Using the CLI

The Cloutility API Client also includes a CLI client based on the Cobra library. To use the CLI, you can run the binary with the appropriate flags. For example, to get a list of all the consumers / consumption-units in the default business-unit:

```
cloutility-api-client consumers list 
```

The cli client will look for a configuration file in the same directory as the binary namned `cloutility-api-client.yaml`. There is an example configuration file included in the repository which may be used as base.

For a full list of commands and flags, you can run:

```
cloutility-api-client --help
```

## Contributing

Contributions to the Cloutility API Client project are welcome. To contribute, please fork the repository, make your changes, and submit a pull request.

## License

Cloutility API Client is released under the BSD License. See `LICENSE` for more information.
