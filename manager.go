package shortcode

import (
    "errors"
)

var (
    ErrPlugInNotFound = errors.New("shortcode : PlugIn not found")
    ErrAlreadyExist = errors.New("shortcode : Already exist") 
)

type Manager interface {
    AddPlugIn(plugInName string, plugIn PlugIn)
    Generate(plugInName string, data interface{}) (shortCode string, err error)
    Execute(code string) error
    Revoke(code string) error
}

type TempManager interface {
    Save(key string, shortCode string) error
    Remove(key string) error
    Get(key string) (shortCode *string, err error)
}

type PlugIn interface {
	Reserve(shortCode string, data interface{}) error
	Execute(shortCode string) error
    Revoke(shortCode string) error
}