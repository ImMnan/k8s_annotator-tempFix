
FROM --platform=linux/amd64 alpine:latest

WORKDIR /usr/local/bin
COPY k8s-annotator .

# Set environment variables
ENV SHIP_ID=""
ENV HARBOUR_ID=""
ENV NAMESPACE=""

# Create a non-root user and group
RUN addgroup -g 1337 -S appgroup && adduser -u 1337 -S appuser -G appgroup \
    && chown appuser:appgroup /usr/local/bin/k8s-annotator \
    && chmod +x /usr/local/bin/k8s-annotator

USER appuser

CMD ["/usr/local/bin/k8s-annotator"]
