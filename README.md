<h1 align="center"> 🌐User System API (DEV)</h1>

Uma API para o gerenciamento de usuários, feito em Go. Simple e Segúra, contém autenticação usando JWT e hash de 
senhas 

#### Funcionalidades

- Criação e exclusão de usuários
- Autenticação de usuários
- Atualização de dados dos usuários
 
<!--- Proteção de rotas com JWT-->
<!--- Validações para prevenir entradas inválidas-->

#### Estrutura do Projeto
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

<!--#### Instalação e Configuração-->
<!---->
<!--1 - Clone o Repositório-->
<!--```bash-->
<!--git clone https://github.com/jhenriquem/go-user-system.git-->
<!--cd go-user-system-->
<!--```-->
<!--2 - Configure as Variáveis de Ambiente-->
<!---->
<!--Crie um arquivo .env no diretório raiz e configure as seguintes variáveis:-->
<!---->
<!--```bash-->
<!--PORT=ports_do_servidor-->
<!--POSTGRES_HOST=localhost-->
<!--POSTGRES_USER=seu_usuario-->
<!--POSTGRES_PASSWORD=sua_senha-->
<!--POSTGRES_NAME=seu_banco-->
<!--JWT_SECRET=sua_chave_secreta-->
<!--```-->
<!---->
<!--3 - Crie e execute a imagem docker -->
<!--```bash-->
<!--docker-compose up --build-->
<!--```-->
