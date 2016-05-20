Go-ShortCode
====

Golang short code library
----------------------------

ShortCode is a tiny library for generate short ticket for user.
The ticket can be use for execution in the ways you need.

Using it, you can implement your own plug-in and regis to ShortCodeManager.

By default, I've implemented shortcode in 4-digit (a-z,A-Z,0-9) and using storage of MongoDB.

### Example

````go
import (
    "github.com/benzsuankularb/go-shortcode"
)

type CustomPlugIn struct {}

func (this *CustomPlugIn) Reserve(shortCode string, data interface{}) error {
    //Implement, Regis income shortcode with your data
}

func (this *CustomPlugIn) Execute(shortCode string) error {
    //Implement, Checkout your shortcode data and do something
}

func (this *CustomPlugIn) Revoke(shortCode string) error {
    //Implement, Delete the shortcode's data.
}
````

````go
import (
    "github.com/benzsuankularb/go-shortcode"
    "github.com/benzsuankularb/go-shortcode/generator"
)

// New shortcode generator, You can implement your own, See Generator interface.
gener := generator.NewGenerator()

// Create a ShortCodeManager
manager := shortcode.NewManager(gener)

// Init your custom plugin.
userPlugIn := CustomPlugIn{}
supportPlugIn := CustomPlugIn{}

// Add it to manager registry
manager.AddPlugIn("create_user", userPlugIn)
manager.AddPlugIn("make_support", supportPlugIn)

// Use it. This will call your plugin Reserve method
shortCode, err := manager.Generate("create_user", map[string]string{"username": "any"})

// Use it. This will call your plugin Execute method
shortCode, err := manager.Execute(shortCode)

// Use it. This will call your plugin Revoke method
shortCode, err := manager.Revoke(shortCode)
````

### Author

benzsuankularb
benzsk130@gmail.com