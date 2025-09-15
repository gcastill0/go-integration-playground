# Go Integration Playground – Assessment Project

This repository contains the implementation of the **Go Skills and Integration Assessment**.  
The project documents solutions with Go (Golang) by solving four practical integration tasks:

## Table of Contents

1. [Task 1 – JSON Validation](#task-1--json-validation)  
2. [Task 2 – HTTP Server](#task-2--http-server)  
3. [Task 3 – REST API Client](#task-3--rest-api-client)  
4. [Task 4 – Maverics Service Extension](#task-4--maverics-service-extension)  
5. [Final Notes](#final-notes)


Each section below explains the task, shows how to run and test the solution, provides sample results, and closes with a short conclusion.

Quick estimates for reading:

| Reading mode           | Estimate        |
| ---------------------- | --------------- |
| Fast skim              | \~6–8 minutes   |
| Typical technical read | \~8–15 minutes  |
| Careful read           | \~15–20 minutes |

Each exercise in this project is presented with both a summary and a conclusion for quick understanding. For readers who want more depth, each task also includes a very detailed explanation with code snippets, analysis, and references to files in the repository. This structure makes it easy to skim quickly or dive deeper as needed.

For questions or feedback, please contact me using [this form](https://gcastill0.github.io/#contact).

<br>

## Prerequisites

- **Go** ≥ 1.22 (developed and tested with Go 1.25.1)  
- **Git** for cloning the repository  
- Internet connection (required for Task 3 and Task 4 testing)  

Install dependencies:

```bash
git clone https://github.com/gcastill0/go-integration-playground.git
cd go-integration-playground
```

Each task is located under the `/tasks` directory. You can run them individually.

<br>

## Task 1 – JSON Validation

### Summary
The goal of this task is to treat a malformed JSON document, explain errors in terms of the JSON grammar, and produce a corrected, well-formed version. We take a layered, testable approach: first validate syntax, then show evidence.

### Analysis

The document presents an array of three objects with two syntax errors in the JSON grammar which affects the document structure. 

1. After the data array opening with a left bracket (`[`) symbol, the first object is missing a curly brace (`{`) before the first `"id"` key. 

    The snippet shown below shows the first object in the original file - [See lines 1–24 of jdoc_original.json](https://github.com/gcastill0/go-integration-playground/blob/main/tasks/task1-jsonfix/data/jdoc_original.json#L1-L24). Note specifically that the first object is missing the opening handlebar (`{`) which should be located in line 2.

    ```json
    01  [
    02    
    03      "id": 1,
    04      "name": "Leanne Graham",
    05      "username": "Bret",
    06      "email": "Sincere@april.biz",
    07      "address": {
    08          "street": "Kulas Light",
    09          "suite": "Apt. 556",
    10          "city": "Gwenborough",
    11          "zipcode": "92998-3874",
    12          "geo": {
    13              "lat": "-37.3159",
    14              "lng": "81.1496"
    15          }
    16      },
    17      "phone": "1-770-736-8031 x56442",
    18      "website": "hildegard.org",
    19      "company": {
    20          "name": "Romaguera-Crona",
    21          "catchPhrase": "Multi-layered client-server neural-net",
    22            "bs": "harness real-time e-markets"
    23        }
    24    },
    ```

<br>

2. The data structure never closes the array with a right bracket (`]`) symbol after the third object; the file closes on the last curly brace (`}`) which belongs to the last object.

    The following snippet shows the last object in the dictionary from original file - [See lines 48–71 of jdoc_original.json](https://github.com/gcastill0/go-integration-playground/blob/main/tasks/task1-jsonfix/data/jdoc_original.json#L48-L71). Note specifically that the array is missing the closing bracker (`]`) which should be located in line 71.

    ```json
    48    {
    49    "id": 3,
    50   "name": "Clementine Bauch",
    51    "username": "Samantha",
    52    "email": "Nathan@yesenia.net",
    53    "address": {
    54      "street": "Douglas Extension",
    55        "suite": "Suite 847",
    56        "city": "McKenziehaven",
    57        "zipcode": "59590-4157",
    58        "geo": {
    59            "lat": "-68.6102",
    60            "lng": "-47.0653"
    61        }
    62    },
    63    "phone": "1-463-123-4447",
    64    "website": "ramiro.info",
    65    "company": {
    66        "name": "Romaguera-Jacobson",
    67        "catchPhrase": "Face to face bifurcated interface",
    68        "bs": "e-enable strategic applications"
    69    }
    70  }
    71
    ```

<br>

### Solution

There are two minimal fixes: 

1. insert a left handle bar, or left curly brace (`{`) on line 2, right after the opening left bracket (`[`) on line 1, and 

2. add the closing right bracker, (`]`), at the end of the array on line 71.


### Conclusion
This exercise confirmed the ability to recognize structural problems in JSON documents and fix them for use in APIs or data pipelines.

<br>

## Task 2 – HTTP Server

### Summary

A simple HTTP server was created that listens on **port 8080** and supports two endpoints:

- `GET /ping` → returns `{"message":"pong"}`
- `POST /echo` → echoes the posted JSON payload

### How to Test

Start the server with the following command:

```bash
go run ./tasks/task2-httpserver
```

1. To evaluate the `GET` verb against the `/ping` endpoint, you use the following `curl` command in a second terminal:

    ```bash
    curl -ivf http://127.0.0.1:8080/ping
    ```

    The expected response shoud be the following string:

    ```json
    {"message":"pong"}
    ```


1. To evaluate the `POST` verb against the `/echo` endpoint, you use the following `curl` command in a second terminal:

    ```bash
    curl -X POST http://localhost:8080/echo \
    -H "Content-Type: application/json" \
    -d '{"name":"test"}'
    ```

    The expected response shoud be the following string:

    ```
    {"name":"test"}
    ```

### Conclusion

This task demonstrates basic Go web server development, JSON handling, and request/response processing.

<br>

## Task 3 – REST API Client

### Summary
A Go client was built to fetch user details from the [JSONPlaceholder API](https://jsonplaceholder.typicode.com/users/2). The client retrieves a user by ID, unmarshals the JSON, and prints selected fields.

### How to Test
Run the client:

If not done yet, you need to install the Maverics service extension as follows:

```bash
go get github.com/strata-io/service-extension/orchestrator
```

To run the REST API Client, use the following command:

```bash
go run ./tasks/task3-restclient \
https://jsonplaceholder.typicode.com/users/2
```

Expected Result
```json
{
  "id": 2,
  "email": "Shanna@melissa.tv",
  "phone": "010-692-6593 x09125",
  "company": {
    "name": "Deckow-Crist",
    "catchPhrase": "Proactive didactic contingency",
    "bs": "synergize scalable supply-chains"
  }
}
```

### Conclusion
This confirms the ability to consume REST APIs, parse JSON, and handle structured data in Go.

<br>

## Task 4 – Maverics Service Extension

### Summary
The Maverics orchestrator was extended with a custom **service extension**. The extension fetches a user (ID 2) from the JSONPlaceholder API (Task 3) and injects the user’s email address into an outbound HTTP header.

The header `CUSTOM-EMAIL` should appear in outbound requests.

### How to Test

1. **Using a shim**. In this exercise, we use a shim, or development harness, to mimic or stand in for the real system. Open two terminals to begin this testing exercise.

    On the first terminal, invoke the shim to test the service extension.

    ```bash
    go run ./tasks/task3-restclient https://jsonplaceholder.typicode.com/users/2
    ```

    On the second terminal, invoke the execution by requesting data from the dev harness endpoint `/headers`

    ```bash
    curl -i http://127.0.0.1:8080/headers
    ```

    When the extension executes within the harness, it fetches the mock user, extracts the email, builds the `http.Header`, and we immediately inspect the response to ensure that `CUSTOM-EMAIL` is present and accurate.

    Sample Output

    ```bash
    HTTP/1.1 200 OK
    Content-Type: application/json
    Custom-Email: Shanna@melissa.tv
    Date: Mon, 15 Sep 2025 17:21:32 GMT
    Content-Length: 51

    {"headers":{"CUSTOM-EMAIL":["Shanna@melissa.tv"]}}
    ```
<br>

2. **Maverics Orchestrator**. During this enrichment phase, the service extension `CreateEmailHeader` is invoked. The extension calls on a predefined URL API and extracts the user’s email and returns it as an `http.Header` to Maverics. The orchestrator then merges this header, specifically, the `CUSTOM-EMAIL` header set to the fetched email into the outbound request and forwards it to the upstream application.

    We can demonstrate this running the Maverics Orchestrator locally. Here is an example:

    ```bash
    ./maverics_darwin_arm64 -config ./maverics.json
    ```

    The Maverics Orchestrator functions as an identity-aware reverse proxy that sits between the client and the target application. It enforces access policies and enriches requests before they reach the application. When a request is received, Maverics evaluates the policy, establishes or resumes a session, and exposes that context to your service extension.

    <details><summary>See detailed output sample</summary>

    ```bash
    ts=2025-09-15T17:52:44.738704Z msg="initializing config filesystem to OS filesystem"
    ts=2025-09-15T17:52:44.739179Z msg="initializing telemetry service" orchestratorID=123456ABCDE
    ts=2025-09-15T17:52:44.739827Z level=info msg="starting Maverics" version=2025.09.3 date=2025-09-11T18:11:29Z
    ts=2025-09-15T17:52:44.739849Z level=info msg="usage of the Maverics Identity Orchestrator is covered by the following terms and conditions https://www.strata.io/legal/enterprise-master-license-subscription-agreement/"
    ts=2025-09-15T17:52:44.739892Z level=info msg="loading configuration from filesystem"
    ts=2025-09-15T17:52:44.739928Z level=info msg="loaded config './maverics.json' from filesystem"
    ts=2025-09-15T17:52:44.739931Z level=info msg="successfully loaded configuration"
    ts=2025-09-15T17:52:44.740393Z level=info msg="using default set of secure ciphers for TLS configuration 'ALMongooseApp'"
    ts=2025-09-15T17:52:44.740657Z level=info msg="cipher enabled for TLS configuration 'maverics'" name=TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA
    ts=2025-09-15T17:52:44.740661Z level=info msg="successfully parsed configuration" version=21
    ts=2025-09-15T17:52:44.741352Z level=info service="Cache Manager" msg="requesting cache" namespace=/app/oidc constraints="{\"Availability\":\"Local\",\"Max Items\":\"1000000\"}"
    ts=2025-09-15T17:52:44.741368Z level=info service="Cache Manager" msg="found cache factory for the given constraints" constraints="{\"Availability\":\"Local\",\"Max Items\":\"1000000\"}" factory_name=local_default
    ts=2025-09-15T17:52:44.74144Z level=info service="InMemory Cache (Basic Unbounded)" msg="starting service" stability=Stable
    ts=2025-09-15T17:52:44.744787Z level=debug msg="loading HTTP server address from environment"
    ts=2025-09-15T17:52:44.744798Z level=debug msg="loading HTTP server TLS key pair from environment"
    ts=2025-09-15T17:52:44.745036Z level=info msg="using default set of secure ciphers for TLS configuration 'maverics'"
    ts=2025-09-15T17:52:44.745046Z level=debug msg="telemetry service starting"
    ts=2025-09-15T17:52:44.745286Z level=debug msg="telemetry service started successfully"
    ts=2025-09-15T17:52:44.745327Z level=info service="HTTP Observability" msg="starting service" stability=Stable
    ts=2025-09-15T17:52:44.745382Z level=info service="Session Manager" msg="starting service" stability=Stable
    ts=2025-09-15T17:52:44.74539Z level=info service="Session Manager" service="Session Store (Using Cache: InMemory Cache (Memory Bound))" msg="starting service" stability=Stable stability=Stable
    ts=2025-09-15T17:52:44.745394Z level=info service="Session Manager" service="InMemory Cache (Memory Bound)" msg="starting service" stability=Stable stability=Stable
    ts=2025-09-15T17:52:44.745403Z level=info service="Session Manager" service="InMemory Cache (Basic Unbounded)" msg="starting service" stability=Stable stability=Stable
    ts=2025-09-15T17:52:44.745414Z level=info service="Cache Manager" msg="starting service" stability=Stable
    ts=2025-09-15T17:52:44.745424Z level=info service="Job Scheduler" msg="job scheduler started" workerCount=100 maxRetries=3 maxQueueSize=100000
    ts=2025-09-15T17:52:44.745431Z level=info msg="initializing OIDC Auth Server" name=oidcProvider
    ts=2025-09-15T17:52:44.745434Z level=debug msg="initializing wellKnown endpoint for OIDC Auth Server" name=oidcProvider endpoint=login.microsoftonline.com/.well-known/openid-configuration
    ```

    </details>

    In our implementation, we create a route to `https://httpbin.org/anything`. The request is sent to the orchestrator, which presents its service endpoint on the localhost via port 443.
    
    When we reach out to the URL endpoint, the request is routed through the Orchestrator, and the custom extension, `CreateEmailHeader`, injects the custom HTTP header in the relay to the upstream route.

    This is the request invocation:

    ```bash
    curl -ik https://127.0.0.1/anything
    ```

    This is the response from `https://httpbin.org/anything`. Note the `Custom-Email` header that the service is supposed to catch. Also, note that IP address `123.456.789.001` is fake and it has been replaced on purpose. 

    ```bash
    HTTP/2 200 
    access-control-allow-credentials: true
    access-control-allow-origin: *
    content-type: application/json
    date: Mon, 15 Sep 2025 18:02:40 GMT
    server: gunicorn/19.9.0
    content-length: 386

    {
    "args": {}, 
    "data": "", 
    "files": {}, 
    "form": {}, 
    "headers": {
        "Accept": "*/*", 
        "Accept-Encoding": "gzip",
        "Custom-Email": "Shanna@melissa.tv", 
        "Host": "httpbin.org", 
        "User-Agent": "curl/8.7.1", 
        "X-Amzn-Trace-Id": "Root=1-68c854c0-71472a5423b04cc118a5631f"
    }, 
    "json": null, 
    "method": "GET", 
    "origin": "127.0.0.1, 123.456.789.001", 
    "url": "https://httpbin.org/anything"
    }
    ```


### Conclusion
This task ties together Go coding, REST integration, and product extension. It demonstrates how to augment a commercial orchestrator with custom logic using Go.

<br>

## Final Notes

- Each task builds on the previous one, showing growth from basic JSON validation to full product integration.  
- Testing steps are lightweight and reproducible.  
- The project emphasizes **practical integration skills**, not just Go syntax.  

<br>
<br>

---

JSON Sample: Original author<br>
GO Snippets: Original author<br>
New Go Code: Apache License © 2025 G Castillo -see [LICENSE](/LICENSE).<br>
This README and notes: CC BY 4.0 © 2025 G Castillo.