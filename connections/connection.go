package connections

import (
	"iota/webmamgiota/web/metadata"

	"github.com/giota"
)

//NewConnection establishes a connection with the given provider and the seed
func NewConnection(provider, seed string) (*Connection, error) {
	return &Connection{
		api:      giota.NewAPI(provider, nil),
		seed:     metadata.Seed,
		security: 3,
		mwm:      metadata.MWM,
	}, nil
}

type Connection struct {
	api      *giota.API
	seed     string
	security int
	mwm      int64
}

func (c *Connection) SendToApi(trs []giota.Transfer) (giota.Bundle, error) {
	seed, err := giota.ToTrytes(c.seed)
	if err != nil {
		return nil, err
	}
	_, bestPow := giota.GetBestPoW()
	return giota.Send(c.api, seed, c.security, trs, c.mwm, bestPow)
}

/*
The way the library does this, it makes use of findTransactions to
figure out if there is any address in the Tangle that has already been used.
If findTransactions returns associated transactions, the key index is
simply incremented, a new address generated and findTransactions called
again until it returns null.
*/
func (c *Connection) FindTransactions(req giota.FindTransactionsRequest) ([]giota.Transaction, error) {
	found, err := c.api.FindTransactions(&req)
	if err != nil {
		return nil, err
	}
	return c.ReadTransactions(found.Hashes)
}

func (c *Connection) ReadTransactions(tIDs []giota.Trytes) ([]giota.Transaction, error) {
	found, err := c.api.GetTrytes(tIDs)
	return found.Trytes, err
}
