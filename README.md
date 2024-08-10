# Godates BE Services

Author: Miftakhul Aziz \
Email: mftakhullaziz@gmail.com

## API Documentation

###### Postman Link
###### https://documenter.getpostman.com/view/6097899/2sA3XLF4jf

##### API Specifications Details
##### User Register

API: https://godating-dealls-service.onrender.com/godating-dealls/api/authenticate/register \
Method: POST \
Detail: This api for registered new users \
Request Body:
```
{
    "email": "david.alaba@gmail.com",
    "username": "david.alaba",
    "password": "Test1234",
    "full_name": "David Alaba"
}
```
Response Body:
```
{
    "status_code": 201,
    "is_success": true,
    "message": "Created account successfully",
    "request_at": "2024-06-10 00:00:00",
    "data": {
        "account_id": 1,
        "email": "andreas.iniesta@gmail.com",
        "username": "andreasiniesta",
        "password": "$2a$04$EWQBDmD7MHy/XzplO1BIiuq9IUlZ1WGtXFtWYCHwU.Y4GD5S.0I6C"
    },
    "total_data": 1
}
``` 

##### User Login

API: https://godating-dealls-service.onrender.com/godating-dealls/api/authenticate/login \
Method: POST \
Detail: This api for login new users \
Request Body:
```
{
    "email": "andreas.iniesta@gmail.com",
    "password": "Test1234"
}

OR

{
    "username": "andreas.iniesta",
    "password": "Test1234"
}
```
Response Body:
```
{
    "status_code": 200,
    "is_success": true,
    "message": "Login account successfully",
    "request_at": "2024-06-10 18:20:31",
    "data": {
        "username": "andreasiniesta",
        "email": "andreas.iniesta@gmail.com",
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJhY2NvdW50X2lkIjoxLCJlbWFpbCI6ImFuZHJlYXMuaW5pZXN0YUBnbWFpbC5jb20iLCJ1c2VybmFtZSI6IiIsImV4cCI6MTcxODEzMDAzMH0.YERE3wC7pY9AYblZs8F72IALJcTx95H1MCYf370lnWg"
    },
    "total_data": 1
}
``` 

##### User Logout

API: https://godating-dealls-service.onrender.com/godating-dealls/api/authenticate/logout \
Method: POST \
Detail: This api for logout new users \
Request Header:
```
Authorization: Bearer access token (REQUIRED)
```
Response Body:
```
{
    "status_code": 200,
    "is_success": true,
    "message": "Logout account successfully",
    "request_at": "2024-06-10 00:00:00",
    "data": {
        "message": "User successfully logged out"
    },
    "total_data": 1
}
``` 

##### User See Others User Daily

API: https://godating-dealls-service.onrender.com/godating-dealls/api/daily-accounts \
Method: POST \
Detail: This api for see users list with maximum 10 users in for user regular and for user premium is unlimited, and this twice user will be not found on 1 day \
Request Header:
```
Authorization: Bearer access token (REQUIRED)
```
Response Body:
```
{
    "status_code": 200,
    "is_success": true,
    "message": "Get users view successfully",
    "request_at": "2024-06-10 18:22:59",
    "data": [
        {
            "user_id": 9,
            "account_id": 9,
            "full_name": "Natashya Bella",
            "username": "natashya.bella",
            "photos": [],
            "videos": [],
            "age": 0,
            "gender": "",
            "address": "",
            "bio": "",
            "verified": false
        },
        {
            "user_id": 5,
            "account_id": 5,
            "full_name": "Toni Kroos",
            "username": "toni.kroos",
            "photos": [],
            "videos": [],
            "age": 0,
            "gender": "",
            "address": "",
            "bio": "",
            "verified": false
        },
        {
            "user_id": 15,
            "account_id": 15,
            "full_name": "Ondo",
            "username": "ondo",
            "photos": [],
            "videos": [],
            "age": 0,
            "gender": "",
            "address": "",
            "bio": "",
            "verified": false
        }
    ],
    "total_data": 3
}
``` 

##### User Update Profile

