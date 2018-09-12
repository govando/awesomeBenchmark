package main

import (

	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"

	"os"
	_"time"
	_"fmt"
	"strconv"
	"strings"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"sync"

	"fmt"
	"encoding/json"
)

//  Benchmark de latencia de comunicación
//  operacion + colecion (comm o bench) + size request + request
//
//  query_bson := bson.M{"data": data, "cmp1": i}

type Client struct {
	num_cliente int //identificador
	timestamp 	uint64

	totalOp		clientData
}

func (c *Client) testComm_emptyCount(wg *sync.WaitGroup)  {

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()

	//borrar y crear el archivo de datos
	f, err := os.Create( "./data/Comm/count/dataCommCount"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	header, err :=  json.Marshal(Header{N_pruebas : n_pruebas, Tamannos:tamannos})
	_, err = f.WriteString(fmt.Sprintf(string(header) ))

	for _, size := range tamannos {
		f.WriteString("\nn;Tamano(bytes);Tiempo(ms)\n")
		coll := mgoSession.DB(db).C(fmt.Sprintf("%v_%d",collCommm,size))
		data := strings.Repeat("a", int(size))
		query_bson := bson.M{"_id": data}
		for i := 0; i < n_pruebas; i++ {
			total_time := count_op(query_bson , coll)
			if i > 3 {
				_, err = f.WriteString(fmt.Sprintf("%d;%d;%d\n",i-3,size,total_time))
				times = append(times, float64(total_time)/float64(1000000))
			}
			c.totalOp.CountComm+=1
		}
		uso_tiempos(times) //agregar metodos aquí
	}

	wg.Done()
}

func (c *Client) testComm_emptyFind(wg *sync.WaitGroup)  {

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	//borrar y crear el archivo de datos
	f, err := os.Create("./data/Comm/find/dataCommFind"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	header, err :=  json.Marshal(Header{N_pruebas : n_pruebas, Tamannos:tamannos})
	_, err = f.WriteString(fmt.Sprintf(string(header) ))

	for _, size := range tamannos {
		f.WriteString("\nn;Tamano(bytes);Tiempo(ms)\n")
		coll := mgoSession.DB(db).C(fmt.Sprintf("%v_%d",collCommm,size))
		data := strings.Repeat("a", int(size))
		query_bson := bson.M{"b": data}
		for i := 0; i < n_pruebas; i++ {
			total_time:= find_op(query_bson , coll)
			if i > 3 {
				_, err = f.WriteString(fmt.Sprintf("%d;%d;%d\n",i-3,size,total_time))
				times = append(times, float64(total_time)/float64(1000000))
			}
			c.totalOp.FindComm+=1
		}
		uso_tiempos(times) //agregar metodos aquí
	}
	wg.Done()
}

func (c *Client) testComm_emptyFindId(wg *sync.WaitGroup)  {

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	//borrar y crear el archivo de datos
	f, err := os.Create("./data/Comm/findId/dataCommFindId"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	header, err :=  json.Marshal(Header{N_pruebas : n_pruebas, Tamannos:tamannos})
	_, err = f.WriteString(fmt.Sprintf(string(header) ))

	for _, size := range tamannos {
		f.WriteString("\nn;Tamano(bytes);Tiempo(ms)\n")
		coll := mgoSession.DB(db).C(fmt.Sprintf("%v_%d",collCommm,size))
		data := strings.Repeat("a", int(size))
		query_bson := bson.M{"_id": data}
		for i := 0; i < n_pruebas; i++ {
			total_time := findId_op(query_bson , f, size, collCommm, coll)
			if i > 3 {
				_, err = f.WriteString(fmt.Sprintf("%d;%d;%d\n",i-3,size,total_time))
				times = append(times, float64(total_time)/float64(1000000))
			}
			c.totalOp.FindIdComm+=1
		}
		uso_tiempos(times) //agregar metodos aquí
	}
	wg.Done()
}

func (c *Client) testComm_emptyUpdate(wg *sync.WaitGroup)  {

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()

	//borrar y crear el archivo de datos
	f, err := os.Create("./data/Comm/update/dataCommUpdate"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	header, err :=  json.Marshal(Header{N_pruebas : n_pruebas, Tamannos:tamannos})
	_, err = f.WriteString(fmt.Sprintf(string(header) ))

	for _, size := range tamannos {
		f.WriteString("\nn;Tamano(bytes);Tiempo(ms)\n")
		coll := mgoSession.DB(db).C(fmt.Sprintf("%v_%d",collCommm,size))
		data := strings.Repeat("a", int(size))
		query_bson := bson.M{"_id": data}
		update_bson := bson.M{"_id": "0"}
		for i := 0; i < n_pruebas; i++ {
			total_time := update_op(query_bson ,update_bson,  coll)
			if i > 3 {
				_, err = f.WriteString(fmt.Sprintf("%d;%d;%d\n",i-3,size,total_time))
				times = append(times, float64(total_time)/float64(1000000))
			}
			c.totalOp.UpdateComm+=1
		}
		uso_tiempos(times) //agregar metodos aquí
	}
	wg.Done()
}

func  (c *Client) testComm_emptyDelete(wg *sync.WaitGroup)  {

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()

	//borrar y crear el archivo de datos
	f, err := os.Create("./data/Comm/delete/dataCommDelete"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	header, err :=  json.Marshal(Header{N_pruebas : n_pruebas, Tamannos:tamannos})
	_, err = f.WriteString(fmt.Sprintf(string(header) ))

	for _, size := range tamannos {
		f.WriteString("\nn;Tamano(bytes);Tiempo(ms)\n")
		coll := mgoSession.DB(db).C(fmt.Sprintf("%v_%d",collCommm,size))
		data := strings.Repeat("a", int(size))
		query_bson := bson.M{"_id": data}
		for i := 0; i < n_pruebas; i++ {
			total_time := delete_op(query_bson, coll)
			if i > 3 {
				_, err = f.WriteString(fmt.Sprintf("%d;%d;%d\n",i-3,size,total_time))
				times = append(times, float64(total_time)/float64(1000000))
			}
			c.totalOp.DeleteComm+=1
		}
		uso_tiempos(times) //no hace nada aun
	}
	wg.Done()
}

func  (c *Client) testComm_emptyInsert(wg *sync.WaitGroup){

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()

	//borrar y crear el archivo de datos
	f, err := os.Create("./data/Comm/insert/dataCommInsert"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()

	header, err :=  json.Marshal(Header{N_pruebas : n_pruebas, Tamannos:tamannos})
	_, err = f.WriteString(fmt.Sprintf(string(header) ))

	for _, size := range tamannos {
		f.WriteString("\nn;Tamano(bytes);Tiempo(ms)\n")
		coll := mgoSession.DB(db).C(fmt.Sprintf("%v_%d",collCommm,size))
		data := strings.Repeat("a", int(size))
		query_bson := bson.M{"_id": data}
		for i := 0; i < n_pruebas; i++ {
			total_time := insert_op(query_bson, coll)
			if i > 3 {
				_, err = f.WriteString(fmt.Sprintf("%d;%d;%d\n",i-3,size,total_time))
				times = append(times, float64(total_time)/float64(1000000))
			}
			c.totalOp.InsertComm+=1
		}
		uso_tiempos(times) //no hace nada aun
	}
	wg.Done()
}



func createFile(filename string) *os.File {

	var f *os.File
	var err error
	var header []byte

	if(idColl==0) {
		f, err = os.Create(filename)
		header, err =  json.Marshal(Header{N_pruebas : n_pruebas, Tamannos:tamannos})
		_, err = f.WriteString(fmt.Sprintf(string(header) ))
	} else {
		f, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	}
	check(err)

	return f
}

// Benchmarks de latencias por operacion
// Modificar aquí los query_bson
/*------------------------------------------------*/

func  (c *Client) InsertSinColl(wg *sync.WaitGroup){

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()

	filename := "./data/Query/insert/dataInsertCreateColl" + strconv.Itoa(c.num_cliente)
	f := createFile(filename)
	defer f.Close()

	//for _, size := range tamannos {

	size:=tamannos[idColl]
	data := strings.Repeat("a", int(size))
	f.WriteString("\nn;Tamano(bytes);Tiempo(ms)\n")

	for i := 0; i < n_pruebas; i++ {
		coll := mgoSession.DB(db).C(fmt.Sprintf("%v_%d",collBench,size))
		query_bson := bson.M{"data":data ,"cliente":c.num_cliente,"cmp1": i, "cmp2": i, "cmp3": i}
		total_time := insert_op(query_bson, coll)
		dropColl(coll)
		if i > 3 {
			_, _ = f.WriteString(fmt.Sprintf("%d;%d;%d\n",i-3,size,total_time))
			times = append(times, float64(total_time)/float64(1000000))
		}
		c.totalOp.Insert+=1.
	}
	uso_tiempos(times)
	//}
	wg.Done()
}


func  (c *Client) InsertOne(wg *sync.WaitGroup){

	var times[] float64
//	var wg2 sync.WaitGroup
//	var wg3 sync.WaitGroup

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()

	//borrar y crear el archivo de datos
	filename := "./data/Query/insert/dataInsertCreateDoc"+strconv.Itoa(c.num_cliente)
	f := createFile(filename)
	defer f.Close()

	//for _, size := range tamannos {

	size:=tamannos[idColl]

	data := strings.Repeat("a", int(size))
	f.WriteString("\nn;Tamano(bytes);Tiempo(ms)\n")
	coll := mgoSession.DB(db).C(fmt.Sprintf("%v_%d",collBench,size))

	for i := 0; i < n_pruebas; i++ {
		query_bson := bson.M{"data":data ,"cliente":c.num_cliente,"cmp1": i, "cmp2": i, "cmp3": i}
		total_time := insert_op(query_bson, coll)
		if i > 3 {
			_, _ = f.WriteString(fmt.Sprintf("%d;%d;%d\n",i-3,size,total_time))
			times = append(times, float64(total_time)/float64(1000000))
		}
		c.totalOp.Insert+=1.
	}
	dropColl(coll)

	uso_tiempos(times)
	//}
	wg.Done()
}


func  (c *Client) InsertErrorPK(wg *sync.WaitGroup){

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()


	//borrar y crear el archivo de datos
	filename := "./data/Query/insert/dataInsertErrorPK"+strconv.Itoa(c.num_cliente)
	f := createFile(filename)
	defer f.Close()

	//for _, size := range tamannos {

	size:=tamannos[idColl]
	coll := mgoSession.DB(db).C(fmt.Sprintf("%v_%d",collBench,size))
	data := strings.Repeat("a", int(size))
	f.WriteString("\nn;Tamano(bytes);Tiempo(ms)\n")
	for i := 0; i < n_pruebas; i++ {
		coll = mgoSession.DB(db).C(fmt.Sprintf("%v_%d",collBench,size))
		query_bson := bson.M{"_id":1, "data": data ,"cliente":c.num_cliente,"cmp1": i, "cmp2": i, "cmp3": i}
		total_time := insert_op(query_bson, coll)

		if i > 3 {
			_, _ = f.WriteString(fmt.Sprintf("%d;%d;%d\n",i-3,size,total_time))
			times = append(times, float64(total_time)/float64(1000000))
		}
		c.totalOp.Insert+=1.
	}
	dropColl(coll)
	uso_tiempos(times)

		//}
	wg.Done()
}



func  (c *Client)  InsertToTest(wg *sync.WaitGroup)  {
	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()


	size:=tamannos[idColl]
	coll := mgoSession.DB(db).C(fmt.Sprintf("%v_%d",collBench,size))
	dropColl(coll)

	data := strings.Repeat("a", int(size))

	for i := 0; i < n_pruebas; i++ {
		query_bson := bson.M{"_id": i, "data": data, "cliente": c.num_cliente, "cmp1": i, "cmp2": i, "cmp3": i}
		insert_op(query_bson, coll)

		c.totalOp.Insert+=1
	}
	wg.Done()
}

func  (c *Client) FindOne(wg *sync.WaitGroup){

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()

	//borrar y crear el archivo de datos

	filename := "./data/Query/findId/dataFind"+strconv.Itoa(c.num_cliente)
	f := createFile(filename)
	defer f.Close()

	//for _, size := range tamannos {

	size:=tamannos[idColl]
	coll := mgoSession.DB(db).C(fmt.Sprintf("%v_%d",collBench,size))
	//data := strings.Repeat("a", int(size))
	f.WriteString("\nn;Tamano(bytes);Tiempo(ms)\n")


	for i := 0; i < n_pruebas; i++ {
		//query_bson := bson.M{"data": data,"cliente":c.num_cliente, "cmp1": i}
		query_bson := bson.M{"cmp1": i}
		total_time := find_op(query_bson,coll)
		if i > 3 {
			f.WriteString(fmt.Sprintf("%d;%d;%d\n",i-3,size,total_time))
			times = append(times, float64(total_time)/float64(1000000))
		}
		c.totalOp.Find+=1
	}

	wg.Done()
}

func  (c *Client) FindIdOne(wg *sync.WaitGroup){

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()

	//borrar y crear el archivo de datos
	filename := "./data/Query/findId/dataFindId"+strconv.Itoa(c.num_cliente)
	f := createFile(filename)
	defer f.Close()

	size:=tamannos[idColl]
	coll := mgoSession.DB(db).C(fmt.Sprintf("%v_%d",collBench,size))
	//data := strings.Repeat("a", int(size))

	f.WriteString("\nn;Tamano(bytes);Tiempo(ms)\n")
	for i := 0; i < n_pruebas; i++ {
		//query_bson := bson.M{"data": data, "cmp1": i}
		query_bson := bson.M{"_id": i}
		total_time := findId_op(query_bson, f, size,collBench, coll)
		if i > 3 {
			f.WriteString(fmt.Sprintf("%d;%d;%d\n",i-3,size,total_time))
			times = append(times, float64(total_time)/float64(1000000))
		}
		c.totalOp.FindId+=1
	}

	wg.Done()
}

func (c *Client) UpdateOne(wg *sync.WaitGroup){

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()

	filename := "./data/Query/update/dataUpdate"+strconv.Itoa(c.num_cliente)
	f := createFile(filename)
	defer f.Close()

	size:=tamannos[idColl]
	coll := mgoSession.DB(db).C(fmt.Sprintf("%v_%d",collBench,size))
	data := strings.Repeat("b", int(size))


	f.WriteString("\nn;Tamano(bytes);Tiempo(ms)\n")
	for i := 0; i < n_pruebas; i++ {
		query_bson := bson.M{"cmp1": i}
		update_bson :=bson.M{"data":data ,"cliente":c.num_cliente,"cmp1": i, "cmp2": i, "cmp3": i}
		total_time := update_op(query_bson,update_bson, coll)
		if i > 3 {
			f.WriteString(fmt.Sprintf("%d;%d;%d\n",i-3,size,total_time))
			times = append(times, float64(total_time)/float64(1000000))
		}
		c.totalOp.Update+=1
	}

	wg.Done()
}

func (c *Client) UpdateOneId(wg *sync.WaitGroup){

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()

	filename := "./data/Query/update/dataUpdateId"+strconv.Itoa(c.num_cliente)
	f := createFile(filename)
	defer f.Close()

	size:=tamannos[idColl]
	coll := mgoSession.DB(db).C(fmt.Sprintf("%v_%d",collBench,size))
	data := strings.Repeat("b", int(size))


	f.WriteString("\nn;Tamano(bytes);Tiempo(ms)\n")
	for i := 0; i < n_pruebas; i++ {
		query_bson := bson.M{"_id": i}
		update_bson :=bson.M{"data":data ,"cliente":c.num_cliente,"cmp1": i, "cmp2": i, "cmp3": i}
		total_time := update_op(query_bson,update_bson, coll)
		if i > 3 {
			f.WriteString(fmt.Sprintf("%d;%d;%d\n",i-3,size,total_time))
			times = append(times, float64(total_time)/float64(1000000))
		}
		c.totalOp.Update+=1
	}

	wg.Done()
}


func (c *Client) DeleteOne(wg *sync.WaitGroup){

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()

	//borrar y crear el archivo de datos
	filename := "./data/Query/delete/dataDelete"+strconv.Itoa(c.num_cliente)
	f := createFile(filename)
	defer f.Close()

	size:=tamannos[idColl]
	coll := mgoSession.DB(db).C(fmt.Sprintf("%v_%d",collBench,size))
	//data := strings.Repeat("b", int(size))

	f.WriteString("\nn;Tamano(bytes);Tiempo(ms)\n")
	for i := 0; i < n_pruebas; i++ {
		query_bson := bson.M{"cmp1": i}
		total_time := delete_op(query_bson,  coll)
		if i > 3 {
			f.WriteString(fmt.Sprintf("%d;%d;%d\n",i-3,size,total_time))
			times = append(times, float64(total_time)/float64(1000000))
		}
		c.totalOp.Delete+=1
	}

	wg.Done()
}

func (c *Client) DeleteOneId(wg *sync.WaitGroup){

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()

	//borrar y crear el archivo de datos
	filename := "./data/Query/delete/dataDeleteId"+strconv.Itoa(c.num_cliente)
	f := createFile(filename)
	defer f.Close()

	size:=tamannos[idColl]
	coll := mgoSession.DB(db).C(fmt.Sprintf("%v_%d",collBench,size))
	//data := strings.Repeat("b", int(size))

	f.WriteString("\nn;Tamano(bytes);Tiempo(ms)\n")
	for i := 0; i < n_pruebas; i++ {
		query_bson := bson.M{"_id": i}
		total_time := delete_op(query_bson, coll)
		if i > 3 {
			f.WriteString(fmt.Sprintf("%d;%d;%d\n",i-3,size,total_time))
			times = append(times, float64(total_time)/float64(1000000))
		}
		c.totalOp.Delete+=1
	}

	wg.Done()
}


func (c *Client) DeleteOneStatic(wg *sync.WaitGroup){

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()

	//borrar y crear el archivo de datos
	filename := "./data/Query/delete/dataDeleteStatic"+strconv.Itoa(c.num_cliente)
	f := createFile(filename)
	defer f.Close()

	size:=tamannos[idColl]
	coll := mgoSession.DB(db).C(fmt.Sprintf("%v_%d",collBench,size))
	data := strings.Repeat("a", int(size))

	f.WriteString("\nn;Tamano(bytes);Tiempo(ms)\n")
	for i := 0; i < n_pruebas; i++ {
		query_bson := bson.M{"cmp1": i}
		total_time := delete_op(query_bson, coll)

		query_bson = bson.M{"_id": i, "data": data, "cliente": c.num_cliente, "cmp1": i, "cmp2": i, "cmp3": i}
		insert_op(query_bson, coll)

		if i > 3 {
			f.WriteString(fmt.Sprintf("%d;%d;%d\n",i-3,size,total_time))
			times = append(times, float64(total_time)/float64(1000000))
		}
		c.totalOp.Delete+=1
	}

	wg.Done()
}

func (c *Client) DeleteOneIdStatic(wg *sync.WaitGroup){

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()

	//borrar y crear el archivo de datos
	filename := "./data/Query/delete/dataDeleteIdStatic"+strconv.Itoa(c.num_cliente)
	f := createFile(filename)
	defer f.Close()

	size:=tamannos[idColl]
	coll := mgoSession.DB(db).C(fmt.Sprintf("%v_%d",collBench,size))
	data := strings.Repeat("a", int(size))

	f.WriteString("\nn;Tamano(bytes);Tiempo(ms)\n")
	for i := 0; i < n_pruebas; i++ {
		query_bson := bson.M{"_id": i}
		total_time := delete_op(query_bson, coll)

		query_bson = bson.M{"_id": i, "data": data, "cliente": c.num_cliente, "cmp1": i, "cmp2": i, "cmp3": i}
		insert_op(query_bson, coll)

		if i > 3 {
			f.WriteString(fmt.Sprintf("%d;%d;%d\n",i-3,size,total_time))
			times = append(times, float64(total_time)/float64(1000000))
		}
		c.totalOp.Delete+=1
	}

	wg.Done()
}

func  (c *Client) Count(wg *sync.WaitGroup){

	var times[] float64

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()


	//borrar y crear el archivo de datos
	f, err := os.Create("./data/Query/count/dataCount"+strconv.Itoa(c.num_cliente))
	check(err)
	defer f.Close()


	header, err :=  json.Marshal(Header{N_pruebas : n_pruebas, Tamannos:tamannos})
	_, err = f.WriteString(fmt.Sprintf(string(header) ))

	for _, size := range tamannos {
		coll := mgoSession.DB(db).C(fmt.Sprintf("%v_%d",collBench,size))
		data := strings.Repeat("a", int(size))
		f.WriteString("\nn;Tamano(bytes);Tiempo(ms)\n")
		for i := 0; i < n_pruebas; i++ {
			query_bson := bson.M{"data": data,"cliente":c.num_cliente, "cmp1": i}
			total_time := count_op(query_bson, coll)
			if i > 3 {
				_, err = f.WriteString(fmt.Sprintf("%d;%d;%d\n",i-3,size,total_time))
				times = append(times, float64(total_time)/float64(1000000))
			}
			c.totalOp.Count+=1
		}
		uso_tiempos(times)
	}
	wg.Done()
}

type StrPrueba struct {
	Id	int	`json:"id"`
	IdPrueba	int	`json:"idPrueba"`
	NombrePrueba	string	`json:"nombrePrueba"`
	TiempoComunicacion	int64  `json:"tiempoComunicacion"`
	TiempoComputacion	int64  `json:"tiempoComputacion"`
	PesoDatoBytes	uint32	`json:"pesoDatoBytes"`
	CantElemsColeccion int 	`json:"cantElemsColeccion"`
	IdAfectado int 	`json:"idAfectado"`
}

func  (c *Client) GeneraTraza(wg *sync.WaitGroup){

	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()

	filename := "./data/Query/traza/dataTraza"+strconv.Itoa(c.num_cliente)
	f := createFile(filename)
	defer f.Close()

	coll := mgoSession.DB(db).C("traza")
	dropColl(coll)

	cont := 0
	acumSize:=uint32(0)
	size:=uint32(0)
	var query_bson bson.M
	var data string

	for cont < cantDatosIniciales {
		size =uint32(random2(100,16000000)) //tamannos[ random2(0, 29) ]
		fmt.Println("datos iniciales => ",cont,") - size: ", size)
		acumSize=acumSize + size
		data = strings.Repeat("a", int(size))
		query_bson = bson.M{"_id": cont, "data": data, "cliente": c.num_cliente, "cmp1": cont, "cmp2": cont, "cmp3": cont}
		insert_op(query_bson, coll)
		cont=cont+1
	}
	fmt.Println("Carga total => GB ", (float32(acumSize)/float32(1024*1024*1024)) )

	var doc Doc
	tmpComu:=int64(0)
	tmpComp:=int64(0)
	pruebas:=0
	cmpBusqueda:=0
	var update_bson bson.M
	idAfectado:=0

	for pruebas < n_pruebas{
		op := random2(1,11)
		if op==8 || op==9 {
			op=op+2
		}
		//fmt.Println("nom : ",nombrePruebas[op], "pruebas :",pruebas)
		switch op{
			case 1: //InsertCreateColl
				size =uint32(random2(100,16000000))
				acumSize=acumSize + size
				coll = mgoSession.DB(db).C("traza_aux")
				dropColl(coll)

				data = strings.Repeat("a", int(size))
				query_bson = bson.M{"_id": cont, "data": data, "cliente": c.num_cliente, "cmp1": cont, "cmp2": cont, "cmp3": cont}
				tmpComp = insert_op(query_bson, coll)
				fmt.Println("cont:",cont," query_bson:", query_bson["_id"], " coll:",coll)
				idAfectado=cont

				query_bson = bson.M{"_id": data}
				tmpComu = count_op(query_bson, coll)

				dropColl(coll)

			case 2: //InsertCreateDoc
				size =uint32(random2(100,16000000))
				coll = mgoSession.DB(db).C("traza")
				data = strings.Repeat("a", int(size))
				query_bson = bson.M{"_id": cont, "data": data, "cliente": c.num_cliente, "cmp1": cont, "cmp2": cont, "cmp3": cont}
				tmpComp = insert_op(query_bson, coll)
				fmt.Println("cont:",cont," query_bson:", query_bson["_id"], " coll:",coll)
				idAfectado=cont

				query_bson = bson.M{"_id": data}
				tmpComu = count_op(query_bson, coll)

				acumSize=acumSize + size
				cont=cont+1

			case 3:	//InsertErrorPK
				size =uint32(random2(100,16000000))
				acumSize=acumSize + size
				coll = mgoSession.DB(db).C("traza")
				data = strings.Repeat("a", int(size))
				query_bson = bson.M{"_id": random2(0,cont-1), "data": data, "cliente": c.num_cliente, "cmp1": cont, "cmp2": cont, "cmp3": cont}
				tmpComp = insert_op(query_bson, coll)
				idAfectado=cont

				query_bson = bson.M{"_id": data}
				tmpComu = count_op(query_bson, coll)

			case 4://Find
				coll = mgoSession.DB(db).C("traza")
				cmpBusqueda = random2(0,cont-1)
				query_bson = bson.M{"cmp1": cmpBusqueda}
				tmpComp = find_op(query_bson, coll)
				idAfectado=cmpBusqueda

				_,doc = find_op2(query_bson, coll)
				size = uint32( len(doc.Data) )
				data = doc.Data
				query_bson = bson.M{"_id": data}
				tmpComu = count_op(query_bson, coll)

			case 5: //FindId
				coll = mgoSession.DB(db).C("traza")
				cmpBusqueda = random2(0,cont-1)
				query_bson = bson.M{"_id": cmpBusqueda}
				fmt.Println("query_bson:", query_bson, " coll:",coll)
				idAfectado=cmpBusqueda

				tmpComp = find_op(query_bson, coll)

				_,doc = find_op2(query_bson, coll)
				size = uint32( len(doc.Data) )
				data = doc.Data
				query_bson = bson.M{"_id": data}
				tmpComu = count_op(query_bson, coll)

			case 6: //Update
				coll = mgoSession.DB(db).C("traza")

				cmpBusqueda = random2(0,cont-1)
				query_bson = bson.M{"cmp1": cmpBusqueda}
				fmt.Println("query_bson:", query_bson, " coll:",coll)
				idAfectado=cmpBusqueda

				_,doc = find_op2(query_bson, coll)

				size = uint32( len(doc.Data) )
				data = doc.Data
				update_bson =bson.M{"data":data ,"cliente":c.num_cliente,"cmp1": cmpBusqueda, "cmp2": cmpBusqueda, "cmp3": cmpBusqueda}
				tmpComp = update_op(query_bson,update_bson, coll)

				query_bson = bson.M{"_id": data}
				tmpComu = count_op(query_bson, coll)

			case 7: //UpdateId
				coll = mgoSession.DB(db).C("traza")

				cmpBusqueda = random2(0,cont-1)
				query_bson = bson.M{"_id": cmpBusqueda}
				idAfectado=cmpBusqueda

				_,doc = find_op2(query_bson, coll)
				size = uint32( len(doc.Data) )
				data = doc.Data
				update_bson =bson.M{"data":data ,"cliente":c.num_cliente,"cmp1": cmpBusqueda, "cmp2": cmpBusqueda, "cmp3": cmpBusqueda}
				tmpComp = update_op(query_bson,update_bson, coll)

				query_bson = bson.M{"_id": data}
				tmpComu = count_op(query_bson, coll)

			//case 8: //Delete
			//case 9: //DeleteId

			case 10: //DeleteStatic
				coll = mgoSession.DB(db).C("traza")

				cmpBusqueda = random2(0,cont-1)
				query_bson = bson.M{"cmp1": cmpBusqueda}
				idAfectado=cmpBusqueda
				_,doc = find_op2(query_bson, coll)
				size = 55
				//data = doc.Data
				tmpComp = delete_op(query_bson, coll)

				query_bson = bson.M{"_id":  strings.Repeat("a", 55)}
				tmpComu = count_op(query_bson, coll)

				insert_op(doc, coll)


			case 11: //DeleteIdStatic
				coll = mgoSession.DB(db).C("traza")

				cmpBusqueda = random2(0,cont-1)
				query_bson = bson.M{"_id": cmpBusqueda}
				idAfectado=cmpBusqueda

				_,doc = find_op2(query_bson, coll)
				size = 55
				//data = doc.Data
				tmpComp = delete_op(query_bson, coll)

				query_bson =  bson.M{"_id":  strings.Repeat("a", 55)}
				tmpComu = count_op(query_bson, coll)


				insert_op(doc, coll)
		}


		var prb []byte
		prb,_ =  json.Marshal(StrPrueba{ Id:pruebas,
		 								IdPrueba:op,
		 								NombrePrueba :nombrePruebas[op],
		 								TiempoComunicacion:tmpComu,
										TiempoComputacion:tmpComp,
										PesoDatoBytes:size,
										CantElemsColeccion:cont,
										IdAfectado:idAfectado})
		f.WriteString(fmt.Sprintf(string(prb)))
		fmt.Println(string(prb))
		pruebas=pruebas+1
	}

	fmt.Println("Carga total => GB ", (float32(acumSize)/float32(1024*1024*1024)) )

	wg.Done()
}


