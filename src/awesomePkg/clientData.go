package main

import "fmt"
/*
type clientData_interface interface {
	resumenComm()
}
*/
type clientData struct {
	CountComm	int
	FindComm	int
	FindIdComm	int
	UpdateComm	int
	DeleteComm	int
	InsertComm	int

	Count	int
	Find	int
	FindId	int
	Update	int
	Delete	int
	Insert	int

}


func resumenComm(clientes[] Client)  {
	fmt.Println("--------  COMUNICACION  -------------")
	fmt.Println("Operaciones totales de los clientes:")
	fmt.Printf("Client\tCount\tFind\tFindId\tInsert\tUpdate\tDelete\n")
	for i, _ := range clientes{
		fmt.Printf("%d\t%d\t%d\t%d\t%d\t%d\t%d\n",i,clientes[i].totalOp.CountComm,
			clientes[i].totalOp.FindComm, clientes[i].totalOp.FindIdComm,
			clientes[i].totalOp.InsertComm, clientes[i].totalOp.UpdateComm,
			clientes[i].totalOp.DeleteComm)
	}
	fmt.Println("--------------------------")
}
func resumenBenchmark(clientes[] Client)  {
	fmt.Println("---------  BENCHMARK  --------------")
	fmt.Println("Operaciones totales de los clientes:")
	fmt.Printf("Client\tCount\tFind\tFindId\tInsert\tUpdate\tDelete\n")
	for i, _ := range clientes{
		fmt.Printf("%d\t%d\t%d\t%d\t%d\t%d\t%d\n",i,clientes[i].totalOp.Count,
			clientes[i].totalOp.Find, clientes[i].totalOp.FindId,
			clientes[i].totalOp.Insert, clientes[i].totalOp.Update,
			clientes[i].totalOp.Delete)
	}
	fmt.Println("--------------------------")
}