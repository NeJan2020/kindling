define exec-command
$(1)
endef

.PHONY: collector
collector: libkindling
	cd collector/ && ./collector-version-build.sh

.PHONY: libkindling
libkindling:
	mkdir -p probe/build && cd probe/build && cmake -DBUILD_DRIVER=OFF -DCMAKE_INSTALL_PREFIX=/usr/lib64/ -DPROBE_VERSION=0.1.1dev .. && make && make install
	mkdir -p collector/docker/libso/. && cp -rf probe/build/src/libkindling.so collector/docker/libso/. && cp -rf probe/build/src/libkindling.so /usr/lib64/.