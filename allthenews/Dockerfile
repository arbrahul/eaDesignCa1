FROM golang:1.13-alpine AS build
WORKDIR /src/
COPY . /src/
RUN CGO_ENABLED=0 go build -o /bin/allthenews

FROM scratch
COPY --from=build /bin/allthenews /bin/allthenews
COPY *.html /
ENTRYPOINT ["/bin/allthenews"]
