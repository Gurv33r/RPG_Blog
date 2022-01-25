FROM node:16-alpine3.15

WORKDIR /usr/src/frontend

COPY package.json .

RUN npm i

COPY . ./

RUN npm run build

EXPOSE 3000

CMD ["npm", "start"]