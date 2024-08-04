# KDot Info API

Golang service that serves as an HTTP API for KDot info.

## Install

```sh
git clone git@github.com:jseashell/kdot-info.git
cd kdot-info
make
```

## Endpoints

> DynamoDB can be seeded with [https://github.com/jseashell/lyrics-db-seeder](https://github.com/jseashell/lyrics-db-seeder)

### /random

Fetches a random Kendrick Lamar song from DynamoDB

```sh
curl https://<api-id>.execute-api.us-east-1.amazonaws.com/random

# TODO response
```

## Deploy

1. [Configure the AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html) on your local workstation.
1. Run the deploy script

    ```sh
    make deploy
    ```

## 3rd party libraries

- [aws-lambda-go](https://github.com/aws/aws-lambda-go) - AWS Lambda SDK for the Go programming language.
- [google/uuid](https://github.com/google/uuid) -  RFC-4122 compliant UUID module by Google.

## Disclaimer

Repository contributors are not responsible for costs incurred by AWS services.

## License

This software is distributed under the terms of the [MIT License](/LICENSE)