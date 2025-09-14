# tests/task3_restclient_no_args.bash
#!/usr/bin/env bash
set -euo pipefail

CLIENT="./tasks/task3-restclient"  # entry path for `go run`

out="$(mktemp)"; err="$(mktemp)"
cleanup() { rm -f "$out" "$err"; }
trap cleanup EXIT

# Run with no args; capture exit code + outputs
set +e
go run "$CLIENT" >"$out" 2>"$err"
code=$?
set -e

last="$(tail -n1 "$err" | tr -d '\r')"
if [[ "$last" =~ exit\ status\ ([0-9]+) ]]; then
  code="${BASH_REMATCH[1]}"
fi

# Assertions
if [[ "$code" -ne 2 ]]; then
  echo "FAIL: expected exit code 2, got $code"
  echo "stderr:"; sed -n '1,200p' "$err" || true
  echo "stdout:"; sed -n '1,200p' "$out" || true
  exit 1
fi

if ! grep -qi 'usage:' "$err" || ! grep -q '<URL>' "$err"; then
  echo "FAIL: expected usage message with <URL> on stderr"
  echo "stderr:"; sed -n '1,200p' "$err" || true
  exit 1
fi

if [[ -s "$out" ]]; then
  echo "FAIL: expected empty stdout for usage case"
  echo "stdout:"; sed -n '1,200p' "$out" || true
  exit 1
fi

echo "ARGS: go run $CLIENT"
echo "PASS: no-args -> usage on stderr and exit $code."
echo ""; sed -n '1,200p' "$err" || true
