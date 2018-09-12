#!/usr/bin/env bash

cantClientes=1
nPruebas=400
ind_pruebas_ejecutar=12
time_to_sleep=5

#restart mongo
sudo killall mongod

sleep $time_to_sleep

#init mongod
sudo mongod -v --logpath /data/mongo/db/logdb.log --dbpath /data/mongo/db  --fork

if ((ind_pruebas_ejecutar <= 0)) ; then

    ############# prueba de comunicación: 1 cliente, operacion countComm ################

    echo " ./awesomePkg $cantClientes countComm "
    ./awesomePkg $cantClientes countComm

fi

if ((ind_pruebas_ejecutar <= 1)) ; then

    ############# insertSinColl #################################
    for((pruebaSizeColl=0;pruebaSizeColl<30; pruebaSizeColl++))
    do
            #restart mongo
            sudo killall mongod

            sleep $time_to_sleep

            #init mongod
            sudo mongod -v --logpath /data/mongo/db/logdb.log --dbpath /data/mongo/db  --fork

            echo " ./awesomePkg $cantClientes insertSinColl $nPruebas $pruebaSizeColl "
            ./awesomePkg $cantClientes insertSinColl $nPruebas $pruebaSizeColl

    done

fi

if ((ind_pruebas_ejecutar <= 2)) ; then

    ############# insertSinDoc #################################
    for((pruebaSizeColl=0;pruebaSizeColl<30; pruebaSizeColl++))
    do
            #restart mongo
            sudo killall mongod

            sleep $time_to_sleep

            #init mongod
            sudo mongod -v --logpath /data/mongo/db/logdb.log --dbpath /data/mongo/db  --fork

            echo " ./awesomePkg $cantClientes insertSinDoc $nPruebas $pruebaSizeColl "
            ./awesomePkg $cantClientes insertSinDoc $nPruebas $pruebaSizeColl

    done

fi

if ((ind_pruebas_ejecutar <= 3)) ; then

    ############# insertErrorPK #################################
    for((pruebaSizeColl=0;pruebaSizeColl<30; pruebaSizeColl++))
    do
            #restart mongo
            sudo killall mongod

            sleep $time_to_sleep

            #init mongod
            sudo mongod -v --logpath /data/mongo/db/logdb.log --dbpath /data/mongo/db  --fork

            echo " ./awesomePkg $cantClientes insertErrorPK $nPruebas $pruebaSizeColl "
            ./awesomePkg $cantClientes insertErrorPK $nPruebas $pruebaSizeColl

    done

fi

if ((ind_pruebas_ejecutar <= 4)) ; then

    ############# inserToTest - findOne #################################
    for((pruebaSizeColl=0;pruebaSizeColl<30; pruebaSizeColl++))
    do
            #restart mongo
            sudo killall mongod

            sleep $time_to_sleep

            #init mongod
            sudo mongod -v --logpath /data/mongo/db/logdb.log --dbpath /data/mongo/db  --fork

            echo " ./awesomePkg $cantClientes insertToTest $nPruebas $pruebaSizeColl "
            ./awesomePkg $cantClientes insertToTest $nPruebas $pruebaSizeColl


            echo " ./awesomePkg $cantClientes findOne $nPruebas $pruebaSizeColl "
            ./awesomePkg $cantClientes findOne $nPruebas $pruebaSizeColl

    done

fi

if ((ind_pruebas_ejecutar <= 5)) ; then

    ############# inserToTest - findIdOne #################################
    for((pruebaSizeColl=0;pruebaSizeColl<30; pruebaSizeColl++))
    do
            #restart mongo
            sudo killall mongod

            sleep $time_to_sleep

            #init mongod
            sudo mongod -v --logpath /data/mongo/db/logdb.log --dbpath /data/mongo/db  --fork

            echo " ./awesomePkg $cantClientes insertToTest $nPruebas $pruebaSizeColl "
            ./awesomePkg $cantClientes insertToTest $nPruebas $pruebaSizeColl


            echo " ./awesomePkg $cantClientes findIdOne $nPruebas $pruebaSizeColl "
            ./awesomePkg $cantClientes findIdOne $nPruebas $pruebaSizeColl

    done

fi


if ((ind_pruebas_ejecutar <= 6)) ; then

    ############# inserToTest - updateOne #################################
    for((pruebaSizeColl=0;pruebaSizeColl<30; pruebaSizeColl++))
    do
            #restart mongo
            sudo killall mongod

            sleep $time_to_sleep

            #init mongod
            sudo mongod -v --logpath /data/mongo/db/logdb.log --dbpath /data/mongo/db  --fork

            echo " ./awesomePkg $cantClientes insertToTest $nPruebas $pruebaSizeColl "
            ./awesomePkg $cantClientes insertToTest $nPruebas $pruebaSizeColl


            echo " ./awesomePkg $cantClientes updateOne $nPruebas $pruebaSizeColl "
            ./awesomePkg $cantClientes updateOne $nPruebas $pruebaSizeColl

    done

fi

