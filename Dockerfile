# A named stage: "build"
FROM golang:1.15 AS build 

WORKDIR /src
COPY go.* ./
RUN go mod download

COPY . /src
RUN go build -o main

# A new stage: "run"
FROM gcr.io/distroless/base-debian10:nonroot AS run

# Copy the binary from stage build
COPY --from=build /src/main /

ENTRYPOINT ["/main"]