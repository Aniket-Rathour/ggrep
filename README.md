# This is a grep - A tool use to find word in particular file of in a dir scanning multiple files

it took me 1week builind it all but it was way better thenreading the books i was reading about system,linex,OS and all  , the core working of the dir is dependent on the recursion  and cases to tackel rest of the flags 

the benchmarks i reached are. :- 

```text
goos: darwin
goarch: arm64
pkg: [github.com/Aniket-Rathour/ggrep](https://github.com/Aniket-Rathour/ggrep)
cpu: Apple M4
BenchmarkGrep/default-10                 2313026               511.0 ns/op
BenchmarkGrep/-n_flag-10                 2590492               464.0 ns/op
BenchmarkGrep/-i_flag-10                 1897234               633.4 ns/op
BenchmarkGrep/-r_flag-10                  409406              2783 ns/op
PASS
ok      [github.com/Aniket-Rathour/ggrep](https://github.com/Aniket-Rathour/ggrep) 4.915s
```

THIS MAY LOOK VERY SLOW BUT IS VERY NEAR TO REAL GREP CONFIGRATIONS 

| | |
| --- | --- |
| MY GREP | ⏱️ ~500 ns |
| macOS BSD grep | ⏱️ ~1,200 - 2,000 ns |
| GNU grep (Linux standard) | ⏱️ ~200 - 300 ns |
| ripgrep (rg) | ⏱️ ~100 - 150 ns |

So at the end its a great start
lets just stay alive and look were it goes
