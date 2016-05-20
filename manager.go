package shortcode

import (
    "errors"
)

func NewManager(generator Generator) *ShortCodeManager {
	return &ShortCodeManager{
		generator,
        map[string]PlugIn{},
	}
}

var (
    ErrPlugInNotFound = errors.New("plugin-manager : PlugIn not found") 
)

type ShortCodeManager struct {
	Generator
	plugIns 	map[string]PlugIn
}

func (this *ShortCodeManager) AddPlugIn(plugInName string, plugIn PlugIn) {
	this.plugIns[plugInName] = plugIn
}

func (this *ShortCodeManager) Reserve(plugInName string, data interface{}) (shortCode string, err error) {
    plugIn, ok := this.plugIns[plugInName]
    if !ok {
        err = ErrPlugInNotFound
        return
    }
    shortCode, err = this.Generator.Reserve(plugInName)
    if err != nil {
        return
    }
    defer func() {
        if r := recover(); r != nil {
            this.Generator.Release(shortCode)
            shortCode = ""
        }
    }()
	err = plugIn.Reserve(shortCode, data)
    if err != nil {
        panic(err)
    }
    return
}

func (this *ShortCodeManager) Execute(code string) error {
    plugInName, err := this.Generator.GetPlugInName(code)
    if err != nil {
        return err
    }
    plugIn, ok := this.plugIns[plugInName]
    if !ok {
        return ErrPlugInNotFound
    }
    err = plugIn.Execute(code)
    if err != nil {
        return err
    }
    return nil
}

func (this *ShortCodeManager) Revoke(code string) error {
    plugInName, err := this.Generator.GetPlugInName(code)
    if err != nil {
        return err
    }
    plugIn, ok := this.plugIns[plugInName]
    if !ok {
        return ErrPlugInNotFound
    }
    err = plugIn.Revoke(code)
    if err != nil {
        return err
    }
    return nil
}

type PlugIn interface {
	Reserve(shortCode string, data interface{}) error
	Execute(shortCode string) error
    Revoke(shortCode string) error
}