protoc:
	rm -rf proto/deposit
	mkdir ./proto/deposit
	@PATH=$PATH:$GOPATH/bin/ protoc --go_out=. proto/*.proto
