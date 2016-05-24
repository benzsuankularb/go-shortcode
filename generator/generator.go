package generator

import (
    . "github.com/benzsuankularb/go-shortcode"
    "gopkg.in/mgo.v2"
    "strconv"
    "math/rand"
)

type mapper struct {
    ShortCode   string  `_id`
    PlugInName  string  `plugin`
}

type MongoGenerator struct {
    collection *mgo.Collection
}

func NewGenerator(c *mgo.Collection) Generator {
    return &MongoGenerator{c}
}

func (this *MongoGenerator) Reserve(plugInName string) (string, error) {
    for {
        newCode := RandomCode()
        if isUsable := this.IsAvailable(newCode); isUsable {
            err := this.tryReserve(newCode, plugInName)
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

func (this *MongoGenerator) tryReserve(code string, plugInName string) error {
    if !this.IsAvailable(code) {
        return ErrNotAvailable
    }
    err := this.collection.Insert(mapper{code, plugInName})
    if err != nil {
        return ErrDatabase
    }
    return nil
}

//TODO TTL reserving
func (this *MongoGenerator) Release(code string) error {
    err := this.collection.RemoveId(code)
    if err != nil {
        return ErrDatabase
    }    
    return nil
}

func (this *MongoGenerator) GetPlugInName(code string) (string, error) {
    result := mapper{}
    q := this.collection.FindId(code)
    c, err := q.Count()
    if err != nil {
        return "", ErrDatabase
    } else if c == 0 {
        return "", ErrNotFound
    }
    err = q.One(&result)
    if err != nil {
        return "", ErrDatabase
    }
    return result.PlugInName, nil
}

func (this *MongoGenerator) IsAvailable(code string) bool {
    c, err := this.collection.FindId(code).Count()
    return err == nil && c == 0
}

func RandomCode() string {
	result := ""
	for i := 0; i < 4 ; i++ {
		result += randomDigit()
	}
	return result
}

func randomDigit() string {
	r := rand.Intn( 35 )
	if r < 10 {
		return strconv.Itoa(r)
	}
	return string(r + 87)
}