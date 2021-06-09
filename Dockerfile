FROM golang:1.5 as build
RUN apt update && apt -y install upx
WORKDIR /
COPY main.go .
RUN go build -ldflags "-linkmode external -extldflags -static" -a main.go && strip main && upx -9 main

FROM scratch
COPY --from=build /main /main
EXPOSE 8000
CMD ["/main"]
