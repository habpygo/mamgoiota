/*
MIT License
Copyright (c) 2017 Harry Boer, Jonah Polack

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/giota/mamgoiota"
	"github.com/iotaledger/mamgoiota/connections"
)

var address = "RQP9IFNFGZGFKRVVKUPMYMPZMAICIGX9SVMBPNASEBWJZZAVDCMNOFLMRMFRSQVOQGUVGEETKYFCUPNDDWEKYHSALY"
var seed = "SIERTBRUINSISBEZIGOMEENRONDJESAMENMETWIMAMENTTEMAKENOMZODESUBSIDIERONDTEKRIJGENH9"

func SendHandler(w http.ResponseWriter, r *http.Request) {
	//"https://testnet140.tangle.works"
	c, err := connections.NewConnection("http://node02.iotatoken.nl:14265", seed)
	if err != nil {
		panic(err)
	}

	msgTime := time.Now().UTC().String()
	message := "Testmessage by hopefully you ;-) on: " + msgTime

	id, err := mamgoiota.Send(address, 0, message, c)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Sent Transaction: %v\n", id)
}
