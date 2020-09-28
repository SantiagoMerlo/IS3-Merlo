FROM node:13.12.0-alpine

WORKDIR /usr/src/app

COPY package*.json ./

ENV PORT_EXPOSE 3000

RUN npm install

COPY . .

EXPOSE $PORT_EXPOSE

CMD npm start