# Sistema de Temperatura por CEP

## Visão Geral
Este projeto consiste em dois serviços:
- **Serviço A**: Recebe um CEP via POST, valida o formato e encaminha para o Serviço B.
- **Serviço B**: Recebe o CEP, busca a cidade correspondente, consulta a temperatura e retorna os dados formatados.


## Como Rodar o Projeto

1. Clone o repositório.
2. Navegue até a pasta do projeto.
3. Adicione sua chave da WeatherAPI no arquivo `main.go` do Serviço B.
4. Execute o comando `docker-compose up`.

## Endpoints

### Serviço A
- **POST /cep**
  - Body: `{ "cep": "29902555" }`
  - Respostas:
    - `200 OK`: `{ "city": "São Paulo", "temp_C": 28.5, "temp_F": 83.3, "temp_K": 301.6 }`
    - `422 Unprocessable Entity`: `{ "message": "invalid zipcode" }`

### Serviço B
- **POST /weather**
  - Body: `{ "cep": "29902555" }`
  - Respostas:
    - `200 OK`: `{ "city": "São Paulo", "temp_C": 28.5, "temp_F": 83.3, "temp_K": 301.6 }`
    - `422 Unprocessable Entity`: `{ "message": "invalid zipcode" }`
    - `404 Not Found`: `{ "message": "can not find zipcode" }`
