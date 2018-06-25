package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"runtime"
	"strconv"
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
	fmt.Println("La aplicación usará:\n",runtime.NumCPU()," CPU \n",nClientes," clientes" )

}

func parserInput(input[] string){
	var error error

	if len(input) > 1 {
		nClientes, error = strconv.Atoi(input[1])
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

func sd()  {

}

func checkQuery(query* mgo.Query)  {

}