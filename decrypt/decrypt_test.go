package decrypt

import (
	"encoding/hex"
	"testing"
)

func TestOpen(t *testing.T) {

	_, err := Open([]byte{}, []byte{1})
	if err == nil {
		t.Fatal("Should detect an empty data slice")
	}
	_, err = Open([]byte{1}, []byte{})
	if err == nil {
		t.Fatal("Should detect an empty password")
	}
	encryptedAndNonce, err := hex.DecodeString("a30750ef173c84441bb27052667558d337d24f46e03c02740cbc20a35722254d5e228cfe48")
	clear, err := Open(encryptedAndNonce, []byte("0123456789"))
	if err != nil {
		t.Fatal(err)
	}
	if string(clear) != "abcdefghi" {
		t.Fatal("Should be able to decrypt data")
	}
}
