FROM golang:1.15-alpine as backend
LABEL backend=blue-dashboard
WORKDIR /src

ADD ./backend .

RUN GOOS=linux go build -o ./build/server ./cmd/server/


FROM npmjs/npm-docker-baseline:12-alpine as front
LABEL front=blue-dashboard

WORKDIR /src
ADD ./front .

RUN npm install
RUN npm run build


FROM nginx:alpine

COPY --from=backend /src/build/ /
COPY --from=front /src/build/ /usr/share/nginx/html
COPY entrypoint.sh entrypoint.sh
RUN chmod +x entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]