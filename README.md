# TodoList - A simple todo list app
---
### Daftar Isi

- [Tentang](#tentang)
- [Panduan Menjalankan](#panduan-menjalankan)
- [Struktur Project](#struktur-project)
---

### Tentang
TodoList merupakan RestFull API untuk mencatat activity dan todo dari pengguna yang dibuat menggunakan bahasa pemrograman Go. RestFull API ini dibuat dengan menggunakan Clean Architecture dan menggunakan database MySQL.

----
### Panduan Menjalankan

Proses menjalankan project dapat dilakukan dengan 3 cara yaitu menjalankan file main secara langsung maupun menggunakan Makefile. Berikut merupakan penjelasan cara menjalankan proyek ini:
+ Cara Pertama
Untuk menjalankan secara langsung file main.go dapat menggunakan command sebagai berikut:
    ```bash
    go run ./app/main.go
    ```
+ Cara Kedua
Untuk menjalankan menggunakan Makefile dapat menggunakan command sebagai berikut:
    ```bash
    Make run
    ```
+ Cara Ketiga
Untuk menjalankan menggunakan nodemon dapat menggunakan command sebagai berikut:
    ```bash
    nodemon --exec go run ./app/main.go
    ```
    Atau dapat menggunakan Makefile yang telah menyediakan command nodemon:

    ```bash
    Make run-nodemon
    ```
    Untuk detail lengkap cara menjalankan menggunakan Makefile dapat dilihat difile Makefile.

----
### Struktur Project
Secara sederhana project ini dibuat dengan menggunakan struktur Clean Architecture yang dipopulerkan oleh Uncle Bob. Berikut merupakan struktur Clean Architecture versi Uncle Bob:


![image](https://user-images.githubusercontent.com/13291041/102681893-84326980-4208-11eb-8f84-2959e03b89d8.png)


Dari Struktur tersebut dilakukan penyesuaian, dikarenakan dalam Rest API masih menggunakan API yang sederhana, maka struktur project akan terlihat seperti berikut:
| Layer                | Directory      |
|----------------------|----------------|
| Frameworks & Drivers | Infrastructures|
| Interface            | Interfaces     |
| Usecases             | Usecases       |
| Entities             | Domain         |
