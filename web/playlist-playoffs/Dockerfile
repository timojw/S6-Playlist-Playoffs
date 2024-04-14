FROM node:latest as build-step

# WORKDIR /playlist-playoffs

# COPY package.json package-lock.json ./

# RUN npm install

# COPY . .

# RUN npm run build

# RUN ls -la dist

# RUN rm -rf dist

# FROM nginx:latest

# COPY --from=build-step /dist/playlist-playoffs /usr/share/nginx/html

# # COPY --from=build-step /usr/local/playlist-playoffs/dist/playlist-playoffs /usr/share/nginx/html

# COPY ./nginx.conf /etc/nginx/nginx.conf

# EXPOSE 80

# CMD ["nginx", "-g", "daemon off;"]