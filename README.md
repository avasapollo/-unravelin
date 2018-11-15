# unravelin

I used `dep` for the dependencies management, you can find the doc about that here `https://github.com/golang/dep`.

The make file run the command `dep ensure` to manage the dependencies of the project.

The command `make all` inside the project will run:
- `dep ensure` to import the dependencies 
- `go test` to run the tests inside the packages
- `go build` to create the executable `unravelin-api`


The server will expose the endpoints on `:8080` port.

Api
method: `post`
url: `http://localhost:8080/v1/form`
content-type: `application/json` (is required) 
 
```
curl -X POST \
  http://localhost:8080/v1/form \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'postman-token: be747a7a-ad88-c438-8b7e-7ddcf763e5a0' \
  -d '{
	  "eventType": "copyAndPaste",
  "websiteUrl": "https://ravelin.com",
  "sessionId": "123123-123123-123123123",
  "pasted": true,
  "formId": "inputCardNumber"
}'

```