API: https://godating-dealls-service.onrender.com/godating-dealls/api/users \
Method: PATCH \
Detail: This api for update profile users \
Request Header:
```
Authorization: Bearer access token (REQUIRED)
```
Request Body:
```
{
    "full_name": "Andreas Iniesta",
    "gender": "P",
    "address": "Jakarta Selatan, Kebayoran Lama",
    "bio": "How to make money!!!",
    "date_of_birth": "1994-01-09"
}
```
Response Body:
```
{
    "status_code": 201,
    "is_success": true,
    "message": "Patch user successfully",
    "request_at": "2024-06-10 18:24:31",
    "data": {
        "user_id": 1,
        "full_name": "Andreas Iniesta",
        "gender": "P",
        "address": "Jakarta Selatan, Kebayoran Lama",
        "bio": "How to make money!!!",
        "date_of_birth": "1994-01-09",
        "age": 30,
        "account_id": 1,
        "updated_at": "2024-06-10 18:24:31"
    },
    "total_data": 1
}
``` 

##### User Actions Swipe From See Users List Daily

API: https://godating-dealls-service.onrender.com/godating-dealls/api/swipes \
Method: POST \
Detail: This api for swipe from user see on list, when user reguler just in 10 swipe every day and for premium user is unlimited in every day \
Request Header:
```
Authorization: Bearer access token (REQUIRED)
```
Request Body:
```
{
    "action_type": "right", // left or right
    "account_id_swipe": 7
}
```
Response Body:
```
{
    "status_code": 200,
    "is_success": true,
    "message": "Get users view successfully",
    "request_at": "2024-06-11 01:45:44",
    "data": {
        "message": "Account Liked!"
    },
    "total_data": 1
}
``` 

##### User Get Packages

API: https://godating-dealls-service.onrender.com/godating-dealls/api/packages \
Method: GET \
Detail: This api for list packages for purchase to premium account \
Request Header:
```
Authorization: Bearer access token (REQUIRED)
```
Response Body:
```
{
    "status_code": 200,
    "is_success": true,
    "message": "Get packages successfully",
    "request_at": "2024-06-10 18:46:48",
    "data": [
        {
            "package_id": 1,
            "package_name": "Basic Package",
            "description": "Access to basic features for one month",
            "package_duration_in_monthly": 1,
            "price": 99999,
            "unlimited_swipes": true,
            "status": true
        },
        {
            "package_id": 2,
            "package_name": "Standard Package",
            "description": "Access to standard features for three months",
            "package_duration_in_monthly": 3,
            "price": 249999,
            "unlimited_swipes": true,
            "status": true
        },
        {
            "package_id": 3,
            "package_name": "Premium Package",
            "description": "Access to all features including unlimited swipes for six months",
            "package_duration_in_monthly": 6,
            "price": 499999,
            "unlimited_swipes": true,
            "status": true
        }
    ],
    "total_data": 3
}
``` 

##### User Purchase Packages

API: https://godating-dealls-service.onrender.com/godating-dealls/api/purchase-package \
Method: POST \
Detail: This api for purchasing package to make user verified as premium badge \
Request Header:
```
Authorization: Bearer access token (REQUIRED)
```
Request Body:
```
{
    "package_id": 1,
    "price": 99999,
    "package_name": "Basic Package",
    "package_duration_in_monthly": 1,
    "unlimited_swipes": true
}
```
Response Body:
```
{
    "status_code": 200,
    "is_success": true,
    "message": "Purchase packages successfully",
    "request_at": "2024-06-10 20:55:34",
    "data": {
        "package_id": 1,
        "price": 99999,
        "message": "Purchased package successfully"
    },
    "total_data": 1
}
``` 

##### User Check Quota Swipe Daily

API: https://godating-dealls-service.onrender.com/godating-dealls/api/quota \
Method: GET \
Detail: This api for purchasing package to make user verified as premium badge \
Request Header:
```
Authorization: Bearer access token (REQUIRED)
```
Response Body:
```
{
    "status_code": 200,
    "is_success": true,
    "message": "Fetch quota successfully",
    "request_at": "2024-06-10 18:38:05",
    "data": {
        "total_quotas": "8",
        "swipe_count": 2
    },
    "total_data": 1
}

OR for premium user

{
    "status_code": 200,
    "is_success": true,
    "message": "Fetch quota successfully",
    "request_at": "2024-06-10 20:59:38",
    "data": {
        "total_quotas": "Unlimited Until 2024-07-10 20:55:35",
        "swipe_count": 100
    },
    "total_data": 1
}
``` 

##### User Get Account Details

