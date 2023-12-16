<div align="center">
    <img src="https://www.leisurebyte.com/wp-content/uploads/2022/11/alchemy-of-souls-season-2-1.jpg"/>
</div>

## Usage
running this command to run service
```
docker-compose up -d db
```
and then running this command to check database can connection
```
docker logs -f db
```
last running this command to running app
```
docker-compose up -d app
```

## List Service

| Feature             | Status              | 
|---------------------|---------------------|
| Authentication      | :heavy_check_mark:  |
| Event               | :heavy_check_mark:  |
| Payment             | Ongoing             |
| User                | :heavy_check_mark:  |
| Dockerize           | :heavy_check_mark:  |
| Frontend            | Ongoing             |
| Order               | Ongoing             |
| Monitoring          | Ongoing             |
| File Handler        | Ongoing             |
| Webhook             | Ongoing             |

## API Documentation

#### Create User - All User
- Method : POST
- Endpoint : `/user/register/`
- Header :
    - Content-Type : `application/json`
    - Accept: `application/json`
- Body :
```json
{
    "first_name": "string",
    "last_name": "string",
    "email": "string",
    "password": "string",
    "address": "long",
    "telephone": "string",
    "role_id": "uint"
}
```
- Response :
```json
{
    "code": "number",
    "message": "string",
    "data": {
        "first_name": "string",
        "last_name": "string",
        "email": "string",
        "password": "string",
        "address": "long",
        "telephone": "string",
        "role": "string"
    }
}
```
#### Login User - All User
- Method : POST
- Endpoint : `/user/login/`
- Header :
    - Content-Type : `application/json`
    - Accept: `application/json`
- Body :
```json
{
    "email": "string",
    "password": "string",
}
```
- Response :
```json
{
    "code": "number",
    "message": "string",
    "data": {
        "first_name": "string",
        "last_name": "string",
        "email": "string",
        "token": "string"
    }
}
```
#### Validate Token
- Method : POST
- Endpoint : `/user/validate-token`
- Header :
    - Content-Type : `header`
- Body :
```http
Set-Cookie: "token"
```
- Response :
```json
{
    "code": "number",
    "message": "string"
}
```

## Authentication
> Semua api wajib di tambahkan authentikasi dari hasil login user
- Header :
    - Authorization: Token JWT

#### FindByID Event - All User
- Method : GET
- Endpoint : `/api/event/{id}`
- Header :
    - Accept: `application/json`
- Response :
```json
{
    "code": "number",
    "message": "string",
    "data": {
        "name": "string",
        "desc": "string",
        "price": "int",
        "participant": "array"
    }
}
```
#### Create Order - All User
- Method : POST
- Endpoint : `/api/order/`
- Header :
    - Accept : `application/json`
    - Content-Type : `application/json`
```json
{
    "user_id": "int",
    "event_id": "int",
    "delegasi": "array ( optional )"
}

```
- Response :
```json
{
    "code": "number",
    "message": "string",
    "data": {
        "uuid": "string",
        "first_name": "string",
        "last_name": "string",
        "email": "string",
        "address": "string",
        "telephone": "string",
        "event_name": "string",
        "event_price": "int"
    }
}
```
#### Checkout Order - All User
- Method : POST
- Endpoint : `/api/order/checkout/`
- Header :
    - Accept : `application/json`
    - Content-Type : `application/json`
```json
{
    "order_id": "int",
}

```
- Response :
```json
{
    "code": "number",
    "message": "string",
    "data": {
        "snap_url": "string",
        "token": "string",
        "name": "string",
        "event": "string",
        "price": "int"
    }
}
```

#### Create Submission Task - All User
- Method : POST
- Endpoint : `/api/event/{id}/submission`
- Header :
    - Accept : `application/json`
    - Content-Type : `application/json`
```json
{
    "name": "string",
    "user_id": "int",
    "event_id": "int",
    "file": "string"
}

```

- Response :
```json
{
    "code": "number",
    "message": "string",
    "data": {
        "name": "string",
        "event": "string",
        "filename": "string"
    }
}
```

#### Create Event - Admin
- Method : POST
- Endpoint : `/api/event/`
- Header :
    - Content-Type : `application/json`
    - Accept: `application/json`
- Body :
```json
{
    "name": "string",
    "desc": "string",
    "price": "int"
}
```
- Response :
```json
{
    "code": "number",
    "message": "string",
    "data": {
        "name": "string",
        "desc": "string",
        "price": "int",
        "participant": "array"
    }
}
```

#### Update Event - Admin
- Method : PUT
- Endpoint : `/api/event/{id}`
- Header :
    - Content-Type : `application/json`
    - Accept: `application/json`
- Body :
```json
{
    "name": "string",
    "desc": "string",
    "price": "int"
}
```
- Response :
```json
{
    "code": "number",
    "message": "string",
    "data": {
        "name": "string",
        "desc": "string",
        "price": "int",
        "participant": "array"
    }
}
```
#### FindByID Role - Admin
- Method : GET
- Endpoint : `/api/role/{id}`
- Header :
    - Content-Type : `application/json`

- Response :
```json
{
    "code": "number",
    "message": "string",
    "data": {
        "name": "string",
        "description": "text",
        "user": "array",
    }
}
```
#### Update Role - Admin
- Method : PUT
- Endpoint : `/api/role/{id}`
- Header :
    - Content-Type : `application/json`
    - Accept: `application/json`
- Body :
```json
{
    "name": "string",
    "description": "text"
}
```
- Response :
```json
{
    "code": "number",
    "message": "string",
    "data": {
        "name": "string",
        "description": "text",
        "user": "array",
    }
}
```