FROM node:16-alpine
RUN apk add --no-cache python2 g++ make
WORKDIR /app
COPY package*.json .
RUN yarn install --production
RUN yarn global add serve
COPY . .
RUN npm run build
CMD ["serve", "-s", "build"]
EXPOSE 3000