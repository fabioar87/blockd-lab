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
