FROM bitnami/kafka:latest

# Copy your script into the container
COPY create-topics.sh /docker-entrypoint-initdb.d/create-topics.sh

# Make sure the script is executable
RUN chmod +x /docker-entrypoint-initdb.d/create-topics.sh
