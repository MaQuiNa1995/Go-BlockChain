# Que es Blockchain

Es una base de datos publica que está distribuida en múltiples nodos

Todos los datos que entren deben de ser confiable por todos los nodos

Podrías por ejemplo tener un 49% de los nodos que produjesen datos erróneos o malintencionados y la red podría recuperarse de ese desajuste

Un Blockchain implica multiples bloques que contienen la información que queremos en nuestra base de datos

Struct de un blockchain
```
type blockChain struct {
	blocks []*block
}
```
En este struct básicamente tenemos un slice de punteros de bloques

# Que es el Block

Básicamente son los objetos que conforman un blockchain este tiene que tener 3 básicos como mínimo

Atributos
* Hash del propio bloque
* Hash del último bloque creado (Es el que nos permite enlazar bloques)
* El dato que guardamos pueden ser imagenes textos numeros etc 
* 
Struct de un bloque básico
```
type block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}
```

## Hash

Para el calculo del hash como standard se usa el algoritmo de encriptado SHA-256 debido a su equilibrio entre coste computacional y solidez si quieres aprender mas sobre este algoritmo de encriptación: [Pincha Aqui](https://academy.bit2me.com/sha256-algoritmo-bitcoin/)

Para calcular el Hash usaremos este método:
```
func (b *block) CalculateHash() {
// Explicado mas abajo
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
  // Lo encriptamos 
	hash := sha256.Sum256(info)
  // Creamos una copia y se la asignamos al hash del struct
	b.Hash = hash[:]
}
```
### Función bytes.Join

Lo que hace es juntar los slice de datos que se le pasan como primer parámetro `[][]byte{b.Data, b.PrevHash}` (Pueden ser cualquier cantidad) teniendo como separador el 2º parámetro `[]byte{}` (En este caso vacío)

un ejemplo:
```
1º Parámetro
AAAA
BBBB
2º Parámetro
CC
Resultado:
AAAACCBBBB
```
[Documentación función bytes.Join](https://www.includehelp.com/golang/bytes-join-function-with-examples.aspx)

## Creación de un bloque

Para crear un bloque deberías de llamar a esta función para asegurarte de que se calcula el hash por lo tanto sería conveniente hacer "privado" el struct del bloque para que nadie pueda instanciar un bloque de otra forma que no sea a traves de esta función 

```
func CreateBlock(data string, prevHash []byte) *block {
// Creamos normal un struct
	block := &block{
		Hash:     []byte{},
    // Aqui adicionalmente pasamos de string a bytes
		Data:     []byte(data),
		PrevHash: prevHash,
	}
  // Llamamos a la función que creamos previamente
	block.CalculateHash()
  // Lo retornamos
	return block
}
```

## Añadiendo un bloque a la blockchain

para añadir un bloque a la blockchain debemos usar la anterior funcion

```
// Recibimos una blockchain
func (chain *blockChain) AddBlock(data string) {
// Cogemos el ultimos bloque
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
  // A traves de la función de antes creamos el nuevo bloque
	newBlock := CreateBlock(data, prevBlock.Hash)
  // Se lo añadimos a la blockchain
	chain.Blocks = append(chain.Blocks, newBlock)
}
```
### El bloque "Genesis"

Como hemos visto siempre referenciamos al hash del anterior bloque pero que pasa con el primer bloque este es imposible que pueda tener ningun hash previo ya que este es el primero, a este bloque se le llama "Genesis Block" que representa el primer bloque de la blockchain

Lo crearemos a traves de este método:
```
func Genesis() *block {
	return CreateBlock("Genesis", []byte{})
}
```
## Creación de una blockchain

Para crear la blockchain debemos usar la función anterior de tal manera:
```
func InitBlockChain() *blockChain {
	return &blockChain{[]*block{Genesis()}}
}
```
Si quieres ver el estado del repositorio hasta ahora ve a [este commit del tag v1.0.0](https://github.com/MaQuiNa1995/Go-BlockChain/tree/d07ba85dabcd2d7a4e5eea4cde54cc016571908d)

