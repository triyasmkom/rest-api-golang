# Unjuk Ketrampilan Prakerja


## 1. Prepare Project

### Insialisasi Project

```shell
# Buat folder project
mkdir "server-golang"

# masuk ke folder project
cd "server-golang"

# Inisialisasi Project
go mod init server-golang
```


### Instalasi package Golang

```shell
# install framework echo
go get github.com/labstack/echo/v4

# install orm gorm
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql

# install env
go get github.com/joho/godotenv

# bcrypt
go get golang.org/x/crypto/bcrypt

# jwt
go get "github.com/labstack/echo-jwt/v4"

# menghapus package
go clean -i github.com/codegangsta/gin

```

## Dokumentasi

Postman : https://documenter.getpostman.com/view/26067457/2s9YC7RWBR

## Menjalankan program 

### Menjalankan file

```shell
go run main.go
```

### Auto restart

```shell
# Download package
go get github.com/codegangsta/gin

# Mengguankannya
gin -i run main.go

# Menghentikan Sementara
ctrl+c

# Menghapus permanen
go clean -i github.com/codegangsta/gin
```
