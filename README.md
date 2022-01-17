# Golang microservices assignment

## Building a Payment Gateway

### Background
E-Commerce is experiencing exponential growth and merchants who sell their goods or
services online need a way to easily collect money from their customers.
We would like to build a payment gateway, an API based application that will allow a
merchant to offer a way for their shoppers to pay for their product. You can use the software
language of your choice. Source code with no binary must be provided.

### Requirement
The software that you will develop will allow a merchant to process payments through a REST API.

It should simulate the following payment flow:

- A merchant requests an authorisation through the `/authorize` call. This call contains
the customer credit card data as well as an amount and currency. It will return a
unique ID that will be used in all next API calls.
- The `/void` call will cancel the whole transaction without billing the customer. No
further action is possible once a transaction is voided.
- The `/capture` call will capture the money on the customer bank. It can be called
multiple times with an amount that should not be superior to the amount authorised in
the first call. For example, a 10£ authorisation can be captured 2 times with a 4£ and
6£ order.
- The `/refund` call will refund the money taken from the customer bank account. It can
be also called multiple times with an amount that can not be superior to the captured
amount. For example, a 10£ authorisation, with a 5£ capture, can only be refunded of
5£. Once a refund has occurred, a capture cannot be made on the specific
transaction.

In order to test specific edge case, your application must trigger errors based on the credit
card number:

- 4000 0000 0000 0119: authorisation failure
- 4000 0000 0000 0259: capture failure
- 4000 0000 0000 3238: refund failure

Your application should handle multiple transactions at the same time, at all different
transaction stages.Endpoints

/authorize

- Input
  - Credit card data
  - Card number
  - Expiry month and date
  - CVV
  - Amount and currency
- Output
  - Unique ID
  - Success or error
  - Amount and currency available

/capture

- Input
  - Authorization unique ID
  - Amount
- Output
  - Success or error
  - Amount and currency available

/void

- Input
  - Authorization unique ID
- Output
  - Success or error
  - Amount and currency available

/refund
- Input
  - Authorization unique ID
  - Amount
- Output
  - Success or error
  - Amount and currency available

### Consideration
Include whatever documentation/notes you feel is appropriate. This should include some
details of assumptions made, areas you feel could be improved.

### Extra mile bonus points
In addition to the above, time permitting, consider the following suggestions for taking your
implementation a step further:
- Implement Luhn check on credit card number
- Client (merchant) authentication
- Application logging
- Containerization
- Anything else you feel may benefit your solution from a technical perspective

## RUN with podman

```sh
# lint
~/go/bin/golangci-lint run ./...

# use podman

# build image
podman build -t paygw .

# run container for all in one
podman run --rm -d --name paygw -p 8080:8080 -p 8090:8090 paygw
# OR
# run container one per service transport
podman run --rm -d --name paygw-grpc -e PAYGW_SERVERS=grpc -p 8090:8090 paygw
podman run --rm -d --name paygw-restgw -e GRPC_ADDRESS=$(hostname -I | grep -o "^[0-9.]*"):8090 -e PAYGW_SERVERS=restgw -p 8080:8080 paygw

# test with rest ui clients 
# test/http-client/rest.http

# check and close containers
podman container ps
podman stop paygw
podman stop paygw-restgw
podman stop paygw-grpc

```

## RUN with docker

```sh
# lint
~/go/bin/golangci-lint run ./...

# use docker

# build image
docker build -t paygw .

# run container for all in one
docker run --rm -d --name paygw -p 8080:8080 -p 8090:8090 paygw
# OR
# run container one per service transport
docker run --rm -d --name paygw-grpc -e PAYGW_SERVERS=grpc -p 8090:8090 paygw
docker run --rm -d --name paygw-restgw -e GRPC_ADDRESS=$(hostname -I | grep -o "^[0-9.]*"):8090 -e PAYGW_SERVERS=restgw -p 8080:8080 paygw
# run container one per service transport on mac
docker run --rm -d --name paygw-grpc -e PAYGW_SERVERS=grpc -p 8090:8090 paygw
docker run --rm -d --name paygw-restgw -e GRPC_ADDRESS=$(hostname):8090 -e PAYGW_SERVERS=restgw -p 8080:8080 paygw

# test with rest ui clients 
# test/http-client/rest.http

# check and close containers
docker container ps
docker stop paygw
docker stop paygw-restgw
docker stop paygw-grpc

```

## RUN with makefile

```sh
# everything at once
make podman-scratch
# OR
make podman-separate-scratch
# OR
make docker-scratch

# sequence with podman
make podman-build
make podman-run
make curl-test
make podman-stop

# sequence with podman-separate
make podman-separate-build
make podman-separate-run
make curl-test
make podman-separate-stop


# sequence with docker
make docker-build
make docker-run
make curl-test
make docker-stop

```
