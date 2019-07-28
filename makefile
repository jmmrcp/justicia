version=3.1.21
date=$(shell date "+(%d %B %Y)")
exec=justicia

.PHONY: all

all:
	@echo " make <cmd>"
	@echo ""
	@echo "commands:"
	@echo " build          - runs go build"
	@echo " build_version  - runs go build with ldflags version=${version} & date=${date}"
	@echo ""

build: clean
	@go build -v -o ${exec}

build_version: clean
	@go build -v -ldflags='-s -w -X "main.version=${version}" -X "main.date=${date}"' -o ${exec}

clean:
	@rm -f ${exec}

check_version:
	@if [ -a "${exec}" ]; then \
		echo "${exec} already exists"; \
		exit 1; \
	fi;