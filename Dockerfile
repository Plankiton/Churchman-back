FROM archlinux
LABEL maintainer Yaks Souza <pl4nk1ton@gmail.com>

RUN mkdir /api
ADD . /api

EXPOSE 8000

WORKDIR /api

CMD ./tmp/main >> out.log 2>> out.log
