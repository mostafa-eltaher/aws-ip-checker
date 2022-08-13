SRC          = $(wildcard */*.go) $(wildcard */*/*.go)
CMDNAME      = aws-ip-checker
BINNAME       = $(CMDNAME)
BINDIR       = bin
DISTDIR      = dist
VERSION      = 0.1.0
GOOS         = $(shell go env GOOS)
GOARCH       = $(shell go env GOARCH)
OS_ARCH      = $(GOOS)_$(GOARCH)
OS_ARCHS_WIN = windows_386 windows_amd64
OS_ARCHS     = darwin_amd64 darwin_arm64 \
linux_386 linux_amd64 linux_arm \
#openbsd_386 openbsd_amd64 \
#freebsd_386 freebsd_amd64 freebsd_arm \
#solaris_amd64

get_os       = $(word 1,$(subst _, ,$(1)))
get_arch     = $(word 2,$(subst _, ,$(1)))

default: build

$(DISTDIR)/$(BINNAME)_$(VERSION)_%.zip: $(BINDIR)/$(VERSION)/%/$(BINNAME).exe | $(DISTDIR)
	zip -j $@ $^

$(DISTDIR)/$(BINNAME)_$(VERSION)_%.tar.gz: $(BINDIR)/$(VERSION)/%/$(BINNAME) | $(DISTDIR)
	tar -zcf $@ -C $(^D) .

$(BINDIR)/$(VERSION)/%/$(BINNAME) $(BINDIR)/$(VERSION)/%/$(BINNAME).exe: $(SRC)
	GOOS=$(call get_os,$(*F)) GOARCH=$(call get_arch,$(*F)) go build -o $@ ./cmd

$(DISTDIR):
	mkdir -p $@

.PHONY: build-all
build-all: $(OS_ARCHS:%=$(BINDIR)/$(VERSION)/%/$(BINNAME)) $(OS_ARCHS_WIN:%=$(BINDIR)/$(VERSION)/%/$(BINNAME).exe)

.PHONY: package-all
package-all: $(OS_ARCHS:%=$(DISTDIR)/$(BINNAME)_$(VERSION)_%.tar.gz) $(OS_ARCHS_WIN:%=$(DISTDIR)/$(BINNAME)_$(VERSION)_%.zip)

.PHONY: build
ifeq ($(GOOS),windows)
build: $(BINDIR)/$(VERSION)/$(OS_ARCH)/$(BINNAME).exe
else
build: $(BINDIR)/$(VERSION)/$(OS_ARCH)/$(BINNAME)
endif

.PHONY: clean
clean:
	rm -rf ./${BINDIR}
	rm -rf ./${DISTDIR}