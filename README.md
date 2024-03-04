# API CRUD em Go com Docker e SQLite

Este projeto é uma API CRUD (Create, Read, Update, Delete) desenvolvida em Go (Golang), que utiliza Docker para a contêinerização e SQLite como banco de dados. A API fornece operações básicas de gerenciamento de usuários, incluindo criação, leitura, atualização e exclusão.

## Objetivo do Projeto:

O objetivo deste projeto é fornecer uma base sólida para o desenvolvimento de APIs CRUD em Go, demonstrando boas práticas de organização de código, integração com banco de dados, e o uso de Docker para garantir ambientes consistentes.

## Funcionalidades Principais:

- **Criar Usuário:** Adicionar um novo usuário ao sistema.
- **Ler Usuários:** Recuperar uma lista de todos os usuários.
- **Ler Usuário por ID:** Obter detalhes de um usuário específico pelo seu ID.
- **Atualizar Usuário:** Modificar informações para um usuário específico.
- **Excluir Usuário:** Remover um usuário do sistema.

## Tecnologias Utilizadas:

- **Go (Golang):** Linguagem de programação eficiente e concisa para o desenvolvimento do backend.
- **Docker:** Facilita a implantação, distribuição e execução consistente da aplicação.
- **SQLite:** Banco de dados leve e incorporado, ideal para projetos menores e desenvolvimento rápido.

## Estrutura do Projeto:

A estrutura do projeto segue uma divisão clara entre os arquivos principais:

- `main.go`: O ponto de entrada da aplicação, onde o servidor HTTP é iniciado e as rotas são definidas.
- `database.go`: Contém a inicialização do banco de dados e as operações relacionadas ao banco de dados.
- `handlers.go`: Implementa os manipuladores (handlers) para cada rota da API.
- `Dockerfile`: Configuração para a construção da imagem Docker.
- `docker-compose.yml`: Configuração do Docker Compose para simplificar o ambiente de desenvolvimento.

## Uso:

1. **Clone o repositório:**

   ```bash
   git clone https://github.com/gatinhodev/crud-go.git
   ```

2. **Navegue até o diretório do projeto:**

   ```bash
   cd crud-go
   ```

3. **Inicie a aplicação com o Docker Compose:**

   ```bash
   docker-compose up
   ```

4. **Interaja com a API usando seu cliente HTTP preferido.**

## Trechos Importantes do Código:

- **Arquivo `main.go`:** Este arquivo é o ponto de entrada da aplicação. Ele inicia o servidor HTTP, configura as rotas e lida com a inicialização do banco de dados.

- **Arquivo `handlers.go`:** Aqui você encontrará a implementação dos manipuladores para cada rota da API. Cada função manipula uma operação CRUD específica.

- **Arquivo `database.go`:** Este arquivo abriga a inicialização do banco de dados e funções auxiliares para operações de banco de dados.

## Contribuições:

Contribuições são bem-vindas! Se você encontrar problemas, tiver sugestões de melhorias ou quiser adicionar novos recursos, sinta-se à vontade para abrir uma issue ou enviar um pull request.

## Licença:

Este projeto está licenciado sob a [Licença GNU Affero General Public License (AGPL)](LICENSE).

---