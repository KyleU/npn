#!/bin/bash
# Content managed by Project Forge, see [projectforge.md] for details.

## Builds the app (or just use make build)

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$dir/../.."

os=${1:-darwin}
arch=${2:-amd64}
fn=${3:-npn}

echo "Building [$os $arch]..."
env GOOS=$os GOARCH=$arch make build-release
mkdir -p "./build/$os/$arch"
mv "./build/release/$fn" "./build/$os/$arch/$fn"
