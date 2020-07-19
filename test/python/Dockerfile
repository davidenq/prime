FROM python:3.9.0a5-alpine3.10
WORKDIR /usr/src/app
RUN pip install  libtcod-cffi
COPY . .
CMD [ "python", "./index.py" ]