FROM npmjs/npm-docker-baseline:12-alpine as builder
LABEL front=blue-dashboard

WORKDIR /src
ADD . .

RUN npm install
RUN yarn build

FROM nginx:alpine
COPY --from=builder /src/build/ /usr/share/nginx/html
