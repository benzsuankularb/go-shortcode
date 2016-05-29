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
    GetData(code string) (plugInName string, data interface{}, err error)
}

type PlugIn interface {
	Reserve(shortCode string, data interface{}) error
	Execute(shortCode string) error
    Revoke(shortCode string) error
    Get(shortCode string) (interface{}, error)
    
    // Incase you have created shortcode. You may not want to create it again 
    Temporary(data interface{}) (*string, error)
}