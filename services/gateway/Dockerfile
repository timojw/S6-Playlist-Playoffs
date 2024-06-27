# KrakenD Gateway Dockerfile
FROM devopsfaith/krakend:2.6.0

# Set environment variables
ENV GAME_HOST=${GAME_HOST}
ENV AUDIENCE=${AUDIENCE}
ENV INFLUX_HOST=${INFLUX_HOST}
ENV INFLUX_PORT=${INFLUX_PORT}
ENV INFLUX_USER=${INFLUX_USER}
ENV INFLUX_USER_PASSWORD=${INFLUX_USER_PASSWORD}

# Install gettext to use envsubst
RUN apk --no-cache add gettext

# Create a non-root user for better security
RUN addgroup -S nonroot && adduser -S nonroot -G nonroot
RUN chown -R nonroot:nonroot /etc/krakend
USER nonroot

# Copy the template configuration file into the image
COPY krakend.template.json /etc/krakend/krakend.template.json

# Set the entrypoint to generate the actual configuration file from the template and run KrakenD
CMD echo "Environment Variables:" && env && \
    echo "Generating krakend.json from krakend.template.json" && \
    envsubst < /etc/krakend/krakend.template.json > /etc/krakend/krakend.json && \
    echo "Generated krakend.json:" && cat /etc/krakend/krakend.json && \
    krakend run -c /etc/krakend/krakend.json
