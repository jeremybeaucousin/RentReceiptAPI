# Docker
## Create container 
> docker run --name RentReceiptGenerator -d -e POSTGRES_PASSWORD=Jonathan -e POSTGRES_USER=rentReceiptGenerator -e POSTGRES_DB=rentReceiptGenerator -e PGDATA=/var/lib/postgresql/data/pgdata -v "rentReceiptGenerator:/var/lib/postgresql/data" -p "5433:5432" postgres

# GO 
## package manager
Initiate package
> go mod init github.com/jeremybeaucousin/RentReceiptAPI

## Download dependencies
> go get

## remove dependencies
> go get github.com/lib/pq@none

## Update package
> go get -u all

# Run app
> go run .

# Google cloud
> gcloud run deploy

## Run with cloud db
> docker run -d -e GOOGLE_APPLICATION_CREDENTIALS=./rent-reiceipt-generator-4e7e69f5f82c.json -e INSTANCE_CONNECTION_NAME='rent-reiceipt-generator:europe-west9:rent-receipt-generator-db' -e DB_NAME='rentReceiptGenerator' -e DB_USER='rentReceiptGenerator' -e DB_PASS='rentReceiptGenerator' -e ORIGIN_RENT_REICEIPT_GENERATOR=http://localhost:3000 --name api_test --rm -p 8081:8080 api:test
a68c3144a58b3a349cff2fbc9bddcc3c4425eb0bd2f64c142ec0723d77bd7d51