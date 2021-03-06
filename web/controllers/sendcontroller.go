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
	"iota/webmamgiota/connections"
	"iota/webmamgiota/web/metadata"
	"net/http"
	"strconv"
)

//SendHandler retrieves the message values from the webpage and sends it to the address given
func SendHandler(w http.ResponseWriter, r *http.Request) {
	//TODO: decide whether all struct fields should be used in the Alert
	data := &struct {
		TransactionID string
		TimeStamp     string
		Success       bool
		Response      bool
	}{
		TransactionID: "",
		Success:       false,
		Response:      false,
	}
	c, err := connections.NewConnection(metadata.Provider, metadata.Seed)

	if err != nil {
		panic(err)
	}

	if r.FormValue("submitted") == "true" {
		newMamMessage := MAMBoardSetup{
			Message: r.FormValue("message"),
		}
		value, err := strconv.ParseInt(r.FormValue("value"), 10, 64)
		if err != nil {
			panic(fmt.Errorf("something went wrong in ParseInt: %v", err))
		}

		newMamMessage.Value = value

		txid, err := connections.Send(metadata.Address, newMamMessage.Value, newMamMessage.Message, c)
		if err != nil {
			panic(fmt.Errorf("Transaction failed: %v", err))
		}

		data.Success = true
		data.Response = true
		data.TransactionID = txid
		//fmt.Println("Data is: ", data)
		fmt.Printf("Sent transaction: %v\n", txid)

	}
	renderTemplate(w, r, "sendmessage.html", data)
}
