package main

import "github.com/Franogales/divide-inversion/second/web"

func main() {
	/* el siguiente comentario solamente inicia el router de gin
	utilizo gin porque permite http1 y http2, y ya tiene el recover de los panic
	*/
	web.ServerRun()
}
