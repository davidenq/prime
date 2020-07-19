FROM node:13.10.1-alpine3.10
WORKDIR /usr/src/app
RUN apk --no-cache --update --virtual build-dependencies add \
  tzdata \
  build-base python \
  && npm install \
  && apk del build-dependencies
RUN npm install -g node-gyp rebuild ffi --production
COPY . .
CMD [ "node", "./index.js" ]