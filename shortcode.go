package shortcode

import (
    "errors"
)

var (
    ErrNotAvailable = errors.New("shortcode : Not available")
    ErrDatabase = errors.New("shortcode : Database error") 
	ErrNotFound = errors.New("shortcode : Not found")
)

type Generator interface {
    Reserve(plugInName string) (string, error)
    GetPlugInName(code string) (string, error)
	Release(code string) error
}