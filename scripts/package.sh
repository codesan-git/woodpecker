#!/usr/bin/env bash
# ─── Build & Package Woodpecker Images ──────────────
# Jalankan di LAPTOP.
# Output: woodpecker-images.tar + docker-compose.yml + .env
set -euo pipefail

WOODPECKER_VERSION="${WOODPECKER_VERSION:-v3}"

echo "📥 Pulling images..."
docker pull "woodpeckerci/woodpecker-server:${WOODPECKER_VERSION}"
docker pull "woodpeckerci/woodpecker-agent:${WOODPECKER_VERSION}"

echo "📦 Saving to woodpecker-images.tar..."
docker save -o woodpecker-images.tar \
    "woodpeckerci/woodpecker-server:${WOODPECKER_VERSION}" \
    "woodpeckerci/woodpecker-agent:${WOODPECKER_VERSION}"

ls -lh woodpecker-images.tar

echo ""
echo "✅ Siap kirim ke VPS:"
echo "   scp woodpecker-images.tar docker-compose.yml .env user@vps:/opt/woodpecker/"
echo ""
echo "Lalu di VPS jalankan:"
echo "   cd /opt/woodpecker"
echo "   docker load -i woodpecker-images.tar"
echo "   docker compose up -d"
