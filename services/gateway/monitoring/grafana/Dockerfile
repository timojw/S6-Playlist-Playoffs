# Dockerfile for Grafana
FROM grafana/grafana:9.1.2

ENV INFLUXDB_HOST=${INFLUXDB_HOST}
ENV INFLUXDB_PORT=${INFLUXDB_PORT}
ENV INFLUXDB_USER=${INFLUXDB_USER}
ENV INFLUXDB_USER_PASSWORD=${INFLUXDB_USER_PASSWORD}

# Expose port 3000
EXPOSE 3000

# Copy dashboard configurations
COPY ./dashboards/all.yml /etc/grafana/provisioning/dashboards/all.yml
COPY ./krakend /var/lib/grafana/dashboards/krakend

# Copy datasource template
COPY ./datasources/all.template.yml /etc/grafana/provisioning/datasources/all.template.yml

USER root

RUN apk add --no-cache gettext

# Set the entrypoint to run envsubst when the container starts
ENTRYPOINT ["sh", "-c", "envsubst < /etc/grafana/provisioning/datasources/all.template.yml > /etc/grafana/provisioning/datasources/all.yml && exec grafana-server"]

USER grafana
