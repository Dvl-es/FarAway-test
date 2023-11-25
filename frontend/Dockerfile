FROM node:18-alpine

WORKDIR /faraway-test/

COPY public/ /faraway-test/public
COPY src/ /faraway-test/src
COPY package.json tsconfig.json /faraway-test/

RUN npm install

EXPOSE 3000
CMD ["npm", "start"]