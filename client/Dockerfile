# Client Dockerfile
FROM node:16-alpine as build-stage

WORKDIR /app

COPY package.json pnpm-lock.yaml ./
RUN npm install -g pnpm && pnpm install

COPY . .
RUN pnpm run build

FROM node:16-alpine as production-stage

WORKDIR /app

RUN npm install -g pnpm

COPY --from=build-stage /app/dist /app/dist
COPY --from=build-stage /app/node_modules /app/node_modules
COPY --from=build-stage /app/package.json /app/package.json

EXPOSE 4173
CMD ["pnpm", "run", "preview", "--host"]
