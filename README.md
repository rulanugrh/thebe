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
> Sebelum melakukan query ke dalam database harus membuat 2 role yaitu `Administrator` dan `Peserta` ke dalam database. Serta Untuk melakukan query bisa `register` lalu `login` terlebih dahulu.

#### Create Role
- Method : POST
- Endpoint : `/role/`
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
#### Create User
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
#### Login User
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


## Authentication
> Semua api wajib di tambahkan authentikasi dari hasil login user
- Header :
    - Authorization: Token JWT

### Create Event
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

#### FindByID Event
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
### Update Event
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
#### FindByID Role
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
#### Update Role
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