package main

import (
	"sync"
	"os"
)

func main() {

	var wg sync.WaitGroup
	var admin Admin
	parserInput(os.Args) // nClientes (default:10)
	environment() // Imprimo Nro CPUs a utilizar

	clientes := make([]Client, nClientes)
	for i, _ := range clientes{
		clientes[i].num_cliente = i
	}

	// Lógica de ejecución //

/*

    admin.createCommDB() //si no existe, se crea coleccion para testear comunicación
    //--- Pruebas de comunicación para cada primitiva
	for i, _ := range clientes{
		go clientes[i].testComm_emptyCount(&wg)
	}
	wg.Wait()

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

	resumenComm(clientes)
*/

	//pruebas de benchmark de primitivas
	admin.cleanCollection(collBench)
	wg.Add(nClientes)
	for i, _ := range clientes{
		go clientes[i].InsertOne(&wg)
	}
	wg.Wait()

	wg.Add(nClientes)
	for i, _ := range clientes{
		go clientes[i].Count(&wg)
	}
	wg.Wait()
	wg.Add(nClientes)
	for i, _ := range clientes{
		go clientes[i].FindOne(&wg)
	}
	wg.Wait()
	wg.Add(nClientes)
	for i, _ := range clientes{
		go clientes[i].FindIdOne(&wg)
	}
	wg.Wait()
	wg.Add(nClientes)
	for i, _ := range clientes{
		go clientes[i].UpdateOne(&wg)
	}
	wg.Wait()
	wg.Add(nClientes)
	for i, _ := range clientes{
		go clientes[i].DeleteOne(&wg)
	}
	wg.Wait()

	resumenBenchmark(clientes)
}

