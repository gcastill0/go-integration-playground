# tests/task3_restclient_invalid_url.bash
#!/usr/bin/env bash
set -euo pipefail

CLIENT="./tasks/task3-restclient"  # entry path for `go run`
BAD_URL="https://jsonplaceholder.typicode.coms/users/2" # note the extra `s` on `.com`

out="$(mktemp)"; err="$(mktemp)"
cleanup() { rm -f "$out" "$err"; }
trap cleanup EXIT

# Run with no args; capture exit code + outputs
set +e
go run "$CLIENT" "$BAD_URL">"$out" 2>"$err"
code=$?
set -e

last="$(tail -n1 "$err" | tr -d '\r')"
if [[ "$last" =~ exit\ status\ ([0-9]+) ]]; then
  code="${BASH_REMATCH[1]}"
fi

# Assertions
if [[ "$code" -ne 3 ]]; then
  echo "FAIL: expected exit code 3, got $code"
  echo "stderr:"; sed -n '1,200p' "$err" || true
  echo "stdout:"; sed -n '1,200p' "$out" || true
  exit 1
fi

if ! grep -qi 'Preflight:' "$err" || ! grep -q 'Get' "$err"; then
  echo "FAIL: expected Preflight message with Get on stderr"
  echo "stderr:"; sed -n '1,200p' "$err" || true
  exit 1
fi

if [[ -s "$out" ]]; then
  echo "FAIL: expected empty stdout for usage case"
  echo "stdout:"; sed -n '1,200p' "$out" || true
  exit 1
fi

echo "ARGS: go run $CLIENT $BAD_URL"
echo "PASS: Preflight -> no such host on stderr and exit $code."
echo ""; sed -n '1,200p' "$err" || true
