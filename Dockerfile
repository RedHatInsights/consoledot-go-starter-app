FROM registry.access.redhat.com/ubi8/go-toolset:1.18.9-13 as builder

USER 0

WORKDIR /app
RUN mkdir -p /app/bin

COPY . .

RUN go mod download 
RUN go build -o server main.go 

# I don't like this
# I'd rather go mod download but I can't seem to get swag to work that way
RUN wget https://github.com/swaggo/swag/releases/download/v1.8.12/swag_1.8.12_Linux_x86_64.tar.gz
RUN tar xvf swag_1.8.12_Linux_x86_64.tar.gz
RUN chmod +x swag
RUN ./swag init

FROM builder as final

WORKDIR /app
EXPOSE 8080
CMD ["./server"]