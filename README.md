# User Service

## Prerequisites

Sebelum memulai, pastikan kamu telah menginstal:

- [Docker Desktop](https://www.docker.com/)
- Git
- [pgAdmin](https://www.pgadmin.org/) (untuk mengelola PostgreSQL)
- PostgreSQL

## Cara Menjalankan Proyek

### 1. Clone Repository

```sh
git clone https://github.com/NDN-RG-Apps-Platform/docker-oplib.git
cd <repository-folder>
```

### 2. Jalankan Docker Compose

```sh
docker compose up -d --build
```

**Opsi:**
- `-d` (detached mode): Menjalankan container di background.
- `--build`: Membangun ulang image sebelum menjalankan container.

### 3. Cek Container yang Berjalan

```sh
docker ps
```

### 4. Menghentikan Container

```sh
docker compose down
```

### 5. Menghapus Volume (Opsional, untuk menghapus data database)

```sh
docker compose down -v
```

## Konfigurasi Database

- **Akses pgAdmin**: [http://localhost:5051](http://localhost:5051)
- **PostgreSQL berjalan di port**: `5432`
- **Pastikan konfigurasi database** di `config.json` sesuai dengan yang ada di `docker-compose.yml`

### Setup pgAdmin

1. Akses pgAdmin melalui browser: [http://localhost:5051](http://localhost:5051)
2. Gunakan kredensial berikut untuk login:
   - **Email**: `admin@pgadmin.com`
   - **Password**: `admin`
3. Setelah login, buat server baru sesuai kebutuhan proyek.

## Menjalankan Aplikasi

Setelah server database dikonfigurasi, jalankan aplikasi dengan perintah berikut:

```sh
make watch-prepare
make watch
```

> **Catatan:** Pastikan server database sudah dibuat sebelum menjalankan perintah di atas.

Untuk informasi lebih lanjut terkait setup, hubungi **Aliza**.

