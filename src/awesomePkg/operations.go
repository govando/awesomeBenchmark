package main

import (
	"os"
	"time"
	"gopkg.in/mgo.v2"
	"fmt"
	"strings"
)



func insert_op(query_bson interface{}, coll* mgo.Collection) int64  {
	tini := time.Now()
	err := coll.Insert(query_bson)
	total_time := time.Since(tini).Nanoseconds()
	//_, err = f.WriteString(fmt.Sprintf("%d\t%f\n",size,float64(total_time)/float64(1000000)))
	if err!=nil {
		if strings.Contains(err.Error(),"duplicate key")==false {
			fmt.Println(" ===> ",err)
			check(err)
		}  else{
			//fmt.Println("error insert : ",err)
		}
	}
	return total_time
}

func find_op(query_bson interface{},  coll* mgo.Collection) int64 {
	var query *mgo.Query

	tini := time.Now()

	err := coll.Find(query_bson).One(&query)
	fmt.Println("--------------------------->query:",query)

	total_time := time.Since(tini).Nanoseconds()
	//_, err = f.WriteString(fmt.Sprintf("%d\t%f\n",size,float64(total_time)/float64(1000000)))
	check(err)
	checkQuery(query)
	return total_time
}

func find_op2(query_bson interface{},  coll* mgo.Collection) (int64, Doc) {
	//var query *mgo.Query
	var doc Doc
	tini := time.Now()
	err := coll.Find(query_bson).One(&doc)
	if (err != nil) {
		fmt.Println("error:", err, " ==> ", doc)
	}
	//err := coll.Find(query_bson).One(&query)

	total_time := time.Since(tini).Nanoseconds()
	//_, err = f.WriteString(fmt.Sprintf("%d\t%f\n",size,float64(total_time)/float64(1000000)))

	check(err)
	//checkQuery(query)
	return total_time, doc
}

func findId_op(query_bson interface{}, f *os.File, size uint32, collection string, coll* mgo.Collection) int64 {
	var results interface{}
	tini := time.Now()
	coll.FindId(query_bson).One(&results)
	total_time := time.Since(tini).Nanoseconds()
	//_, err := f.WriteString(fmt.Sprintf("%d\t%f\n",size,float64(total_time)/float64(1000000)))
	//check(err)
	return total_time
}

func update_op(query_bson interface{},update_bson interface{}, coll* mgo.Collection) int64  {
	tini := time.Now()
	err := coll.Update(query_bson, update_bson)
	total_time := time.Since(tini).Nanoseconds()
	//_, err = f.WriteString(fmt.Sprintf("%d\t%f\n",size,float64(total_time)/float64(1000000)))
	check(err)
	return total_time
}


func delete_op(query_bson interface{}, coll* mgo.Collection) int64 {
	tini := time.Now()
	_, err := coll.RemoveAll(query_bson)
	total_time := time.Since(tini).Nanoseconds()
	//_, err = f.WriteString(fmt.Sprintf("%d\t%f\n",size,float64(total_time)/float64(1000000)))
	check(err)
	return total_time
}

func count_op(query_bson interface{}, coll* mgo.Collection) int64 {
	tini := time.Now()
	_, err := coll.Find(query_bson).Count()
	total_time := time.Since(tini).Nanoseconds()
	//_, err = f.WriteString(fmt.Sprintf("%d\t%f\n",size,float64(total_time)/float64(1000000)))
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


	exist := 0

	sesion := mgoSession.DB(db)

	names, err := sesion.CollectionNames()
	if err != nil {
		fmt.Println("Failed to get coll names: %v", err)
		return
	}

	// Simply search in the names slice, e.g.
	for _, name := range names {
		if name == collCommm {
			fmt.Println("The collection exists!")
			exist=1
			break
		}
	}

	if exist == 0{
		// Create a Collection
		err := sesion.C(collCommm).Create(&mgo.CollectionInfo{
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

func cleanColl(collName string){
	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()

	coll := mgoSession.DB(db).C(collName)
	coll.RemoveAll(nil)
}

func dropColl(coll *mgo.Collection)  {
	coll.DropCollection()
}

func tryConnection()  {

	mgoSession, err := mgo.Dial(host)
	for err!=nil  {
		fmt.Println("tratando de conectar")
		mgoSession, err = mgo.Dial(host)
		time.Sleep(1 * time.Second)
	}

	mgoSession.Close()
}

