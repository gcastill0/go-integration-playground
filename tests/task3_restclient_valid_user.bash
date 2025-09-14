# tests/task3_restclient_invalid_user.bash
#!/usr/bin/env bash
set -euo pipefail

CLIENT="./tasks/task3-restclient"  # entry path for `go run`
API_URL="https://jsonplaceholder.typicode.com/users" # note the extra `s` on `.com`
USER_ID="2"

out="$(mktemp)"; err="$(mktemp)"
cleanup() { rm -f "$out" "$err"; }
trap cleanup EXIT

# Run with no args; capture exit code + outputs
set +e
go run "$CLIENT" "$API_URL/$USER_ID">"$out" 2>"$err"
code=$?
set -e

last="$(tail -n1 "$err" | tr -d '\r')"
if [[ "$last" =~ exit\ status\ ([0-9]+) ]]; then
  code="${BASH_REMATCH[1]}"
fi

# Assertions
if [[ "$code" -ne 0 ]]; then
  echo "FAIL: expected exit code 0, got $code"
  echo "stderr:"; sed -n '1,200p' "$err" || true
  echo "stdout:"; sed -n '1,200p' "$out" || true
  exit 1
fi

# Stderr should generally be empty on success; warn (don’t fail) if not.
if [[ -s "$err" ]]; then
  echo "NOTE: stderr not empty on success:"
  sed -n '1,200p' "$err" || true
fi

# Expect a JSON body containing known fields for user 2
grep -q '"id": 2' "$out" || { echo 'FAIL: expected `"id": 2` in stdout'; exit 1; }
grep -q '"email": "Shanna@melissa.tv"' "$out" || { echo 'FAIL: expected email in stdout'; exit 1; }
grep -q '"phone": "010-692-6593 x09125"' "$out" || { echo 'FAIL: expected phone in stdout'; exit 1; }
grep -q '"company":' "$out" || { echo 'FAIL: expected company object in stdout'; exit 1; }
grep -q '"name": "Deckow-Crist"' "$out" || { echo 'FAIL: expected company name in stdout'; exit 1; }

echo "ARGS: go run $CLIENT $API_URL/$USER_ID"
echo "PASS: Preflight -> $USER_ID found and exit $code."
echo ""; sed -n '1,200p' "$out" || true
