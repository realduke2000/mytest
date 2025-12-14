FROM ubuntu:22.04
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
        ca-certificates \
        curl \
        vim \
        bash && \
    rm -rf /var/lib/apt/lists/*
WORKDIR /app

COPY rolex-srv /app/rolex-srv
COPY srv.yaml /app/srv.yaml
RUN chmod +x /app/rolex-srv
CMD ["/app/rolex-srv"]