package shortcode

type FakeTempManager struct {
    storage map[string]string
}

func NewFakeTempManager() TempManager {
    return &FakeTempManager{
        map[string]string{},
    }
}

func (this *FakeTempManager) Save(key string, shortCode string) error {
    _, ok := this.storage[key]
    if ok {
        return ErrAlreadyExist
    }
    this.storage[key] = shortCode
    return nil
}

func (this *FakeTempManager) Get(key string) (*string, error) {
    code, ok := this.storage[key]
    if !ok {
        return nil, nil
    }
    return &code, nil
}

func (this *FakeTempManager) Remove(key string) error {
    delete(this.storage, key)
    return nil
}