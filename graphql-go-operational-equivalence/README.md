# GraphQL Go Operational Equivalence

Example application for the operational equivalence excercise.

## Getting Started

Running it:
```bash
./bin/start.sh
```

## Output

```bash
Query results:
{
    "data": {
        "ID": "d983b9d9-681c-4059-b5a3-5329d1c6f82d",
        "boolean": true,
        "enum": "SECOND_ENUM",
        "float": 20.01,
        "int": 20,
        "node": {
            "id": "user-1",
            "name": "John Doe"
        },
        "object": {
            "name": "Name of the object instance."
        },
        "objectList": [
            {
                "name": "First object in list"
            },
            {
                "name": "Second object in list"
            },
            {
                "name": "Third object in list"
            }
        ],
        "objectWithArguments": {
            "name": "Name of the object with arguments instance, id: 1"
        },
        "product": {
            "id": "product-1",
            "name": "GraphQL Book"
        },
        "productNonNull": {
            "id": "product-non-null-1",
            "name": "GraphQL Book Non-Null"
        },
        "searchResult": {
            "id": "user-1",
            "name": "John Doe"
        },
        "string": "str",
        "stringList": [
            "first string",
            "second string",
            "third string"
        ],
        "user": {
            "id": "user-1",
            "name": "John Doe"
        },
        "userNonNull": {
            "id": "user-non-null-1",
            "name": "John Doe Non-Null"
        }
    }
}

Mutation results:
{
    "data": {
        "createProduct": {
            "id": "product-1",
            "name": "GraphQL Guide",
            "price": 49.99
        },
        "createUser": {
            "id": "user-1",
            "name": "Alice"
        },
        "deleteProduct": "Product with id: product-2 deleted successfully",
        "deleteUser": "User with id: user-2 deleted successfully",
        "updateProduct": {
            "id": "product-1",
            "name": "GraphQL Guide Updated",
            "price": 59.99
        },
        "updateUser": {
            "id": "user-1",
            "name": "Alice Updated"
        }
    }
}

Subscription results:
{
    "data": {
        "productAdded": {
            "id": "product-1",
            "name": "New Product Added",
            "price": 0
        },
        "userAdded": {
            "id": "user-1",
            "name": "New User Added"
        }
    }
}
```
