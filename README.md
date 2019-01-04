# Consumer Driven Testing

## Description

Exploring how to use [pact](https://docs.pact.io/) for better service tests using contracts.

Intention is to better understand how the tool works and see if it is a viable option for replacing integration tests

## Structure

The repository is split into an [api](./api) and a [consumer](./consumer).

The api is a basic REST backend API that will be queried by the consumer.

The consumer is a client that wants information from the backend api.

Both are currently written in Go for speed, but would like to rewrite the consumer in JS or Swift to understand
the different [pact libraries](https://docs.pact.io/implementation_guides)

## License

MIT
