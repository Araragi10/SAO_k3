# Rudeus Telegram Bot Project
# Copyright (C) 2021 wotoTeam, ALiwoto
# This file is subject to the terms and conditions defined in
# file 'LICENSE', which is part of the source code.

FROM golang:1.15-alpine AS golang

COPY --from=wcsiu/tdlib:1.7-alpine /usr/local/include/td /usr/local/include/td
COPY --from=wcsiu/tdlib:1.7-alpine /usr/local/lib/libtd* /usr/local/lib/
COPY --from=wcsiu/tdlib:1.7-alpine /usr/lib/libssl.a /usr/local/lib/libssl.a
COPY --from=wcsiu/tdlib:1.7-alpine /usr/lib/libcrypto.a /usr/local/lib/libcrypto.a
COPY --from=wcsiu/tdlib:1.7-alpine /lib/libz.a /usr/local/lib/libz.a
RUN apk add build-base

WORKDIR /myApp

COPY . .

RUN go build --ldflags "-extldflags '-static -L/usr/local/lib -ltdjson_static -ltdjson_private -ltdclient -ltdcore -ltdactor -ltddb -ltdsqlite -ltdnet -ltdutils -ldl -lm -lssl -lcrypto -lstdc++ -lz'" -o /tmp/rudeus01 main.go

FROM gcr.io/distroless/base:latest
COPY --from=golang /tmp/rudeus01 /rudeus01
ENTRYPOINT [ "/rudeus01" ]




#FROM gcc:9.2

#ENV DEBIAN_FRONTEND noninteractive

#RUN apt-get update && apt-get install -y cmake libgtest-dev libboost-test-dev && rm -rf /var/lib/apt/lists/* 

#RUN git clone https://github.com/tdlib/td.git

#RUN cd td

#RUN rm -rf build

#RUN mkdir build

#RUN cd build

#RUN cmake -DCMAKE_BUILD_TYPE=Release -DCMAKE_INSTALL_PREFIX:PATH=../tdlib ..

#RUN cmake --build . --target install

#RUN cd ..

#RUN cd ..

#RUN ls -l td/tdlib

#FROM golang:1.16

#RUN git clone https://github.com/Araragi10/SAO_k3.git

#WORKDIR Rudeus01

#COPY --from=wcsiu/tdlib:1.7-alpine /usr/local/include/td /usr/local/include/td
#COPY --from=wcsiu/tdlib:1.7-alpine /usr/local/lib/libtd* /usr/local/lib/
#COPY --from=wcsiu/tdlib:1.7-alpine /usr/lib/libssl.a /usr/local/lib/libssl.a
#COPY --from=wcsiu/tdlib:1.7-alpine /usr/lib/libcrypto.a /usr/local/lib/libcrypto.a
#COPY --from=wcsiu/tdlib:1.7-alpine /lib/libz.a /usr/local/lib/libz.a
# RUN apk add build-base


#RUN make

#CMD .docker_build/rudeus01