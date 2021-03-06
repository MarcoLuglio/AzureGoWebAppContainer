# Azure pipelines tem uma task só para aplicações Go
# https://docs.microsoft.com/en-us/azure/devops/pipelines/tasks/build/go?view=azure-devops

# https://hub.docker.com/_/golang
FROM golang:alpine AS builder
LABEL maintainer = "Marco Luglio <marcodejulho@gmail.com>"

#ENV GOOS "linux"
COPY . .

WORKDIR /go/src
#RUN go install -v ./... \
RUN go get -d -v \
	&& go build -v -o /go/bin/main

###############################################

FROM alpine

COPY --from=builder /go/bin/main /bin/main

# user execute /bin/main
RUN chmod u+x /bin/main

WORKDIR /bin
EXPOSE 80
ENTRYPOINT [ "/bin/main" ]
CMD [""]