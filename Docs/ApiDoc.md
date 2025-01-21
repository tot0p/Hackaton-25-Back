# API Doc

## Table of Contents

- [Ping](#ping)
- [Create User](#create-user)
- [Login](#login)
- [profile](#profile)
- [Travels](#travels)
- [Add Travel](#add-travel)
- [Calculate CO2](#calculate-co2)


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
  "token": "jwt token",
  "exp" : 2772000
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

```json
{
  "error": "Internal server error"
}
```

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
    "password": "password"
}
```

**Success Response**

**Code** : `200 OK`

**Content** :
```json
{
    "token": "jwt token",
    "exp" : 2772000
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

```json
{
  "error": "Internal server error"
}
```


## profile

Get user profile

**URL** : `/profile`

**Method** : `GET`

**Headers** : `Authorization: Bearer

**Success Response**

**Code** : `200 OK`

**Content** :
```json
{
    "username": "test"
}
```

**Error Response**

**Code** : `401 UNAUTHORIZED`

**Content** :

```json
{
  "error": "Unauthorized"
}
```

**Code** : `500 INTERNAL SERVER ERROR`

**Content** :

```json
{
  "error": "Internal server error"
}
```

## Travels

Get user travels

**URL** : `/travels`

**Method** : `GET`

**Headers** : `Authorization: Bearer

**Success Response**

**Code** : `200 OK`

**Content** :
```json
[
  {
    "uuid": "f9e086d3-5c90-4eb6-a91f-993702704ba5",
    "userUUID": "9a0973a1-6020-49c0-89c9-4e01fb094fa3",
    "date": "2025-01-20T00:00:00Z",
    "startLocation": "Paris",
    "endLocation": "Madrid",
    "distance": 30004,
    "duration": 3000,
    "co2": 405.4,
    "transportType": "Voiture"
  }
]
```

**Error Response**

**Code** : `401 UNAUTHORIZED`

**Content** :

```json
{
  "error": "Unauthorized"
}
```

**Code** : `500 INTERNAL SERVER ERROR`

**Content** :

```json
{
  "error": "Internal server error"
}
```

## Add Travel

Add a new travel

**URL** : `/travels`

**Method** : `POST`

**Headers** : `Authorization: Bearer

**Data constraints**

```json
{
    "startLocation": "[start location in plain text]",
    "endLocation": "[end location in plain text]",
    "distance": "[distance in meters]",
    "duration": "[duration in seconds]",
    "co2": "[co2 in grams]",
    "transportType": "[transport type in plain text]"
}
```

**Data example**

```json
{
    "startLocation": "Paris",
    "endLocation": "Madrid",
    "distance": 30004,
    "duration": 3000,
    "co2": 405.4,
    "transportType": "Voiture"
}
```

**Success Response**

**Code** : `201 CREATED`

**Content** :
```json
{
  "message": "Travel added"
}
```

**Error Response**

**Code** : `401 UNAUTHORIZED`

**Content** :

```json
{
  "error": "Unauthorized"
}
```

**Code** : `500 INTERNAL SERVER ERROR`

**Content** :

```json
{
  "error": "Internal server error"
}
```

## Calculate CO2

Calculate CO2 for a travel

**URL** : `/calculate-co2`

**Method** : `POST`

**Headers** : `Authorization: Bearer

**Data constraints**

```json
{
  "itinerary": [
    {
      "address": "[address in plain text]"
    },
    {
      "address": "[address in plain text]"
    }
  ]
}
```

OR 

```json
{
  "itinerary": [
    {
      "lat": 41.9028,
      "lng": 12.4964
    },
    {
      "address": "[address in plain text]"
    }
  ]
}
```

OR 

```json
{
  "itinerary": [
    {
      "lat": 48.8566,
      "lng": 2.3522
    },
    {
      "lat": 41.9028,
      "lng": 12.4964
    }
  ]
}
```

**Data example**

```json
{
  "itinerary": [
    {
      "address": "Paris"
    },
    {
      "address": "Madrid"
    }
  ]
}
```

**Success Response**

**Code** : `200 OK`

**Content** :
```json
{
  "bikingFootprint": 11052.801487141147,
  "busFootprint": 88422.41189712918,
  "carFootprint": 132633.61784569375,
  "planeFootprint": 165792.02230711718,
  "trainFootprint": 33158.40446142344,
  "truckFootprint": 176844.82379425835,
  "walkingFootprint": 11052.801487141147
}
```

**Error Response**

**Code** : `401 UNAUTHORIZED`

**Content** :

```json
{
  "error": "Unauthorized"
}
```


**Code** : `500 INTERNAL SERVER ERROR`

**Content** :

```json
{
  "error": "Internal server error"
}
```


