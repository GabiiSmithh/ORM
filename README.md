# Object-Relational Mapper (ORM) para MongoDB

- Foi desenvolvido um ORM simplificado para MongoDB na linguagem Go, utiliando recursos da linguagem como structs, modularização e interfaces.
- Foi implementado um CRUD básico (com as operações de Create, Read, Update e Delete), mapeamento de objeto para documento, geração da estrutura do banco de dados, e, por fim, suporte a ...
- Foi utilizado o sistema operacional Windows para o desenvolvimento.

## Estrutura do projeto

- O projeto possui a seguinte estrutura de pastas:

```
ORM/
├── config/
│   └── config.go               # Configurar a conexão com o banco de dados
├── model/                      # Definição das structs com o mapeamento dos documentos no banco de dados
│   └── pessoa.go
│   └── livro.go
│   └── produto.go
├── orm/
│   ├── crud.go                 # Funções CRUD genéricas e reutilizáveis para os modelos
│   └── generator.go            # Criação de índices únicos no banco
│   └── query.go                # Buscas no banco utilizando filtros e ordenação
├── utils/                              
│   └── handle                  # Adicionar dados nos documentos de cada modelo
│       └── handleLivro.go
│       └── handlePessoa.go
│       └── handleProduto.go
├── main.go
├── go.mod                      # Dependências da lingaugem 
├── go.sum                      # Integridade das dependências (gerado automaticamente)
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
  3. Crie um projeto ou clone este repositório: (https://github.com/GabiiSmithh/ORM.git)
    ```bash
    git clone
    ```
  4. Acesse o diretório desse projeto e instale os drivers do MongoDB:
    ```bash
    go get go.mongodb.org/mongo-driver/mongo
    go get go.mongodb.org/mongo-driver/mongo/options
    ```

- Após isso, você pode executar o programa com o seguinte comando:
    ```bash
    go run main.go
    ```