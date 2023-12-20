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
| Payment             | :heavy_check_mark:  |
| User                | :heavy_check_mark:  |
| Dockerize           | :heavy_check_mark:  |
| Frontend            | Ongoing             |
| Order               | Ongoing             |
| Monitoring          | Ongoing             |
| File Handler        | Ongoing             |
| Webhook             | :heavy_check_mark:  |

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
    "name": "string",
    "email": "string",
    "password": "string",
    "address": "long",
    "telephone": "string"
}
```
- Response :
```json
{
    "code": "number",
    "message": "string",
    "data": {
        "name": "string",
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
        "name": "string",
        "email": "string",
        "token": "string"
    }
}
```
#### Refresh Token - All User
- Method : POST
- Endpoint : `/user/refresh-token/`
- Header :
    - Content-Type : `application/json`
    - Accept: `application/json`
- Body :
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

#### Validate Token - All User
- Method : POST
- Endpoint : `/user/validate-token`
- Header :
    - Content-Type : `header`
- Body :
```json
{
    "token": "string",
}
```
- Response :
```json
{
    "code": "number",
    "message": "string"
}
```

#### NotificationSteream Midtrans 
- Method : POST
- Endpoint : `/midtrans/payment-callback/`
- Header :
    - Content-Type : `application/json`
    - Accept: `application/json`
- Response :
```json
response dari midtrans, sebagai callback untuk ke midtrans
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

#### User Logout - All User
- Method : DELETE
- Endpoint : `/user/logout/`
- Header :
    - Content-Type : `application/json`
    - Accept: `application/json`
- Body :
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

#### Create Order - All User
- Method : POST
- Endpoint : `/api/order/`
- Header :
    - Accept : `application/json`
    - Content-Type : `application/json`
- Body :
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
        "name": "string",
        "email": "string",
        "address": "string",
        "telephone": "string",
        "event_name": "string",
        "event_price": "int"
    }
}
```

#### FindByUUID Order - All User
- Method : GET
- Endpoint : `/api/order/{uuid}`
- Header :
    - Accept : `application/json`
    - Content-Type : `application/json`

- Response :
```json
{
    "code": "number",
    "message": "string",
    "data": {
        "uuid": "string",
        "name": "string",
        "email": "string",
        "address": "string",
        "telephone": "string",
        "event_name": "string",
        "event_price": "int"
    }
}
```

#### FindByUserID Order - All User (  all history user )
- Method : GET
- Endpoint : `/api/order/{user_id}`
- Header :
    - Accept : `application/json`
    - Content-Type : `application/json`

- Response :
```json
{
    "code": "number",
    "message": "string",
    "data": {
        "uuid": "string",
        "name": "string",
        "email": "string",
        "address": "string",
        "telephone": "string",
        "event_name": "string",
        "event_price": "int"
    }
}
```

#### FindByUserIDDetail Order - All User ( detail history )
- Method : GET
- Endpoint : `/api/order/{user_id}/{uuid}`
- Header :
    - Accept : `application/json`
    - Content-Type : `application/json`

- Response :
```json
{
    "code": "number",
    "message": "string",
    "data": {
        "uuid": "string",
        "name": "string",
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

#### Handling Status Midtrans - Admin
- Method : GET
- Endpoint : `/api/order/{uuid}/status`
- Header :
    - Accept : `application/json`
    - Content-Type : `application/json`
- Response :
```json
response by midtrans
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