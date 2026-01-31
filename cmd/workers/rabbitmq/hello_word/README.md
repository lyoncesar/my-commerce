# Propósito

Registro de anotações de forma geral sobre o RabbitMQ.

# Índice
- [Propósito](#propósito)
- [Índice](#índice)
- [RabbitMQ](#rabbitmq)
  - [Protocolo AMQP](#protocolo-amqp)
  - [Message Broker](#message-broker)
    - [Exchange](#exchange)
      - [Tipos de Exchange](#tipos-de-exchange)
        - [Exchange padrão](#exchange-padrão)
        - [Direct exchange](#direct-exchange)
    - [Queue](#queue)
    - [Bindings](#bindings)
  - [Message](#message)
    - [Publisher](#publisher)
    - [Consumer](#consumer)
    - [Criar Publisher/Consumer](#criar-publisherconsumer)
    - [Message meta-data](#message-meta-data)
    - [Message Acknowledgements](#message-acknowledgements)

# RabbitMQ

O RabbitMQ é um message broker, o objetivo desse sistema é garantir que as mensagens publicadas nele cheguem até os seus consumidores.

## Protocolo AMQP

O RabbitMQ usa o protocolo de mensagens [AMQP](https://en.wikipedia.org/wiki/Advanced_Message_Queuing_Protocol) (Advanced Message Queuing Protocol), que é um protocolo de camada de aplicação de padrão aberto para middleware orientado a mensagens.

As características que definem o AMQP são orientação a mensagens, enfileiramento, roteamento (incluindo `point-to-point` e `publish-and-subscribe`), confiabilidade e segurança.

```
AMQP 0-9-1 (Advanced Message Queuing Protocol) is a messaging protocol that enables conforming client applications to communicate with conforming messaging middleware brokers.
```

Exchanges, queues e bindings são denomidados `entidades AMQP` (AMQP entities).

O protocolo AMQP fornece ferramentas para que todas as configurações do broker sejam feitas pelas aplicações que se conectam a ele, sem a necessidade de acessar o painel do RabbitMQ para configurar um broker administrador.

As aplicações podem declarar as entidades AMQP que elas precisam, definir esquemas de roteamento e escolher quando deletar as entidades quando não são mais usadas.

## Message Broker

O sistema de Message Broker do RabbitMQ é composto por:

- Exchanges
  - Bindings
- Queues

### Exchange

A Exchange é uma entidade AMQP para onde as mensagens são enviadas.
Elas fazem o roteamento das mensagens para zero ou mais filas.

O algoritmo de roteamento usado depende do tipo de Exchange e das regras, chamadas `bindings`.

#### Tipos de Exchange

| Tipo | Padrão de nome pré-declarado |
| ---- | ---------------------------- |
| Direct | String vazia / amq.direct |
| Funout | amq.funout |
| Topic | amq.topic |
| Headers | amq.match/amq.headers(RabbitMQ) |

Outros atributos normalmente usados na declaração de uma exchange são:

- **Name**

    Nome da exchange.

- **Durability**

    durabilidade, indica se a exchange continua existindo quando o broker for reiniciado.

- **Auto-delete**

    Se a exchange deve ser removida quando a última mensagem sair.

- **Arguments**

    Opcional, usado por plugins e features específicas do broker.

Exchanges podem ser duráveis(`durable`) ou transitórias(`transient`).

Exchanges duráveis continuam existindo quando o broker for reiniciado, enquanto as transitórias não (elas precisam ser redeclaradas quando o broker volta a ficar disponível).

Nem todos os cenários e casos de uso precisam de exchanges duráveis.

##### Exchange padrão

Ao criar uma exchange com uma string vazia no nome, ela será criada com o tipo `direct` e define as routing keys com o nome das filas conectadas a ela.

##### Direct exchange

A exchange do tipo `direct` faz a entrega da mensagem para as filas com base na chave de roteamento da mensagem (`routing key`).

### Queue

- Se conecta a uma Exchange para receber cópias das mensagens.
- Armazena cópias das mensagens.
- Armazena os registros de leitura das mensagens.

### Bindings

São as configurações das Exchanges, nelas ficam as definições de como a Exchange deve enviar cópias das mensagens recebidas para as filas conectadas nela.

## Message

Mensagens são dados enviados para um Message Broker.

- Producers
- Consumers

Publisher -> Message -> Exchange -> Bindings -> Queue -> Consumer

### Publisher

- Envia mensagens para uma exchange.
- Mais de um Publisher pode se conectar a mesma Exchange para enviar mensagens.

### Consumer

- Faz a leitura das mensagens de uma ou mais filas.

### Criar Publisher/Consumer

O processo de criação segue as etapas a seguir.

1. Criar conexão com o RabbitMQ.
2. Criar channel do Go para processar mensagens em goroutines.
3. Declarar fila (usa uma fila existente ou cria uma nova, caso não exista).
4. Cria um Publisher/Consumer do channel.

**Criar conexão com o RabbitMQ**

```go
conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
failOnError(err, "Failed to connect to RabbitMQ")
defer conn.Close()
```

**Criar channel do Go**

```go
ch, err := conn.Channel()
failOnError(err, "Failed to open a channel")
defer ch.Close()
```

**Declarar fila**

```go
q, err := ch.QueueDeclare(
    "hello",
    false,
    false,
    false,
    false,
    nil,
)
failOnError(err, "Failed to declare a queue")
```

**Cria um Publisher**

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

body := "Hello World!"

err = ch.PublishWithContext(ctx,
    "",
    q.Name,
    false,
    false,
    amqp.Publishing{
        ContentType: "text/plain",
        Body:        []byte(body),
    })
failOnError(err, "Failed to publish a message")
```

**Cria um Consumer**

```go
msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
failOnError(err, "Failed to register a consumer")
```
Consumir mensagens

```go
# channel para receber mensagens lidas
var forever chan struct{}

go func() {
    for d := range msgs {
        log.Printf("Received a message: %s", d.Body)
    }
}()

log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

# retorna conteúdo do channel
<-forever
```

### Message meta-data

São atributos das mensagens que podem ser usados pelo Broker

### Message Acknowledgements

Os consumidores podem sinalizar ao broker que uma mensagem foi completamente processada para indicar que ela pode ser completamente removida da fila.

Esse processo é chamado de `Acknowledgement`.