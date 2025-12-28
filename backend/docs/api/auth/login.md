# Login

Used to collect a Token for a registered User.

**URL** : `/api/login/`

**Method** : `POST`

**Auth required** : NO

**Data constraints**

```json
{
    "email": "[valid email address]",
    "password": "[password in plain text]"
}
```

**Data example**

```json
{
    "email": "user@example.com",
    "password": "abcd1234"
}
```

## Success Response

**Code** : `200 OK`

**Content example**

```json
{
    "message": "Login successfully",
    "data": {
      "id": 1,
      "name": "Muhammad Idrus",
      "email": "user@example.com"
      },
    "token": "93144b288eb1fdccbe46d6fc0f241a51766ecd3d"
}
```

## Error Response

**Condition** : If 'email' and 'password' combination is wrong.

**Code** : `400 BAD REQUEST`

**Content** :

```json
{
  "message": "Invalid email or password"
}
```