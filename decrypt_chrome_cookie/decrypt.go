package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/user"
	"path"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/pbkdf2"
)

//GetCookies returns decrypted chrome cookies on linux
func GetCookies(hostname string) (ret []*http.Cookie) {
	usr, err := user.Current()
	if err != nil {
		log.Panic(err)

	}
	cookiesDBPath := path.Join(usr.HomeDir, ".config", "google-chrome", "Default", "Cookies")
	db, err := sql.Open("sqlite3", cookiesDBPath)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	rows, err := db.Query("select name, encrypted_value from cookies where host_key like ?", hostname)
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name, encryptedValue string
		rows.Scan(&name, &encryptedValue)

		//Decrypt cookie value
		//On Mac, replace MY_PASS with your password from Keychain
		//On Linux, replace MY_PASS with 'peanuts'
		myPass := []byte("peanuts")
		//Default values used by both Chrome and Chromium in OSX and Linux
		salt := []byte("saltysalt")
		iv := []byte("                ")
		length := 16

		//1003 on Mac, 1 on Linux
		iterations := 1

		key := pbkdf2.Key(myPass, salt, iterations, length, sha1.New)
		block, err := aes.NewCipher(key)
		if err != nil {
			log.Panic(err)

		}
		mode := cipher.NewCBCDecrypter(block, iv)
		valueBytes := make([]byte, 1024)
		mode.CryptBlocks(valueBytes, []byte(encryptedValue[3:]))
		value := string(trim(valueBytes))

		log.Println(value)

		ret = append(ret, &http.Cookie{Name: name, Value: value})
	}
	return
}

func trim(value []byte) []byte {
	l := len(value) - 1
	for ; value[l] == 0 && l >= 0; l-- {
	}
	return value[:l+1]
}

func main() {
	log.SetOutput(os.Stdout)
	for _, c := range GetCookies("%wiki.n.miui.com%") {
		log.Println(c)
	}
}
