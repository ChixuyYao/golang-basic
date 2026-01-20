FROM ubuntu:20.04
LABEL authors="chixuy"
COPY golang /app/golang
WORKDIR /app
CMD ["/app/golang"]