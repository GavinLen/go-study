package xtao

import (
	"fmt"
	"encoding/json"
	"net/http"
	"strings"
)

func AuthHeaderFromClient(req *http.Request) (error, *UserAccountInfo) {
	signedHeaders := []string{}

	// 如果 request 的 Header 中包含
	if _, exists := req.Header["Content-Type"]; exists {
		signedHeaders = append(signedHeaders, "Content-Type")
	}

	signedHeaders = append(signedHeaders, AccountHeader)

	options := Options {
		SignedHeaders: signedHeaders,
		SecretKey: KeyLocator(func(apiKey string) (string, string) {
			AesAkey := NewAesEncrypt(BIOFLOW_API_KEY_SEED)
			AesSkey := NewAesEncrypt(BIOFLOW_SECURITY_KEY_SEED)

			user, err := AesAkey.Decrypt(apiKey)
			if err != nil {
				return "", ""
			}

			if strings.ToUpper(user) == "ROOT" {
				ServerLogger.Debugf("user:%s,Skey:%s \n",user, globalAdminSkey)
				return globalAdminSkey, user
			}

			skey, err := AesSkey.Encrypt(user)
			if err != nil {
				return "", user
			}
			return skey, user
		}),
	}

	err, user, skey := HMACAuth(options, req)
	if err != nil {
		ServerLogger.Errorf("authenticate failure: %s\n",
			err.Error())
		return err, nil
	}

	ServerLogger.Debugf("user:%s,Skey:%s \n",user, skey)

	dkey := skey
	if len(dkey) < 16 {
		dkey = fmt.Sprintf("%16s", skey)
	}

	aesEnc := NewAesEncrypt(dkey)
	enAccountHeader := req.Header.Get(AccountHeader)
	accountHeader, err := aesEnc.Decrypt(enAccountHeader)
	if err != nil {
		ServerLogger.Errorf("Decrypt header failure: %s\n",
			err.Error())
		return err, nil
	}

	account := UserAccountInfo{}
	err = json.Unmarshal([]byte(accountHeader), &account)

	if err != nil {
		ServerLogger.Errorf("Failed to parse JSON account:%s\n",
			accountHeader)
		return err, nil
	}

	return nil, &account
}
