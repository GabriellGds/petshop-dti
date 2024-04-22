# petshop-dti

## Descrição

Este projeto é uma aplicação web que ajuda os usuários a encontrar o melhor petshop com base na quantidade de cães pequenos e grandes e na data desejada. O projeto conta com 3 petshops:

1. Meu canino feliz: Está distante 2km do canil. Em dias de semana o banho para cães pequenos custa R$20,00 e o banho dem cães grands custa R$40,00. Durante os finais de semana o preço dos banhos é aumentado em 20%.

2. Vai Rex: Está localizado na mesma avenida do canil, a 1,7km. O preço do banho para dias úteis em cães pequenos é R$15,00 e em cães grandes é R$50,00. Durante os finais de semana o preço para cães pequenos é R$20,00  e para os grandes é R$55,00.

3. ChowChawgas: Fica a 800m do canil. O preço do banho é o mesmo em todos os dias da semana. Para cães pequenos custa R$30 e para cães grandes R$45,00.

A aplicação sempre vai retornar o petshop com o preço mais acessível. Em caso de empate o petshop retornado será o petshop mais próximo.


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
**Navegue até o diretorio do projeto**
```
cd petshop-dit
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
**Baixe as dependências**
```
npm install
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

