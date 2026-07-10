#!/usr/bin/env bash
# ─── Load & Run Woodpecker on VPS ──────────────────
# Jalankan di VPS setelah file dikirim dari laptop.
set -euo pipefail

echo "📥 Loading images..."
docker load -i woodpecker-images.tar

echo "🚀 Starting Woodpecker..."
docker compose up -d

echo ""
echo "✅ Cek status:"
docker compose ps

echo ""
echo "🌐 Akses:"
grep WOODPECKER_HOST .env 2>/dev/null | cut -d= -f2 || echo "  (cek .env untuk URL)"
