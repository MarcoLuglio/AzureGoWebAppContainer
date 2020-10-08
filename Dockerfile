# Azure pipelines tem uma task só para aplicações Go
# https://docs.microsoft.com/en-us/azure/devops/pipelines/tasks/build/go?view=azure-devops

# https://hub.docker.com/_/golang
# usar FROM golang:latest se o alpine não funcionar direito
FROM golang:alpine
LABEL maintainer = "Marco Luglio <marcodejulho@gmail.com>"
ENV GOOS = linux
#RUN
WORKDIR /src
COPY . .

#for AWS lambda
#RUN go get -d -v ./...
#RUN go install -v ./...
#RUN go build -v -o ./...
RUN go build -v ./main.go
ENTRYPOINT [ "./main" ]
EXPOSE 80
#CMD ["app"]
