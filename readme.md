# Bank MNC API

## The functionality of the API includes :

- Login: do log in customer and if customer does not exist, then got reject.

- Payment: customer that has been login can do the payment. No max & min limit amount for transfer Transfer only do for registered customer.

- Logout: do logout for the logged in custome.

## How to run

```sh
cd api-bank
go run main.go
```

**API Endpoint**

```sh
POST /login
```

```sh
POST /payment
```

```sh
POST /logout
```

**For login example you can use :**

```sh
username : user1
password : secret
```

OR

```sh
username : user4
password : password
```

**For payment you must input :**

Authorization Headers :

```sh
Authorization : jwt_token
```

Request body form-data

```sh
to_account : username
amount : transfer_amount
merchant : merchant_name
```

**For logout you must input :**

Authorization Headers :

```sh
Authorization : jwt_token
```

## API documentation :

https://documenter.getpostman.com/view/27682343/2s93z88iWx
