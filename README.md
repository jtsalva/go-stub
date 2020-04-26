# go-stub
Declarative bare-bones stub server. Inspired by [stubby4j](https://github.com/azagniotov/stubby4j)

## Key features
* Split stubs across multiple files in one directory
* Supports YAML anchors and aliases for reuse of common stub attributes
* Easily enable CORS for a given origin via a command line flag
* Regex path matching via [gorilla mux](https://github.com/gorilla/mux#matching-routes)
* Regex query and headers matching

## Basic example
```yaml
# stubs-folder/example.stub.yml
- request:
    methods: [GET]
    path: /api/user
  response:
    status: 200
    headers:
      Content-Type: application/json
    file: "path/to/user.json"

- request:
    methods: [PUT, PATCH]
    path: /api/user
  response:
    status: 204
    latency: 100
    
- request:
    methods: [GET, PUT, POST, PATCH, DELETE]
    path-prefix: /api/admin
  response:
    status: 403
    body: forbidden
```

## Getting started
1. Download an executable from the [releases page](https://github.com/jtsalva/go-stub/releases). Alternatively, clone this repo and [build from source](#build-from-source).
2. Simplify the name of the binary e.g. `go-stub`
3. Run `./go-stub -p 8080 -d "path/to/stubs-folder"`

## YAML stub fields
| Field                 | Description                                                                           |             Example             |                  Required                 |
|-----------------------|---------------------------------------------------------------------------------------|:-------------------------------:|:-----------------------------------------:|
| `request`             | declare what request to respond to                                                    | [basic example](#basic-example) |                  required                 |
| `request.methods`     | list of HTTP methods                                                                  |         [GET, PUT, POST]        |                  required                 |
| `request.path-prefix` | path prefix to match. Overriden by `request.path`                                     |                /                | required if missing `request.path-prefix` |
| `request.path`        | specific path to match. Overwrites `request.path-prefix`                              |          /path/to/stub          |     required if missing `request.path`    |
| `request.query`       | key-value map of query string parameters to match. Values support regex.              |           animal: dog           |                  optional                 |
| `request.headers`     | key-value map of headers to match. Values support regex.                              |      Connection: keep-alive     |                  optional                 |
|                                                                                                                                                                                             |
|                                                                                                                                                                                             |
|                                                                                                                                                                                             |
| `response`            | declare the contents of the response                                                  | [basic example](#basic-example) |                  required                 |
| `response.status`     | status code to respond with                                                           |               200               |                  required                 |
| `response.headers`    | key-value map of headers to be included in the response                               |  Content-Type: application/json |                  optional                 |
| `response.body`       | value of response body. Overriden by `response.file`                                  |           hello world           |                  optional                 |
| `response.file`       | path to file containing the contents of the response body. Overwrites `response.body` |        /path/to/data.json       |                  optional                 |
| `response.latency`    | duration in milliseconds to wait before responding                                    |               250               |                  optional                 |

## Command line flags
| Long                | Short | Description                                                                                                                                                               |                 Usage                 | Required |
|---------------------|-------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:-------------------------------------:|----------|
| `directory`         | `d`   | Path to directory containing YAML stub files                                                                                                                              | `./go-stub -d "path/to/stubs-folder"` | required |
| `port`              | `p`   | PortPort to serve stubs from                                                                                                                                              |          `./go-stub -p 8080`          | required |
| `cors-allow-origin` | `o`   | Allow CORS from specified origin, this will,automatically apply headers 'Access-Control-Allow-Origin', 'Access-Control-Allow-Methods', and'Access-Control-Expose-Headers' |           `./go-stub -o "*"`          | optional |
| `write-timeout`     | N/A   | Server write timeout duration (default: 15s)                                                                                                                              |    `./go-stub --write-timeout 30s`    | optional |
| `read-timeout`      | N/A   | Server read timeout duration (default: 15s)                                                                                                                               |     `./go-stub --read-timeout 1m`     | optional |
| `idle-timeout`      | N/A   | Server idle timeout duration (default: 1m0s)                                                                                                                              |    `./go-stub --idle-timeout 1m30s`   | optional |
| `disable-color`     | N/A   | Disable color in console output                                                                                                                                           |      `./go-stub --disable-color`      | optional |

## Build from source
1. Have a working [go installation](https://golang.org/doc/install)
2. Clone this repo `git clone github.com/jtsalva/go-stub`
3. Go into the repo folder and run `./build.sh` to build for the most common platforms
4. Find the binaries in the bin folder, named like so `{os}-{architecture}-go-stub`

If you only want to build for your current platform, just run `go build -o go-stub *.go`