package main

import "github.com/elcorderista/godesdecero/middleware"

func main() {

	/* estado, texto := variables.ConviertoTexto(1588)
	fmt.Println(texto)
	fmt.Println(estado)
	if os := runtime.GOOS; os == "Linux." || os == "darwin" {
		fmt.Println("Esto no es Windows")
		fmt.Println(os)
	} else {
		fmt.Println("Esto es Windows")
		fmt.Println(os)
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es linux")
	case "darwin":
		fmt.Println("Esto es darwing")
	default:
		fmt.Printf("%s \n", os)
	}

	numero, texto := ejercicios.Ejercicio01Resuelto("101")
	fmt.Println(numero)
	fmt.Println(texto)

	teclado.IngresoNumeros()


	iteraciones.Iterar()
	iteraciones.IterarReverso()
	iteraciones.IterarSaltos()
	fmt.Println(ejercicios.Multiplicar())
	files.GrabaTabla()

	files.SumaTabla()

	files.LeoFile()

	funciones.Calculos()

	funciones.LlamarClouser()

	funciones.Exponencia(2)*/

	//arreglos_slices.MuestroArreglos()
	//arreglos_slices.MuestraSlice()
	//arreglos_slices.Capacidad()
	//mapas.MostrarMapas()
	//users.AltaUsuario()
	//defer_panic.VemosDefer()

	//defer_panic.EjemploPanic()
	//canal1 := make(chan bool)
	//go goroutines.MiNombreLento("elcorderista Medrano", canal1)
	//
	//defer func() {
	//	<-canal1
	//}()

	//We indicate wait for chanel 1 to terminate
	//webserver.MiWebServer()
	middleware.Middleware()
}
