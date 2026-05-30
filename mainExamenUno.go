// Gael Santiago Barrera Medrano - 25760069
package main

import (
	"fmt"
	"math/rand"
)

// hacer una, carta tiene color y numero
type Carta struct {
	Color string
	Valor string
}

// hacer un jugador con nombre y sus cartas en la mano
type Jugador struct {
	Nombre string
	Mano   []Carta
}

// estructura del juego tiene con el mazo, los jugadores, la pila para descartar las cartas y de quien es el turno
type Juego struct {
	Mazo         []Carta
	Jugadores    []Jugador
	PilaDescarte []Carta
	Turno        int
}

// Aqui se crean todas las cartas y las mezcla
func (j *Juego) CrearMazo() {
	colores := []string{"Rojo", "Verde", "Azul", "Amarillo"}
	numeros := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for _, color := range colores {
		for _, numero := range numeros {
			j.Mazo = append(j.Mazo, Carta{Color: color, Valor: numero})
		}
	}

	// mezclar el mazo
	rand.Shuffle(len(j.Mazo), func(i, k int) {
		j.Mazo[i], j.Mazo[k] = j.Mazo[k], j.Mazo[i]
	})
}

// reparte 5 cartas a cada jugador
func (j *Juego) RepartirCartas() {
	for i := 0; i < 5; i++ {
		for k := range j.Jugadores {
			carta := j.RobarCarta()
			j.Jugadores[k].Mano = append(j.Jugadores[k].Mano, carta)
		}
	}

	// la primera carta del mazo va a la pila de descarte para empezar
	j.PilaDescarte = append(j.PilaDescarte, j.RobarCarta())
}

// saca la carta de hasta arriba del mazo
func (j *Juego) RobarCarta() Carta {
	carta := j.Mazo[0]
	j.Mazo = j.Mazo[1:] // quitamos la primera carta del mazo
	return carta
}

// revisa si una carta se puede jugar comparandola con la carta de la pila
func cartaEsValida(cartaJugador Carta, cartaPila Carta) bool {
	mismoColor := cartaJugador.Color == cartaPila.Color
	mismoNumero := cartaJugador.Valor == cartaPila.Valor
	return mismoColor || mismoNumero
}

// muestra las cartas del jugador con numeros para elegir
func mostrarCartas(jugador Jugador) {
	fmt.Println("Tus cartas:")
	for i, carta := range jugador.Mano {
		fmt.Println(" ", i+1, ")", carta.Color, carta.Valor)
	}
}

// el turno de un jugador
func (j *Juego) JugarTurno() {
	jugador := &j.Jugadores[j.Turno]
	cartaEnMesa := j.PilaDescarte[len(j.PilaDescarte)-1] // la carta de hasta arriba de la pila

	fmt.Println("\n----------------------------------------")
	fmt.Println("Turno de:", jugador.Nombre)
	fmt.Println("Carta en mesa:", cartaEnMesa.Color, cartaEnMesa.Valor)
	mostrarCartas(*jugador)

	// buscamos si tiene alguna carta que pueda jugar
	puedeJugar := false
	for _, carta := range jugador.Mano {
		if cartaEsValida(carta, cartaEnMesa) {
			puedeJugar = true
			break
		}
	}

	// si no puede jugar, roba una carta
	if !puedeJugar {
		fmt.Println("No tienes cartas para jugar, robas una carta...")
		nuevaCarta := j.RobarCarta()
		jugador.Mano = append(jugador.Mano, nuevaCarta)
		fmt.Println("Robaste:", nuevaCarta.Color, nuevaCarta.Valor)
		return
	}

	// el jugador elige que carta jugar
	var eleccion int
	for {
		fmt.Print("Elige el numero de la carta que quieres jugar: ")
		fmt.Scan(&eleccion)
		eleccion-- // restamos 1 porque el jugador ve desde 1 pero el slice empieza en 0

		// verificamos que la eleccion este en rango
		if eleccion < 0 || eleccion >= len(jugador.Mano) {
			fmt.Println("Numero invalido, intenta de nuevo")
			continue
		}

		cartaElegida := jugador.Mano[eleccion]

		// verificamos que la carta sea valida
		if !cartaEsValida(cartaElegida, cartaEnMesa) {
			fmt.Println("Esa carta no se puede jugar, elige otra")
			continue
		}

		// se juega la carta, la quitamos de la mano y la ponemos en la pila
		j.PilaDescarte = append(j.PilaDescarte, cartaElegida)
		jugador.Mano = append(jugador.Mano[:eleccion], jugador.Mano[eleccion+1:]...)
		fmt.Println("Jugaste:", cartaElegida.Color, cartaElegida.Valor)
		break
	}
}

// cambia al siguiente jugador
func (j *Juego) SiguienteTurno() {
	j.Turno = (j.Turno + 1) % len(j.Jugadores)
}

// revisa si algun jugador ya no tiene cartas (gano)
func (j *Juego) HayGanador() bool {
	for _, jugador := range j.Jugadores {
		if len(jugador.Mano) == 0 {
			fmt.Println("\n", jugador.Nombre, "gano el juego!, felicidades prro")
			return true
		}
	}
	return false
}

func main() {
	juego := Juego{
		Jugadores: []Jugador{
			{Nombre: "Jugador 1"},
			{Nombre: "Jugador 2"},
		},
	}

	juego.CrearMazo()
	juego.RepartirCartas()

	// el juego sigue hasta que alguien gane
	for {
		juego.JugarTurno()

		if juego.HayGanador() {
			break
		}

		juego.SiguienteTurno()
	}
}
