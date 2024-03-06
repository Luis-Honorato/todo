# GoTodoAPI

Olá, que bom que você está aqui 😄, este projeto foi realizado com fins de estudo, fiz com o intuito de aprender goLang e aplicar meus conhecimentos em uma API real.

Espero que você goste do aplicativo e que você me dê um feedback sobre oque achou.
Qualquer coisa meus contatos estão no README.md do meu perfil do Github, vai ser um prazer conversar com você!

Agora fique com as especificações e o detalhamento do projeto:

## Funcionalidades
   - Criar um novo todo
   - Retornar um todo baseado no ID
   - Retornar uma lista de todos
   - Atualizar um todo com base nos atributos e no ID
   - Deletar um todo baseado no ID

## Tecnologias Utilizadas

   - GoLang: Linguagem de programação poderosa e eficiente para desenvolvimento de aplicativos escaláveis.
   - Gin: Framework web para Go que facilita a criação de APIs de forma rápida e simples.
   - GORM: Biblioteca ORM para Go que simplifica a interação com o banco de dados, neste caso, utilizamos SQLite.

## Como usar
### Pré-requisitos

   - Go instalado em sua máquina
   - SQLite instalado (ou pode usar outra base de dados suportada pelo GORM)

### Passos para Execução

   1. Clone o repositório:

```bash
git clone https://github.com/Luis-Honorato/todo_go_api
```

   2. Navegue até o diretório do projeto:

```bash
cd todo_go_api
```

  3. Instale as dependências:

```bash 
go mod tidy
```

   4. Execute o aplicativo:

```bash
go run main.go
```

A API estará em execução em `http://localhost:8080/api/v1/`

## Endpoints
### GET /todos

Retorna uma lista com todos os todos.
### GET /todo:id

Retorna um todo baseado no ID fornecido.
### POST /todo:id

Cria um novo todo com base nos dados fornecidos no corpo da solicitação.
### PUT /todo:id

Atualiza um todo existente com base no ID fornecido e nos dados fornecidos no corpo da solicitação.
### DELETE /todo:id

Deleta um todo baseado no ID fornecido.

![image](https://github.com/Luis-Honorato/todo_go_api/assets/90717674/0bba0edb-d3eb-427f-8c30-5a54553ebbd5)


## Todo

### A struct Todo representa uma tarefa em nossa aplicação (é tipo uma classe ou um model do banco). Ela possui os seguintes campos:

   - gorm.Model: Este campo é uma estrutura embutida fornecida pelo GORM que adiciona campos padrão ao modelo, como ID, CreatedAt, UpdatedAt e DeletedAt, para facilitar operações comuns de banco de dados, como criação, atualização e exclusão. Isso ajuda a manter o controle de registros no banco de dados.

   - Title (string): Este campo armazena o título da tarefa. É uma descrição curta e significativa que identifica a tarefa.

   - Description (string): Este campo contém uma descrição mais detalhada da tarefa. Pode incluir informações adicionais sobre a tarefa, como instruções específicas ou contexto relevante.

   - Finished (bool): Este campo indica se a tarefa está concluída ou não. Se for true, significa que a tarefa está marcada como concluída; caso contrário, será false. Isso permite que os usuários rastreiem o status de conclusão da tarefa.

     ![image](https://github.com/Luis-Honorato/todo_go_api/assets/90717674/fd307167-1d98-4ccc-9e6f-e3b55c160762)

## Contribuição

Contribuições são bem-vindas! Se você deseja melhorar este projeto ou resolver problemas, sinta-se à vontade para abrir uma issue ou enviar um pull request.

## Licença

Este projeto está sob a licença MIT.

Obrigado por ler até aqui 💙
