package wtclient

import (
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/lightningnetwork/lnd/keychain"
	"github.com/lightningnetwork/lnd/watchtower/wtdb"
)

// Tower represents the info about a watchtower server that a watchtower client
// needs in order to connect to it.
type Tower struct {
	// ID is the unique, db-assigned, identifier for this tower.
	ID wtdb.TowerID

	// IdentityKey is the public key of the remote node, used to
	// authenticate the brontide transport.
	IdentityKey *btcec.PublicKey

	// Addresses is an AddressIterator that can be used to manage the
	// addresses for this tower.
	Addresses *AddressIterator
}

// NewTowerFromDBTower converts a wtdb.Tower, which uses a static address list,
// into a Tower which uses an address iterator.
func NewTowerFromDBTower(t *wtdb.Tower) (*Tower, error) {
	addrs, err := newAddressIterator(t.Addresses...)
	if err != nil {
		return nil, err
	}

	return &Tower{
		ID:          t.ID,
		IdentityKey: t.IdentityKey,
		Addresses:   addrs,
	}, nil
}

// ClientSession represents the session that a tower client has with a server.
type ClientSession struct {
	// ID is the client's public key used when authenticating with the
	// tower.
	ID wtdb.SessionID

	wtdb.ClientSessionBody

	// Tower holds the pubkey and address of the watchtower.
	Tower *Tower

	// SessionKeyECDH is the ECDH capable wrapper of the ephemeral secret
	// key used to connect to the watchtower.
	SessionKeyECDH keychain.SingleKeyECDH
}
