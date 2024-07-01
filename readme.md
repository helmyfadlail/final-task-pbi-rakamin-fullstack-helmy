# Image Profile API

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)

Image Profile API adalah sebuah API sederhana yang memungkinkan pengguna untuk mendaftarkan akun, login, mengunggah gambar profil, dan menghapus gambar profil. Proyek ini dibangun menggunakan Go, Gin Gonic, GORM, JWT, dan beberapa pustaka lain untuk menangani validasi dan hashing password.

## Prasyarat

Pastikan Anda telah menginstal perangkat lunak berikut:

- Go 1.18 atau versi lebih baru
- MySQL atau PostgreSQL

## Struktur Folder

```
.
├── app
│   ├── photo.go
│   └── user.go
├── controllers
│   ├── photoController.go
│   └── userController.go
├── database
│   ├── db.go
│   └── migrate.go
├── helpers
│   ├── bcrypt.go
│   └── jwt.go
├── middlewares
│   └── authMiddleware.go
├── models
│   ├── photoModel.go
│   └── userModel.go
├── router
│   └── router.go
├── go.mod
├── go.sum
├── main.go
└── readme.md
```

## Clone Repositori

```
$ git clone https://github.com/helmyfadlail/image-profile-api.git
$ cd image-profile-api
```

## Instal Dependensi

```
$ go mod tidy
```

## Menjalankan Aplikasi Sekaligus Migrasi Database

```
$ go run main.go
```
