cd internal/contracts/$1
solc --abi $1.sol -o build --overwrite
solc --bin $1.sol -o build --overwrite
LOWERCASE="$(tr [A-Z] [a-z] <<< "$1")"
abigen --bin=./build/$1.bin --abi=./build/$1.abi --pkg=$LOWERCASE --out=$1.go
cd ../../..
INPUT=$(tr -d aeiou <<< "$LOWERCASE")
go run cmd/main.go -dp=$INPUT
