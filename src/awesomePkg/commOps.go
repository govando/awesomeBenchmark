package main

/*
import (
	"gopkg.in/mgo.v2"
	"os"
	"strings"
	"time"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"bytes"
)


//'n' pruebas de comunicacion para cada tamaño
func commTest_emptyCount(size uint32, f *os.File)  {


	//var buffer bytes.Buffer
	var times[] float64
	f.WriteString("\nn\tTamano(bytes)\tTiempo(ms)\n")

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()


	conn := mgoSession.DB(db).C(coll)

	data := strings.Repeat("a", int(size))
	for i := 0; i < n_pruebas; i++ {
		tini := time.Now()
		count, _ = conn.Find(bson.M{"_id": data}).Count()
		total_time = time.Since(tini).Nanoseconds()

		_, err = f.WriteString(fmt.Sprintf("%d\t%d\t%f\n",i,size,float64(total_time)/float64(1000000)))
		check(err)

		//no almaceno los mayores a un milisegundos (outliers de 20+ms)
		if total_time < 1000000 {
			times = append(times,float64(total_time)/float64(1000000))
		}
	}

	average(times)
	fmt.Println("----------------------")
}

func commTest_emptyFind(size uint32, f *os.File)  {


	//var buffer bytes.Buffer
	var times[] float64
	f.WriteString("\nn\tTamano(bytes)\tTiempo(ms)\n")

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()


	conn := mgoSession.DB(db).C(coll)

	data := strings.Repeat("a", int(size))
	for i := 0; i < n_pruebas; i++ {
		tini := time.Now()
		conn.Find(bson.M{"_id": data})
		total_time = time.Since(tini).Nanoseconds()

		_, err = f.WriteString(fmt.Sprintf("%d\t%d\t%f\n",i,size,float64(total_time)/float64(1000000)))
		check(err)

		//no almaceno los mayores a un milisegundos (outliers de 20+ms)
		if total_time < 1000000 {
			times = append(times,float64(total_time)/float64(1000000))
		}

	}



	average(times)
	fmt.Println("----------------------")
}

func commTest_emptyUpdate(size uint32, f *os.File)  {


	//var buffer bytes.Buffer
	var times[] float64
	f.WriteString("\nn\tTamano(bytes)\tTiempo(ms)\n")

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()


	conn := mgoSession.DB(db).C(coll)

	data := strings.Repeat("a", int(size))
	for i := 0; i < n_pruebas; i++ {
		tini := time.Now()
		_, err := conn.UpdateAll(bson.M{"_id": data},bson.M{"cmp": "1"})
		total_time = time.Since(tini).Nanoseconds()

		_, err = f.WriteString(fmt.Sprintf("%d\t%d\t%f\n",i,size,float64(total_time)/float64(1000000)))
		check(err)

		//no almaceno los mayores a un milisegundos (outliers de 1ms+)
		if total_time < 1000000 {
			times = append(times,float64(total_time)/float64(1000000))
		}

	}


	average(times)
	fmt.Println("----------------------")
}

func commTest_emptyDelete(size uint32, f *os.File)  {


	//var buffer bytes.Buffer
	var times[] float64
	f.WriteString("\nn\tTamano(bytes)\tTiempo(ms)\n")

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()


	conn := mgoSession.DB(db).C(coll)

	data := strings.Repeat("a", int(size))
	for i := 0; i < n_pruebas; i++ {
		tini := time.Now()
		_, err = conn.RemoveAll(bson.M{"_id": data})
		total_time = time.Since(tini).Nanoseconds()

		_, err = f.WriteString(fmt.Sprintf("%d\t%d\t%f\n",i,size,float64(total_time)/float64(1000000)))
		check(err)

		//no almaceno los mayores a un milisegundos (outliers de 20+ms)
		if total_time > 1000000 {
			times = append(times,float64(total_time)/float64(1000000))
		}

	}

	//_, err = f.WriteString(buffer.String())
	//check(err)

	average(times)
	fmt.Println("----------------------")
}

//se rechaza la inserción del elemento debido a restriccion de Validación de la colección
func commTest_emptyInsert(size uint32, f *os.File)  {



	var times[] float64
	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()


	conn := mgoSession.DB(db).C(coll)

	data := strings.Repeat("a", int(size))
	doc := &Doc{
		Data:     data,
	}

	f.WriteString("\nn\tTamano(bytes)\tTiempo(ms)\n")
	for i := 0; i < n_pruebas; i++ {
		tini := time.Now()
		err = conn.Insert(doc)
		total_time = time.Since(tini).Nanoseconds()

		_, err = f.WriteString(fmt.Sprintf("%d\t%d\t%f\n",i,size,float64(total_time)/float64(1000000)))
		check(err)

		//no almaceno los mayores a un milisegundos (outliers de 20+ms)
		if total_time < 1000000 {
			times = append(times,float64(total_time)/float64(1000000))
		}

	}

	average(times)
	fmt.Println("----------------------")
}




//

func commTest_emptyBulkFind(size uint32, f *os.File) {

	var buffer bytes.Buffer
	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	f.WriteString("\nn\tTamano(bytes)\tTiempo(ms)\n")
	data := strings.Repeat("a", (int)(size))

	bulk := mgoSession.DB(db).C(coll).Bulk()

	for i := 0; i < n_pruebas; i++ {

		update := bson.M{"_id": data}
		selector := bson.M{"_id": " "}

		bulk.Update(selector, update)

		count++
		//bulk.Run si: se tienen 1000 elementos || el tamaño no debe superar los 16MB || se deben enviar menos de 1000 elementos
		if  i%1000==0 || 16000000<(count+1)*int (tamannos[i]) || i+1 == n_pruebas {
			tini := time.Now()
			_, err = bulk.Run()
			total_time = time.Since(tini).Nanoseconds()
			check(err)
			_, err = buffer.WriteString(fmt.Sprintf("%d\t%d\t%f\n",i,size,float64(total_time)/float64(1000000)))
			check(err)
			count=0
		}

	}

}

*/

/*


func commTest_emptyBulkFind(size uint32, f *os.File) {

	var buffer bytes.Buffer
	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	f.WriteString("\nn\tTamano(bytes)\tTiempo(ms)\n")
	data := strings.Repeat("a", (int)(size))

	bulk := mgoSession.DB(db).C(coll).Bulk()

	for i := 0; i < n_pruebas; i++ {

		update := bson.M{"_id": data}
		selector := bson.M{"_id": " "}

		bulk.Update(selector, update)

		count++
		//bulk.Run si: se tienen 1000 elementos || el tamaño no debe superar los 16MB || se deben enviar menos de 1000 elementos
		if  i%1000==0 || 16000000<(count+1)*int (tamannos[i]) || i+1 == n_pruebas {
			tini := time.Now()
			_, err = bulk.Run()
			total_time = time.Since(tini).Nanoseconds()
			check(err)
			_, err = buffer.WriteString(fmt.Sprintf("%d\t%d\t%f\n",i,size,float64(total_time)/float64(1000000)))
			check(err)
			count=0
		}

	}

}


 */