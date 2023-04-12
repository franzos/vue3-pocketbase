# Vue 3, Vite and PocketBase

This repo explores the integration of a Vue 3 frontend with a minimal PocketBase backend.

**Features**

- Registration and login on Vue 3 frontend
- News collection example

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

Install client dependencies:

```bash
pnpm install
```

Install PocketBase

```bash
wget https://github.com/pocketbase/pocketbase/releases/download/v0.14.3/pocketbase_0.14.3_linux_amd64.zip
unzip pocketbase_0.14.3_linux_amd64.zip
```

Run client and server:

```bash
./pocketbase serve
cd client; pnpm run dev
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
- [ ] Explore 3rd party auth
- [ ] Cleanup client
- [ ] Move server into `./server`
- [ ] Tests
- [ ] Fix Docker permissions