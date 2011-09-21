Simple/Fluent config file reader:

eg.conf
-----
#config comments are ignored
no value are ignored too
int_val = 345
bool_val = True
string_val = any string. only single line though.


# empty lines above are ignored too.
--

f := NewConfigo("eg.conf")
f.Load()

f.Get("int_val").asInt() => 345
f.Get("int_val").asString() => "345"
f.Get("bool_val").asBool() => true
f.Get("no_key").Default("OOPS").asString() => "OOPS"
Simple config file reader.

