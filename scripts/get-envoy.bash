#!/bin/bash
set -euo pipefail

PATH="$PATH:$(go env GOPATH)/bin"
export PATH

_project_root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)/.."
_envoy_version=1.24.0
_dir="$_project_root/pkg/envoy/files"

for _target in linux-arm64; do

  #container_id=$(podman create "docker.io/thegrandpkizzle/envoy:$_envoy_version")
  #podman cp "$container_id:/usr/local/bin/envoy" "$_dir/envoy-$_target"
  #podman rm "$container_id"

  #curl \
  #  --silent \
  #  --fail \
  #  --show-error \
  #  --compressed \
  #  --location \
  #  --time-cond "$_dir/envoy-$_target.sha256" \
  #  --output "$_dir/envoy-$_target.sha256" \
  #  "$_url.sha256" &

  echo "$_envoy_version" >"$_dir/envoy-$_target.version"
done

wait
