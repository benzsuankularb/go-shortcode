package shortcode

func NewManager(generator Generator) *CoreManager {
	return &CoreManager{
		generator,
        map[string]PlugIn{},
	}
}

type CoreManager struct {
	Generator
	plugIns 	map[string]PlugIn
}

func (this *CoreManager) AddPlugIn(plugInName string, plugIn PlugIn) {
	this.plugIns[plugInName] = plugIn
}

func (this *CoreManager) Generate(plugInName string, data interface{}) (shortCode string, err error) {
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

func (this *CoreManager) Execute(code string) error {
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
    err = this.Generator.Release(code)
    if err != nil {
        // TODO log
    }
    return nil
}

func (this *CoreManager) Revoke(code string) error {
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
    err = this.Generator.Release(code)
    if err != nil {
        // TODO log
    }
    return nil
}