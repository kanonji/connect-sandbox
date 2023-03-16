FROM node:18-alpine as builder
WORKDIR /build
COPY ./mock-server/package*.json .
RUN npm install

FROM node:18-alpine as runner
RUN apk add --no-cache git ca-certificates tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    echo "Asia/Tokyo" > /etc/timezone
COPY --from=builder /build/node_modules /node_modules
RUN chown node:node -R /node_modules
EXPOSE 3000 
WORKDIR /workdir/mock-server
CMD rm -fR ./node_modules && cp -a /node_modules . && npm start
