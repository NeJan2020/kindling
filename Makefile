GitCommit=${shell git rev-parse --short HEAD || echo unknow}
HARBOR ?= 
AgentTag=${shell git describe --tags --dirty}
mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
DIR := $(dir $(mkfile_path))
LIB_PATH = $(LIBRARY_PATH):${DIR}collector/docker/libso/

export LIB_PATH

define exec-command
$(1)
endef

.PHONY: collector

collector: libkindling
	@echo Build Env:
	@echo "	LD_LIBRARY_PATH=$(LIB_PATH)"
	@echo Libkindling Info:
	@echo "	Path=${shell ldconfig ${DIR}collector/docker/libso/ -p |grep kindling |cut -d '>' -f 2}"
	@echo Repository Info:
	@echo "	Agent Commit: ${GitCommit}"
	@echo "	Last Tag: ${AgentTag}"
	cd collector && go build -o docker/kindling-collector -ldflags="-X 'github.com/Kindling-project/kindling/collector/version.CodeVersion=${GitCommit}'" ./cmd/kindling-collector/

.PHONY: libkindling
libkindling:
	mkdir -p probe/build && cd probe/build && cmake -DBUILD_DRIVER=OFF -DPROBE_VERSION=0.1.1dev .. && make 
	mkdir -p collector/docker/libso/. && cp -rf probe/build/src/libkindling.so collector/docker/libso/.

.PHONY: agent-image
agent-image: collector
	docker build -t  ${HARBOR}kindling-agent:${AgentTag} -f collector/docker/Dockerfile collector/docker \
		--label "AgentTag=${AgentTag}" \
		--label "CommitId=${GitCommit}" 
