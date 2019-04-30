package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
)

/*Open decrypts a []byte encrypter with gcm/aes256*/
func Open(data, password []byte) ([]byte, error) {

	s256 := sha256.Sum256(password)
	block, err := aes.NewCipher(s256[:])
	if err != nil {
		return []byte{}, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return []byte{}, err
	}

	nonce, cipherData := data[:aesGCM.NonceSize()], data[aesGCM.NonceSize():]
	clearData, err := aesGCM.Open(nil, nonce, cipherData, nil)
	if err != nil {
		return []byte{}, err
	}
	return clearData, nil
}
