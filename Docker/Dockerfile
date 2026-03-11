# golang image used to build application
FROM golang:1.26-alpine AS builder  

# ENV <-- if i want specify the not secret environment variables, i can put them here

# builder workdirectory
WORKDIR /app

# copy inside the work directory the needed files to download dependencies
COPY go.mod go.sum ./

# download dependencies
RUN go mod download

# copy all source code
COPY . .

# compile the binary code for linux 64-bit
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# use the light base image version
FROM alpine:latest

# chose the work directory to run the code
WORKDIR /root/

# copy the compiled code inside the work directory
COPY --from=builder /app/main .

# container port
EXPOSE 8080

# run command
CMD ["./main"]