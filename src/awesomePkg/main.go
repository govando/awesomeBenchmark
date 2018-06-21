package main


func main() {

	//var c
//    client()

	clientes := make([]Client, 1)

	for i, _ := range clientes{
		clientes[i].num_cliente = i
	}


    create_CommTestColl()

	//clientes[0].testComm_emptyCount()
    //testComm_emptyFind()
    //testComm_emptyUpdate()
    //testComm_emptyDelete()

    //pruebas de comunicación
    clientes[0].testComm_emptyCount()
	clientes[0].testComm_emptyFind()
	clientes[0].testComm_emptyFindId()
	clientes[0].testComm_emptyUpdate()
	clientes[0].testComm_emptyDelete()
	clientes[0].testComm_emptyInsert()

	//pruebas de benchmark
	clientes[0].InsertOne()
	clientes[0].Count()
	clientes[0].FindOne()
	clientes[0].FindIdOne()
	clientes[0].UpdateOne()
	clientes[0].DeleteOne()



    //

	//err := c.Insert(&Person{Name: "Ale", Phone: "+55 53 1234 4321", Timestamp: time.Now()},
	//	&Person{Name: "Cla", Phone: "+66 33 1234 5678", Timestamp: time.Now()})
/*
	data := strings.Repeat("a", 64)
	fmt.Println("len: ", len(data))
	var q Document
	//q.Id = 0
	q.Data = data

	time.Sleep(10)
	tini := time.Now()
	query,_ := conn.Find(bson.M{"_id": data}).Count()
	//conn.C(coll).Find(q)
	total :=  time.Since(tini).Nanoseconds()
	fmt.Println("time connection: ", total)
	fmt.Println("----------------------")

	fmt.Print(query)
	//fmt.Println("sizeof uint32: ", unsafe.Sizeof(data), " bytes \n sizeof data: ", unsafe.Sizeof(data), " bytes")
*/


}

func create_empty_ix_collection(){

}

