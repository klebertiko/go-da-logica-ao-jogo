### GoLang ###

- Verificando onde esta a variavel GOROOT
$ go env GOROOT

- Verificando onde esta a variavel GOPATH
$ go env GOPATH

- Criando variavel GOPATH
$ export GOPATH=$HOME/<your-path>/go
$ source <your-shell>

- Todo projeto do GO deve ser feito abaixo do GOPATH dentro de um dir src
    - Crie um diretorio src dentro de GOPATH
    $ mkdir $GOPATH/src

Crie um diretorio para seu projeto/programa
$ mkdir $GOPATH/src/estudos

Rodando programas go
$ go run $GOPATH/src/estudos/main.go

Criando um binario e instalando na pasta $GOPATH/bin
$ go install $GOPATH/src/estudos/main.go
$ ../../../bin/main

Criando um binario
$ go build $GOPATH/src/estudos/main.go
$ ./main

- DEP Gerenciador de dependencias

- Iniciar gerenciamento de dependencias em um projeto existente
$ cd $GOPATH/src/estudos
$ dep init 

- Com o comando acima ele cria os seguintes arquivos e o diretorio vendor
$ ls
$ Gopkg.toml Gopkg.lock vendor/

- Adicionar/atualizar dependencias com o dep (sao baixadas no diretorio vendor)
$ dep ensure -add github.com/foo1/bar1 github.com/foo2/bar2
$ dep ensure -update github.com/foo1/bar1

Visualizando dependencias
$ dep status -dot | dot -T png | display