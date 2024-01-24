# ugc-5-6

## How to run

Prepare two terminal. One one terminal, go to folder `account-service`:

```shell
cd account-service
```

And execute:

```shell
make run
```

And on the other, go to folder `api-gateway`:

```shell
cd api-gateway
```

And execute:

```shell
make run
```

Make sure that you already have swaggo installed. If not, you can use:

```shell
make swag-install
```

. Now that both applications are running, go to your web browser and go to:

```shell
localhost:8080/swagger/index.html
```
