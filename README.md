# Volume Assignment

**Story:** There are over 100,000 flights a day, with millions of people and cargo being transferred around the world. With so many people and different carrier/agency groups, it can be hard to track where a person might be. In order to determine the flight path of a person, we must sort through all of their flight records.

**Goal:** To create a simple microservice API that can help us understand and track how a particular person's flight path may be queried. The API should accept a request that includes a list of flights, which are defined by a source and destination airport code. These flights may not be listed in order and will need to be sorted to find the total flight paths starting and ending airports.

## Usage

Download and install [Docker](https://www.docker.com)

```bash
docker build -t "volume-assignment" -f Dockerfile .
docker run -p "8080:8080" rollee
```

Alternatively, you can use `make` command if it's installed in your system

```bash
make build-docker
make run-docker
```

Or you can build and run from source, download and install [go](https://go.dev/dl/) **1.18+** version

```
go run cmd/server/main.go
```

To set different port than **8080**, set following env variable: **PORT**
To set server read timeout, set following env variable: **SERVER_READ_TIMEOUT** (in seconds)
To set server write timeout, set following env variable: **SERVER_WRITE_TIMEOUT** (in seconds)

## Tests

To run tests

```
go test -v ./...
```

## API

After running the service, it will listen on http 8080 port and you can query following endpoint:

### POST - `/calculate`

-------------------------------------------
Example body:
```
[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]
```

-------------------------------------------

**Success response**:

If request contains **valid** payload and there is start and end airports then endpoint will return status code **200** and following body:
```
["SFO", "EWR"]
```

**Bad request response**:

If request contains **invalid** payload then endpoint will return status code **400** and following body:
```
Bad Request
```

**Internal server error response**:

If service fails to find start and end airports or there is a failure in service internals then endpoint will return status code **500** and following body:
```
Internal Server Error
```
