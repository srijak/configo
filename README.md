### Configo: the fluent config file reader for go

### Installation:
```
goinstall github.com/srijak/configo
```


### Usage:


#### Option A: Create a struct and load configs into it using one line

``` go
type EgConfigStruct struct {
  // Properties to be populated using
  // the config file should be exportable
  Astring string
  Somenum int
  Abool bool
}

ec := EgConfigStruct{}
NewConfigo("./test_data/hydrate.conf").Hydrate(&ec)
// Now:
// ec.Astring => abra
// ec.Somenum => 23
// ec.Abool => true
```

#### Option B: manually access items

``` go
f := configo.NewConfigo("eg.conf")
f.Load()

f.Get("Somenum").AsInt() => 345
f.Get("Somenum").AsString() => "345"
f.Get("Abool").AsBool() => true
// can assign defaults
f.Get("NoKey").Default("OOPS").AsString() => "OOPS"
```

#### Option C: mix and match

``` go
ec := EgConfigStruct{}
f:= NewConfigo("./test_data/hydrate.conf").Hydrate(&ec)

// f.Get("Somenum").AsInt() => 345
// ec.Somenum => 345
```

#### Example config file 
<pre>
# comments should be preceeded by #s and are ignored
no value lines are ignored too
Somenum = 345
Abool = True
Astring = any string. only single line though.
 
# empty lines above are ignored too.
</pre>
