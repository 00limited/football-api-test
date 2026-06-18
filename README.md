# ⚽ Football Team Management API

REST API backend untuk manajemen tim sepakbola amatir, dibangun menggunakan **Go + GIN Framework**.

> Submission untuk **AYO Software Developer Technical Test 2026**

---

## 🛠️ Tech Stack

| Komponen   | Teknologi                   |
| ---------- | --------------------------- |
| Language   | Go 1.21+                    |
| Framework  | GIN v1.9+                   |
| Database   | PostgreSQL 15+              |
| ORM        | GORM v2 (soft delete)       |
| Auth       | JWT (golang-jwt/jwt v5)     |
| Password   | bcrypt                      |
| Validation | go-playground/validator v10 |
| Config     | godotenv                    |

---

## 📁 Struktur Project

```
football-api/
├── cmd/
│   └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── middleware/
│   │   └── auth.go
│   ├── models/
│   │   ├── admin.go
│   │   ├── team.go
│   │   ├── player.go
│   │   ├── match.go
│   │   ├── match_result.go
│   │   └── goal.go
│   ├── handlers/
│   │   ├── auth_handler.go
│   │   ├── team_handler.go
│   │   ├── player_handler.go
│   │   ├── match_handler.go
│   │   ├── result_handler.go
│   │   └── report_handler.go
│   ├── services/
│   │   ├── team_service.go
│   │   ├── player_service.go
│   │   ├── match_service.go
│   │   ├── result_service.go
│   │   └── report_service.go
│   ├── repositories/
│   │   ├── team_repository.go
│   │   ├── player_repository.go
│   │   ├── match_repository.go
│   │   ├── result_repository.go
│   │   └── report_repository.go
│   ├── dto/
│   │   ├── request/
│   │   └── response/
│   └── router/
│       └── router.go
├── storage/
│   └── logos/
├── .env.example
├── go.mod
├── go.sum
└── README.md
```

---

## ⚙️ Setup & Instalasi

### Prasyarat

- Go 1.21+
- PostgreSQL 15+
- Git

### 1. Clone Repository

```bash
git clone https://github.com/00limited/football-api-test.git
cd football-api-test
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Konfigurasi Environment

```bash
cp .env.example .env
```

Edit file `.env` sesuaikan dengan konfigurasi lokal Anda:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=football_db

JWT_SECRET=your-super-secret-key-minimum-32-chars
JWT_EXPIRY_HOURS=24

STORAGE_PATH=./storage/logos
APP_PORT=8080
```

### 4. Buat Database

```bash
createdb football_db
```

> Migrasi tabel berjalan **otomatis** saat aplikasi pertama kali dijalankan (GORM AutoMigrate).

### 5. Jalankan Aplikasi

```bash
go run cmd/main.go
```

Aplikasi berjalan di: `http://localhost:8080`

---

## 🔐 Autentikasi

API menggunakan **JWT Bearer Token**.

1. Register admin baru via `POST /api/v1/auth/register`
2. Login via `POST /api/v1/auth/login` → dapatkan token
3. Sertakan token di setiap request header:

```
Authorization: Bearer <token>
```

---

## 📡 API Endpoints

Base URL: `http://localhost:8080/api/v1`

### Auth

| Method | Endpoint         | Auth | Keterangan             |
| ------ | ---------------- | ---- | ---------------------- |
| POST   | `/auth/register` | ❌   | Register admin         |
| POST   | `/auth/login`    | ❌   | Login, dapat JWT token |

### Teams

| Method | Endpoint          | Auth | Keterangan              |
| ------ | ----------------- | ---- | ----------------------- |
| GET    | `/teams`          | ✅   | List semua tim          |
| POST   | `/teams`          | ✅   | Tambah tim baru         |
| GET    | `/teams/:id`      | ✅   | Detail tim              |
| PUT    | `/teams/:id`      | ✅   | Update tim              |
| DELETE | `/teams/:id`      | ✅   | Hapus tim (soft delete) |
| POST   | `/teams/:id/logo` | ✅   | Upload logo tim         |

### Players

| Method | Endpoint                 | Auth | Keterangan                 |
| ------ | ------------------------ | ---- | -------------------------- |
| GET    | `/teams/:teamId/players` | ✅   | List pemain dalam tim      |
| POST   | `/teams/:teamId/players` | ✅   | Tambah pemain              |
| GET    | `/players/:id`           | ✅   | Detail pemain              |
| PUT    | `/players/:id`           | ✅   | Update pemain              |
| DELETE | `/players/:id`           | ✅   | Hapus pemain (soft delete) |

### Matches

