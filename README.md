Claro, aqui está uma versão mais organizada e compreensível do seu README:
API de Pagamentos

Esta API tem como objetivo integrar diversos aplicativos (plataformas) com um único gateway de pagamento principal. A ideia é simplificar a mudança de gateway de pagamento, caso seja necessário, sem causar problemas nos serviços já existentes.
Como executar a aplicação

Para rodar a aplicação, execute o seguinte comando no terminal:

```bash
go run .
```

# Rotas

## Platform

- GET /platforms: Retorna todas as plataformas cadastradas.
- POST /platform: Cria uma nova plataforma.
- GET /platform/:id: Retorna uma plataforma específica com base no ID.
- GET /platform/slug/:slug: Retorna uma plataforma específica com base no slug.
- PUT /platform/:id: Atualiza os detalhes de uma plataforma existente.
- DELETE /platform/:id: Exclui uma plataforma existente.

## Payments
- GET /payments: Retorna todos os pagamentos registrados.
- POST /payment: Cria um novo pagamento e retorna um código PIX QRCode para o pagamento.
- GET /payment/:id: Retorna um pagamento específico com base no ID.
- PUT /payment/:id: Atualiza os detalhes de um pagamento existente.
- DELETE /payment/:id: Exclui um pagamento existente.
