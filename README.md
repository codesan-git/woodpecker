# Woodpecker CI + GitHub Actions 🐦 + 🐙

**CI/CD otomatis penuh!** GitHub Actions CI, Woodpecker CD — tanpa registry external.

## Arsitektur (Production)

```
 Push ke main
      │
      ├──→ 🐙 GitHub Actions (cloud)      : build + test (otomatis)
      │
      └──→ 🐦 Woodpecker (VPS/public)    : terima webhook otomatis
                                            → clone repo
                                            → docker build
                                            → restart container
                                            → health check
```

> Semua otomatis — nggak perlu klik apa-apa. Push → CI jalan + CD jalan.

---

## 🚀 Setup Production (VPS)

### 1. Siapkan VPS

- Ubuntu 22.04 (atau yang baru)
- Docker + Docker Compose terinstall
- Domain (misal: `ci.namakamu.com`) pointing ke IP VPS

### 2. Setup Nginx + SSL

```bash
# Install Nginx & Certbot
apt install nginx certbot python3-certbot-nginx

# Dapatkan SSL
certbot --nginx -d ci.namakamu.com
```

Konfigurasi Nginx (`/etc/nginx/sites-available/woodpecker`):
```nginx
server {
    listen 443 ssl;
    server_name ci.namakamu.com;

    location / {
        proxy_pass http://127.0.0.1:8000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

### 3. GitHub OAuth App

Buka [github.com/settings/developers](https://github.com/settings/developers) → New OAuth App:

| Field | Value |
|---|---|
| App name | Woodpecker |
| Homepage URL | `https://ci.namakamu.com` |
| Callback URL | `https://ci.namakamu.com/authorize` |

Copy Client ID & Client Secret.

### 4. Deploy Woodpecker

```bash
# Clone project ini ke VPS
git clone git@github.com:USERNAME/woodpecker-demo.git
cd woodpecker-demo

# Isi .env
vim .env
```

```env
WOODPECKER_ADMIN=github-username-kamu
WOODPECKER_HOST=https://ci.namakamu.com
WOODPECKER_GITHUB_CLIENT=Ov23li...
WOODPECKER_GITHUB_SECRET=...
WOODPECKER_AGENT_SECRET=rahasia-agent-woodpecker-2024
```

```bash
docker compose up -d
# → https://ci.namakamu.com
```

### 5. Aktifkan Repo

1. Buka `https://ci.namakamu.com` → login GitHub
2. Repositories → cari `woodpecker-demo` → **Activate**
3. Woodpecker otomatis pasang webhook di repo GitHub

### 6. Selesai!

Tiap `git push` ke main:
- ✅ GitHub Actions jalanin build + test
- ✅ Woodpecker otomatis deploy via webhook

---

## 📋 Pipeline Detail

### GitHub Actions (`.github/workflows/ci.yml`)
```
📦 Build → 🧪 Test → 🔍 Lint
```

### Woodpecker (`.woodpecker.yml`)
```
📥 Clone repo → 🐳 Build image → 🔄 Restart container → ❤️ Health check
```

> Tidak perlu Docker Hub / registry. Woodpecker build image langsung di VPS dari source code.

---

## 🔧 Test Lokal (Development)

Buat coba-coba di laptop:

```bash
# Ganti WOODPECKER_HOST di .env jadi:
WOODPECKER_HOST=http://localhost:8000

docker compose up -d
# → http://localhost:8000
# Trigger pipeline manual dari UI (karena localhost nggak bisa terima webhook)
```

---

## 🧹 Cleanup

```bash
docker compose down -v
```
