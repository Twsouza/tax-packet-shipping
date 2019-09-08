# Tax Packet Shipping

Tax Packet Shipping is an API to calculate and track shipping tax from a packet/prouct.

## Installation

Tax Packet Shipping requires [Golang](https://golang.org/) 1.11+ to run.

We use Go Modules so all you got do is.

```
$ go install
```

It'll get all dependecies and install tha API, than you just have to run it.

### Config File

There are some configurations that depends of file called `config.json`. After build the api, you have to as param to the executable e.g `./tax-packet-shipping config.json`.
We already provided an example in the source code, just rename it as yours configurations.

### For production environments...

Tax Packet Shipping is ready for [Docker](https://docs.docker.com/get-started/) installed. If it's already installed, run in the app folder

```
$ docker-compose up -d --build
```

### Dependecies used

Tax Packet Shipping uses a fewer dependecies, here are what we used it.

| Dependecie |
| ------ |
| [Gorilla Mux](https://github.com/gorilla/mux/) |
| [MongoDB Driver](https://godoc.org/go.mongodb.org/mongo-driver/mongo) |
