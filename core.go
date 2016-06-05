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

func (this *CoreManager) Generate(plugInName string, data interface{}) (string, error) {
    plugIn, ok := this.plugIns[plugInName]
    if !ok {
        return "", ErrPlugInNotFound
    }
    
    pTempCode, err := plugIn.Temporary(data)
    if err != nil {
        return "", err
    }
    if pTempCode != nil {
        return *pTempCode, nil
    }
    shortCode, err := this.Generator.Reserve(plugInName)
    if err != nil {
        return "", err
    }
	err = plugIn.Reserve(shortCode, data)
    if err != nil {
        this.Generator.Release(shortCode)
        return "", err 
    }
    return shortCode, nil
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

func (this *CoreManager) GetData(code string) (plugInName string, data interface{}, err error) {
    plugInName, err = this.Generator.GetPlugInName(code)
    
    if err != nil {
        return
    }
    plugIn, ok := this.plugIns[plugInName]
    if !ok {
        err = ErrPlugInNotFound
        return
    }
    data, err = plugIn.Get(code)
    return
}