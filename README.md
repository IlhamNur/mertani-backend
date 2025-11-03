# 🌱 Mertani IoT Management Dashboard

Sistem manajemen **Device & Sensor IoT** dengan integrasi multi-service:  
- Backend Go untuk **Device Service** + pengiriman data terjadwal & retry.  
- Backend Java (Quarkus) untuk **Sensor Service**.  
- Frontend berbasis **HTML + Bootstrap + Chart.js** untuk dashboard visualisasi data dan monitoring log pengiriman. 

---

## 🧩 Arsitektur Sistem

```text
┌──────────────────────────┐
│      Frontend Web        │
│  (HTML + Bootstrap + JS) │
│     └─ Chart.js Visuals  │
└────────────┬─────────────┘
             │ REST API
┌────────────┴─────────────┐
│     Device Service       │
│ (Golang + Gin + GORM)    │
│   • CRUD Device          │
│   • Scheduler & Retry    │
│   • Delivery Log API     │
└────────────┬─────────────┘
             │ REST API
┌────────────┴─────────────┐
│     Sensor Service       │
│ (Java + Quarkus + JPA)   │
│   • CRUD Sensor          │
│   • Data per Device      │
└────────────┬─────────────┘
             │
       PostgreSQL Database
```

---

## ⚙️ Fitur Utama

| Modul | Fitur | Teknologi |
|--------|--------|------------|
| **Device Service (Go)** | CRUD Device, Pengiriman data terjadwal, Retry otomatis dengan backoff, Penyimpanan log ke `delivery_logs` | Go, Gin, GORM, cron |
| **Sensor Service (Quarkus)** | CRUD Sensor, relasi ke Device, integrasi REST API | Java, Quarkus, Hibernate ORM |
| **Frontend Dashboard** | Statistik device/sensor, Grafik visualisasi, Monitoring delivery log, Auto-refresh tiap 15 detik | HTML, Bootstrap, Chart.js |
| **Database** | Tabel `devices`, `sensors`, `delivery_logs` | PostgreSQL |
---

## 🚀 Panduan Menjalankan Proyek

### 1️⃣ Jalankan Database  
Pastikan kamu berada di folder `infra/` dan sistem membaca root `.env`.

```bash
cd infra
docker compose --env-file ../.env up -d
```

> Database PostgreSQL akan berjalan di `localhost:5432`  
> Username/password diambil dari file `.env` global.

---

### 2️⃣ Jalankan Device Service (Go)
```bash
cd device-service
go mod tidy
go run main.go
```

Server berjalan di:  
📍 `http://localhost:${DEVICE_PORT}` (default `8080`)

Endpoint penting:
| Endpoint | Method | Deskripsi |
|-----------|---------|------------|
| `/devices` | GET/POST/PUT/DELETE | CRUD Device |
| `/delivery-logs` | GET | Lihat log pengiriman data |

---

### 3️⃣ Jalankan Sensor Service (Quarkus)
```bash
cd sensor-service
./mvnw quarkus:dev
```

Server berjalan di:  
📍 `http://localhost:${SENSOR_PORT}` (default `8082`)

Endpoint penting:
| Endpoint | Method | Deskripsi |
|-----------|---------|------------|
| `/sensors` | GET/POST/PUT/DELETE | CRUD Sensor |

---

### 4️⃣ Jalankan Frontend Dashboard
Masuk ke folder `frontend/`, lalu buka `index.html` dengan **Live Server (VS Code)**  
atau jalankan langsung di browser:

```
file:///path/to/frontend/index.html
```

Dashboard dapat diakses di:  
🌐 `http://localhost:5500/frontend/index.html` (jika via Live Server)

---

## 🧩 File `.env` (Konfigurasi Global)

📄 `/.env`
```env
# --- Database ---
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=admin
DB_NAME=mertani
DB_SSLMODE=disable

# --- Device Service (Go) ---
DEVICE_PORT=8080

# --- Sensor Service (Quarkus) ---
SENSOR_PORT=8082

# --- CORS ---
ALLOWED_ORIGINS=http://localhost:5500,http://127.0.0.1:5500
```

---

## 🧠 Struktur Folder (Terbaru)

```
mertani-backend/
│
├── .env                      # Global environment config untuk semua service
│
├── device-service/           # Backend Go
│   ├── main.go
│   ├── controllers/
│   ├── models/
│   └── workers/
│
├── sensor-service/           # Backend Quarkus (Java)
│   ├── src/main/java/com/mertani/
│   ├── src/main/resources/application.properties
│   └── pom.xml
│
├── frontend/                 # Web Dashboard
│   ├── index.html
│   ├── devices.html
│   ├── sensors.html
│   └── js/
│
└── infra/                    # Docker Compose & Database
    └── docker-compose.yml
```

---

## 💻 Teknologi yang Digunakan

| Layer | Stack |
|-------|-------|
| **Frontend** | HTML5, Bootstrap 5, JavaScript ES6, Chart.js |
| **Backend (Device)** | Go 1.22+, Gin, GORM, robfig/cron, godotenv |
| **Backend (Sensor)** | Java 17+, Quarkus, Hibernate ORM, SmallRye Config |
| **Database** | PostgreSQL 14 |
| **DevOps** | Docker Compose (menggunakan `--env-file ../.env`) |

---

## 🧾 Lisensi
Proyek ini dibuat sebagai bagian dari **Technical Test Backend Developer – Mertani**.  
Seluruh kode bersifat open untuk keperluan evaluasi.

---
