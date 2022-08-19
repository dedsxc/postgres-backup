########################
####  BUILD GO APP   ###
########################
FROM golang:1.19-alpine AS build

WORKDIR /app

COPY . /app/

RUN CGO_ENABLED=0 go build -o postgres-backup 

################################
## BUILD POSTGRES WITH GO APP ##
################################
FROM alpine:3.16

ARG UID="1000"
ARG USER="user"

WORKDIR /app

COPY --from=build /app/postgres-backup /app

RUN apk update && \
    apk add --no-cache postgresql14-client=14.5-r0 && \
    adduser -D -H -u $UID $USER && \
    chown -R $USER:$USER /app

USER $USER

ENTRYPOINT [ "/app/postgres-backup" ] 