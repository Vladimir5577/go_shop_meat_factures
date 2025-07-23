Install dependencies:
    $ go get -u github.com/go-chi/chi/v5
    $ go get github.com/go-playground/validator/v10

// ===================================
Live reload;
    $ go install github.com/air-verse/air@latest
    binary file will be in the path:
        ~/go/bin/air

Init .air.toml file:
    $ ~/go/bin/air init  --- it will create .air.toml file, put you credentials in it
    in this case in line (need to specify where main.go file and where to build binary)
    cmd = "go build -o ./tmp/main cmd/main.go"

Run project with air:
    $ ~/go/bin/air


// ====================================

fmt.Printf("%#v", myStruct)

POST /products

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
        },
        {
            "product_id": 6,
            "amount": 2
        },
        {
            "product_id": 10,
            "amount": 7
        }
    ],
    "comment": "My order!"
}