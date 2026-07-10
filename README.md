# Woodpecker CI Demo 🐦

CI/CD pipeline dengan **Woodpecker CI** + **GitHub OAuth** + **Docker Build & Push**.

## 🗂 Struktur Proyek

```
woodpecker/
├── docker-compose.yml      # Woodpecker Server + Agent
├── .env                     # GitHub OAuth credentials (kamu isi sendiri)
├── Dockerfile               # Multi-stage build (golang → scratch, ~3 MB)
├── .dockerignore
├── .woodpecker.yml          # Pipeline CI: build → test → lint → docker push
├── main.go                  # Go HTTP server + frontend
├── main_test.go             # Unit test
├── go.mod / go.sum
├── static/
│   └── index.html           # Frontend halaman status
└── README.md
```

## 🚀 3 Langkah Setup

### 1️⃣ Bikin OAuth App di GitHub (2 menit)

1. Buka [github.com/settings/developers](https://github.com/settings/developers)
2. Klik **"New OAuth App"**
3. Isi form:
   | Field                          | Value                                |
   |--------------------------------|--------------------------------------|
   | Application name               | `Woodpecker Local`                   |
   | Homepage URL                   | `http://localhost:8000`              |
   | Authorization callback URL     | `http://localhost:8000/authorize`    |
4. Klik **"Register application"**
5. Klik **"Generate a new client secret"**
6. **Copy Client ID & Client Secret**

### 2️⃣ Isi File `.env`

Edit `.env` dengan credential dari GitHub:

```env
WOODPECKER_ADMIN=github-username-kamu
WOODPECKER_GITHUB_CLIENT=Ov23li...          # ← paste Client ID
WOODPECKER_GITHUB_SECRET=...                # ← paste Client Secret
WOODPECKER_AGENT_SECRET=rahasia-agent-woodpecker-2024
```

### 3️⃣ Jalankan Woodpecker

```bash
docker compose up -d
```

Buka **http://localhost:8000** → login pakai akun GitHub kamu.

---

## 🧪 Coba Aplikasi (Tanpa CI)

```bash
# Jalankan Go server (frontend + API)
go run main.go
# → Buka http://localhost:8080
```

```json
GET /health
{
  "status": "ok",
  "timestamp": "2026-01-01T00:00:00Z",
  "version": "dev"
}
```

## 🐳 Build Docker Image (Manual)

```bash
# Build & run
docker build -t woodpecker-demo .
docker run -d -p 8080:8080 --name demo woodpecker-demo

# Image size hanya ~3 MB (from scratch!)
docker images woodpecker-demo
```

## 📋 Pipeline Workflow

```
┌──────────┐   ┌──────────┐   ┌──────────┐   ┌─────────────────┐
│  Build   │→  │   Test   │→  │   Lint   │→  │ Docker Publish  │
│ go build │   │ go test  │   │ go vet   │   │ build + push 🐳 │
└──────────┘   └──────────┘   └──────────┘   └─────────────────┘
   selalu         selalu         selalu        hanya push main
                                              & manual trigger
```

## 🔐 Setup Secrets di Woodpecker

Untuk step **Docker push**, tambahkan secrets di Woodpecker UI:

1. Buka **http://localhost:8000** → pilih repo
2. **Settings → Secrets**
3. Tambah:

| Nama Secret       | Value                             |
|-------------------|-----------------------------------|
| `docker_username` | Username Docker Hub kamu          |
| `docker_password` | [Docker Access Token](https://hub.docker.com/settings/security) |

Lalu edit `repo` di `.woodpecker.yml`:

```yaml
- name: docker-publish
  image: woodpeckerci/plugin-docker-buildx
  settings:
    repo: namakamu/woodpecker-demo    # ← ganti ini
```

---

## 🔧 Trigger Pipeline

Setelah repo terhubung di Woodpecker:

```bash
# Via CLI
brew install woodpecker-cli
export WOODPECKER_SERVER=http://localhost:8000
export WOODPECKER_TOKEN=token-kamu
woodpecker-cli pipeline start woodpecker-demo

# Atau klik "Run Pipeline" di Web UI
```

---

## 🧹 Cleanup

```bash
docker compose down         # stop semua
docker compose down -v      # stop + hapus semua data
```
