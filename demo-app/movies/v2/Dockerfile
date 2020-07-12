FROM node:12.18-alpine
LABEL maintainer="Rahul Rai <rahul@rahul-rai.com> & Tarun Pabbi <tarun.pabbi@outlook.com>"
ENV NODE_ENV production
WORKDIR /usr/src/app
COPY ["package.json", "package-lock.json*", "npm-shrinkwrap.json*", "./"]
RUN npm install --production --silent && mv node_modules ../
COPY . .
EXPOSE 8080
CMD node index.js