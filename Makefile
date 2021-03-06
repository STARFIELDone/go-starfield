# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: gest android ios gest-cross evm all test clean
.PHONY: gest-linux gest-linux-386 gest-linux-amd64 gest-linux-mips64 gest-linux-mips64le
.PHONY: gest-linux-arm gest-linux-arm-5 gest-linux-arm-6 gest-linux-arm-7 gest-linux-arm64
.PHONY: gest-darwin gest-darwin-386 gest-darwin-amd64
.PHONY: gest-windows gest-windows-386 gest-windows-amd64

GOBIN = ./build/bin
GO ?= latest
GORUN = env GO111MODULE=on go run

gest:
	$(GORUN) build/ci.go install ./cmd/gest
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gest\" to launch gest."

all:
	$(GORUN) build/ci.go install

android:
	$(GORUN) build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/gest.aar\" to use the library."
	@echo "Import \"$(GOBIN)/gest-sources.jar\" to add javadocs"
	@echo "For more info see https://stackoverflow.com/questions/20994336/android-studio-how-to-attach-javadoc"
	
ios:
	$(GORUN) build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/Gest.framework\" to use the library."

test: all
	$(GORUN) build/ci.go test

lint: ## Run linters.
	$(GORUN) build/ci.go lint

clean:
	env GO111MODULE=on go clean -cache
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/kevinburke/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go get -u github.com/golang/protobuf/protoc-gen-go
	env GOBIN= go install ./cmd/abigen
	@type "npm" 2> /dev/null || echo 'Please install node.js and npm'
	@type "solc" 2> /dev/null || echo 'Please install solc'
	@type "protoc" 2> /dev/null || echo 'Please install protoc'

# Cross Compilation Targets (xgo)

gest-cross: gest-linux gest-darwin gest-windows gest-android gest-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/gest-*

gest-linux: gest-linux-386 gest-linux-amd64 gest-linux-arm gest-linux-mips64 gest-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/gest-linux-*

gest-linux-386:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/gest
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/gest-linux-* | grep 386

gest-linux-amd64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/gest
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gest-linux-* | grep amd64

gest-linux-arm: gest-linux-arm-5 gest-linux-arm-6 gest-linux-arm-7 gest-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/gest-linux-* | grep arm

gest-linux-arm-5:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/gest
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/gest-linux-* | grep arm-5

gest-linux-arm-6:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/gest
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/gest-linux-* | grep arm-6

gest-linux-arm-7:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/gest
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/gest-linux-* | grep arm-7

gest-linux-arm64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/gest
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/gest-linux-* | grep arm64

gest-linux-mips:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/gest
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/gest-linux-* | grep mips

gest-linux-mipsle:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/gest
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/gest-linux-* | grep mipsle

gest-linux-mips64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/gest
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/gest-linux-* | grep mips64

gest-linux-mips64le:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/gest
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/gest-linux-* | grep mips64le

gest-darwin: gest-darwin-386 gest-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/gest-darwin-*

gest-darwin-386:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/gest
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/gest-darwin-* | grep 386

gest-darwin-amd64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/gest
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gest-darwin-* | grep amd64

gest-windows: gest-windows-386 gest-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/gest-windows-*

gest-windows-386:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/gest
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/gest-windows-* | grep 386

gest-windows-amd64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/gest
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/gest-windows-* | grep amd64
