# API Doc

## Table of Contents

- [Ping](#ping)
- [Create User](#create-user)
- [Login](#login)


## Ping

Check if the server is running

**URL** : `/ping`

**Method** : `GET`

**Success Response**

**Code** : `200 OK`

**Content** :
```json
{
    "message": "pong"
}
```

## Create User

Create a new user

**URL** : `/user`

**Method** : `POST`

**Data constraints**

```json
{
    "username": "[username in plain text]",
    "password": "[password in plain text]"
}
```

**Data example**

```json
{
    "username": "user",
    "password": "password"
}
```

**Success Response**

**Code** : `201 CREATED`

**Content** :
```json
{
  "message": "User created"
}
```

**Error Response**

**Code** : `400 BAD REQUEST`

**Content** :
```json
{
  "error": "User already exist"
}
```

**Code** : `500 INTERNAL SERVER ERROR`

**Content** :
plain text actually

## Login

Login with a user

**URL** : `/login`

**Method** : `POST`

**Data constraints**

```json
{
    "username": "[username in plain text]",
    "password": "[password in plain text]", 
    "Device": "[device name in plain text]"
}
```

**Data example**

```json
{
    "username": "user",
    "password": "password",
    "Device": "phone"
}
```

**Success Response**

**Code** : `200 OK`

**Content** :
```json
{
    "token": "1e6397b5-3bf2-4f89-bbba-XXXXXXXXXXXX",
    "username": "test",
    "device": "phone"
}
```

**Error Response**

**Code** : `400 BAD REQUEST`

**Content** :
```json
{
  "error": "Invalid password or username"
}
```

**Code** : `500 INTERNAL SERVER ERROR`

**Content** :
plain text actually