API: https://godating-dealls-service.onrender.com/godating-dealls/api/account-details \
Method: GET \
Detail: This api for get user account details \
Request Header:
```
Authorization: Bearer access token (REQUIRED)
```
Response Body:
```
{
    "status_code": 200,
    "is_success": true,
    "message": "Fetch account detail successfully",
    "request_at": "2024-06-10 19:26:14",
    "data": {
        "account_data": {
            "user_id": 1,
            "account_id": 1,
            "full_name": "Andreas Iniesta",
            "username": "andreasiniesta",
            "email": "andreas.iniesta@gmail.com",
            "age": 30,
            "gender": "P",
            "address": "Jakarta Selatan, Kebayoran Lama",
            "bio": "How to make money!!!",
            "verified": false,
            "date_of_birth": "1994-01-09T00:00:00Z"
        },
        "account_view": {
            "total_swipe_like": 7,
            "total_swipe_passed": 2,
            "total_views": 100 
        }
    },
    "total_data": 1
}
``` 

##### User View Profile Other User

API: https://godating-dealls-service.onrender.com/godating-dealls/api/account-view \
Method: POST \
Detail: This api for view user account others \
Request Header:
```
Authorization: Bearer access token (REQUIRED)
```
Request Body:
```
{
    "account_id_view": 7
}
```
Response Body:
```
{
    "status_code": 200,
    "is_success": true,
    "message": "View account successfully",
    "request_at": "2024-06-10 19:28:02",
    "data": {
        "account_id": 7,
        "user_name": "frenkie.dejong",
        "email": "frenkie.dejong@gmail.com",
        "verified": false,
        "full_name": "Frenkie De Jong"
    },
    "total_data": 1
}
``` 

## Architecture Service

![img.png](docs/img/clean-architecture.png)

###### This project building using implement clean architecture, with details

###### Infrastructure layer
Layer ini merupakan bagian terluar dari struktur proyek yang mencakup frameworks dan driver. Semua proses yang terjadi di layer ini termasuk dalam eksternal dependency, seperti third party library atau database. Pada proyek ini, saya membuat lapisan infrastruktur yang berisi dependency ke database, Redis, atau library eksternal lainnya. Keuntungan dari penggunaan layer ini adalah bahwa jika terjadi perubahan pada database atau pihak ketiga lainnya, kita tidak perlu mengubah banyak kode, bahkan mungkin tidak perlu melakukan perubahan di package core. Perubahan hanya terjadi pada implementasi fungsi-fungsi tertentu.

###### Delivery
Lapisan ini mewakili layer eksternal pada clean architecture, yaitu layer interface adapters. Pada proyek ini, lapisan ini direpresentasikan oleh package presenter dan controller yang diwakili oleh package handler. Lapisan ini bertanggung jawab untuk berinteraksi dengan dunia luar, seperti menerima HTTP requests dan memberikan responses.

###### Core
Core dari clean architecture terdiri dari enterprice business rules dan application business rules. Application business rules biasanya disebut sebagai use case atau proses bisnis dari aplikasi. Pada proyek ini, use case akan memproses alur data dengan memanggil entity-entity dari enterprise business rules. Enterprise business rules merupakan layer yang mengatur perilaku mereka sendiri dan bertanggung jawab untuk mengatur logika dan perilaku dari entity-entitynya sendiri.

###### Domain
Domain dalam clean architecture mengacu pada bagian dari sistem yang berisi aturan bisnis inti, logika, dan entitas utama yang membentuk inti dari aplikasi. Pada proyek ini, package domain berisi representasi dari entitas dan data penting dalam aplikasi. Meskipun tidak ada logika bisnis yang diimplementasikan di sini, domain mendefinisikan format dan jenis data yang digunakan dalam sistem ini. Representasi data yang konsisten dan jelas adalah kunci untuk memastikan bahwa aplikasi dapat berfungsi dengan baik dan berkomunikasi dengan benar di semua lapisan.

###### Common
Layer "common" seringkali berfungsi sebagai tempat untuk menyimpan kode yang digunakan secara umum di seluruh aplikasi. Pada proyek ini, package common digunakan untuk mendefinisikan utilitas, fungsi bantuan, konstanta, dan fungsi lain yang digunakan secara luas di berbagai bagian dari aplikasi.

#### Diagram Database
    Berikut ini adalah relational database diagram

![db.png](docs/img/diagram_database.png)

## How to Run
    
##### Requirement:
Golang\
MySQL\
Redis

clone this project: git clone https://github.com/mftakhullaziz/godating-dealls.git \
after clone, checkout branch development to running mode localhost

how to run, if your device support makefile its will be easy, example:
command\
```makefile
make build/service -> this for build project to binary go\
make run/service -> this for running binary go service\
make brun/service -> this for build and running
```

if not have make file just run:
```makefile
go run main.go
```