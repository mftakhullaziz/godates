## Technical Test - Software Engineer at Dealls

Author: Miftakhul Aziz \
Email: mftakhullaziz@gmail.com

### API Documentation

###### Postman Link
###### https://documenter.getpostman.com/view/6097899/2sA3XLF4jf

### Architecture Service

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

### Diagram Database
    Berikut ini adalah relational database diagram

![db.png](docs/img/diagram_database.png)

### Tech Stack
    Golang
    JWT
    MySQL
    Redis

### How to Run
    
    Use makefile
    
    make build/service -> this for build project to binary go
    make run/service -> this for running binary go service
    make brun/service -> this for build and running
