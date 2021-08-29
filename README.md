### Requirement
* Install `gcc`
* Install [docker][1] on MacOS or [docker][2] on Ubuntu
* Install `solc` by using docker command 
```
docker pull ethereum/solc:0.6.12
```
* Install `abigen`
```
  go get github.com/ethereum/go-ethereum
  cd $GOPATH/pkg/mod/github.com/ethereum/go-ethereum/
  make
  make devtools
  ```

### How to run
* We can generate the ABI from a solidity source file.
```azure
docker run --rm -v $(pwd):/root ethereum/solc:0.6.12 --abi /root/Store.sol -o /root/build
```
* Convert the ABI to a Go file that we can import. This new file will contain all the available methods the we can use to interact with the smart contract from our Go application.
```
abigen --abi=./build/ERC20.abi --pkg=token --out=erc20.go
```
* Let's start project
```
go run .
```
### Project structure
```
├── README.md
├── gen // generated from abi files
├── build // build contract to abi
├── client // setup client
├── config // includes function get variable from enviroment variables
├── contract // includes files contract
├── vendor // packages
├── .env
├── main.go
```

[1]: https://docs.docker.com/docker-for-mac/install/
[2]: https://docs.docker.com/engine/install/ubuntu/
[3]: https://linuxize.com/post/how-to-install-gcc-compiler-on-ubuntu-18-04/




