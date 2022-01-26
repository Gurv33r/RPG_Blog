FROM node:17-alpine

RUN mkdir /frontend

COPY ./package.json /frontend

WORKDIR /frontend

RUN npm i

COPY . ./

