#!/usr/bin/env bash

set -euo pipefail

REPO="adarshbimal/docklab"
VERSION="v0.1.0"

echo "========================================"
echo "      Installing DockLab"
echo "========================================"

# Check Docker
if ! command -v docker >/dev/null 2>&1; then
    echo
    echo "Docker is not installed."
    echo "Please install Docker first:"
    echo "https://docs.docker.com/engine/install/"
    exit 1
fi

# Detect architecture
ARCH=$(uname -m)

case "$ARCH" in
    x86_64)
        FILE="docklab-linux-amd64"
        ;;
    aarch64|arm64)
        FILE="docklab-linux-arm64"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

URL="https://github.com/$REPO/releases/download/$VERSION/$FILE"

echo "Downloading DockLab..."
curl -L "$URL" -o docklab

chmod +x docklab

echo "Installing to /usr/local/bin..."
sudo mv docklab /usr/local/bin/docklab
sudo install -Dm755 docklab /usr/local/bin/docklab
echo
echo "Installation complete!"
echo

docklab --help

echo
echo "Try:"
echo "  docklab up"