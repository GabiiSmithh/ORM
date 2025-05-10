# Object-Relational Mapper (ORM) para MongoDB

- Foi desenvolvido um ORM simplificado para MongoDB na linguagem Go, utiliando recursos da linguagem como structs, modularização e interfaces.
- Foi implementado um CRUD básico (com as operações de Create, Read, Update e Delete), mapeamento de objeto para documento, geração da estrutura do banco de dados, e, por fim, suporte a ...
- Foi utilizado o sistema operacional Windows para o desenvolvimento.

## Estrutura do projeto

- O projeto possui a seguinte estrutura de pastas:

```
go-mongo-orm/
├── config/                      # Conexão e configuração do banco
│   └── config.go
├── model/                       # Definição pura do modelo User
│   └── user.go
├── orm/                         # Lógica de banco: CRUD + geração de índices únicos (migração)
│   ├── crud.go
│   └── generator.go
├── schema/                      # Mapeamento + conversão de tipos
│   └── user_schema.go
├── utils/                       
│   └── ...
├── main.go                      # Arquivo para testar o ORM
├── go.mod                       # Gerenciador de dependências do Go
├── go.sum                       # 
```
## Especificidades da linguagem Go

- A linguagem Go é uma linguagem compilada, tipada estaticamente e com suporte a concorrência. Ela é projetada para ser simples, eficiente e fácil de ler. Algumas de suas características exploradas incluem:
  - Structs e tags
  - Modularização
  - Funções com contexto (para controle de tempo de execução)

## Como Executar o Programa

- Primeiro, crie o projeto em sua máquina. Para isso, siga os passos abaixo:
  1. Instale o Go em sua máquina (https://golang.org/doc/install)
  2. Instale o MongoDB em sua máquina (https://www.mongodb.com/try/download/community)
  3. Crie um projeto ou clone este repositório: 
    ```bash
    git clone
    ```
  4. Instale os drivers do MongoDB:
    ```bash
    go get go.mongodb.org/mongo-driver/mongo
    go get go.mongodb.org/mongo-driver/mongo/options
    ```

- Após isso, você pode executar o programa com o seguinte comando:
    ```bash
    go run main.go
    ```

## APAGAR DPS

- config.go -> abstrai e centraliza a configuração da conexão do banco de dados. Funcionalidades do go usadas: contexto, modularização, logs, variável global. Essa estrutura permite reaproveitamento da conexão em outros arquivos (crud.go).
- user.go -> define estrutura do domínio da aplicação (usuário).
- crud.go -> implemeta operações básicas de CRUD (create, read, update, delete) sobre o User.
- generator.go -> cria índices únicos no banco. Garante a integridade dos dados, impede duplicidade.
- user_schema.go -> mapeamento de dados, ou seja, representaçãodos dados do banco em structs. Converte os dados do banco para a aplicação e vice-versa. Possui o modelo lógico (aplicação), o modelo de armazenamento (banco) e a conversão entre eles.