FROM golang:1.21-alpine AS gobuilder
ENV CGO_ENABLED 0
COPY . /app
WORKDIR /app
RUN go build -o onycom .

FROM scratch
COPY --from=gobuilder /app/onycom /
COPY ./templates /app/templates

ENV TEMPLATE_DIR=/app/templates/*

CMD ["/onycom"]
EXPOSE 8080