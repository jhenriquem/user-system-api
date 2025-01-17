<h1 align="center"> 🌐User System API </h1>

Uma API para o gerenciamento de usuários, feito em Go. Simple e segura. Fornece endpoints para registro, autenticação e manipulação de dados de usuários (como visualização, atualização e exclusão), Também utiliza middlewares para proteger rotas sensíveis, como manipulação de dados de usuários.

### Funcionalidades

- Criação e exclusão de usuários
- Autenticação de usuários
- Recuperação de dados do usuário
- Atualização de dados dos usuários
- Proteção de rotas com JWT
- Validações para prevenir entradas inválidas

### Tecnologias 
- Linguagem Go 
- Containerização usando Docker 
- PostGreSQL 
- Pacote [gorilla/mux](https://github.com/gorilla/mux) para roteamento 

### Estrutura do Projeto
```bash
.
├── controllers   # Handlers para as rotas
├── middleware    # Middleware para as rotas
├── routes        # Configuração das rotas
├── database      # Configurações do banco de dados
├── types         # Definições de tipos e estruturas
├── server        # Configuração e ponto de entrada da aplicação
├── go.mod        # Gerenciamento de dependências
└── go.sum        # Lockfile das dependências
```

### Endpoints


#### Rotas públicas
| Método | Endpoints | Descrição |
|--------|-----------|-----------|
| POST   |`api/users/` | Registro de um novo usuário 
| GET    | `api/users/auth` | Autenticação de usuários

#### Rotas privadas (protegidas por autenticação jwt)
Todas as rotas requerem um token jwt em seu cabeçalho, obtido atravéz da autenticação( Endpoint `api/users/auth`) 

| Método | Endpoints | Descrição |
|--------|-----------|-----------|
| GET    | `api/user` | Recuperar as informações do usuário
| PUT    |`api/user` | Atualizar as informações do usuário
| DELETE |`api/user` | Apagar p usuário


### Instalação e Configuração

##### Requisitos
- [Linguagem Go](https://go.dev/)
- [Docker](https://www.docker.com/) 

1 - Clone o Repositório
```bash
git clone https://github.com/jhenriquem/go-user-system.git
cd go-user-system
```
2 - Configure as Variáveis de Ambiente

Crie um arquivo .env no diretório raiz e configure as seguintes variáveis:

```bash
PORT=porta_do_servidor
POSTGRES_HOST=localhost
POSTGRES_USER=seu_usuario
POSTGRES_PASSWORD=sua_senha
POSTGRES_NAME=seu_banco
JWT_SECRET=sua_chave_secreta
```

3 - Crie e execute a imagem docker 
```bash
docker-compose up --build
```
