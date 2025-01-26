FROM node:latest

ENV SSE_SERVER_HOST="localhost:8080"
ENV CLIENT_PORT=8082

WORKDIR /client

RUN npm init -y
RUN npm install express ejs


COPY ./ .

CMD ["node", "client.js"]


