# Blockdaemon Technical Assessment

TODO: Provide the main requirement to run the application
For this demo purpose, it was tried to avoid third part library to implement the API service.
A custom multiplexer function was used to handle routes mapping them to the respective handlers.

## API: Transaction records

Sample and format:

```shell
   curl --header "Content-Type: application/json" \
-X POST http://<fqdn>api/models/ \ --data '{"transactionId":"0f7e46df-c685-4df9-9e23-e75e7ac8ba7a","amount": "99.99" ,"timestamp":"2009-09-28T19:03:12Z"}'
```

## API Service 

The API service is a `Golang` web application, using the version 1.22. More information about the service project
can be found [here](./api/README.md). The API will insert a new transaction in the following format:

```json
{
  "transactionId": "<UUID format>",
  "amount": "<Float format>",
  "timestamp": "<String format>"
}
```

