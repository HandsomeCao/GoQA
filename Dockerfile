FROM golang:alpine

WORKDIR /app

COPY . /app/

ENV GOPROXY https://goproxy.io

RUN echo "https://mirror.tuna.tsinghua.edu.cn/alpine/v3.11/main" > /etc/apk/repositories &&\
	echo "https://mirror.tuna.tsinghua.edu.cn/alpine/v3.11/community" >> /etc/apk/repositories &&\
	echo "https://mirror.tuna.tsinghua.edu.cn/alpine/edge/testing" >> /etc/apk/repositories &&\
	apk update &&\
	apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime &&\
	echo "Asia/Shanghai" > /etc/timezone &&\
	apk del tzdata &&\
	apk add gcc g++ --no-cache&&\
	go build .

EXPOSE 8080

ENTRYPOINT ["./XgfyQA"]
