# Vue 3, Vite and PocketBase

This repo explores the integration of a Vue 3 frontend with a minimal PocketBase backend.

**Features**

- Registration and login on Vue 3 frontend
- Form submissions (with basic spam check, blocklist and email forwarding)

**Technologies**

- PocketBase
- Vue 3, Vite, Pinia, Naive UI

## Setup
### Production

For a quick trial, 

1. Create a `prod.env` file with `PB_ENCRYPTION_KEY`
2. Run docker: `docker-compose up -d`.

Backend data is stored at `./pb_data`. You can change this path in `docker-compose.yml`.

- Admin UI: http://localhost:8090/_/
- REST API: http://localhost:8090/api/
- Vue 3 frontend: http://localhost:4173/

If you want to go live, you'll need to setup a reverse proxy like nginx.

### Development

Install client dependencies and run:

```bash
cd client
pnpm install
pnpm run dev
```

Run client and server:

```bash
cd server
# guix environment --pure --ad-hoc go
go mod download && go mod verify
go run main.go serve
```

#### Hot reload

Hot reload with `air`:

```bash
# guix environment --pure --ad-hoc go
go install github.com/cosmtrek/air@latest
export GOBIN=/home/$(whoami)/go/bin
export PATH=$PATH:$GOBIN
```

Create a `.air.toml` file roughly like this (substitute the path to your `pb_data`):

```toml
# .air.toml
[build]
  cmd = "go build ."
  bin = "tmp/main serve --dir /home/franz/git/gpt-pocketbase/server/pb_data"
  log = "air.log"

[watch]
  includes = ["."]
  excludes = ["tmp", "vendor"]

[logger]
  level = "info"
  output = "air.log"

[on_error]
  command = ""
  output = "/dev/null"

[on_shutdown]
  command = ""
  output = "/dev/null"
```

Run `air`:

```bash
air
```

## Migratons

Here's how-to create and run database migrations:

```bash
$ docker exec -it b7452e7e4782 ./pocketbase migrate create "news"
? Do you really want to create migration "pb_migrations/1681321777_news.js"? Yes
Successfully created file "pb_migrations/1681321777_news.js
```

The cool thing is, you can also simply create a new collection trough via admin panel 127.0.0.1:8090/_/#/collections and PocketBase will create the migration-file in `./pb_migrations` for you. More on this here: https://pocketbase.io/docs/migrations/

## Todo

- [X] Explore migrations
- [X] Turn into something useful
- [X] Update docker file
- [ ] Explore 3rd party auth
- [ ] Cleanup client
- [X] Move server into `./server`
- [ ] Tests
- [ ] Replace logs ([issue](https://github.com/pocketbase/pocketbase/discussions/1781))
- [ ] Add support for organizations
- [ ] Support global and organization-specific blocklists

## API

Here's what a form submission could look like:

```bash
curl -X POST \
  -H "Content-Type: multipart/form-data" \
  -F "name=John Doe" \
  -F "email=johndoe@example.com" \
  -F "phone=1234567890" \
  -F "message=Hello, this is a test message" \
  http://localhost:8090/api/submit/fe9bkipw3ppus5a:5eh2ztnh02cu0v9
```

Here's a spammy example:

```bash
curl -X POST \
    -H "Content-Type: multipart/form-data" \
    -F "name=John" \
    -F "email=abc" \
    -F "phone=abc" \
    -F "message=Wanna buy stuff? Pharma buy. Buy Pharma. Buy pharma. Check this link http://spamsite.com and this one https://spamsite2.com" \
    http://localhost:8090/api/submit/fe9bkipw3ppus5a:5eh2ztnh02cu0v9
```