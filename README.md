<p align="center"></p>

## About Api

This microservice is used to manage companies records and provides endpoints to add, update, get and delete companies.

In this api JWT token is used for authenticate user, so user needs to be registered first via user registration api.

Healthcheck endpoint is provided to check the health of application before making request.


## Api resources:

Api Documentation : https://documenter.getpostman.com/view/8736410/2s8YeuKWHN

Postman Collection : https://www.getpostman.com/collections/5630d2a75a18f89c3288

## Prerequisites:

The application requires docker to be installed before running this application.

## Run application:

Clone the application source code on your local machine and run below command at root of the application. This command
will build and run microservice in docker container.

```
sudo docker-compose up
```

After successfully running application you can check health of the microservice at http://localhost:5000/v1/healthcheck

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
