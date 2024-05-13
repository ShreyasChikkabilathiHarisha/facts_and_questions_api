# facts_and_questions_api

Steps to run

1. Initialize go mod from the repository root
`go mod init facts_and_questions_api`

2. Run the service from the repository root
`go run main.go`

3. To make the HTTP request, curl with the appropriate params or use an http client like postman to send the requests to `http://127.0.0.1:8383`


Some example http requests are as follows:

- `curl -X GET -d "{\"question\": \"q4\", \"answer\": \"a4\"}" http://localhost:8383/create`
- `curl -X GET -d "{\"question\": \"q1\"}" http://localhost:8383/fetch`
- `curl -X GET -d "{}" http://localhost:8383/fetch`