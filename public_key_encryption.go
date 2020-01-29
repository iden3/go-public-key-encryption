package pkencryption

import (
	"encoding/hex"
	"fmt"

	"github.com/jamesruan/sodium"
)

func ImportBoxPublicKey(pkHex string) (*sodium.BoxPublicKey, error) {
	pkBytes, err := hex.DecodeString(pkHex)
	if err != nil {
		return nil, err
	}
	if len(pkBytes) != 32 {
		return nil, fmt.Errorf("Expected 32 bytes for public key")
	}
	return &sodium.BoxPublicKey{Bytes: sodium.Bytes(pkBytes)}, nil
}

func ExportBoxPublicKey(pk *sodium.BoxPublicKey) string {
	return hex.EncodeToString(pk.Bytes[:])
}

func ImportBoxKP(kpHex string) (*sodium.BoxKP, error) {
	kpBytes, err := hex.DecodeString(kpHex)
	if err != nil {
		return nil, err
	}
	if len(kpBytes) != 64 {
		return nil, fmt.Errorf("Expected 32 bytes for key pair")
	}
	pk := sodium.BoxPublicKey{Bytes: sodium.Bytes(kpBytes[:32])}
	sk := sodium.BoxSecretKey{Bytes: sodium.Bytes(kpBytes[32:])}
	return &sodium.BoxKP{PublicKey: pk, SecretKey: sk}, nil
}

func ExportBoxKP(kp *sodium.BoxKP) string {
	return hex.EncodeToString(kp.PublicKey.Bytes[:]) + hex.EncodeToString(kp.SecretKey.Bytes[:])
}

func GenKP() sodium.BoxKP {
	return sodium.MakeBoxKP()
}

func Encrypt(pk *sodium.BoxPublicKey, data []byte) []byte {
	return sodium.Bytes(data).SealedBox(*pk)
}

func Decrypt(kp *sodium.BoxKP, encData []byte) ([]byte, error) {
	return sodium.Bytes(encData).SealedBoxOpen(*kp)
}
