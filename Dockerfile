# Build Stage
# First pull Golang image
FROM golang:1.17 as build-env

# Set environment variable
ENV APP_NAME project-web-service
ENV CMD_PATH main.go
 
# Copy application data into image
COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME
 
# Budild application
RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH

FROM golang:1.17

# Set environment variable
ENV APP_NAME project-web-service
 
# Copy only required data into this image
COPY --from=build-env /$APP_NAME .

COPY .env .
 
# Expose application port
EXPOSE 8999
 
# Start app
CMD ./$APP_NAME