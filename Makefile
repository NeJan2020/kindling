define exec-command
$(1)
endef

.PHONY: collector
collector: libkindling
	cd collector/ && ./collector-version-build.sh

.PHONY: libkindling
libkindling:
	mkdir -p probe/build && cd probe/build && cmake -DBUILD_DRIVER=OFF -DPROBE_VERSION=0.1.1dev .. && make
	cp -rf ./src/libkindling.so ../../collector/docker/libso
	cp -rf ./src/libkindling.so /usr/lib64/