package main

import (
	"database/sql"
	"log"
	"runtime"
)

// -------------------------------------------------
//     ___ _           _     ___
//    / __| |_  ___ __| |__ | __|_ _ _ _ ___ _ _
//   | (__| ' \/ -_) _| / / | _|| '_| '_/ _ \ '_|
//    \___|_||_\___\__|_\_\ |___|_| |_| \___/_|
func checkErr(err error) (isErr bool){
	defer func() { r := recover(); if r != nil { log.Print("Error detected:", r) } }()
	isErr = false
	if err != nil {
		isErr = true
		myLog(runtime.Caller(1))
		log.Panic(err)
	}
	return
}


// -------------------------------------------------
//     ___ _           _  DB ___
//    / __| |_  ___ __| |__ | __|_ _ _ _ ___ _ _
//   | (__| ' \/ -_) _| / / | _|| '_| '_/ _ \ '_|
//    \___|_||_\___\__|_\_\ |___|_| |_| \___/_|
func checkDbErr(err error,db *sql.DB) (isErr bool){
	defer func() { r := recover(); if r != nil { log.Print("DB Error detected:", r) } }()
	isErr = false
	if err != nil {
		isErr = true
		myLog(runtime.Caller(1))
		db.Ping()  //Return the Database Error object
		log.Panic(err)
	}
	return
}

//   ----------------------------------
//    __  __         _
//   |  \/  |       | |
//   | \  / |_   _  | |     ___   __ _
//   | |\/| | | | | | |    / _ \ / _` |
//   | |  | | |_| | | |___| (_) | (_| |
//   |_|  |_|\__, | |______\___/ \__, |
//            __/ |               __/ |
//           |___/               |___/
func myLog(msg ...interface{}) {
	if conf.DEBUG {
		log.Print(msg)
	}
}