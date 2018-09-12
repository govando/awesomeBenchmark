package main

import (

	"fmt"
	"gopkg.in/mgo.v2"
)

type Admin struct {
	num_cliente int //identificador
	timestamp uint64
}
/*
func (c *Client) testComm_emptyCount()  {

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	conn := mgoSession.DB(db).C(coll)

	//borrar y crear el archivo de datos
	f, err := os.Create( "./data/Comm/count/dataCommCount"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	for _, size := range tamannos {data := strings.Repeat("a", int(size))
		query_bson := bson.M{"_id": data}
		for i := 0; i < n_pruebas; i++ {
			total_time := count_op(query_bson , f, size, collCommm, conn)
			times = append(times, float64(total_time)/float64(1000000))
		}
		uso_tiempos(times) //agregar metodos aquí
	}
}
*/

func (c *Admin) createCommDB()  {
	fmt.Println("Admin: creando DB para pruebas de comunicación")
	create_CommTestColl()
}

func (c * Admin) cleanCollection(coll string){
	cleanColl(coll)
}

func (c * Admin) cleanDB()  {
	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()

	fmt.Println("Admin: Borrando colleciones")
	for _, size := range tamannos {
			coll := mgoSession.DB(db).C(fmt.Sprintf("%v_%d", collBench, size))
			dropColl(coll)
	}

}

func (c *Admin) createIndex()  {

}