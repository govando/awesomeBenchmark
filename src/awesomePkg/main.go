package main

import (
	"sync"
	"os"
	"strings"
	"fmt"
)

//input programa:
// param1: nro clientes  param2: nro de operacion
func main() {

	var wg sync.WaitGroup
	var admin Admin
	parserInput(os.Args) // nClientes (default:10)
	environment()        // Imprimo Nro CPUs a utilizar
	tryConnection()      //esperar hasta que mongo este levantado

	flgEjecuta := 0
	clientes := make([]Client, nClientes)
	for i, _ := range clientes {
		clientes[i].num_cliente = i
	}

	// Lógica de ejecución //
	//--- Pruebas de comunicación para cada primitiva

	if strings.Compare("countComm", Operacion) == 0 {
		admin.createCommDB() //si no existe, se crea coleccion para testear comunicación
		wg.Add(nClientes)
		for i, _ := range clientes {
			go clientes[i].testComm_emptyCount(&wg)
		}
		wg.Wait()
		resumenComm(clientes)
		flgEjecuta = 1
	}

	/*
	wg.Add(nClientes)
	for i, _ := range clientes{
		go clientes[i].testComm_emptyFind(&wg)
	}
	wg.Wait()

	wg.Add(nClientes)
	for i, _ := range clientes{
		go clientes[i].testComm_emptyFindId(&wg)
	}
	wg.Wait()


	wg.Add(nClientes)
	for i, _ := range clientes{
		go clientes[i].testComm_emptyUpdate(&wg)
	}
	wg.Wait()

	wg.Add(nClientes)
	for i, _ := range clientes{
		go clientes[i].testComm_emptyDelete(&wg)
	}
	wg.Wait()

	wg.Add(nClientes)
	for i, _ := range clientes{
		go clientes[i].testComm_emptyInsert(&wg)
	}
	wg.Wait()
*/


	if strings.Compare("insertToTest", Operacion) == 0 {
		admin.cleanDB()
		fmt.Println("insert - To - Test (size) => ",tamannos[idColl])
		wg.Add(nClientes)
		for i, _ := range clientes {
			go clientes[i].InsertToTest(&wg)
		}
		wg.Wait()
		flgEjecuta = 2
	}


	//
	if strings.Compare("insertSinColl", Operacion) == 0 {
		admin.cleanDB()
		fmt.Println("Insert sin Colección (size) => ",tamannos[idColl])
		wg.Add(nClientes)
		for i, _ := range clientes {
			go clientes[i].InsertSinColl(&wg)
		}
		wg.Wait()
		flgEjecuta = 2
	}
	if strings.Compare("insertSinDoc", Operacion) == 0 {
		admin.cleanDB()
		fmt.Println("Insert sin Documento (size) => ",tamannos[idColl])
		wg.Add(nClientes)
		for i, _ := range clientes {
			go clientes[i].InsertOne(&wg)
		}
		wg.Wait()
		flgEjecuta = 2
	}
	if strings.Compare("insertErrorPK", Operacion) == 0 {
		admin.cleanDB()
		fmt.Println("insertErrorPK (size) => ",tamannos[idColl])
		wg.Add(nClientes)
		for i, _ := range clientes {
			go clientes[i].InsertErrorPK(&wg)
		}
		wg.Wait()
		flgEjecuta = 2
	}

	if strings.Compare("count", Operacion) == 0 {
		wg.Add(nClientes)
		fmt.Println("Count (size) => ",tamannos[idColl])
		for i, _ := range clientes {
			go clientes[i].Count(&wg)
		}
		wg.Wait()
		flgEjecuta = 2
	}

	if strings.Compare("findOne", Operacion) == 0 {
		wg.Add(nClientes)
		fmt.Println("findOne (size) => ",tamannos[idColl])
		for i, _ := range clientes {
			go clientes[i].FindOne(&wg)
		}
		wg.Wait()
		flgEjecuta = 2
	}

	if strings.Compare("findIdOne", Operacion) == 0 {
		wg.Add(nClientes)
		fmt.Println("findIdOne (size) => ",tamannos[idColl])
		for i, _ := range clientes {
			go clientes[i].FindIdOne(&wg)
		}
		wg.Wait()
		flgEjecuta = 2
	}

	if strings.Compare("updateOne", Operacion) == 0 {
		wg.Add(nClientes)
		fmt.Println("updateOne (size) => ",tamannos[idColl])
		for i, _ := range clientes {
			go clientes[i].UpdateOne(&wg)
		}
		wg.Wait()
		flgEjecuta = 2
	}

	if strings.Compare("updateOneId", Operacion) == 0 {
		wg.Add(nClientes)
		fmt.Println("updateOneId (size) => ",tamannos[idColl])
		for i, _ := range clientes {
			go clientes[i].UpdateOneId(&wg)
		}
		wg.Wait()
		flgEjecuta = 2
	}

	if strings.Compare("deleteOne", Operacion) == 0 {
		wg.Add(nClientes)
		fmt.Println("deleteOne (size) => ",tamannos[idColl])
		for i, _ := range clientes {
			go clientes[i].DeleteOne(&wg)
		}
		wg.Wait()
		flgEjecuta = 2
	}

	if strings.Compare("deleteOneId", Operacion) == 0 {
		wg.Add(nClientes)
		fmt.Println("deleteOneId (size) => ",tamannos[idColl])
		for i, _ := range clientes {
			go clientes[i].DeleteOneId(&wg)
		}
		wg.Wait()
		flgEjecuta = 2
	}

	if strings.Compare("deleteOneStatic", Operacion) == 0 {
		wg.Add(nClientes)
		fmt.Println("deleteOne Static (size) => ",tamannos[idColl])
		for i, _ := range clientes {
			go clientes[i].DeleteOne(&wg)
		}
		wg.Wait()
		flgEjecuta = 2
	}

	if strings.Compare("deleteOneIdStatic", Operacion) == 0 {
		wg.Add(nClientes)
		fmt.Println("deleteOneId Static (size) => ",tamannos[idColl])
		for i, _ := range clientes {
			go clientes[i].DeleteOneId(&wg)
		}
		wg.Wait()
		flgEjecuta = 2
	}

	if strings.Compare("GeneraTrazaPrueba", Operacion) == 0 {
		wg.Add(nClientes)
		fmt.Println("GeneraTrazaPrueba => ")
		for i, _ := range clientes {
			go clientes[i].GeneraTraza(&wg)
		}
		wg.Wait()
		flgEjecuta = 3
	}



	if (flgEjecuta == 2) {
		resumenBenchmark(clientes)
	} else if (flgEjecuta == 0){
		fmt.Println("NO INGRESO CORRECTAMENTE LOS PARAMETROS DEL SISTEMA, NO SE EJECUTO EL BENCHMARK")
	}
}

