FROM alpine:3.14
RUN apk add --no-cache unzip wget

WORKDIR /app

RUN wget https://github.com/pocketbase/pocketbase/releases/download/v0.14.3/pocketbase_0.14.3_linux_amd64.zip && \
    unzip pocketbase_0.14.3_linux_amd64.zip && \
    rm pocketbase_0.14.3_linux_amd64.zip

EXPOSE 8090

CMD ["./pocketbase", "serve"]
