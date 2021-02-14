FROM node:14.15.1-alpine
WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

COPY ./package.json ./
COPY ./yarn.lock ./
RUN npm install yarn
RUN yarn 

COPY . ./
EXPOSE 3000
CMD ["yarn", "start"]