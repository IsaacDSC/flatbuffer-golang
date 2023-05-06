# flatbuffer-golang


### Initial commands

*Compile using command*
```sh
flatc -b user.fbs users.txt
```

*Re-compile using command*

```sh 
flatc --raw-binary -t --strict-json user.fbs -- users.bin
```


*Gerar lib para sua linguagem*
```sh
flatc --go user.fbs
```