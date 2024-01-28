# Build the application from source
FROM golang:1.21 AS build-stage

WORKDIR /

COPY . .
RUN apt update
RUN apt install npm -y
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.1/tailwindcss-linux-x64
RUN chmod +x tailwindcss-linux-x64
RUN go mod download
RUN npm install

RUN mkdir assets/dist
RUN mkdir assets/static

RUN templ generate
RUN ./tailwindcss-linux-x64 -i ./assets/tailwind.css -o ./assets/dist/styles.css --minify
RUN	CGO_ENABLED=0 GOOS=linux go build -o .out/gqs cmd/gqs/main.go 
RUN	CGO_ENABLED=0 GOOS=linux go build -o .out/generator cmd/generator/main.go
# RUN CGO_ENABLED=0 GOOS=linux go build -o /gqs cmd/gqs
# RUN CGO_ENABLED=0 GOOS=linux go build -o /generator cmd/generator
# Deploy the application binary into a lean image
# FROM gcr.io/distroless/base-debian11 AS build-release-stage
FROM alpine:3.11.3 AS build-release-stage
WORKDIR /

COPY --from=build-stage .out/gqs /
COPY --from=build-stage .out/generator /
COPY ./assets /assets
COPY ./fullchain.pem /
COPY ./privkey.pem /

# COPY ./init.sh /init.sh
# RUN ["chmod", "+x", "/init.sh"]

ENTRYPOINT ["/bin/sh", "-c", "./generator -source psql && ./gqs -prod"]