# Buscador por CEP
Este código em Go foi desenvolvido para demonstrar a utilização de duas APIs diferentes para obter informações de endereços a partir de um CEP. O programa compara o tempo de resposta das APIs ViaCEP e BrasilAPI, exibindo os dados da API que responde mais rapidamente dentro de um prazo de 1 segundo.

## Estrutura do Código
O código consiste em um arquivo principal, main.go, que contém a lógica principal do programa. Há também a definição de uma estrutura Address para armazenar os dados do endereço e duas funções principais: main() e getAddress().

- main(): Inicia o programa, define o CEP a ser consultado, inicia duas goroutines para chamar as APIs simultaneamente, aguarda a conclusão das goroutines e exibe os resultados.

- getAddress(): Realiza uma solicitação HTTP para uma API específica, decodifica os dados JSON da resposta e envia os resultados por meio de um canal.


## Como Rodar o Código
Clone o repositório:

```bash
git clone https://github.com/booscaaa/desafio-multithreading-go-expert-pos.git
```
## Navegue até o diretório do projeto:

```bash
cd desafio-multithreading-go-expert-pos
```

## Execute o código Go:

```bash
go mod tidy
go run main.go
```

## Exemplo de Uso
Ao rodar o programa, ele enviará solicitações para as APIs ViaCEP e BrasilAPI com o CEP fixo 99150000(Marau), exibindo os dados da API que responde mais rapidamente dentro do prazo de 1 segundo.

## Exemplo de saída bem-sucedida:

```bash
go run main.go

API mais rápida: https://brasilapi.com.br/api/cep/v1/99150000
Os dados são: map[cep:99150000 city:Marau service:correios state:RS]
```
## Observações
O tempo limite para resposta é configurado para 1 segundo. Se nenhuma das APIs responder dentro desse prazo, o programa exibirá uma mensagem de timeout.