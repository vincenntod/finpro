[![GO Language](https://travis-ci.org/DATA-DOG/go-sqlmock.svg)](https://travis-ci.org/DATA-DOG/go-sqlmock)
[![GoDoc](https://godoc.org/github.com/DATA-DOG/go-sqlmock?status.svg)](https://go.dev/doc/)

# API Dashboard Transactions BRI CERIA Using Golang

## About Project

**Revamp Ceria Repayment Service** is a project that will be used by participants to retrieve transaction data for customers who will make payments.
Ceria is a digital loan for financing transactions through e-commerce, online travel sites or ride sharing.
Ceria from Bank BRI has been registered and supervised by the Financial Services Authority.

## Requirements

- Install Go Version: 1.2.0 or higher
- Text Editor : **GoLand(_Recommended_)** or Visual Studio Code
- Platform API : **POSTMAN**
- Install **PosgreSQL** 15.3 or **MySQL** 8

## How Run

- Git Clone into your repository
- Create Database name `dbceria` in your local
- Run SQL Command

```
CREATE TABLE account (
    id integer NOT NULL,
    name character varying(200) NOT NULL,
    phone character varying(13) NOT NULL,
    role character varying(20) DEFAULT 'admin'::character varying NOT NULL,
    password character varying(200) NOT NULL,
    email character varying(100) NOT NULL
);

CREATE TABLE transaction (
    id integer NOT NULL,
    oda_number character varying,
    bank_account_no character varying,
    billing_cycle_date character varying,
    payment_due_date date,
    overflow_amount double precision,
    bill_amount double precision,
    principal_amount double precision,
    interest_amount double precision,
    total_fee_amount double precision,
    status character varying,
    payment_method character varying,
    auto_debet_counter integer,
    created_at date DEFAULT CURRENT_TIMESTAMP,
    updated_at date,
    payment_amount double precision,
    billing_gen_date date
);

```

- Insert dummy values into your database
- Open Project using your text editor
- run command

```
    go mod tidy
```

- Run application using code in command

```
    go run main.go
```

- Make sure your database connected in terminal

## List Endpoints

### Account

1. GET /api/data-user
2. GET /api/data-user/:id
3. PUT /api/data-user/:id
4. DELETE /api/data-user/:id
5. POST /send-email-registration/:email
6. POST /send-email/:email
7. POST /compare-verification-code
8. PUT /edit-password/
9. POST /login
10. POST /create-user

### Export CSV

1. GET /api/export-transaction

### Get Transaction

1. GET /api/get-transactions/
2. GET /api/get-transaction-status/:status/
3. GET /api/get-TransactionDate/:start/:end/
4. GET /api/get-TransactionStatusDate/:status/:start/:end/
5. GET /api/get-transactions-limit/:id

### Sample Request

Dont forget add Header Authorization with Token

```
//get transaction by status and date (start - end)
localhost:8080/api/get-TransactionStatusDate/SUCCESS/12-12-2022/14-12-2022/

//export CSV
localhost:8080/api/export-transaction?status=SUCCESS&start_date=12-12-2022&end_date=14-12-2022

//get user
localhost:8080/api/data-user/1

```

### Architecture and Design

this service using onion architecture, there are 5 layers
from inner to outer which are entity, repository, use case,
controller, and request handler. the usage and responsibility of
each layer are follow:

1. **Entity**: this layer contains the domain model or entities
   of the system. These are the core objects that
   represent the business concepts and rules.
2. **Repository**: This layer provides an interface for the
   application to access and manipulate the entities.
   It encapsulates the data access logic and provides
   a way to abstract the database implementation details.
3. **Use case** : This layer contains the business logic
   or use cases of the system. It defines the operations
   that can be performed on the entities and orchestrates
   the interactions between the entities and the repository layer.
4. **Controller**: This layer handles the HTTP requests and
   responses. It maps the incoming requests to the appropriate
   use case and returns the response to the client.
5. **Request handler**: This layer is responsible for handling
   the incoming HTTP requests and passing them on to
   the controller layer.

### Unit Test

1. Using Mockery and Gomock
2. Testify
3. SQL Mock

## License

**PT BANK RAKYAT INDONESIA**

[PT SATKOMINDO](https://www.satkomindo.com/)
