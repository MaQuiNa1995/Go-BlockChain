package model

type block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

func CreateBlock(data string, prevHash []byte) *block {

	block := &block{
		Hash:     []byte{},
		Data:     []byte(data),
		PrevHash: prevHash,
		Nonce:    0, // inicializamos el nonce a 0
	}

	// creamos la prueba de trabajo
	pow := NewProof(block)

	// Y la iniciamos
	nonce, hash := pow.Run()

	// Cuando hayamos completado la prueba de trabajo
	// podremos rellenar el nonce y el hash obtenido
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func Genesis() *block {
	return CreateBlock("Genesis", []byte{})
}
