# Register

Used to collect a Token for a registered User.

**URL** : `/api/register/`

**Method** : `POST`

**Auth required** : NO

**Data constraints**

```json
{
    "username": "[valid email address]",
    "password": "[password in plain text]"
}
```

**Data example**

```json
{
    "username": "user@example.com",
    "password": "abcd1234"
}
```

## Success Response

**Code** : `201 Created`

**Content example**

```json
{
    "message": "Register successfully",
    "data": {
      "id": 1,
      "name": "Muhammad Idrus",
      "email": "user@example.com"
      }
}
```

## Error Response

**Condition** : If 'email' is already in use

**Code** : `409 CONFLICT`

**Content** :

```json
{
  "message": "Email is already in use"
}
```