# Daily Coding Problem: Problem #903 [Easy]

This is the solution to problem 903 in the Daily Coding Problem newsletter: https://www.dailycodingproblem.com/

This is the description of the problem: 

"This problem was asked by Microsoft.

Given an array of numbers and a number k, determine if there are three entries in the array which add up to the specified number k. For example, given [20, 303, 3, 4, 25] and k = 49, return true as 20 + 4 + 25 = 49."

## How to run
- navigate to the problem folder and run `go run .`
- a server will be started listening in port 8000
- the following endpoints are available:
  - POST /add-item/{item}
  - GET /get-items
  - DELETE /delete-all
  - GET /find-three/{sum}

The `rinkeby/rinkeby.go` file contains three constants that you need to provide:

- rinkeby URL: You can get one at https://infura.io/
- contract Address: you can copy this from the Remix IDE after deploying the contract to Rinkeby
- keyPassword: necessary to authenticate with the network using your private key

As well as these three, you need to send the information about your private key as the JSON body of the request. A JSON body would look something like this:
```
{
    "address":"...",
    "crypto":{
        "cipher":"...",
        "ciphertext":"...",
        "cipherparams":{
            "iv":"..."
        },
        "kdf":"...",
        "kdfparams":{
            "dklen": ..,
            "n": ...,
            "p": ...,
            "r": ...,
            "salt":"..."
        },
        "mac":"..."
    },
    "id":"...",
    "version": ...
}
```

Copy your metamask private key to a file, and run this command: `geth account import path/to/private/key/file
`. It will create a new file in `~/Library/Ethereum/keystore`, which contains the JSON you need. 


# Links
- a good quick intro to testing and contracts in Go: https://medium.com/hackernoon/a-step-by-step-guide-to-testing-and-deploying-ethereum-smart-contracts-in-go-9fc34b178d78
- online IDE for compiling/deploying contracts: https://remix.ethereum.org/
- API links to connect to Rinkeby: https://infura.io/
- Daily Coding Problem newsletter: https://www.dailycodingproblem.com/