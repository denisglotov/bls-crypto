package test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.digiu.ai/blockchainlaboratory/bls-crypto/bls"
)

func Test_EncryptDecrypt(t *testing.T) {
	pass := "password"
	priv0, _ := bls.GenerateRandomKeyForTest()

	protectedKey, err := priv0.Encrypt(pass)
	require.NoError(t, err)
	t.Log(protectedKey)

	decr, err := bls.Decrypt([]byte(protectedKey), pass)
	require.NoError(t, err)
	require.Equal(t, decr, priv0.Marshal())

}
