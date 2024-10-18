# K. Lyrics API

Golang service that serves as an HTTP API for [K. Lyrics](https://kdotlyrics.com).

## Install

```sh
git clone git@github.com:jseashell/kdot-lyrics-api.git
cd kdot-lyrics-api
make
```

## Endpoints

> DynamoDB can be seeded with [https://github.com/jseashell/lyrics-db-seeder](https://github.com/jseashell/lyrics-db-seeder)

### /random

Fetches a random Kendrick Lamar song from DynamoDB.

```sh
curl https://<api-id>.execute-api.us-east-1.amazonaws.com/random-song

# TODO response
```

## 3rd party libraries

- [aws-lambda-go](https://github.com/aws/aws-lambda-go) - AWS Lambda SDK for the Go programming language.
- [google/uuid](https://github.com/google/uuid) -  RFC-4122 compliant UUID module by Google.

## Disclaimer

Repository contributors are not responsible for costs incurred by AWS services.

## License

This software is distributed under the terms of the [MIT License](/LICENSE)
