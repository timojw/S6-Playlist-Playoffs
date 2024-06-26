# Stage 1: Build the Angular application
FROM node:18-alpine AS build
RUN apk update && apk add git
RUN mkdir -p /app
WORKDIR /app

COPY package.json .
COPY package-lock.json .

RUN npm install --legacy-peer-deps
COPY . .
RUN npm run build

FROM nginx:alpine

COPY --from=build /app/dist/login-demo /usr/share/nginx/html
COPY nginx/conf.d/default.conf /etc/nginx/conf.d/default.conf
COPY certs/tls.crt /etc/tls/tls.crt
COPY certs/tls.key /etc/tls/tls.key

EXPOSE 80 443

CMD ["nginx", "-g", "daemon off;"]
