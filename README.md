# MultiOpQueryerFileUpload

`go run main.go` 

Leverages the example [serviceUpload](https://github.com/nautilus/gateway/tree/master/examples/fileupload) starting it in the background (port 5000).  
Introspects and then starts the gateway using a Multi Op Queryer (port 4040).

```bash 
curl localhost:4040/graphql \
  -F operations='{ "query": "mutation ($someFile: Upload!) { upload(file: $someFile) }", "variables": { "someFile": null } }' \
  -F map='{ "0": ["variables.someFile"] }' \
  -F 0=@README.md
```

Resulting error when using the multiOpQueryer is: 
```json
{"data":{},"errors":[{"extensions":null,"message":"no file found in request","path":["upload"]}]}
```
