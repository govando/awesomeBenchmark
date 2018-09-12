package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"runtime"
	"strconv"
	"math/rand"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func uso_tiempos(times[] float64){
	//de aqui llamar otras funciones que usen tiempos

}


func environment()  {
	fmt.Println("La aplicación usará:\n",runtime.NumCPU()," CPU \n",nClientes," clientes\nOperacion a realizar: ",Operacion )

}

func parserInput(input[] string){
	var error error

	if len(input) == 2 {
		nClientes, error = strconv.Atoi(input[1])
		check(error)
	} else if len(input) == 3 {
		nClientes, error = strconv.Atoi(input[1])
		check(error)
		Operacion = input[2]
		check(error)
	} else if len(input) == 4 {
		nClientes, error = strconv.Atoi(input[1])
		check(error)
		Operacion = input[2]
		check(error)
		n_pruebas,error = strconv.Atoi(input[3])
		n_pruebas += 4
		check(error)
	} else if len(input) == 5 {
		nClientes, error = strconv.Atoi(input[1])
		check(error)
		Operacion = input[2]
		check(error)
		n_pruebas,error = strconv.Atoi(input[3])
		n_pruebas += 4
		check(error)
		idColl,error = strconv.Atoi(input[4])
		check(error)
	}


}

func average(times []float64)  {

	var avg,value float64
	var i int
	for i, value = range times {
		avg += value
	}
	fmt.Printf("total: %f \n",avg/float64(i+1))
}

func randNumbers()  {

}

func sd()  {

}

func checkQuery(query* mgo.Query)  {

}

func random(min, max uint32) uint32 {
	//rand.Seed(time.Now().Unix())
	return uint32(rand.Float32()*float32(max) + float32(min));
}

func random2(min, max int) int {
	//rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min + 1
}