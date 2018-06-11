# Go Boilerplate

Este repositório foi criado para facilitar o desenvolvimento de aplicações microservico REST, trazendo uma estrutura base para desenvolvimento com base em design patterns e estruturas de código bem estruturadas na área de programação.
Seu principal intuito foi criar uma estrutura backend para ser utilizada junto com ferramentas SPA como Vue, React, Angular, etc, porem sua aplicação tambem pode servir para criação de somente API REST.

## Dependências do Projeto

* [Gorm (Banco de Dados)](https://github.com/jinzhu/gorm)
* [Gin (HTTP Router)](https://github.com/gin-gonic/gin)
* [Ozzo Validation (Validação de campos)](https://github.com/go-ozzo/ozzo-validation)
* [Viper (Auxiliar de Configuração via Flags, Variaveis de Ambiente e Arquivos)](https://github.com/spf13/viper)
* [Configor (Leitor de arquivos YAML, JSON, etc)](https://github.com/jinzhu/configor)
* [Gin JWT (Middleware para validação com JWT)](https://github.com/appleboy/gin-jwt)
* [Gin Contrib - Cors (Middleware para liberação de CORS)](https://github.com/gin-contrib/cors)
* [Gin Contrib - Gzip (Middleware para Gzip)](https://github.com/gin-contrib/gzip)
* [Gin Contrib - Location (Middleware para pegar URL do projeto)](https://github.com/gin-contrib/location)

## Dependências Auxiliares

Bibliotecas que não foram adicionadas ao projeto porem podem ser utilizadas para auxiliar casos especificos:

* [Logrus (Criação de logs)](https://github.com/sirupsen/logrus)
* [Boilr (Criação de boilerplates via CLI)](https://github.com/tmrts/boilr)

## Projetos Inspirados

Lista dos projetos onde certas partes de código foram copiadas ou retiradas ideias para contrução da estrutura:

* [Go Base](https://github.com/dhax/go-base)
* [Alloy](https://github.com/olliecoleman/alloy)
* [Golang gin Realworld Example App](https://github.com/gothinkster/golang-gin-realworld-example-app)
* [Go RESTful Application Starter Kit](https://github.com/qiangxue/golang-restful-starter-kit)
