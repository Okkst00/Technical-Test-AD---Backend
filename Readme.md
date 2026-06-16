# Backend API (Go) - Layered Architecture

Project ini dibangun menggunakan Golang dengan pendekatan **Layered Architecture** untuk menjaga codebase tetap rapi, modular, dan mudah dikembangkan.

## Struktur Folder

├── config
├── handler
├── helper
├── middleware
├── model
├── repository
├── routes
├── service
└── utils

## Penjelasan Layer

- **config**
  Layer konfigurasi aplikasi seperti database, environment, dan inisialisasi service.

- **handler (Presentation Layer)**
  Bertanggung jawab menerima HTTP request dan mengembalikan response. Tidak berisi business logic.

- **service (Business Layer)**
  Tempat utama business logic berjalan. Mengatur alur proses aplikasi.

- **repository (Data Layer)**
  Bertanggung jawab untuk akses database (query, CRUD, transaction).

- **model**
  Struktur data (struct) dan representasi entity yang digunakan di aplikasi.

- **routes**
  Mengatur routing endpoint dan mapping ke handler.

- **middleware**
  Layer perantara request (auth, logging, validation, dll).

- **helper**
  Fungsi bantu yang reusable di berbagai layer.

- **utils**
  Utility umum seperti hashing, formatting, encryption, dll.

## Arsitektur Flow

Client → Routes → Middleware → Handler → Service → Repository → Database

## Tujuan

Menggunakan Layered Architecture agar setiap layer memiliki tanggung jawab yang jelas (Separation of Concerns), sehingga:

- kode lebih mudah di-maintain
- mudah di-scale
- lebih mudah di-test

## Tech Stack

- Golang
- MySQL
- REST API
