FROM node:alpine AS builder

WORKDIR /chat-app

COPY . .

RUN npm install && \
    npm run build

FROM nginx:alpine

COPY --from=builder /chat-app/dist/chat-app/* /usr/share/nginx/html/
