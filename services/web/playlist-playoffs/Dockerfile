FROM node:18-alpine AS build
RUN apk update && apk add git
RUN mkdir -p /app
WORKDIR /app

COPY package.json .
COPY package-lock.json .

RUN npm install --legacy-peer-deps
COPY . .
RUN npm run build

ENV API_URL ${API_URL}

FROM node:18-alpine
RUN mkdir -p /app
WORKDIR /app
COPY package.json .
COPY package-lock.json .
RUN npm install --production --legacy-peer-deps
COPY --from=build ./app/dist/login-demo ./dist
COPY ./server.js .
COPY ./api-server.js .
COPY ./auth_config.json .
ENV NODE_ENV=production
ENV SERVER_PORT=3000
ENV API_SERVER_PORT=3001
EXPOSE 3000
EXPOSE 3001

CMD ["npm", "run", "prod"]
