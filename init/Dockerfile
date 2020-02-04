FROM alpine:3.3

RUN apk update \
    && apk upgrade \
    && apk add bash curl iptables \
    && rm -rf /var/cache/apk/*

# Copy the initialization script into the container.
COPY init.sh /usr/local/bin/

# Mark the initialization script as executable.
RUN chmod +x /usr/local/bin/init.sh

# Start the initialization script on container startup.
ENTRYPOINT ["init.sh"]
