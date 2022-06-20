# Desenvolvimento

É necessário alterar o GO package de `tonussi/hermes` para `xyz/hermes`.

Para alterar o package busque todas as ocorrências de `tonussi/hermes` e altere para `xyz/hermes`, também é necessário criar um repositório no Github `xyz/hermes`.

Depois de alterar o nome do pacote GO, remova os arquivos go.mod e go.sum e execute as linhas a seguir:

```sh
go mod init github.com/tonussi/hermes
go mod tidy
go get -u all
```

Para desenvolver é necessário instalar Docker Compose, Docker, Go, Vscode, make.

Uma vez instalados é possível apenas executar `make build_debug_hermes`, para gerar imagem de contêiner para um Hermes com depurador Delve embutido.

Depois disso é possível acionar *Run e Debug Mode* selecionando a opção: `Debug Hermes (Attach)`.

O Vscode irá iniciar o Hermes em modo depuração.

O arquivo Docker Compose irá gerar um volume de dados para o BoltDB, um banco de dado usado internamente pelo Raft (interno ao Hermes).

Por padrão o Hermes irá escutar o endereço `localhost:8000` e enviar para o endereço `localhost:8001`. Ou seja, o servidor da aplicação alvo poderá escutar o endereço `localhost:8001`. A menos que seja necessário alterar os endereços e portas, no caso se o servidor da aplicação precisa escutar outra porta, necessariamente, basta alterar o endereço de envio do Hermes, via parâmetros.
