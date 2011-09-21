Simple/Fluent config file reader:


Config file: eg.conf =>
<pre>
# comments should be preceeded by #s and are ignored
no value lines are ignored too
int_val = 345
bool_val = True
string_val = any string. only single line though.
 
# empty lines above are ignored too.
</pre>

Access code:

``` go

f := NewConfigo("eg.conf")
f.Load()

f.Get("int_val").AsInt() => 345
f.Get("int_val").AsString() => "345"
f.Get("bool_val").AsBool() => true
f.Get("no_key").Default("OOPS").AsString() => "OOPS"

```
