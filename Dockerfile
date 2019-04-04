FROM golang

# Add Maintainer Info
LABEL maintainer="Brandon Piner <brandon2255p@gmail.com>"

RUN go get google.golang.org/grpc
RUN go get github.com/google/uuid
RUN go get github.com/looplab/eventhorizon 

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/restaurantonthego/restaurant-svc

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
#RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

ENV MONGO_HOST=localhost:27017

EXPOSE 50051

CMD ["restaurant-svc"]