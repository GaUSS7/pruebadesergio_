package main

import (
	"fmt"
	"strings"
)

func main() {
	/* en este juego tines que saber la cantida de laetras que lleva los
	nombres como tambien sin tienn una cantidad de letras en el y aparte
	tienes que saber si hay suphijos y pregijos*/
	questionString1 := "Mi nombre es Gabriel"
	questionString2 := "Las Salamandras no son serpientes marinas"
	questionString3 := "Gabriel Jose"

	//priemra respuesta
	asnwer1 := len(questionString1)

	//segunda respuesta
	asnwer2 := strings.Count(questionString2, "S")
	asnwer2_1 := strings.Count(questionString2, "s")
	asnwer2_2 := asnwer2 + asnwer2_1

	//tercera respuesta
	asnwerPrefi := strings.HasPrefix(questionString3, "Gabriel")
	asnwerSufi := strings.HasSuffix(questionString3, "Jose")

	//respuestas  del jugador de turno
	score := 0 //puntos del jugador
	respuesta := 0
	respuesta2 := 0
	var respuesta3 bool
	var respuesta3_1 bool

	fmt.Println("Pregunta N*1: Cuantos letras hay, contando los espacios, de: ", questionString1)
	fmt.Scanln(&respuesta)
	if respuesta == asnwer1 {
		fmt.Println("felicidades, es correcto, tiene: ", asnwer1, " letras")
		score++
	} else {
		fmt.Println("Lo sentimos. no es correcto. Tu ingresaste: ", respuesta, " Y el resultado era: ", asnwer1)
		fmt.Println("No te precupes aun te quedan 2 preguntas mas.")
	}
	fmt.Print("\n")

	fmt.Println("Pregunta N*2: Cuantas veces se repite  le letra ( s ) minuscula y mayuscula en: ", questionString2)
	fmt.Scanln(&respuesta2)
	if asnwer2_2 == respuesta2 {
		fmt.Println("felicidades, es correcto, la letra se llega a repetir: ", asnwer2_2, " veces")
		score++
	} else {
		fmt.Println("Lo sentimos. no es correcto. Tu ingresaste: ", respuesta2, " Y el resultado era: ", asnwer2_2)
		fmt.Println("No te precupes aun te queda 1 preguntas mas.")
	}
	fmt.Print("\n")

	fmt.Println("Pregunta N*3: Es hora de la ultima preguanta, y la pregunta es, este nombre tiene Sufijo: ", questionString3)
	fmt.Println(" 1.) S1  o  0)No :")
	fmt.Scanln(&respuesta3)
	fmt.Println("Y tambien llevara Prefijo ? : ")
	fmt.Println(" 1.) S1  o  0)No :")
	fmt.Scanln(&respuesta3_1)
	if respuesta3 == asnwerPrefi && respuesta3_1 == asnwerSufi {
		fmt.Println("felicidades es correcto.")
		score++
	} else {
		fmt.Println("Lo sentimos. no es correcto.")
	}

	fmt.Print("\n")
	fmt.Println("")
	if score == 3 {
		fmt.Print("PUNTAJE MAXIMO. AMAZING")
	} else if score == 2 {
		fmt.Print("Tu puntaje es de 2 puntos. No esta mal, puedes hacerlo mejor")
	} else if score == 1 {
		fmt.Print("TU puntaje es de de solo 1 punto, no desanimes, lo puedes hacer mejor")
	} else {
		fmt.Print("Tu puedes...")
	}
}
