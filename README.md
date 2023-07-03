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