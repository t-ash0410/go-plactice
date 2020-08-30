# ベースとなるDockerイメージ指定
FROM golang:latest

#vscode user create
ARG USERNAME=vscode
RUN useradd -m $USERNAME
RUN gpasswd -a $USERNAME sudo
RUN echo "$USERNAME:${USERNAME}_pass" | chpasswd

# ENV PATH $GOROOT:/usr/local/go
# ENV PATH $GOPATH:/usr/local/go/bin/go

RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/derekparker/delve/cmd/dlv
