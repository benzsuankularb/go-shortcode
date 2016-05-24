package main

import (
    "github.com/benzsuankularb/go-shortcode"
    "github.com/benzsuankularb/go-shortcode/generator"
    "gopkg.in/mgo.v2"
)


    
	

func main() {
    
    session, err := mgo.Dial("127.0.0.1")
    if err != nil {
		panic(err)
	}
    db := session.DB("test")
    c := db.C("test")
    c.DropCollection()
    if err != nil {
		panic(err)
	}
    
    // New shortcode generator, You can implement your own, See Generator interface.
    gener := generator.NewGenerator(c)

    // Create a ShortCodeManager
    _ = shortcode.NewManager(gener)
}