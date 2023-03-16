FROM node:18-alpine
RUN apk add --no-cache git ca-certificates tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    echo "Asia/Tokyo" > /etc/timezone
EXPOSE 3000 
WORKDIR /go/src/app/mock-server
CMD npm install && npm start
