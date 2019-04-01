FROM golang

# Add Maintainer Info
LABEL maintainer="Brandon Piner <brandon2255p@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/brandon2255p/restaurant/restaurant-svc
RUN go get google.golang.org/grpc
# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
#RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

EXPOSE 50051

CMD ["restaurant-svc"]