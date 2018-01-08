# mamgoiota

This is a small project to implement Masked Authenticated Messaging on the IOTA tangle with Golang.

This project is still under construction (see TODO) with the aim to get IoT sensors and devices to send MAMs.
Name of the project will change to `webmamgoiota` as we want to have a nice interface to send, receive and get a list of past messages send to a particular address.

## Install

It is assumed that you have Golang installed. You also need to install the Go library API for IOTA which you can download at:

```javascript
go get -u github.com/iotaledger/giota
```

After that you can download the mamgoiota package.

```javascript
go get -u github.com/habpygo/mamgoiota
```

To be able to do testing and assertions you have to install the `stretchr` package

```javascript
go get -u github.com/stretchr/testify
```


## Sending MAMs to the IOTA tangle with Go

### API

#### Create a new Connection
```go
import "github.com/iotaledger/mamgoiota"

func main(){
    c, err := mamgoiota.NewConnection("someNodeURL", "yourSeed")
    if c != nil && err == nil{
        fmt.Println("Connection is valid")
    }
}
```
If you don't have a nodeURL try out one from: http://iotasupport.com/lightwallet.shtml

If you don't have a seed yet, follow the description here: https://iota.readme.io/docs/securely-generating-a-seed

Please keep in mind that you may NEVER loose this seed nor give it to anybody else, because the seed is the connection to your funds!




#### Send a MAM to the IOTA tangle from the CLI
```go
import "github.com/iotaledger/mamgoiota"

func main(){
    c, err := mamgoiota.NewConnection("someNodeURL", "yourSeed")
    if err != nil{
        panic(err)
    }
    id, err := Send("the receiving address", 0, "your stringified message", c)
    if err != nil{
        panic(err)
    }
    fmt.Printf("Send to the Tangle. TransactionId: %v\n", id)
}
```
After sending, you find your transaction here https://thetangle.org giving the TransactionId

If you want to transfer value aswell (here 100 IOTA) call the send method like this: ```Send("the receiving address", 100, "your stringified message", c)```.

#### Read data from the IOTA tangle from the CLI
Reading all transaction received by a certain adress:
```go
import "github.com/iotaledger/mamgoiota"

func main(){
    c, err := NewConnection("someNodeURL", "")
    if err != nil{
        panic(err)
    }

    ts, err := ReadTransactions("Receiving Address", c)
    if err != nil{
        panic(err)
    }
    for i, tr := range ts {
        t.Logf("%d. %v: %d IOTA, %v to %v\n", i+1, tr.Timestamp, tr.Value, tr.Message, tr.Recipient)
    }
}
```
The seed can be ommitted here, since reading does not require an account



Reading a special transaction by transactionID:
```go
import "github.com/iotaledger/mamgoiota"

func main(){
    c, err := NewConnection("someNodeURL", "")
    if err != nil{
        panic(err)
    }

    tx, err := ReadTransaction("Some transactionID", c)
    if err != nil{
        panic(err)
    }
    t.Logf("%v: %d IOTA, %v to %v\n", tx.Timestamp, tx.Value, tx.Message, tx.Recipient)
}
```
#### Examples webmamgoiota - reading and sending from the webpage
1. From the root run `go run main.go`
2. Open your webbrowser and point to `http://localhost:3000` 
3. Click on the `Query all messages` tab and you should see all messages listed, sorted from youngest to oldest.

If you want to send messages, click back to the `Send messages` tab and write your message in the Text message input field.

Make sure to set `value` to 0 (I assume you have no IOTAs at the address)

#### Examples mamgoiota
These examples won't work anymore on this site. Hopefully we will manage to get this workin with the `iotaledger/iota.lib.go` repository on GitHub.

Check out our [example folder](/example) for a send and a receive example.

To run this, cd into the example folder and edit the `sender/send.go` and `receiver/receive.go` file, set the correct provider and address and you are ready to run.

Start the receiver first: `$ go run receiver/receive.go`. It will check for new messages every 5 seconds, until cancelled.

Then start the sender: `$ go run sender/send.go`.

You can also read all the past transactions, i.e. messages + value,  at the address: `go run history/history.go`.

If you pick up the transaction hash from the Terminal output and paste it into the input field on the site https://thetangle.org you find your transaction.

If the Node is offline try another one, mentioned above.

### TODOs
- [ ] GoDoc
- [ ] Travis
- [ ] Make web-app
- [ ] Read sensor data, e.g. RuuVi tag
- [ ] More Read options
- [X] Read by TransactionId





