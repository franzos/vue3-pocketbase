# Vue 3, Vite and PocketBase

This repo explores the integration of a Vue 3 frontend with a minimal PocketBase backend.

**Features**

- Registration and login on Vue 3 frontend
- News collection example
- Form submissions (with basic spam check and email forwarding)

**Technologies**

- PocketBase
- Vue 3, Vite, Pinia, Naive UI

## Setup

Here's a rough overview of what you need to do, to set this up:

1. Run locally or with Docker
2. Open backend http://127.0.0.1:8090/_/ in your browser and create a admin user
3. Double-check that the migration has run (`News` collection should be listed here: http://127.0.0.1:8090/_/#/collections)
4. Open front-end http://localhost:5173 and enjoy

### Run

For a quick trial, simply use docker: `docker-compose up`.

#### Local

Install client dependencies and run:

```bash
cd client
pnpm install
pnpm run dev
```

Install PocketBase

```bash
wget https://github.com/pocketbase/pocketbase/releases/download/v0.14.3/pocketbase_0.14.3_linux_amd64.zip
unzip pocketbase_0.14.3_linux_amd64.zip
```

Run client and server:

```bash
cd server
go run main.go serve
```

Hot reload with `air`:

```bash
guix environment --pure --ad-hoc go
go install github.com/cosmtrek/air@latest
export GOBIN=/home/$(whoami)/go/bin
export PATH=$PATH:$GOBIN
```

Create a `.air.toml` file roughly like this (substitute the path to your `pb_data`):

```toml
# .air.toml
[build]
  cmd = "go build -o ./tmp/main ."
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

#### Docker

Run Docker container:

```bash
docker-compose up -d
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
- [ ] Update docker file
- [ ] Explore 3rd party auth
- [ ] Cleanup client
- [X] Move server into `./server`
- [ ] Tests
- [ ] Fix Docker permissions

## API

Here's what a form submission looks like:

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