| Method | Endpoint       | Auth | Keterangan                 |
| ------ | -------------- | ---- | -------------------------- |
| GET    | `/matches`     | ✅   | List semua jadwal          |
| POST   | `/matches`     | ✅   | Buat jadwal baru           |
| GET    | `/matches/:id` | ✅   | Detail jadwal              |
| PUT    | `/matches/:id` | ✅   | Update jadwal              |
| DELETE | `/matches/:id` | ✅   | Hapus jadwal (soft delete) |

### Match Results

| Method | Endpoint              | Auth | Keterangan                     |
| ------ | --------------------- | ---- | ------------------------------ |
| POST   | `/matches/:id/result` | ✅   | Input hasil & gol pertandingan |
| GET    | `/matches/:id/result` | ✅   | Lihat hasil pertandingan       |

### Reports

| Method | Endpoint               | Auth | Keterangan                       |
| ------ | ---------------------- | ---- | -------------------------------- |
| GET    | `/reports/matches`     | ✅   | Laporan semua pertandingan       |
| GET    | `/reports/matches/:id` | ✅   | Laporan detail satu pertandingan |

---

## 📋 Contoh Request & Response

### Login

```bash
POST /api/v1/auth/login
Content-Type: application/json

{
  "username": "admin_xyz",
  "password": "P@ssword123"
}
```

```json
{
  "status": "success",
  "message": "Login berhasil",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIs...",
    "expires_at": "2026-07-17T10:00:00Z"
  }
}
```

### Tambah Pemain

```bash
POST /api/v1/teams/1/players
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "Budi Santoso",
  "height_cm": 175.5,
  "weight_kg": 68.0,
  "position": "MIDFIELDER",
  "jersey_number": 10
}
```

> Nilai `position` yang valid: `FORWARD` | `MIDFIELDER` | `DEFENDER` | `GOALKEEPER`

### Input Hasil Pertandingan

```bash
POST /api/v1/matches/1/result
Authorization: Bearer <token>
Content-Type: application/json

{
  "goals": [
    { "player_id": 5, "team_id": 1, "goal_minute": 23 },
    { "player_id": 5, "team_id": 1, "goal_minute": 67 },
    { "player_id": 12, "team_id": 2, "goal_minute": 45 }
  ]
}
```

### Report Detail Pertandingan

```json
{
  "status": "success",
  "data": {
    "match": { "id": 1, "match_date": "2026-07-20", "match_time": "15:30:00" },
    "home_team": { "id": 1, "name": "Persebaya Surabaya" },
    "away_team": { "id": 2, "name": "Persib Bandung" },
    "home_score": 2,
    "away_score": 1,
    "match_status": "HOME_WIN",
    "top_scorer": { "player_id": 5, "name": "Budi Santoso", "total_goals": 2 },
    "home_team_total_wins": 7,
    "away_team_total_wins": 4,
    "goals": [
      { "player_name": "Budi Santoso", "team_name": "Persebaya", "minute": 23 },
      { "player_name": "Andi Wijaya", "team_name": "Persib", "minute": 45 },
      { "player_name": "Budi Santoso", "team_name": "Persebaya", "minute": 67 }
    ]
  }
}
```

> Nilai `match_status`: `HOME_WIN` | `AWAY_WIN` | `DRAW`

---

## ✅ Business Rules

- Semua penghapusan menggunakan **Soft Delete** (`deleted_at`)
- Nomor punggung pemain **unik per tim**
- Tim home dan tim away **tidak boleh sama** dalam satu pertandingan
- Hasil pertandingan hanya bisa diinput **sekali** (status: `SCHEDULED` → `FINISHED`)
- Pencetak gol harus merupakan **pemain dari salah satu tim** yang bertanding
- Skor dihitung **otomatis** dari jumlah data gol, bukan input manual
- `home_team_total_wins` dan `away_team_total_wins` adalah akumulasi dari **seluruh pertandingan FINISHED** yang melibatkan tim tersebut

---

## 📮 Postman Collection

Import file berikut ke Postman untuk mencoba semua endpoint:

📎 [`Football_API.postman_collection.json`](./Football_API.postman_collection.json)

> Set environment variable `base_url = http://localhost:8080` dan `token = <hasil login>` di Postman.

---

## 🗄️ Skema Database

```
admins          → id, username, password, timestamps
teams           → id, name, logo_url, founded_year, address, city, timestamps
players         → id, team_id*, name, height_cm, weight_kg, position, jersey_number, timestamps
matches         → id, match_date, match_time, home_team_id*, away_team_id*, status, timestamps
match_results   → id, match_id* (unique), home_score, away_score, timestamps
goals           → id, match_result_id*, player_id*, team_id*, goal_minute, timestamps

* = Foreign Key | timestamps = created_at, updated_at, deleted_at
```

---

<p align="center">Made with ☕ for AYO Technical Test 2026</p>
