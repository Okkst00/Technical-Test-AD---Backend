## Cara Setup di Local (Step-by-Step)

Berikut langkah-langkah untuk menjalankan project ini di local development:

1. Clone Repository
   git clone https://github.com/Okkst00/Technical-Test-AD---Backend.git
   cd project-name

2. Install Dependency
   Pastikan Go sudah terinstall di device kamu, lalu jalankan:
   go mod tidy

3. Setup Environment
   Buat file .env di root project, lalu isi sesuai konfigurasi berikut
   LOCAL_DB=root:@tcp(127.0.0.1:3306)/shipcommerce_id?parseTime=true
   PAYMENT_SECRET=981231231AGHHA

4. Jalankan Aplikasi
   go run main.go

5. Test API
   Default server biasanya berjalan di:
   http://localhost:8080

## Reason Choose Go

Go dipilih karena kombinasi antara **performance, simplicity, concurrency, dan kemudahan deployment**, yang sangat cocok untuk backend API modern dan scalable.

# Backend API (Go) - Layered Architecture

Project ini dibangun menggunakan Golang dengan pendekatan **Layered Architecture** untuk menjaga codebase tetap rapi, modular, dan mudah dikembangkan.

# Standard Response API, masih sempat diterapkan di api product

    "success": true,
    "message": "success",
    "data": [
        {}
    ]

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

Menggunakan Layered Architecture agar setiap layer memiliki tanggung jawab yang jelas (Separation of Concerns) dengan waktu pengerjaan yang terbatas, sehingga:

- kode lebih mudah di-maintain
- mudah di-scale
- lebih mudah di-test

## Tech Stack

- Golang
- MySQL
- REST API

# Fitur yang selesai

- Manajemen Auth & Regist - JWT + Hash
- Manajemen Produk (belum include search produk)
- Membuat Pesanan
- Manajemen Membership Benefit
- Manajemen Webhook Confirm Sederhana -> Menggunakan Tombol

# Fitur yang belum selesai

- Test Race Condition
- Indempotency
- Webhook
  -- Saat klik 'Bayar', random outcome: 80% success, 20% failed (untuk testing)
  -- System terima webhook callback dari payment gateway
- Manajemen informasi ketika under stock dan gagal transaction prosess (create order, auth, manajemen produk)
- manajemen view order antara admin dan member

<img width="1321" height="596" alt="Techincal_Flow_Dropship" src="https://github.com/user-attachments/assets/71c205c3-3dd6-4610-9fae-01d7550eaff9" />
