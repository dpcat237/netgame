FROM node:latest

WORKDIR /usr/src/app
COPY package*.json ./
ADD . /usr/src/app
RUN yarn global add @vue/cli
RUN yarn install

COPY . .

RUN yarn build
RUN rm -rf ./build
RUN rm -rf ./src

CMD yarn start
