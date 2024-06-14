# Saludos en go 

Este paquete proporciona una forma simple de obtener saludos personalizados en go

## Instalacion 
Ejecute los siguientes camandos para instalar el paquete:
```bash
go get -u github.com/ErnestoMolina/greetings
```
## Uso
Aqui tienes un ejemplo de como utilizar el paquete en tu codigo:

```go
package main

import (
	"fmt"
	"github.com/ErnestoMolina/greetings"
)

func main() {

	message, err := greetings.Hello("Juanito")

	if err != nil {
		fmt.Println("Ocurrio un error:", err)
		retunr
	}

	fmt.Println(message)
}
```
Este ejemplo importa el paquete github.com/ErnestoMolina/greetings y llama a la funcion Hello el cual imprime 
un saludo personalizado, si ocurre un error, se imprime el error.