// Copyright (C) 2019 José Martínez Ruiz <jmmrcp@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// Mlab Schema MLabs database
	Mlab struct {
		ID         primitive.ObjectID `bson:"_id" json:"id"`
		Test       int                `bson:"test" json:"test"`
		Categoria  string             `bson:"categoria" json:"categoria"`
		Temas      []int              `bson:"temas" json:"temas"`
		Titulo     string             `bson:"titulo" json:"titulo,"`
		Ord        int                `bson:"ord" json:"ord"`
		Pregunta   string             `bson:"pregunta" json:"pregunta"`
		Respuestas []string           `bson:"respuestas" json:"respuestas"`
		Articulos  string             `bson:"articulos" json:"articulos"`
		Fecha      time.Time          `bson:"fecha" json:"fecha"`
		Box        int                `bson:"box" json:"box"`
	}
)

// Parse convert Objet to Array
func (mlab *Mlab) Parse() []string {
	data := []string{Decrypt(mlab.Pregunta)}
	mlab.Respuestas[0] = Decrypt(mlab.Respuestas[0])
	data = append(data, mlab.Respuestas...)
	data = append(data, mlab.Articulos, mlab.ID.Hex())
	return data
}

// Decrypt AES 256 cripto
func Decrypt(encryptedString string) (decryptedString string) {

	var keyString = os.Getenv("KEY")

	key, _ := hex.DecodeString(keyString)
	enc, _ := hex.DecodeString(encryptedString)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}
