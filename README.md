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
# Ejecución del código hasta ahora

Si quieres ver el estado del repositorio hasta ahora ve a [este commit](https://github.com/MaQuiNa1995/Go-BlockChain/tree/f8e0bf7d63d0334eac3b3153a8dbbc5e5cb057b9)

Nuestro main sacará por consola esto:

```
Bloque:
        Previous Hash:
        Data in Block: Genesis
        Hash: 81ddc8d248b2dccdd3fdd5e84f0cad62b08f2d10b57f9a831c13451e5c5c80a5

Bloque:
        Previous Hash: 81ddc8d248b2dccdd3fdd5e84f0cad62b08f2d10b57f9a831c13451e5c5c80a5
        Data in Block: 1º Block Despues del Genesis
        Hash: 11f27e8a6ee5b1b8d1eada2d6ce758bd7028d86b47dcac4ac27b202eaeedead2

Bloque:
        Previous Hash: 11f27e8a6ee5b1b8d1eada2d6ce758bd7028d86b47dcac4ac27b202eaeedead2
        Data in Block: 2º Block Despues del Genesis
        Hash: d81bbc87021060a1925f297a98fe3b6236481fe42e82c856bb42ea119b3f72bf

Bloque:
        Previous Hash: d81bbc87021060a1925f297a98fe3b6236481fe42e82c856bb42ea119b3f72bf
        Data in Block: 3º Block Despues del Genesis
        Hash: 3f6628fe789a6518e1dfe77075e9f8f88028def45d281580b9f26d0590ee8317
```

No importa la veces que lo ejecutemos siempre será lo mismo, por lo tanto al ejecutar varias veces este main obtendras varias copias de la blockchain la manera de saber si está corrupto es comparando los hashes de las diferentes copias 

# Qué es Prueba de trabajo / Proof of Work (PoW)

Hay diferentes Algoritmos de consenso (Consensus algorithms / Proof Algorithms)
* Proof Of Work
* Proof Of Steak
* y mas...

# Proof Of Work

Básicamente lo que quiere decir este es que forzamos a la red a realizar trabajo para añadir un bloque a la blockchain

Este "trabajo" es computacional, cuando hablamos de mineros minando Bitcoin nos referimos a esta "Prueba de trabajo" para añadir bloques a la blockchain la razón por la que estos consiguen bitcoins es esencialmente porque potencian a la red a escribir mas rápido ese bloque

Adicionalmente hacen que el dato de los bloques sea mas seguro, el Proof of work viene de la mano de la validación de esa prueba que es cuando un usuario hace el trabajo necesario para añadir ese bloque se requiere que demuestre ese trabajo realizado de ahi el nombre de "Prueba de trabajo" (Proof of Work)

Un concepto importante es que el trabajo realizado debe ser dificil pero la demostración del mismo debe ser relativamente fácil

## Algoritmo de Proof Of Work

Los pasos a seguir van a ser:
* Coger el dato del bloque
* Crear un contador que empiece en el 0 (Llamado nonce) que será incrementado en +1 teóricamente infinitas veces
* Crear el hash del dato + el nonce
* Verificar el hash para ver si cumple determinados requerimientos de aqui viene la llamada "dificultad"
	* Requerimientos:
		* Los primeros bytes deben contener 0

Digamos que quieres escribir un bloque, si el hash de ese nuevo bloque no cumple los requerimientos tendrás que volver a generar el hash ese reintento es a lo que se le llama dificultad teniendo que recrear otro hash para ver si cumple con los requerimientos

En la proof of Work original de bitcoin la especificación se llama "Hash cash"

En la dificultad original se requería que 20 bytes consecutivos del hash fueran 0, con el paso del tiempo esa dificultad se ha incrementado por lo tanto se requiere de mas trabajo para poder escribir un bloque

Podemos aumentar la dificultad por ejemplo haciendo que el requerimiento de 20 ceros pase a 50 

## Dificultad

Para empezar definiremos una constante para definir la dificultad: `const difficulty = 18`

Tendremos una dificultad estática en nuestra prueba de concepto pero si quieres crear un algoritmo de blockchain de verdad tendrás que crear algún tipo de función que incremente la dificultad de poco en poco dentro de un periodo de tiempo largo, básicamente quieres que esto suceda para aumentar el número de mineros en la red y la potencia creciente de los ordenadores que pudieran venir en el futuro. Queremos que el tiempo de minado de un bloque sea uniforme 

## Creando el Struct
```
type ProofOfWork struct {
        // Representa un puntero a un bloque
	Block  *block
	// Representa un puntero que es el requerimiento que hemos definido arriba 
	// para entender esto necesitas saber como los bytes se comportan en el ordenador
	Target *big.Int
}
```
[Big Int en Go](https://golang.org/pkg/math/big/)
[Bytes en Go](https://golang.org/pkg/bytes/)

## Creando Proof of work

```
func NewProof(b *block) *ProofOfWork {
        // Crearemos nuestro target 
	target := big.NewInt(1)
	// Despues tendremos que coger el 256 que es el número de bytes de los hashes
	// y extraer la dificultad de ellos , para despues hacer un left shift (Lsh) de los bytes de ese número
	target.Lsh(target, uint(256-difficulty))
	// Ahora cogemos ese valor al que le hemos hecho el left shifted
	// y lo metemos a nuestro struct
	pow := &ProofOfWork{
		Block: b,
		Target: target,
	}
	// Despues lo retornamos
	return pow
}
```

Con está función desde un puntero a un bloque obtendríamos un puntero a una prueba de trabajo
