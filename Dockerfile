FROM golang as build

WORKDIR /app

# cache go mod deps
COPY go.mod go.sum ./
RUN go mod graph | awk '{if ($1 !~ "@") print $2}' | xargs go get

COPY . .
RUN go build -o svc .


FROM ubuntu as run

COPY --from=build /app/svc /app/svc
WORKDIR /app
ENTRYPOINT [ "/app/svc" ]
