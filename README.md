# 背景说明
- 本项目是根据《[从零到壹构建基于 Fabric-SDK-Go 的Web应用](https://github.com/kevin-hf/kongyixueyuan)》项目改造，所以整个流程基本一致，只是链码部分由于2.0版本与1.x版本很不一致，所以相关的接口和流程都大改。
- 项目使用Fabric2.2版本
- 项目通过[hyperledger/fabric-samples](https://github.com/hyperledger/fabric-samples)进行本地Fabric测试网络的搭建，参考项目的`Makefile`文件
- 根据《[Fabric test-network搭建[solo-1orderer-2peer]](https://blog.csdn.net/shuizhongmose/article/details/109056691)》准备Fabric测试网络环境

# 参考资料
- [Fabric v2.0 智能合约新生命周期模型的常用操作指南](https://blog.csdn.net/ice_fire_x/article/details/105211579)
- [Goland 远程开发](https://blog.csdn.net/shuizhongmose/article/details/108990936)
- [Goland 远程开发调试](https://cloud.tencent.com/developer/article/1682850)

# Fabric-sdk-go开发指南
`fabric-sdk-go`相关的测试参考见`fabric-sdk-go`(`master`版本)的`README.md`文件，现把重要的部分摘抄如下：
- [E2E Test](test/integration/e2e/end_to_end.go): Basic example that uses SDK to query and execute transaction
- [Ledger Query Test](test/integration/pkg/client/ledger/ledger_queries_test.go): Basic example that uses SDK to query a channel's underlying ledger
- [Multi Org Test](test/integration/e2e/orgs/multiple_orgs_test.go): An example that has multiple organisations involved in transaction
- [Dynamic Endorser Selection](test/integration/pkg/fabsdk/provider/sdk_provider_test.go): An example that uses dynamic endorser selection (based on chaincode policy)
- [E2E PKCS11 Test](test/integration/e2e/pkcs11/e2e_test.go): E2E Test using a PKCS11 crypto suite and configuration
- [CLI](https://github.com/securekey/fabric-examples/tree/master/fabric-cli/): An example CLI for Fabric built with the Go SDK.
- More examples needed!

# 项目源码说明
```
├── chaincode       // 链码示例  
│   ├── go.mod      // 链码的依赖库，需要在链码打包阶段打包进`tar`包  
│   ├── go.sum  
│   ├── simplechaincode.go  
│   └── vendor  
├── config.yaml     // 配置文件，参考hyperledger\fabric-sdk-go\test\fixtures\config下面的配置文件  
├── debugm.sh       // goland远程调试一键运行脚本  
├── go.mod          // 本项目的依赖文件，go目前推荐go modules管理依赖库  
├── go.sum  
├── main.go         // 程序入口  
├── Makefile        // 一键启动网络、创建通道、安装链码等  
├── README.md  
├── sdkenv          // fabric sdk相关接口  
│   ├── sdkEnvInfo.go  
│   ├── sdkIntegration  
│   ├── sdkSeetings_test.go  
│   └── sdkSettings.go      // 接口实现代码  
├── service                 // 链码在web端调用的服务层  
│   ├── simpleService.go  
│   └── simpleServiceImpl.go  
├── vendor  
│   └── github.com  
└── web             // WEB的程序   
    ├── controllers     
    ├── static      
    ├── tpl    
    └── webServer.go  
```  

# 完整的运行日志
```bash
$ make 
Stop environment ...
Stopping network
Stopping peer0.org2.example.com ... done
Stopping peer0.org1.example.com ... done
Stopping orderer.example.com    ... done
Removing peer0.org2.example.com ... done
Removing peer0.org1.example.com ... done
Removing orderer.example.com    ... done
Removing network net_test
Removing volume net_orderer.example.com
Removing volume net_peer0.org1.example.com
Removing volume net_peer0.org2.example.com
Removing network net_test
WARNING: Network net_test not found.
Removing volume net_peer0.org3.example.com
WARNING: Volume net_peer0.org3.example.com not found.
No containers available for deletion
Untagged: dev-peer0.org2.example.com-simplecc_1.0.0-82b65341e685b7304ca35298657583e69cba34d5c24c49b647cd4be6b3b2d6d9-91b669db00dead79bd687b13b14491961a8840fa5ef569bdd9ff91bc8dfab1f7:latest
Deleted: sha256:d77f1bd8119d414c8a8e2c8379b01498f7c2b4af4f45503ce421b15c8f455296
Deleted: sha256:199576f93003996aece15a60d5cede60134177bad449780ac65e19e0d3488ab9
Deleted: sha256:a37852035c464a1e11b6f19d16470ccdc3fead90f087c6d8c2b10b54fca09a4c
Deleted: sha256:160e2e2c6b38ad7d57ed06b8192d61a3cda388162096128d52cd249024e07a6a
Untagged: dev-peer0.org1.example.com-simplecc_1.0.0-82b65341e685b7304ca35298657583e69cba34d5c24c49b647cd4be6b3b2d6d9-cd01b6238c4e8b7bf19420cc695aa815398169cb59be91d498a341ed4469a923:latest
Deleted: sha256:15bedab6a52e98007f9c74431844fdab51371c133e36138a62e3f252f78da6d5
Deleted: sha256:a348729929d6090769c54c24b4f49eb32847424c3086ad7499fb43538469498e
Deleted: sha256:f603374c4e2c49cea2d18b7d292dbca2261afd58a577f59c18e120edc2105d06
Deleted: sha256:32a5b15119cd170d543721e2e1633933b1817352a8ff167ff2a5f90830b4d7b2
Environment down
Clean up ...
Stopping network
Removing network net_test
WARNING: Network net_test not found.
Removing volume net_orderer.example.com
WARNING: Volume net_orderer.example.com not found.
Removing volume net_peer0.org1.example.com
WARNING: Volume net_peer0.org1.example.com not found.
Removing volume net_peer0.org2.example.com
WARNING: Volume net_peer0.org2.example.com not found.
Removing network net_test
WARNING: Network net_test not found.
Removing volume net_peer0.org3.example.com
WARNING: Volume net_peer0.org3.example.com not found.
No containers available for deletion
No images available for deletion
Clean up done ...
Build ...
Build done
Start environment ...
Starting nodes with CLI timeout of '5' tries and CLI delay of '3' seconds and using database 'leveldb' with crypto from 'cryptogen'
LOCAL_VERSION=2.2.2
DOCKER_IMAGE_VERSION=2.2.2
/home/verayy/codes/go/src/github.com/hyperledger/fabric-samples/test-network/../bin/cryptogen
Generate certificates using cryptogen tool
Create Org1 Identities
+ cryptogen generate --config=./organizations/cryptogen/crypto-config-org1.yaml --output=organizations
org1.example.com
+ res=0
Create Org2 Identities
+ cryptogen generate --config=./organizations/cryptogen/crypto-config-org2.yaml --output=organizations
org2.example.com
+ res=0
Create Orderer Org Identities
+ cryptogen generate --config=./organizations/cryptogen/crypto-config-orderer.yaml --output=organizations
+ res=0
Generate CCP files for Org1 and Org2
./network.sh: line 212: ./organizations/ccp-generate.sh: Permission denied
/home/verayy/codes/go/src/github.com/hyperledger/fabric-samples/test-network/../bin/configtxgen
Generating Orderer Genesis block
+ configtxgen -profile TwoOrgsOrdererGenesis -channelID system-channel -outputBlock ./system-genesis-block/genesis.block
2020-10-13 16:06:51.912 CST [common.tools.configtxgen] main -> INFO 001 Loading configuration
2020-10-13 16:06:51.961 CST [common.tools.configtxgen.localconfig] completeInitialization -> INFO 002 orderer type: etcdraft
2020-10-13 16:06:51.962 CST [common.tools.configtxgen.localconfig] completeInitialization -> INFO 003 Orderer.EtcdRaft.Options unset, setting to tick_interval:"500ms" election_tick:10 heartbeat_tick:1 max_inflight_blocks:5 snapshot_interval_size:16777216 
2020-10-13 16:06:51.962 CST [common.tools.configtxgen.localconfig] Load -> INFO 004 Loaded configuration: /home/verayy/codes/go/src/github.com/hyperledger/fabric-samples/test-network/configtx/configtx.yaml
2020-10-13 16:06:51.964 CST [common.tools.configtxgen] doOutputBlock -> INFO 005 Generating genesis block
2020-10-13 16:06:51.964 CST [common.tools.configtxgen] doOutputBlock -> INFO 006 Writing genesis block
+ res=0
Creating network "net_test" with the default driver
Creating volume "net_orderer.example.com" with default driver
Creating volume "net_peer0.org1.example.com" with default driver
Creating volume "net_peer0.org2.example.com" with default driver
Creating peer0.org1.example.com ... done
Creating orderer.example.com    ... done
Creating peer0.org2.example.com ... done
CONTAINER ID        IMAGE                               COMMAND             CREATED             STATUS                  PORTS                              NAMES
e4a213e01421        hyperledger/fabric-orderer:latest   "orderer"           4 seconds ago       Up 1 second             0.0.0.0:7050->7050/tcp             orderer.example.com
1ce8ee6f8ec6        hyperledger/fabric-peer:latest      "peer node start"   4 seconds ago       Up 1 second             7051/tcp, 0.0.0.0:9051->9051/tcp   peer0.org2.example.com
be6ffcb1d842        hyperledger/fabric-peer:latest      "peer node start"   4 seconds ago       Up Less than a second   0.0.0.0:7051->7051/tcp             peer0.org1.example.com
Environment up
Start app ...
2020-10-13 16:06:57.369 CST [common.tools.configtxgen] main -> INFO 001 Loading configuration
2020-10-13 16:06:57.425 CST [common.tools.configtxgen.localconfig] Load -> INFO 002 Loaded configuration: /home/verayy/codes/go/src/github.com/hyperledger/fabric-samples/test-network/configtx/configtx.yaml
2020-10-13 16:06:57.425 CST [common.tools.configtxgen] doOutputChannelCreateTx -> INFO 003 Generating new channel configtx
2020-10-13 16:06:57.428 CST [common.tools.configtxgen] doOutputChannelCreateTx -> INFO 004 Writing new channel tx
2020-10-13 16:06:57.489 CST [common.tools.configtxgen] main -> INFO 001 Loading configuration
2020-10-13 16:06:57.551 CST [common.tools.configtxgen.localconfig] Load -> INFO 002 Loaded configuration: /home/verayy/codes/go/src/github.com/hyperledger/fabric-samples/test-network/configtx/configtx.yaml
2020-10-13 16:06:57.552 CST [common.tools.configtxgen] doOutputAnchorPeersUpdate -> INFO 003 Generating anchor peer update
2020-10-13 16:06:57.554 CST [common.tools.configtxgen] doOutputAnchorPeersUpdate -> INFO 004 Writing anchor peer update
2020-10-13 16:06:57.616 CST [common.tools.configtxgen] main -> INFO 001 Loading configuration
2020-10-13 16:06:57.676 CST [common.tools.configtxgen.localconfig] Load -> INFO 002 Loaded configuration: /home/verayy/codes/go/src/github.com/hyperledger/fabric-samples/test-network/configtx/configtx.yaml
2020-10-13 16:06:57.676 CST [common.tools.configtxgen] doOutputAnchorPeersUpdate -> INFO 003 Generating anchor peer update
2020-10-13 16:06:57.679 CST [common.tools.configtxgen] doOutputAnchorPeersUpdate -> INFO 004 Writing anchor peer update
>> 开始创建通道......
>>>> 使用每个org的管理员身份更新锚节点配置...
>>>> 使用每个org的管理员身份更新锚节点配置完成
>> 创建通道成功
>> 加入通道......
>> 加入通道成功
>> 开始打包链码......
>> 打包链码成功
>> 开始安装链码......
>> 安装链码成功
>> 组织认可智能合约定义......
>>> chaincode approved by Org1 peers:
	peer0.org1.example.com:7051
>>> chaincode approved by Org2 peers:
	peer0.org2.example.com:9051
>> 组织认可智能合约定义完成
>> 检查智能合约是否就绪......
LifecycleCheckCCCommitReadiness cc = simplecc, = {map[Org1MSP:true Org2MSP:true]}
LifecycleCheckCCCommitReadiness cc = simplecc, = {map[Org1MSP:true Org2MSP:true]}
>> 智能合约已经就绪
>> 提交智能合约定义......
>> 智能合约定义提交完成
>> 调用智能合约初始化方法......
>> 完成智能合约初始化
>> 通过链码外部服务设置链码状态......
接收到链码事件: &{a37a503883a756a9d52d630031bb0dfbf451c84dd8fad10578000ee0043ad2f3 simplecc eventSetInfo [] 7 peer0.org1.example.com:7051}
a37a503883a756a9d52d630031bb0dfbf451c84dd8fad10578000ee0043ad2f3
>> 设置链码状态完成
>> 启动web服务......
启动Web服务, 监听端口号: 9000
```
