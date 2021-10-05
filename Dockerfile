FROM golang@sha256:3c4de86eec9cbc619cdd72424abd88326ffcf5d813a8338a7743c55e5898734f AS base
WORKDIR /src
ENV CGO_ENABLED=0
COPY go.* ./
RUN go mod download
COPY . ./

FROM base AS build
ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 go build .

FROM gcr.io/distroless/base:nonroot@sha256:56d73a61ea1135c28f2be9afe2be88fc360e5fa1a892d600512a10eb2e028fa5
COPY --from=build /src/stunning-tribble /
ENTRYPOINT [ "/stunning-tribble" ]
