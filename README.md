
**Simple Web Scraper**
=======================

Um raspador web simples escrito em Go que extrai informações de produtos do Mercado Livre.

### Características

* Raspagem de informações de produtos do Mercado Livre usando Colly
* Extração de link, nome e valor do produto
* Tratamento de erros e casos limite
* Fácil de usar e personalizar

### Uso

Para usar este raspador, basta executar o comando:
```bash
go run ./cmd/cli/main.go <search>
```
Substitua `<search>` pelo produto que deseja buscar.

### Exemplo de Saída

```
Search: iPhone 13

Link: https://www.mercadolivre.com.br/iPhone-13-64gb
Nome: iPhone 13 64gb
Valor: 3499.99

Link: https://www.mercadolivre.com.br/iPhone-13-128gb
Nome: iPhone 13 128gb
Valor: 3999.99
```

### Licença

Este projeto é licenciado sob a Licença MIT. Veja o arquivo LICENSE para mais informações.

### Contribuição

Contribuições são bem-vindas! Se tiver alguma ideia ou correção de bug, sinta-se à vontade para abrir uma issue ou enviar um pull request.

### Agradecimentos

* Este projeto utiliza o framework de raspagem web Colly.
* Agradecemos à equipe do Mercado Livre por fornecer uma plataforma excelente para raspagem!

Espero que goste! Se precisar de alguma alteração, basta me informar!