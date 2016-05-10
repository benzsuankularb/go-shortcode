package generator

import (
    . "github.com/benzsuankularb/go-shortcode"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type MongoGenerator struct {
    collection mgo.Collection
}

func NewGenerator(c mgo.Collection) Generator {
    return &MongoGenerator{c}
}

func (this *MongoGenerator) Reserve() (ShortCode, error) {
    for {
        newCode := RandomCode()
        if isUsable := this.IsAvailable(newCode); isUsable {
            err := this.tryReserve(newCode)
            if err == ErrNotAvailable {
                continue
            } else if err == nil {
                return newCode, nil
            } else {
                return "", err
            }
        }
    }
    return "", nil
}

func (this *MongoGenerator) tryReserve(code ShortCode) error {
    if !this.IsAvailable(code) {
        return ErrNotAvailable
    }
    err := this.collection.Insert(bson.M{"_id": code})
    if err != nil {
        return ErrDatabase
    }
    return nil
}

//TODO TTL reserving
func (this *MongoGenerator) Release(code ShortCode) error {
    err := this.collection.RemoveId(code)
    if err != nil {
        return ErrDatabase
    }    
    return nil
}

func (this *MongoGenerator) IsAvailable(code ShortCode) bool {
    c, err := this.collection.FindId(code).Count()
    return err == nil && c == 0
}