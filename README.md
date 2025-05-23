# Sistema de estoque

Sistema para controle simples de estoque com ações de adição, atualização e remoção de itens/produtos.

# Tecnologias

- Backend construído com Golang + Gin + ORM
- Frontend construído com HTML, CSS e Javascript puro

Utilizado Docker para tornar a aplicação portável para máquinas que tenham o Docker instalado e simplificado o processo de build e execução com Makefile, onde com apenas um comando o sistema todo é prontificado para uso.

# Como usar

### Executando a aplicação

Para executar a aplicação, basta executar o comando:
```bash
$ make all
```
Esse comando irá buildar e executar o frontend e backend via Docker. O frontend funcionará na porta 80 e backend na porta 8080.

### Rotas

| Método | Rota | Função |
| ------ | ---- | ------ |
| POST | /produtos | Criar produto |
| GET | /produtos | Listar produtos |
| PUT | /produtos/{id} | Atualizar produto |
| DELETE | /produtos/{id} | Deletar produto |
