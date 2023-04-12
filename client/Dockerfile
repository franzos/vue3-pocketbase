FROM node
RUN npm install -g pnpm

WORKDIR /app

COPY ./package.json ./
COPY ./.npmrc ./

RUN pnpm install

EXPOSE 5173

CMD ["pnpm", "run", "dev", "--host"]
