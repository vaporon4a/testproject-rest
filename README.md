# testproject-rest
REST API

## ТЗ

Добрый день, уважаемый соискатель, данное задание нацелено на выявление вашего реального уровня в разработке на golang, поэтому отнеситесь к нему, как к работе на проекте. Выполняйте его честно и проявите себя по максимуму, удачи!

### API

#### POST /api/v1/wallet

* `walletId`: UUID
* `operationType`: DEPOSIT or WITHDRAW
* `amount`: int

### GET /api/v1/wallets/{WALLET_UUID}

* `WALLET_UUID`: UUID

### Requirements

* Stack: Golang, Postgresql, Docker
* Highload: 1000 RPS per wallet
* Docker container
* Docker-compose
* Test coverage
