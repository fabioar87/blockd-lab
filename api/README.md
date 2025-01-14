## Transaction Service

The _transaction service API_ is a Golang application. The main target is the implementation
of an API with the following features:

* `/api/health`: shall get the API health figures. Important metric to define the K8S probes correctly.
* `/api/transaction`: this is a POST route to add a new transaction record.
* `/api/transactions`: this is a GET route that returns all the transactions stored in the DB repository.

## 3rd Party Libraries

My initial idea, was to build an API using only built-in Golang libraries, custom mux and the native Golang server.
However, the complexity was increasing impacting the final project delivery. This is also true regarding the data layer
implementation. As a POC, I changed the strategy, adding to the project the following opensource libraries:

*  Gorilla Mux: More info can be found [here](https://github.com/gorilla/mux). The criteria of choice was its popularity, flexibility
       and a feature-rich HTTP request router and URL matcher for building REST APIs and web application. It is part of 
       the Gorilla Web Toolkit and provides enhanced routing capabilities over Go's standard net/http package.
*  GORM ORM: More info can be found [here](https://gorm.io/docs/). The criteria of choice was the capacity of fast
    prototyping the data layer abstraction. The rich driver vendors ecosystem, was easy to find a MySQL (DB Engine used
    at the backend) driver and worked very well. 