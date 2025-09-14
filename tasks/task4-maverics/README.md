```bash
./maverics_darwin_arm64 -config ./maverics.json
ts=2025-09-14T17:50:19.005377Z msg="initializing config filesystem to OS filesystem"
ts=2025-09-14T17:50:19.006212Z msg="initializing telemetry service" orchestratorID=1eba7028-ab69-4882-8e3a-ca3ff3b97b8c
ts=2025-09-14T17:50:19.00727Z level=info msg="starting Maverics" version=2025.09.3 date=2025-09-11T18:11:29Z
ts=2025-09-14T17:50:19.007322Z level=info msg="usage of the Maverics Identity Orchestrator is covered by the following terms and conditions https://www.strata.io/legal/enterprise-master-license-subscription-agreement/"
ts=2025-09-14T17:50:19.007349Z level=info msg="loading configuration from filesystem"
ts=2025-09-14T17:50:19.007421Z level=info msg="loaded config './maverics.json' from filesystem"
ts=2025-09-14T17:50:19.007427Z level=info msg="successfully loaded configuration"
ts=2025-09-14T17:50:19.007589Z level=info msg="successfully parsed configuration" version=1
ts=2025-09-14T17:50:19.008419Z level=info msg="starting service" service="HTTP Observability" stability=Stable
ts=2025-09-14T17:50:19.008437Z level=info msg="starting service" service="Session Manager" stability=Stable
ts=2025-09-14T17:50:19.00844Z level=info msg="starting service" service="Session Manager" stability=Stable service="Session Store (Using Cache: InMemory Cache (Memory Bound))" stability=Stable
```

```bash
go run ./tasks/task4-maverics-ext
tasks/task4-maverics-ext/main.go:12:2: no required module provides package github.com/strata-io/service-extension/orchestrator; to add it:
        go get github.com/strata-io/service-extension/orchestrator
```

```bash
go get github.com/strata-io/service-extension/orchestrator

go: downloading github.com/strata-io/service-extension v0.23.0
go: added github.com/strata-io/service-extension v0.23.0
```