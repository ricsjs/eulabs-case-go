# eulabs-case-go

### O que foi utilizado:
- GoLang
- Echo Framework
- Docker
- MySQL
- Postman

### Passos para rodar o projeto:
- Fazer o clone do projeto: git clone https://github.com/ricsjs/eulabs-case-go.git
- Rodar o comando: docker-compose up (irá subir dois containers, um para o servidor, outro para o banco)
- Criar a tabela 'produto' no banco de dados

#### Após realizar os passos acima, já poderá fazer as requisições, deixei um arquivo JSON como base.

### Requisições: 
- GET: localhost/produtos (irá mostrar todos os registros da tabela de produtos)
- GET por ID: localhost/produtos/id (deverá passar como parâmetro o ID do produto que deseja buscar)
- GET por filtro de preço: localhost/produtos/preco/preco1/preco2 (deverá passar como parâmetro o preço inicial e o preço final, a resposta será uma lista de produtos que estão entre essa faixa de preço)
- POST: localhost/produtos (deverá colocar no objeto JSON "nome", "preco", "status" como mostra o arquivo JSON disponibilizado)
- PUT: localhost/produtos/id (deverá passar como parâmetro o ID do produto que deseja atualizar e a mesma estrutura do objeto JSON do método POST)
- DELETE: localhost/produtos/id (deverá passar como parâmetro o ID do produto que deseja deletar)

