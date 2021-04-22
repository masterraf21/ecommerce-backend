
# API Documentation for e-commerce App 
Built using docgen https://github.com/thedevsaddam/docgen


## Indices

* [Buyer](#buyer)

  * [CREATE Buyer](#1-create-buyer)
  * [GET All Buyer](#2-get-all-buyer)
  * [GET Buyer By ID](#3-get-buyer-by-id)

* [Order](#order)

  * [Accept Order](#1-accept-order)
  * [CREATE Order](#2-create-order)
  * [GET Acccepted Order By Seller ID](#3-get-acccepted-order-by-seller-id)
  * [GET Accepted Order By Buyer ID](#4-get-accepted-order-by-buyer-id)
  * [GET All Order](#5-get-all-order)
  * [GET Order By Buyer ID](#6-get-order-by-buyer-id)
  * [GET Order By ID Copy](#7-get-order-by-id-copy)
  * [GET Order By Seller ID](#8-get-order-by-seller-id)
  * [GET Pending Order By Buyer ID](#9-get-pending-order-by-buyer-id)
  * [GET Pending Order By Seller ID](#10-get-pending-order-by-seller-id)

* [Product](#product)

  * [CREATE Product](#1-create-product)
  * [GET All Product](#2-get-all-product)
  * [GET Product By ID](#3-get-product-by-id)
  * [GET Product By Seller ID](#4-get-product-by-seller-id)

* [Seller](#seller)

  * [CREATE Seller](#1-create-seller)
  * [GET All Seller](#2-get-all-seller)
  * [GET Seller By ID](#3-get-seller-by-id)


--------


## Buyer



### 1. CREATE Buyer



***Endpoint:***

```bash
Method: POST
Type: RAW
URL: {{e-commerce}}/buyer
```



***Body:***

```js        
{
    "email": "apa@gmail.com",
    "name": "Aang",
    "password": "huyu",
    "delivery_address": "huhuhuhu"
}
```



### 2. GET All Buyer



***Endpoint:***

```bash
Method: GET
Type: 
URL: {{e-commerce}}/buyer
```



### 3. GET Buyer By ID



***Endpoint:***

```bash
Method: GET
Type: 
URL: {{e-commerce}}/buyer/:id_buyer
```



***URL variables:***

| Key | Value | Description |
| --- | ------|-------------|
| id_buyer | 1 |  |



## Order



### 1. Accept Order



***Endpoint:***

```bash
Method: POST
Type: 
URL: {{e-commerce}}/order/:id_order/accept
```



***URL variables:***

| Key | Value | Description |
| --- | ------|-------------|
| id_order | 1 |  |



### 2. CREATE Order



***Endpoint:***

```bash
Method: POST
Type: RAW
URL: {{e-commerce}}/order
```



***Body:***

```js        
{
    "id_buyer": 1,
    "id_seller": 1,
    "products": [
        {
            "id_product": 1,
            "quantity": 10
        }
    ]
}
```



### 3. GET Acccepted Order By Seller ID



***Endpoint:***

```bash
Method: GET
Type: 
URL: {{e-commerce}}/order/seller/:id_seller/accepted
```



***URL variables:***

| Key | Value | Description |
| --- | ------|-------------|
| id_seller | 1 |  |



### 4. GET Accepted Order By Buyer ID



***Endpoint:***

```bash
Method: GET
Type: 
URL: {{e-commerce}}/order/buyer/:id_buyer/accepted
```



***URL variables:***

| Key | Value | Description |
| --- | ------|-------------|
| id_buyer |  |  |



### 5. GET All Order



***Endpoint:***

```bash
Method: GET
Type: 
URL: {{e-commerce}}/order
```



### 6. GET Order By Buyer ID



***Endpoint:***

```bash
Method: GET
Type: 
URL: {{e-commerce}}/order/buyer/:id_buyer
```



***URL variables:***

| Key | Value | Description |
| --- | ------|-------------|
| id_buyer | 1 |  |



### 7. GET Order By ID



***Endpoint:***

```bash
Method: GET
Type: 
URL: {{e-commerce}}/order/:id_order
```



***URL variables:***

| Key | Value | Description |
| --- | ------|-------------|
| id_order | 2 |  |



### 8. GET Order By Seller ID



***Endpoint:***

```bash
Method: GET
Type: 
URL: {{e-commerce}}/order/seller/:id_seller
```



***URL variables:***

| Key | Value | Description |
| --- | ------|-------------|
| id_seller |  |  |



### 9. GET Pending Order By Buyer ID



***Endpoint:***

```bash
Method: GET
Type: 
URL: {{e-commerce}}/order/buyer/:id_buyer/pending
```



***URL variables:***

| Key | Value | Description |
| --- | ------|-------------|
| id_buyer |  |  |



### 10. GET Pending Order By Seller ID



***Endpoint:***

```bash
Method: GET
Type: 
URL: {{e-commerce}}/order/seller/:id_seller/pending
```



***URL variables:***

| Key | Value | Description |
| --- | ------|-------------|
| id_seller | 1 |  |



## Product



### 1. CREATE Product



***Endpoint:***

```bash
Method: POST
Type: RAW
URL: {{e-commerce}}/product
```



***Body:***

```js        
{
     "product_name": "Hey",
            "description": "Wuwhwuw",
            "price": 123100,
            "id_seller": 1
}
```



### 2. GET All Product



***Endpoint:***

```bash
Method: GET
Type: 
URL: {{e-commerce}}/product
```



### 3. GET Product By ID



***Endpoint:***

```bash
Method: GET
Type: 
URL: {{e-commerce}}/product/:id_product
```



***URL variables:***

| Key | Value | Description |
| --- | ------|-------------|
| id_product | 1123 |  |



### 4. GET Product By Seller ID



***Endpoint:***

```bash
Method: GET
Type: 
URL: {{e-commerce}}/product/seller/:id_seller
```



***URL variables:***

| Key | Value | Description |
| --- | ------|-------------|
| id_seller | 12 |  |



## Seller



### 1. CREATE Seller



***Endpoint:***

```bash
Method: POST
Type: RAW
URL: {{e-commerce}}/seller
```



***Body:***

```js        
{
    "email": "apa@gmail.com",
    "name": "Aang",
    "password": "huyu",
    "delivery_address": "asdadssda"
}
```



### 2. GET All Seller



***Endpoint:***

```bash
Method: GET
Type: 
URL: {{e-commerce}}/seller
```



### 3. GET Seller By ID



***Endpoint:***

```bash
Method: GET
Type: 
URL: {{e-commerce}}/seller/:id_seller
```



***URL variables:***

| Key | Value | Description |
| --- | ------|-------------|
| id_seller |  |  |



---
[Back to top](#e-commerce)
> Made with &#9829; by [thedevsaddam](https://github.com/thedevsaddam) | Generated at: 2021-04-22 10:57:40 by [docgen](https://github.com/thedevsaddam/docgen)
