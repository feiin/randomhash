# random-hash

Generate random hashes with a fully customizable charset and configurable length


It is a GO port of https://github.com/PabloSichert/random-hash

## Usage

```golang
    //default charset abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_

    randomHash := randomhash.New("")
	result, err := randomHash.GenerateHash(10) //setting hash length
	if err != nil {
		t.Error(err)
	}
    fmt.Printf("random hash %s", result) //Dgc2iUZLeR
    
```

```golang
    
    charset := "😁😎😍😇🤓🤔😴😝🤑🤒😭😈👻👽🤖💩🎅💪👈👉👆👇✌✋👌👍👎👐👂👃👣👁"
    randomHash := randomhash.New(charset)
	result, err := randomHash.GenerateHash(5) //setting hash length
	if err != nil {
		t.Error(err)
	}
    fmt.Printf("random hash %s", result) //🤑✋🤒😍🤑
```


## Algorithm for `generateHash`
```
|----------- a1 ----------| ... |------------------------- a(n) ------------------------|
| b1 b2 b3 b4 b5 b6 b7 b8 | ... | b(n-7) b(n-6) b(n-5) b(n-4) b(n-3) b(n-2) b(n-1) b(n) |
|--- c1 ---|                               ...                               |-- c(n) --|

a(n): binary with 8 digits
b(n): binary from RNG
c(n): binary with log2(charset.length) digits

a1 ... a(ceil(hashLength * log2(charset.length) / 8))
b1 ... b(ceil(hashLength * log2(charset.length) / 8) * 8)
c1 ... c(hashLength)
```

```
charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"
length = 4

|------- a1 ------|------- b2 ------|------- c3 ------|
| 0 1 1 0 0 0 1 1 | 0 1 0 1 0 1 0 0 | 0 0 0 0 0 0 0 1 | random bytes
|---- c1 ----|----- c2 ----|----- c3 ----|---- c4 ----|
|----- y ----|----- 1 -----|----- q -----|----- b ----| = "y1qb"
```

```
Example

charset = "😁😎😍😇🤓🤔😴😝🤑🤒😭😈👻👽🤖💩🎅💪👈👉👆👇✌✋👌👍👎👐👂👃👣👁"

length = 3

|------- a1 ------|------- b2 ------|
| 0 0 0 1 1 0 1 1 | 0 0 1 0 1 1 1 0 | random bytes
|--- c1 ---|---- c2 ---|--- c3 --|
|--- 😇 ---|---- 👻 ----|--- ✋ ---| = "😇👻✋"
```

