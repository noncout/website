FROM golang:latest

WORKDIR /home/user/GoLand/website

COPY . /home/user/GoLand/website

# VOLUME [ "/home/user/website" ]

#设置goproxy
RUN go env -w GOPROXY="https://goproxy.cn,direct"

RUN go build .

# 不指定会已这个为默认端口运行
EXPOSE 8000

#启动命令
ENTRYPOINT ["./website"]