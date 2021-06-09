# Rudeus Telegram Bot Project
# Copyright (C) 2021 wotoTeam, ALiwoto
# This file is subject to the terms and conditions defined in
# file 'LICENSE', which is part of the source code.

GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
DOCKER_BUILD=$(shell pwd)/.docker_build
DOCKER_CMD=$(DOCKER_BUILD)/rudeus01

$(DOCKER_CMD): clean
	mkdir -p $(DOCKER_BUILD)
	$(GO_BUILD_ENV) go build -o $(DOCKER_CMD) .
#	$(GO_BUILD_ENV) go build -v --ldflags "-extldflags '-static -L/usr/local/lib -ltdjson_static -ltdjson_private -ltdclient -ltdcore -ltdactor -ltddb -ltdsqlite -ltdnet -ltdutils -ldl -lm -lssl -lcrypto -lstdc++ -lz'" -o $(DOCKER_CMD) .
# --ldflags "-extldflags '-static -L/usr/local/lib -ltdjson_static -ltdjson_private -ltdclient -ltdcore -ltdactor -ltddb -ltdsqlite -ltdnet -ltdutils -ldl -lm -lssl -lcrypto -lstdc++ -lz'"#

clean:
	rm -rf $(DOCKER_BUILD)