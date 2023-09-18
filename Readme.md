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

## 2. Dokumentasi

Postman : https://documenter.getpostman.com/view/26067457/2s9YC7RWBR

## 3. Menjalankan program 

### Menjalankan file

#### 1. Silakan clone di  git@github.com:triyasmkom/server-golang.git
#### 2. Setup file ```.env```, silakan copy dari file ```.env.sample```

```env

DB_USER=username
DB_PASSWORD=password
DB_HOST=host
DB_PORT=port
DB_NAME=database_name
PORT=10000
JWT_KEY=secret
JWT_EXP=72
DEBUG=true

```

#### 3. Buka dokumentasi postman dan lihat API yang tersedia

#### 4. Jalankan perintah berikut ini

```shell
# untuk reload package
go mod tidy

# untuk menjalankan file
go run main.go

# untuk help
go help

go mod help
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


## 4. Catatan Koding

### Akses Interface

```
var email string

if value, ok := verifyJwt.Data.(map[string]interface{}); ok {
    email = value["email"].(string)
}
verifyUser := VerifyUser(email)
```


## 5 Referensi

- framework: https://echo.labstack.com/
- orm: https://gorm.io/index.html
