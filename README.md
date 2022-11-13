<p align="center"></p>

## About Api

This microservice is used to manage companies records. And provided endpoints to add, update, get and delete companies.

In this api JWT token is used for authenticate user, so user needs to be registered first via user registration api.

Healthcheck endpoint is provided to check the health of application before making request.


## Api resources:

Api Documentation : https://documenter.getpostman.com/view/8736410/2s8YeuKWHN

Postman Collection : https://www.getpostman.com/collections/5630d2a75a18f89c3288


## Run application:

```
make run
```

## Run tests:

```
make test
```

## Step to test:

+ Healthcheck api
+ Register user
+ Get authentication token
+ Add company information
+ Get company detail
+ Update company information
+ Delete company
+ Run tests