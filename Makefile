.PHONY: all dev clean build env-up env-down run

all: clean build env-up run

dev: build run

##### BUILD
build:
	@echo "Build ..."
	@go mod vendor
	@cd chaincode && go mod vendor
	@go build
	@echo "Build done"

##### ENV
env-up:
	@echo "Start environment ..."
	@cd ${GOPATH}/src/github.com/hyperledger/fabric-samples/test-network && ./network.sh up
	@echo "Environment up"

env-down:
	@echo "Stop environment ..."
	@cd ${GOPATH}/src/github.com/hyperledger/fabric-samples/test-network && ./network.sh down
	@echo "Environment down"

##### RUN
run:
	@echo "Start app ..."
	@cd ${GOPATH}/src/github.com/hyperledger/fabric-samples/test-network && configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/testchannel.tx -channelID testchannel  -configPath ./configtx/
	@cd ${GOPATH}/src/github.com/hyperledger/fabric-samples/test-network && configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org1MSPanchors.tx -channelID testchannel -asOrg Org1MSP -configPath ./configtx/
	@cd ${GOPATH}/src/github.com/hyperledger/fabric-samples/test-network && configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/Org2MSPanchors.tx -channelID testchannel -asOrg Org2MSP -configPath ./configtx/
	@chmod +x ./fabric-first-go-app
	@./fabric-first-go-app

##### CLEAN
clean: env-down
	@echo "Clean up ..."
	@rm -rf /home/verayy/data/fabric-first-go-app/*
	@cd ${GOPATH}/src/github.com/hyperledger/fabric-samples/test-network && ./network.sh down
	@echo "Clean up done ..."