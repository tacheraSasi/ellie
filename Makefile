VERSION=0.1.9

# UPX installation instructions
# macOS: brew install upx
# Linux: sudo apt-get install upx
# Windows: Download from https://github.com/upx/upx/releases

build:
	make build_android
	make build_linux
	make build_windows
	make build_mac

build_upx:
	@echo 'Checking UPX installation...'
	@which upx >/dev/null 2>&1 || (echo "UPX not found. Please install UPX:" && \
		echo "macOS: brew install upx" && \
		echo "Linux: sudo apt-get install upx" && \
		echo "Windows: Download from https://github.com/upx/upx/releases" && exit 1)
	@echo 'Building with UPX compression...'
	make build_android_upx
	make build_linux_upx
	make build_windows_upx
	make build_mac_upx

build_linux:
	@echo 'building linux binary...'
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ellie
	@echo 'zipping build...'
	tar -zcvf binaries/ellie_linux_amd64.tar.gz ellie
	@echo 'cleaning up...'
	rm ellie     

build_linux_upx:
	@echo 'building linux binary with UPX...'
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ellie
	@echo 'compressing with UPX...'
	upx --best ellie
	@echo 'zipping build...'
	tar -zcvf binaries/ellie_linux_amd64_upx.tar.gz ellie
	@echo 'cleaning up...'
	rm ellie

build_windows:
	@echo 'building windows executable...'
	env GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ellie_windows_amd64.exe
	@echo 'zipping build...'
	zip binaries/ellie_windows_amd64.zip ellie_windows_amd64.exe
	@echo 'cleaning up...'
	rm ellie_windows_amd64.exe

build_windows_upx:
	@echo 'building windows executable with UPX...'
	env GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ellie_windows_amd64.exe
	@echo 'compressing with UPX...'
	upx --best ellie_windows_amd64.exe
	@echo 'zipping build...'
	zip binaries/ellie_windows_amd64_upx.zip ellie_windows_amd64.exe
	@echo 'cleaning up...'
	rm ellie_windows_amd64.exe

build_mac:
	@echo 'building mac binary...'
	env GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o ellie
	@echo 'zipping build...'
	tar -zcvf binaries/ellie_mac_amd64.tar.gz ellie
	@echo 'cleaning up...'
	rm ellie

build_mac_upx:
	@echo 'building mac binary with UPX...'
	env GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o ellie
	@echo 'compressing with UPX...'
	upx --best ellie
	@echo 'zipping build...'
	tar -zcvf binaries/ellie_mac_amd64_upx.tar.gz ellie
	@echo 'cleaning up...'
	rm ellie

build_android:
	@echo 'building android binary'
	env GOOS=android GOARCH=arm64 go build -ldflags="-s -w" -o ellie
	@echo 'zipping build...'
	tar -zcvf binaries/ellie_android_arm64.tar.gz ellie
	@echo 'cleaning up...'
	rm ellie

build_android_upx:
	@echo 'building android binary with UPX...'
	env GOOS=android GOARCH=arm64 go build -ldflags="-s -w" -o ellie
	@echo 'compressing with UPX...'
	upx --best ellie
	@echo 'zipping build...'
	tar -zcvf binaries/ellie_android_arm64_upx.tar.gz ellie
	@echo 'cleaning up...'
	rm ellie

build_test:
	go build -ldflags="-s -w" -o ellie

dependencies:
	@echo 'checking dependencies...'
	go mod tidy

test:
	@echo -e '\nTesting Lexer...'
	@./gotest --format testname ./lexer/ 
	@echo -e '\nTesting Parser...'
	@./gotest --format testname ./parser/
	@echo -e '\nTesting AST...'
	@./gotest --format testname ./ast/
	@echo -e '\nTesting Object...'
	@./gotest --format testname ./object/
	@echo -e '\nTesting Evaluator...'
	@./gotest --format testname ./evaluator/

clean:
	go clean


set_global:
	@echo 'Setting global version to $(VERSION)...'
	@sudo mv ellie /usr/local/bin
