FROM alpine

WORKDIR /app

COPY ./backend/main /app/main

RUN chmod +x /app/main

ENTRYPOINT ["/app/main"]