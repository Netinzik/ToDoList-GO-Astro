# ToDoList App

Este é um projeto de aplicativo de lista de tarefas (ToDoList) com o backend desenvolvido em Golang/Fiber e o frontend utilizando Astro. O aplicativo permite que os usuários criem, visualizem, atualizem e excluam tarefas de forma simples e eficiente.

## Backend (Golang/Fiber)

### Configuração

1. Certifique-se de ter o Go instalado em seu sistema. Se não, você pode baixá-lo em [golang.org](https://golang.org/dl/).

2. Instale as dependências do projeto com o seguinte comando:

    ```bash
    go get -u github.com/gofiber/fiber/v2
    ```

3. Execute o servidor backend:

    ```bash
    cd backend
    go run main.go
    ```

O servidor será iniciado em `http://localhost:8080`.

### Endpoints

- `GET /tasks`: Obtém todas as tarefas.
- `GET /tasks/:id`: Obtém uma tarefa específica por ID.
- `POST /tasks`: Adiciona uma nova tarefa.
- `PUT /tasks/:id`: Atualiza uma tarefa existente.
- `DELETE /tasks/:id`: Exclui uma tarefa.

## Frontend (Astro)

### Configuração

1. Certifique-se de ter o Node.js instalado em seu sistema. Você pode baixá-lo em [nodejs.org](https://nodejs.org/).

2. Instale as dependências do projeto:

    ```bash
    cd frontend
    npm install
    ```

3. Execute o servidor frontend:

    ```bash
    npm run dev
    ```

O servidor será iniciado em `http://localhost:3000`.

### Estrutura do Projeto

- `src/components`: Componentes reutilizáveis.
- `src/pages/Index.astro`: Página principal do aplicativo.
- `src/main.astro`: Arquivo principal do Astro.

## Como Usar

1. Acesse o aplicativo no navegador em `http://localhost:3000`.
2. Utilize a interface intuitiva para adicionar, visualizar, atualizar e excluir tarefas.

## Contribuições

Contribuições são bem-vindas! Sinta-se à vontade para me escrever no e-mail que se encontra na aba formulário.

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).