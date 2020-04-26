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

## Fields
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
