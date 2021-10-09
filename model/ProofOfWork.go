package model

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

const difficulty = 18

type ProofOfWork struct {
	Block  *block
	Target *big.Int
}

func NewProof(b *block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-difficulty))

	pow := &ProofOfWork{
		Block:  b,
		Target: target,
	}

	return pow
}

// Este método será el que genere el hash del bloque
// esta función será la que sustituya a la que hicimos de los bloques CalculateHash()
func (pow *ProofOfWork) InitData(nonce int) []byte {

	// usaremos el bytes.Join
	data := bytes.Join(
		[][]byte{
			// Meteremos el prevHash y el dato
			pow.Block.PrevHash,
			pow.Block.Data,
			// Adicionalmente meteremos el nonce y la dificultad
			// Acordarse de cuando hablamos del algoritmo de proof of work
			// Crear el hash del dato + el nonce
			// Para simplificar las cosas crearemos una nueva
			// función que explicaremos a continuación
			ToHex(int64(nonce)),
			ToHex(int64(difficulty)),
		},
		[]byte{},
	)

	return data
}

// De un int64 obtendremos un slice de bytes simplemente
func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	// Iniciamos el contador
	nonce := 0

	// Haremos una especie de do while (Si venís de otro lenguaje)
	// En este loop prepararemos nuestro dato y
	// luego lo hashearemos a sha-256
	// Seguidamente convertiremos ese Hash a un biginteger
	// Por ultimo compararemos ese biginteger generado con el del target
	// Que estará dentro de nuestro struct de proof of work
	for nonce < math.MaxInt64 {
		// Llamaremos a nuestro InitData para preparar el dato
		data := pow.InitData(nonce)
		// Los hashearemos
		hash = sha256.Sum256(data)

		// Con fines de demostración hacemos un log en pantalla
		fmt.Printf("\r%x", hash)

		// haremos una copia del slice
		intHash.SetBytes(hash[:])

		// Ahora compararemos el hash generado con el del target
		if intHash.Cmp(pow.Target) == -1 {
			// Si el hash generado es menor nos salimos del bucle
			// Ya que esto quiere decir que hemos podido firmar el bloque
			break
		}
		// De otra forma seguimos incrementando el contador para repetir el proceso
		nonce++
	}
	// hacemos un salto de línea para separar trazas
	fmt.Println()

	// retornamos el contador y una copia del slice
	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}
