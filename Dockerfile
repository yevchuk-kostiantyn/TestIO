FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/yevchuk-kostiantyn/TestIO

# Build the Solar-System command inside the container.
# Fetch dependencies using "go get"
RUN go get github.com/gorilla/mux
RUN go get github.com/sirupsen/logrus
RUN go get menteslibres.net/gosexy/redis
RUN go install github.com/yevchuk-kostiantyn/TestIO

# Run the Solar-System command by default when the container starts.
ENTRYPOINT /go/bin/TestIO

# Document that the service listens on port 1961.
EXPOSE 1997