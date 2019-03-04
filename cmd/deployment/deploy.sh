parentdir="$(dirname $(pwd))"
solc --abi $parentdir/smart-contracts/$1/$1.sol -o $parentdir/smart-contracts/$1/build --overwrite
solc --bin $parentdir/smart-contracts/$1/$1.sol -o $parentdir/smart-contracts/$1/build --overwrite
LOWERCASE="$(tr [A-Z] [a-z] <<< "$1")"
abigen --bin=$parentdir/smart-contracts/$1/build/$1.bin --abi=$parentdir/smart-contracts/$1/build/$1.abi --pkg=$LOWERCASE --out=$parentdir/bchain-qlearning/internal/contracts/$1/$1.go
INPUT=$(tr -d aeiou <<< "$LOWERCASE")
go run cmd/main.go -dp=$INPUT
