# petshop-dti

## Descrição

Este projeto é uma aplicação web que ajuda os usuários a encontrar o melhor petshop com base na quantidade de cães pequenos e grandes e na data desejada.

## Premissas Assumidas

1. O usuário fornecerá a data, o número de cães pequenos e o número de cães grandes.
2. A API do backend retornará o melhor petshop com base nos dados fornecidos pelo usuário.
3. O backend espera a data no formato `yyyy/mm/dd`.
4. O backend retornará um erro se nenhum cachorro for fornecido.

## Pré-Requisitos
- [Go](https://go.dev/): The Go programming language.
- [Node](https://nodejs.org/en): Node js

## Como rodar o projeto
**Clone o repositório**
```
https://github.com/GabriellGds/petshop-dti.git
```

**Navegue até a pasta cmd** 
```
cd cmd 
```
**execute o backend**
```
go run main.go
```
**navegue até a pasta petshop**
```
cd ../petshop
```
**execute o fronted**
```
npm run dev
```
**abra o navegador navegador nessa url
```
http://127.0.0.1:5173
```
**Para rodar o teste, excute este comando na raiz do projeto**
```
go test -v ./...
```

## Decisões de Projeto

1. **Frontend em React**
2. **Backend em Golang**
3. **Estilização com CSS**

