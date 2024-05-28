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

## Usuario

#### Post

```bash
    /user

```
#### BODY

```JSON
{
  "email": "example@example.com",
  "password": "password123",
  "name": "John Doe",
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




## Campanha de Anuncio

#### Post

```bash
    /campaign/

```
#### BODY

```JSON
{
    "app_secret": "exampleCampaignAccountID",
    "token": "exampleAdAccountID"
}


```
#### Post

```bash
    /campaign/copy

```
#### BODY

```JSON
{
  "account_id": "123456789",
  "buying_type": "AUCTION",
  "campaign_group_id": "987654321",
  "bid_strategy": "LOWEST_COST_WITHOUT_CAP",
  "bid_amount": 5000,
  "can_use_spend_cap": true,
  "configured_status": "ACTIVE",
  "created_time": "2023-05-20T14:30:00Z",
  "daily_budget": "100000",
  "effective_status": "ACTIVE",
  "id": "1122334455",
  "lifetime_budget": "5000000",
  "name": "My Campaign",
  "objective": "CONVERSIONS",
  "spend_cap": "10000000",
  "start_time": "2023-05-25T00:00:00Z",
  "status": "ACTIVE",
  "stop_time": "2023-06-25T00:00:00Z",
  "updated_time": "2023-05-28T14:30:00Z",
  "special_ad_categories": ["NONE"] //é uma lista de categorias de anúncios especiais associadas à campanha (por exemplo, NENHUMA).
}


```
#### Get

```bash
    /campaign?user_id={iddousuario}

```
- Lista todas as campanhas desse usuario(array)
#### Response

```JSON
[{
  "account_id": "123456789",
  "buying_type": "AUCTION",
  "campaign_group_id": "987654321",
  "bid_strategy": "LOWEST_COST_WITHOUT_CAP",
  "bid_amount": 5000,
  "can_use_spend_cap": true,
  "configured_status": "PAUSED",
  "created_time": "2023-05-20T14:30:00Z",
  "daily_budget": "100000",
  "effective_status": "ACTIVE",
  "id": "1122334455",
  "lifetime_budget": "5000000",
  "name": "My Campaign",
  "objective": "CONVERSIONS",
  "spend_cap": "10000000",
  "start_time": "2023-05-25T00:00:00Z",
  "status": "ACTIVE",
  "stop_time": "2023-06-25T00:00:00Z",
  "updated_time": "2023-05-28T14:30:00Z",
  "special_ad_categories": ["NONE"]
}
]

```


#### Get

```bash
    /campaign?cp_id={iddacampanha}

```
- Lista uma campanha em especifico
#### Response

```JSON
{
  "account_id": "123456789",
  "buying_type": "AUCTION",
  "campaign_group_id": "987654321",
  "bid_strategy": "LOWEST_COST_WITHOUT_CAP",
  "bid_amount": 5000,
  "can_use_spend_cap": true,
  "configured_status": "PAUSED",
  "created_time": "2023-05-20T14:30:00Z",
  "daily_budget": "100000",
  "effective_status": "ACTIVE",
  "id": "1122334455",
  "lifetime_budget": "5000000",
  "name": "My Campaign",
  "objective": "CONVERSIONS",
  "spend_cap": "10000000",
  "start_time": "2023-05-25T00:00:00Z",
  "status": "ACTIVE",
  "stop_time": "2023-06-25T00:00:00Z",
  "updated_time": "2023-05-28T14:30:00Z",
  "special_ad_categories": ["NONE"]
}



```


