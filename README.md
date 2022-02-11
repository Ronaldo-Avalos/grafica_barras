<<<<<<< HEAD
![](img/oco-logo.png)

# Gráfico de Barras

El propósito de este ejercicio es conocer tus habilidades de programación, puedes utilizar el lenguaje que prefieras. Tu solución deberá mostrar el uso de los principios de la programación orientada a objetos. Es decir, deberá estar organizado en clases con métodos y atributos.


## Requerimientos funcionales

### Parte 1
Si los datos son 1, 2, 3, 4, 5, "+" la salida deberá ser algo parecido a:
```
│    +
│   ++
│  +++
│ ++++
│+++++
└─────
 ```
(cada digito conmo parametro sera la altlulra de su barra)

### Ejercicio 2

Modifica el programa para soportar un número variable de ancho y alto para cada dígito.
Por ejemplo, para ancho = 3 y alto = 2 el dígito 2 sería:

```
│    *
│    *
│    *
│*   *
│*  **
│*  **
└─────
 ```
=======
![ICI LOGO](img/ici_logo.png)

# Gráfico de Barras

El propósito de este ejercicio es conocer tus habilidades de programación, puedes utilizar el lenguaje que prefieras. Tu solución deberá mostrar el uso de los principios de la programación orientada a objetos. Es decir, deberá estar organizado en clases con métodos y atributos.


## Requerimientos funcionales

### Parte 1
Si los datos son 1, 2, 3, 4, 5, "+" la salida deberá ser algo parecido a:
```
│    +
│   ++
│  +++
│ ++++
│+++++
└─────
 ```
(cada digito conmo parametro sera la altlulra de su barra)

### Ejercicio 2

Modifica el programa para soportar un número variable de ancho y alto para cada dígito.
Por ejemplo, para ancho = 3 y alto = 2 el dígito 2 sería:

```
│    *
│    *
│    *
│*   *
│*  **
│*  **
└─────
 ```

 ## pantallazos de su ejecución con al menos dos series de argumentos distintos.

## Resultado 1
 ```GO
 
func main() {

	var num = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	var symbol = "█"
	numax := findMax(num)
	grafica(numax, num, symbol)

}
 ```
 ### Resultado de terminal
 ```
 PS C:\Users\Admin\workspace\Golang\grafica_barras> go run grafica_barras.go
 
 |█
 |██
 |███
 |████
 |█████
 |██████
 |███████
 |████████
 |█████████
 |██████████
  ¯¯¯¯¯¯¯¯¯¯
 ```
 ## Resultado 2
 ```
 func main() {

	var num = []int{4, 6, 2, 0, 5, 7, 1}
	var symbol = "▍"
	numax := findMax(num)
	grafica(numax, num, symbol)

}
```
### Resultado de terminal
```
PS C:\Users\Admin\workspace\Golang\grafica_barras> go run grafica_barras.go
       
 |     ▍
 | ▍   ▍
 | ▍  ▍▍
 |▍▍  ▍▍
 |▍▍  ▍▍
 |▍▍▍ ▍▍
 |▍▍▍ ▍▍▍
  ¯¯¯¯¯¯¯
```

>>>>>>> f42bce6 (addimg)
