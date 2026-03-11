# Setup Guide — EXBanka-3-Infrastructure

## To build after cloning

### `make install-buf` — run once, ever

This installs 5 CLI tools globally into your Go `$GOPATH/bin`:

```
buf                       — the proto build tool (replaces raw protoc)
protoc-gen-go             — generates *.pb.go files (Go structs from .proto messages)
protoc-gen-go-grpc        — generates *_grpc.pb.go files (gRPC client/server interfaces)
protoc-gen-grpc-gateway   — generates *.pb.gw.go files (HTTP→gRPC reverse proxy)
protoc-gen-openapiv2      — generates swagger.json from your proto annotations
```

You need these installed before the next step. Without them, `buf generate` fails with "plugin not found".

---

### `make proto` — run after any `.proto` change

Runs `buf generate`, which reads `buf.gen.yaml` and processes all 3 proto files in `proto/`. It populates `gen/proto/` with:

```
gen/proto/
  auth/v1/
    auth.pb.go          ← Go structs: LoginRequest, LoginResponse, Claims…
    auth_grpc.pb.go     ← interfaces: AuthServiceServer, AuthServiceClient
    auth.pb.gw.go       ← HTTP handlers that forward to gRPC
  employee/v1/
    employee.pb.go
    employee_grpc.pb.go
    employee.pb.gw.go
  notification/v1/
    notification.pb.go
    notification_grpc.pb.go
```

**This step is what makes the code compile.** The files in `internal/handler/` and `cmd/server/main.go` import `EXBanka/gen/proto/auth/v1` etc. — those packages don't exist until `buf generate` writes them. If you skip this step, `go build` gives "cannot find module" errors.

---

### `make seed` — run once against a fresh database

Runs `cmd/seed/main.go`, which:
1. Connects to PostgreSQL using `.env` settings
2. Runs `AutoMigrate` — creates all 4 tables (`employees`, `clients`, `permissions`, `tokens`) and the `employee_permissions` join table
3. Inserts the 6 default permissions (`admin`, `employee.create`, `employee.read`, etc.)
4. Creates the first admin account with all permissions:
   - **Email:** `admin@bank.com`
   - **Password:** `Admin123!`
   - Password is hashed with PBKDF2+salt (not stored in plaintext)

Without this step, the database is empty and you can't log in to do anything.

---

### `make run` — start the server

Runs `cmd/server/main.go`, which starts **two listeners in the same process**:

```
:9090  — gRPC server  (used internally by grpc-gateway)
:8080  — HTTP server  (used by the frontend / curl / browser)
```

On startup it also re-runs migrate + seed-permissions (idempotent — safe to repeat), so if you add a new permission to `DefaultPermissions` later, just restarting the server inserts it.

The HTTP server mux has two routes:
- `GET /health` → returns `{"status":"ok"}` — used by Docker health checks
- `/*` → CORS middleware → grpc-gateway mux → forwards to gRPC on `:9090`

---

### Full first-run sequence

```sh
# 1. Install tools (once per machine)
make install-buf

# 2. Generate Go code from proto (once, then again after any .proto edit)
make proto

# 3. Start PostgreSQL (needs to be running before seed/run)
docker-compose up -d postgres   # or run Postgres locally

# 4. Create tables + initial admin
make seed

# 5. Start the server
make run
```

After step 5 the server is live at `http://localhost:8080`. Verify with:

```sh
curl http://localhost:8080/health
# → {"status":"ok","service":"EXBanka"}
```
