# my-react-app (monorepo)

- `FE/` React + Vite frontend
- `BE/` Go backend (net/http)

## Run locally

```sh
# backend (http://localhost:8080)
cd BE && make run

# frontend (http://localhost:5173)
cd FE && yarn install && yarn dev
```

The Vite dev server proxies `/api` to the backend on port 8080.

## Docker

```sh
docker compose up --build
```
