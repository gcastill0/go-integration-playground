# tests/task2_httpserver_01.bash
#!/usr/bin/env bash
set -euo pipefail

PORT="${PORT:-8080}"
RUN_TARGET="./tasks/task2-httpserver" 

# start server
go run "${RUN_TARGET}" >/dev/null 2>&1 &
PID=$!
trap 'kill ${PID} 2>/dev/null || true' EXIT

# synthetic delay waiting for server up
sleep 5

# tiny readiness wait
for _ in {1..30}; do
  curl -vvsSf "http://127.0.0.1:${PORT}/ping" >/dev/null && break
  sleep 1
done

# fetch body and status
BODY="$(curl -sS "http://127.0.0.1:${PORT}/ping")"
STATUS="$(curl -sS -o /dev/null -w '%{http_code}' "http://127.0.0.1:${PORT}/ping")"

# assertions
[ "${STATUS}" = "200" ] || { echo "expected 200, got ${STATUS}"; exit 1; }
if [ "${BODY}" != '{"message":"pong"}' ]; then
  echo "unexpected body: ${BODY@Q}"
  exit 1
fi

echo
echo "OK: /ping returned 200 and response ${BODY}"
