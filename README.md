# Internet shop for selling meat products

## Preparing project
Copy .env.dist to .env
```bash
cp .env.dist .env
```
and write credentials <br />
Rum migrations with goose (database should be created):

```bash
goose up
```

For testing purpose in migration have seeds.

## http endpoints

Register a new user:
>/register - POST 

```
{
    "name": "Bob",
    "password": "12345",
    "phone": "1234567",
    "address": "Moscow"
}
```
Login: 
>/login - POST 
```
{
    "password": "12345",
    "phone": "1234567"
}
```
Get all products
>/products - GET

Create order - only for registered users!!!
> Authorization with bearer token  <br />
>/orders - POST 
```
{
    "products": [
        {
            "product_id": 1,
            "amount": 3
        },
        {
            "product_id": 2,
            "amount": 7
        },
        {
            "product_id": 5,
            "amount": 4
        }
    ],
    "comment": "My order!"
}
```

Get all user orders
```
/orders?user_id=123
```