package main

import (
	"os"
	"time"
	"gopkg.in/mgo.v2"
	"fmt"
)



func insert_op(query_bson interface{}, f *os.File, size uint32, collection string, conn* mgo.Collection) int64  {
	tini := time.Now()
	err := conn.Insert(query_bson)
	total_time := time.Since(tini).Nanoseconds()
	_, err = f.WriteString(fmt.Sprintf("%d\t%f\n",size,float64(total_time)/float64(1000000)))
	check(err)
	return total_time
}

func find_op(query_bson interface{}, f *os.File, size uint32, collection string, conn* mgo.Collection) int64 {
	var query *mgo.Query
	tini := time.Now()
	err := conn.Find(query_bson).One(&query)
	total_time := time.Since(tini).Nanoseconds()
	_, err = f.WriteString(fmt.Sprintf("%d\t%f\n",size,float64(total_time)/float64(1000000)))
	check(err)
	checkQuery(query)
	return total_time
}

func findId_op(query_bson interface{}, f *os.File, size uint32, collection string, conn* mgo.Collection) int64 {
	tini := time.Now()
	query := conn.FindId(query_bson)
	total_time := time.Since(tini).Nanoseconds()
	_, err := f.WriteString(fmt.Sprintf("%d\t%f\n",size,float64(total_time)/float64(1000000)))
	check(err)
	checkQuery(query)
	return total_time
}

func update_op(query_bson interface{},update_bson interface{}, f *os.File, size uint32, collection string, conn* mgo.Collection) int64 {
	tini := time.Now()
	_, err := conn.UpdateAll(query_bson, update_bson)
	total_time := time.Since(tini).Nanoseconds()
	_, err = f.WriteString(fmt.Sprintf("%d\t%f\n",size,float64(total_time)/float64(1000000)))
	check(err)
	return total_time
}


func delete_op(query_bson interface{}, f *os.File, size uint32, collection string, conn* mgo.Collection) int64 {
	tini := time.Now()
	_, err := conn.RemoveAll(query_bson)
	total_time := time.Since(tini).Nanoseconds()
	_, err = f.WriteString(fmt.Sprintf("%d\t%f\n",size,float64(total_time)/float64(1000000)))
	check(err)
	return total_time
}

func count_op(query_bson interface{}, f *os.File, size uint32, collection string, conn* mgo.Collection) int64 {
	tini := time.Now()
	_, err := conn.Find(query_bson).Count()
	total_time := time.Since(tini).Nanoseconds()
	_, err = f.WriteString(fmt.Sprintf("%d\t%f\n",size,float64(total_time)/float64(1000000)))
	check(err)
	return total_time
}

// Crea test.emptyColl con un validador (para no insertar ningún documento)
// Si ya existe, no hace nada
// Uso exclusivo para pruebas de comunicación
func create_CommTestColl()  {

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()

	db := "test"
	coll := "emptyColl"

	exist := 0

	conn := mgoSession.DB(db)

	names, err := conn.CollectionNames()
	if err != nil {
		fmt.Println("Failed to get coll names: %v", err)
		return
	}

	// Simply search in the names slice, e.g.
	for _, name := range names {
		if name == coll {
			fmt.Println("The collection exists!")
			exist=1
			break
		}
	}

	if exist == 0{
		// Create a Collection
		err := conn.C(coll).Create(&mgo.CollectionInfo{
			DisableIdIndex:false, Capped: false, ValidationLevel: "strict", Validator: invalidDoc,
		})
		if err != nil {
			panic(err)
		} else {
			fmt.Println("Empty collection created")
			fmt.Println("----------------------")
		}
	}

	time.Sleep(1000)

}

func cleanColl(coll string){
	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()

	conn := mgoSession.DB(db).C(coll)
	conn.RemoveAll(nil)

}