# Docker
## Create container 
> docker run --name RentReceiptGenerator -d -e POSTGRES_PASSWORD=Jonathan -e POSTGRES_USER=rentReceiptGenerator -e POSTGRES_DB=rentReceiptGenerator -e PGDATA=/var/lib/postgresql/data/pgdata -v "rentReceiptGenerator:/var/lib/postgresql/data" -p "5433:5432" postgres

# GO package manager
Initiate package
> go mod init github.com/jeremybeaucousin/RentReceiptAPI

Download dependencies
> go get
