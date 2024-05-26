FROM node:latest as build-step

# Set the working directory in the container
WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build -- --output-path=./dist/playlist-playoffs --configuration production --progress=false

FROM nginx:latest

RUN rm -rf /usr/share/nginx/html/*

COPY --from=build-step /app/dist/playlist-playoffs /usr/share/nginx/html

COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
