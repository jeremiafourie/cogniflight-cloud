FROM golang:1.24

ARG MODE=dev
ENV MODE=${MODE}

RUN useradd -ms /bin/sh -u 1001 app
USER app

WORKDIR /app
COPY --chown=app:app . .

RUN go mod download && mkdir -p /app/bin

# Production: build once and run the binary
RUN if [ "$MODE" = "prod" ]; then \
      go build -o /app/bin/main . && \
      chown app:app /app/bin/main ; \
    fi

CMD if [ "$MODE" = "prod" ]; then \
      exec /app/bin/main ; \
    else \
      go build -o /app/bin/main . && exec /app/bin/main ; \
    fi
