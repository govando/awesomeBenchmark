package main

import (
	"os"
	"time"
	"gopkg.in/mgo.v2"
	"fmt"
)

var (
	query		*mgo.Query
	count		int
	err			error
	total_time	int64
)

func insert(query_bson interface{}, f *os.File, size uint32, collection string)  {


	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	conn := mgoSession.DB(db).C(collection)

	f.WriteString("\nn\tTamano(bytes)\tTiempo(ms)\n")
	for i := 0; i < n_pruebas; i++ {
		tini := time.Now()
		err = conn.Insert(query_bson)
		total_time = time.Since(tini).Nanoseconds()
		_, err = f.WriteString(fmt.Sprintf("%d\t%d\t%f\n",i,size,float64(total_time)/float64(1000000)))
		check(err)
	}

	times = append(times,float64(total_time)/float64(1000000))


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