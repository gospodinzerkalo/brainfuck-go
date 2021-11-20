# Brainfuck-go

## Usage
1) String args
```shell
go run main.go run --i '>++++++++[-<+++++++++>]<.>>+>-[+]++>++>+++[>[->+++<<+++>]<<]>-----.>->+++..+++.>-.<<+[>[+>+]>>]<--------------.>>.+++.------.--------.>+.>+.'
```
2) With path
```shell
go run main.go run --p './tests/fibint.b'
```
3) Wit file path & input file path
```shell
go run main.go run --p "./tests/Collatz.b" --if "./tests/Collatz.in"
```

## Test
```shell
go test
```