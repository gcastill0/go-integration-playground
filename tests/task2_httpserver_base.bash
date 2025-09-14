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
  curl -vvsSf "http://127.0.0.1:${PORT}/" >/dev/null && break
  sleep 1
done

# fetch body and status
BODY="$(curl -sS "http://127.0.0.1:${PORT}/")"
STATUS="$(curl -sS -o /dev/null -w '%{http_code}' "http://127.0.0.1:${PORT}/")"

# assertions
[ "${STATUS}" = "200" ] || { echo "expected 200, got ${STATUS}"; exit 1; }
if [ "${BODY}" != "Hello, world" ] && [ "${BODY}" != $'Hello, world\n' ]; then
  echo "unexpected body: ${BODY@Q}"
  exit 1
fi

echo
echo "OK: / returned ${STATUS} and response \"${BODY}\""
