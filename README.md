# EnvServer

EnvServer is a lightweight web server built with Go that provides two endpoints to access environment variables. By utilizing Docker Compose for deployment, setting up and running the EnvServer becomes a breeze.

## Installation

Before proceeding, ensure you have the following installed on your system:
- Docker
- Docker Compose

To quickly set up and run the EnvServer, follow these steps:

1. Clone the repository:
```bash
git clone <repository-url>
cd EnvServer
```

2. Start the EnvServer using Docker Compose:
```bash
docker-compose up -d
```

## Endpoints

The EnvServer provides two endpoints to interact with environment variables:

1. `/env`

This endpoint returns plain text containing all the environment variables available on the server.

Example Request:
```
curl http://localhost:8080/env
```
Example Response:
```
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
HOME=/root
GO_VERSION=go1.16.5
...
```

2. `/env/<key>`

This endpoint returns the value of a specific environment variable based on the provided key.

Example Request:
```
curl http://localhost:8080/env/PATH
```
Example Response:
```
/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
```

If the provided key does not exist, the response will indicate that the key is not found.

## Test

To run the automated tests for this project, follow these steps:

1. Install the necessary dependencies by running `go get -d ./...`.
2. Run the tests by running `go test ./...`.
3. If all tests pass, the output should indicate that the tests have passed. If any tests fail, the output will provide information on which tests failed.
