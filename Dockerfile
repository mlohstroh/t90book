FROM golang:1.13-alpine as go-build

ENV PORT=3000

WORKDIR /app
RUN apk add --no-cache git make gcc musl-dev
COPY . .
RUN make install

FROM node:14.17-alpine as assets-build

WORKDIR /app
COPY . .
WORKDIR /app/frontend
RUN npm install yarn \
 && yarn \
 && yarn build

FROM alpine:latest

WORKDIR /app
COPY --from=go-build /bin/video-viewer /bin/video-viewer
COPY --from=assets-build /app/frontend/build /app/public

EXPOSE ${PORT}

CMD /bin/video-viewer
