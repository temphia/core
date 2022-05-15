package ticketer

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"github.com/temphia/core/backend/server/lib/kosher"
)

type Ticketer struct {
	key []byte
}

func New(key string) *Ticketer {
	return &Ticketer{
		key: []byte(key),
	}
}

func (t *Ticketer) NewWithExpiry(expiry int64, paths ...string) string {

	paths = append(paths, fmt.Sprint(expiry))

	mac := newMac(t.key, paths...)

	foot := make([]byte, 8)
	binary.LittleEndian.PutUint64(foot, uint64(expiry))

	mac = append(mac, foot...)

	return hex.EncodeToString(mac)
}

func (t *Ticketer) DecodeWithExpiry(ticket string, paths ...string) (int64, bool) {
	bytes, err := hex.DecodeString(ticket)
	if err != nil {
		return 0, false
	}

	oldMac := bytes[:8]
	foot := bytes[8:]

	exp := int64(binary.LittleEndian.Uint64(foot))

	paths = append(paths, fmt.Sprint(exp))

	mac := newMac(t.key, paths...)

	if !hmac.Equal(mac, oldMac) {
		return 0, false
	}

	return exp, true
}

func newMac(key []byte, paths ...string) []byte {
	mac := hmac.New(sha256.New, key)
	for _, path := range paths {
		mac.Write(kosher.Byte(path))
	}
	return mac.Sum(nil)
}
