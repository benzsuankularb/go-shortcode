package shortcode

import (
    "errors"
    "strconv"
    "math/rand"
)

var (
    ErrNotAvailable = errors.New("shortcode : Not available")
    ErrDatabase = errors.New("shortcode : Database error") 
)

type ShortCode string

func ToShortCode(shortString string) (ShortCode, error) {
    
    return ShortCode(shortString), nil
}

type Generator interface {
    IsAvailable(code ShortCode) bool
    Reserve() (ShortCode, error)
    Release(code ShortCode) error
}

func RandomCode() ShortCode {
	result := ""
	for i := 0; i < 4 ; i++ {
		result += randomDigit()
	}
	return ShortCode(result)
}

func randomDigit() string {
	r := rand.Intn( 35 )
	if r < 10 {
		return strconv.Itoa(r)
	}
	return string(r + 87)
}