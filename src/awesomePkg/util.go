package main

import (
	"fmt"
	"bytes"
	"gopkg.in/mgo.v2"
	"runtime"
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
	fmt.Println("La aplicación usará",runtime.NumCPU()," CPU" )

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

func concat(s1 string, s2 string)  string {
	var buffer bytes.Buffer
	buffer.WriteString(s1)
	buffer.WriteString(s2)
	return buffer.String()
}

func checkQuery(query* mgo.Query)  {

}