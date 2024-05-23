# api


## Instalar

- GOLANG
- DOCKER

## só de rodar a api já sobe o banco de dados
```bash
sudo docker-compose up --build
```
## Caso venha a enfrentar problemas rode 

```bash
 sudo docker-compose down

```

## Endpoints

#### Post

```bash
    /user

```
#### BODY

```JSON
{
  "id": 123,
  "email": "example@example.com",
  "password": "password123",
  "name": "John Doe",
  "cellPhone": "123-456-7890",
  "city": "New York",
  "state": "NY",
  "zipCode": "10001",
  "district": "Downtown",
  "street": "Main Street",
  "streetNumber": "123",
  "complement": "Apt 1A",
  "birthdate": "1990-01-01",
  "gender": "male",
  "cpf": "123.456.789-10",
  "source": 1
}

```
#### Post

```bash
    /user/login

```
#### BODY

```JSON
{
  "email": "example@example.com",
  "password": "password123"
}

```
#### Get

```bash
    /user/me

```
#### Response

```JSON
{
  "id": 123,
  "email": "example@example.com",
  "name": "John Doe",
  "cellPhone": {
    "ddd": "22",
    "number": "4567890"
  },
  "city": "Sample City",
  "state": "Sample State"
}


```


