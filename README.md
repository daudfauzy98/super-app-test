# PT Super App Test
This is a coding test for a company recruitment. This project is using PostgreSQL. I have put a migration and seeder in the `main.go` so it will do the job every time you running the project.

## List of Routes

### Product
### GET /products
#### Response
```javascript
{
    "data": [
        {
            "brand": "Pepsodent",
            "category": "Pasta gigi",
            "id": "12197be3-45a9-45fd-942d-fd67d2a45aca",
            "name": "Pepsodent Sparkling Mint",
            "stocks": 30
        },
        {
            "brand": "Head&Shoulder",
            "category": "Perlalatan mandi",
            "id": "f4fb53f9-e933-4111-abbf-692758b6e0c8",
            "name": "Shampoo",
            "stocks": 25
        }
    ],
    "error": false,
    "message": "success"
}
```

### GET /product/:id
#### Request
* Path param __id__ (product)
#### Response
```javascript
{
    "data": {
        "brand": "Head&Shoulder",
        "category": "Perlalatan mandi",
        "id": "476b8653-34f3-48e3-88e8-9850a6259cd9",
        "name": "Shampoo",
        "stocks": 25
    },
    "error": false,
    "message": "success"
}
```

### POST /product
#### Request
```javascript
{
    "name": "Pepsodent Sparkling Mint",
    "category": "Pasta gigi",
    "brand": "Pepsodent",
    "stocks": [
        {
            "warehouse_loc": "Ngawi",
            "count": 30
        }
    ]
}
```
#### Response
```javascript
{
    "data": {
        "id": "12197be3-45a9-45fd-942d-fd67d2a45aca",
        "name": "Pepsodent Sparkling Mint",
        "category": "Pasta gigi",
        "brand": "Pepsodent",
        "stocks": [
            {
                "id": "40f38da3-44ed-4859-8305-34113072f498",
                "product_id": "12197be3-45a9-45fd-942d-fd67d2a45aca",
                "warehouse_id": "950e4e0b-17c9-4308-bfd8-bc82646710a1",
                "warehouse_loc": "Ngawi",
                "count": 30
            }
        ]
    },
    "error": false,
    "message": "success"
}
```

## Stock
### GET /stock/:id
#### Request
* Path param __id__ (product)
#### Response
```javascript
{
    "data": [
        {
            "id": "83f7bf75-1ab7-4125-bb49-16a1c76a32cb",
            "product_id": "57b30f5c-3522-4794-a334-daed7bdb63fc",
            "warehouse_id": "aee1a7fd-7db5-429d-a8c0-1686f86a8ddb",
            "warehouse_loc": "Ngawi",
            "count": 30
        },
        {
            "id": "027da742-969e-4451-8423-daba0da13c96",
            "product_id": "57b30f5c-3522-4794-a334-daed7bdb63fc",
            "warehouse_id": "1609f714-b4ed-4c09-8715-5281176f1bbb",
            "warehouse_loc": "Malang",
            "count": 25
        }
    ],
    "error": false,
    "message": "success"
}
```
### PUT /stock/:id
#### Request
* Path param __id__ (stock)
* Query param __count__
#### Response
```javascript
{
    "error": false,
    "message": "success"
}
```

## Background Process
### PUT /order/:id
#### Request
* Path param __id__ (product)
* Query param __order__
#### Response
```javascript
{
    "error": false,
    "message": "success"
}
```