if ((ind_pruebas_ejecutar <= 7)) ; then

    ############# inserToTest - updateOneId #################################
    for((pruebaSizeColl=0;pruebaSizeColl<30; pruebaSizeColl++))
    do
            #restart mongo
            sudo killall mongod

            sleep $time_to_sleep

            #init mongod
            sudo mongod -v --logpath /data/mongo/db/logdb.log --dbpath /data/mongo/db  --fork

            echo " ./awesomePkg $cantClientes insertToTest $nPruebas $pruebaSizeColl "
            ./awesomePkg $cantClientes insertToTest $nPruebas $pruebaSizeColl


            echo " ./awesomePkg $cantClientes updateOneId $nPruebas $pruebaSizeColl "
            ./awesomePkg $cantClientes updateOneId $nPruebas $pruebaSizeColl

    done

fi

if ((ind_pruebas_ejecutar <= 8)) ; then

    ############# inserToTest - deleteOne #################################
    for((pruebaSizeColl=0;pruebaSizeColl<30; pruebaSizeColl++))
    do
            #restart mongo
            sudo killall mongod

            sleep $time_to_sleep

            #init mongod
            sudo mongod -v --logpath /data/mongo/db/logdb.log --dbpath /data/mongo/db  --fork

            echo " ./awesomePkg $cantClientes insertToTest $nPruebas $pruebaSizeColl "
            ./awesomePkg $cantClientes insertToTest $nPruebas $pruebaSizeColl


            echo " ./awesomePkg $cantClientes deleteOne $nPruebas $pruebaSizeColl "
            ./awesomePkg $cantClientes deleteOne $nPruebas $pruebaSizeColl

    done

fi

if ((ind_pruebas_ejecutar <= 9)) ; then

    ############# inserToTest - deleteOneId #################################
    for((pruebaSizeColl=0;pruebaSizeColl<30; pruebaSizeColl++))
    do
            #restart mongo
            sudo killall mongod

            sleep $time_to_sleep

            #init mongod
            sudo mongod -v --logpath /data/mongo/db/logdb.log --dbpath /data/mongo/db  --fork

            echo " ./awesomePkg $cantClientes insertToTest $nPruebas $pruebaSizeColl "
            ./awesomePkg $cantClientes insertToTest $nPruebas $pruebaSizeColl


            echo " ./awesomePkg $cantClientes deleteOneId $nPruebas $pruebaSizeColl "
            ./awesomePkg $cantClientes deleteOneId $nPruebas $pruebaSizeColl

    done

fi

if ((ind_pruebas_ejecutar <= 10)) ; then

    ############# inserToTest - deleteOneStatic #################################
    for((pruebaSizeColl=0;pruebaSizeColl<30; pruebaSizeColl++))
    do
            #restart mongo
            sudo killall mongod

            sleep $time_to_sleep

            #init mongod
            sudo mongod -v --logpath /data/mongo/db/logdb.log --dbpath /data/mongo/db  --fork

            echo " ./awesomePkg $cantClientes insertToTest $nPruebas $pruebaSizeColl "
            ./awesomePkg $cantClientes insertToTest $nPruebas $pruebaSizeColl


            echo " ./awesomePkg $cantClientes deleteOneStatic $nPruebas $pruebaSizeColl "
            ./awesomePkg $cantClientes deleteOneStatic $nPruebas $pruebaSizeColl

    done

fi

if ((ind_pruebas_ejecutar <= 11)) ; then

    ############# inserToTest - deleteOneIdStatic #################################
    for((pruebaSizeColl=0;pruebaSizeColl<30; pruebaSizeColl++))
    do
            #restart mongo
            sudo killall mongod

            sleep $time_to_sleep

            #init mongod
            sudo mongod -v --logpath /data/mongo/db/logdb.log --dbpath /data/mongo/db  --fork

            echo " ./awesomePkg $cantClientes insertToTest $nPruebas $pruebaSizeColl "
            ./awesomePkg $cantClientes insertToTest $nPruebas $pruebaSizeColl


            echo " ./awesomePkg $cantClientes deleteOneIdStatic $nPruebas $pruebaSizeColl "
            ./awesomePkg $cantClientes deleteOneIdStatic $nPruebas $pruebaSizeColl

    done

fi


if ((ind_pruebas_ejecutar <= 12)) ; then

    ############# inserToTest - deleteOneIdStatic #################################
    #for((pruebaSizeColl=0;pruebaSizeColl<30; pruebaSizeColl++))
    #do
            #restart mongo
            sudo killall mongod

            sleep $time_to_sleep

            #init mongod
            sudo mongod -v --logpath /data/mongo/db/logdb.log --dbpath /data/mongo/db  --fork

            #echo " ./awesomePkg $cantClientes insertToTest $nPruebas $pruebaSizeColl "
            #./awesomePkg $cantClientes insertToTest $nPruebas $pruebaSizeColl


            echo " ./awesomePkg $cantClientes GeneraTrazaPrueba $nPruebas 0 "
            ./awesomePkg $cantClientes GeneraTrazaPrueba $nPruebas 0

    #done

fi