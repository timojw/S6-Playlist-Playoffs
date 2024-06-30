# Stage 1: Build the Angular application
FROM node:18-alpine AS build
RUN apk update && apk add git
WORKDIR /app

# Copy package files and install dependencies
COPY services/web/playlist-playoffs/package.json services/web/playlist-playoffs/package-lock.json ./
RUN npm install --legacy-peer-deps

# Copy the rest of the application files from the specific directory
COPY services/web/playlist-playoffs .

# Build the Angular application
RUN npm run build

# Stage 2: Set up NGINX
FROM nginx:alpine

COPY --from=build /app/dist/login-demo /usr/share/nginx/html
COPY services/web/playlist-playoffs/auth_config.json /usr/share/nginx/html/auth_config.json
COPY services/web/playlist-playoffs/nginx/conf.d/default.conf /etc/nginx/conf.d/default.conf

# Copy TLS certificates for HTTPS
COPY services/web/playlist-playoffs/certs/tls.crt /etc/nginx/tls/tls.crt
COPY services/web/playlist-playoffs/certs/tls.key /etc/nginx/tls/tls.key

# Expose ports for HTTP and HTTPS
EXPOSE 80 443

# Start NGINX
CMD ["nginx", "-g", "daemon off;"]
