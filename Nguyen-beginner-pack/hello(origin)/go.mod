module hello

go 1.15

require rsc.io/quote v1.5.2
// replace the module path with a greetings directory
replace example.com/greetings => ../greetings

//require example.com/greetings v1.1.0