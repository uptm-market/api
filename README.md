# api

## ROTA PRD

```bash

https://api-production-4147.up.railway.app/

```


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

#### PUT

```bash
    /user/{ID}

```
#### Response

```JSON
{
    "email": "user@example.com",
    "cpf": "123.456.789-00"
}



```


#### PUT

```bash
    /user/{ID}/password

```
#### Response

```JSON
{
    "oldPassword": "oldPassword123",
    "newPassword": "newPassword456"
}




```


## Campanha de Anuncio

#### Post

```bash
    /campaign/?user_id={passar user id aqui}

```
#### BODY

```JSON
{
  "app_secret": "your_campaign_account_id",
  "token": "your_ad_account_id",
  "businessID": ["business_id_1", "business_id_2", "business_id_3"]
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
    /campaign?user_id={iddousuario}?act={passar o act aqui}

```
- Lista todas as campanhas desse usuario(array)
#### Response

```JSON
{
  "data": [
    {
      "buying_type": "AUCTION",
      "effective_status": "ACTIVE",
      "is_budget_schedule_enabled": false,
      "budget_remaining": "0",
      "budget_rebalance_flag": false,
      "objective": "OUTCOME_SALES",
      "smart_promotion_type": "GUIDED_CREATION",
      "status": "ACTIVE",
      "special_ad_category": "NONE",
      "adsets": {
        "data": [
          {
            "billing_event": "IMPRESSIONS",
            "budget_remaining": "25000",
            "daily_budget": "25000",
            "lifetime_budget": "0",
            "id": "120211909128310758"
          },
          {
            "billing_event": "IMPRESSIONS",
            "budget_remaining": "5000",
            "daily_budget": "5000",
            "lifetime_budget": "0",
            "id": "120211909128300758"
          },
          {
            "billing_event": "IMPRESSIONS",
            "budget_remaining": "5000",
            "daily_budget": "5000",
            "lifetime_budget": "0",
            "id": "120211909128280758"
          },
          {
            "billing_event": "IMPRESSIONS",
            "budget_remaining": "5000",
            "daily_budget": "5000",
            "lifetime_budget": "0",
            "id": "120211909128260758"
          },
          {
            "billing_event": "IMPRESSIONS",
            "budget_remaining": "5000",
            "daily_budget": "5000",
            "lifetime_budget": "0",
            "id": "120211909100300758"
          }
        ],
        "paging": {
          "cursors": {
            "before": "QVFIUlJlYVJZAN19LcGpabnhoZAnU5ZAVBOaGZAIbXBHbjF4UWszMjNLSWptR2dRTzFLOUQ3bzVHM1pkblVLLXdSb0tCNUZAMZAzN3NllWRXFVZAHlDNVdoRjJINGlR",
            "after": "QVFIUmtuaEhhZAG1GX2xqNktSSDRpYnhoMko2ekt3NERTVncwSmlwUXpqWkRCci10MHMzendlTE9NR1hNS3RkZAjZAJOTBKNDk1VkNwRm5HVnRPN19MQ3hLZAmlB"
          }
        }
      },
      "insights": {
        "data": [
          {
            "account_currency": "BRL",
            "buying_type": "AUCTION",
            "cost_per_unique_click": "1.042935",
            "cpc": "0.841605",
            "cpm": "50.571217",
            "date_start": "2024-07-08",
            "frequency": "1.288831",
            "reach": "11505",
            "website_ctr": [
              {
                "action_type": "link_click",
                "value": "5.138926"
              }
            ],
            "date_stop": "2024-08-06"
          }
        ],
        "paging": {
          "cursors": {
            "before": "MAZDZD",
            "after": "MAZDZD"
          }
        }
      },
      "id": "120211909100120758"
    },
    {
      "buying_type": "AUCTION",
      "effective_status": "PAUSED",
      "is_budget_schedule_enabled": false,
      "budget_remaining": "0",
      "budget_rebalance_flag": false,
      "objective": "OUTCOME_ENGAGEMENT",
      "smart_promotion_type": "GUIDED_CREATION",
      "status": "PAUSED",
      "special_ad_category": "NONE",
      "adsets": {
        "data": [
          {
            "billing_event": "IMPRESSIONS",
            "budget_remaining": "5000",
            "daily_budget": "5000",
            "lifetime_budget": "0",
            "id": "120211688935480758"
          },
          {
            "billing_event": "IMPRESSIONS",
            "budget_remaining": "5000",
            "daily_budget": "5000",
            "lifetime_budget": "0",
            "id": "120211688385190758"
          }
        ],
        "paging": {
          "cursors": {
            "before": "QVFIUk1rQTVDNzFuakZAndXZACSVU5OFFUZA09YbW5Jb0lhUG42NjFScG9WTW9zM0NNcGRTTUFaR1kyanJydm16TFJKOE45c1hXWTRpTkJmb0RXQjZAZAamZAMV2Fn",
            "after": "QVFIUnhQR3dkeDdibGItT0E4TGM0MXVqNmJZAdHVxYzcta0FYOERCcmdEQzItNHUtcE90SnZAVRGNOd05PYnhoOWlVdmJuLWJfNUtwdU5oblBoUU5pTUJsbk13"
          }
        }
      },
      "insights": {
        "data": [
          {
            "account_currency": "BRL",
            "buying_type": "AUCTION",
            "cost_per_unique_click": "1.215",
            "cpc": "1.07797",
            "cpm": "31.895439",
            "date_start": "2024-07-08",
            "frequency": "1.20996",
            "reach": "3715",
            "website_ctr": [
              {
                "action_type": "link_click",
                "value": "1.89099"
              }
            ],
            "date_stop": "2024-08-06"
          }
        ],
        "paging": {
          "cursors": {
            "before": "MAZDZD",
            "after": "MAZDZD"
          }
        }
      },
      "id": "120211688385150758"
    },
    {
      "daily_budget": "30000",
      "buying_type": "AUCTION",
      "effective_status": "PAUSED",
      "is_budget_schedule_enabled": false,
      "budget_remaining": "30000",
      "budget_rebalance_flag": false,
      "objective": "OUTCOME_SALES",
      "pacing_type": [
        "standard"
      ],
      "smart_promotion_type": "GUIDED_CREATION",
      "status": "PAUSED",
      "special_ad_category": "NONE",
      "bid_strategy": "LOWEST_COST_WITHOUT_CAP",
      "adsets": {
        "data": [
          {
            "billing_event": "IMPRESSIONS",
            "budget_remaining": "0",
            "id": "120211310840660758"
          }
        ],
        "paging": {
          "cursors": {
            "before": "QVFIUjdTbEIyUDY1bXF0eEVULXNKRkc0Wkszc2JTVlFEbWlBaFVhbEdtWFdiSncxMDA1dkpLbE90d0FnQXBuajlqT21wSU5waE1zMjcyUWlWT0ZAuNnlBSGNR",
            "after": "QVFIUjdTbEIyUDY1bXF0eEVULXNKRkc0Wkszc2JTVlFEbWlBaFVhbEdtWFdiSncxMDA1dkpLbE90d0FnQXBuajlqT21wSU5waE1zMjcyUWlWT0ZAuNnlBSGNR"
          }
        }
      },
      "insights": {
        "data": [
          {
            "account_currency": "BRL",
            "buying_type": "AUCTION",
            "cost_per_unique_click": "1.218816",
            "cpc": "1.029222",
            "cpm": "30.993754",
            "date_start": "2024-07-08",
            "frequency": "1.088635",
            "reach": "8236",
            "website_ctr": [
              {
                "action_type": "link_click",
                "value": "2.409101"
              }
            ],
            "date_stop": "2024-08-06"
          }
        ],
        "paging": {
          "cursors": {
            "before": "MAZDZD",
            "after": "MAZDZD"
          }
        }
      },
      "id": "120211310840550758"
    }
  ],
  "paging": {
    "cursors": {
      "before": "QVFIUjJ1UGFQMXItUGUzZAVk0dGtNSEJqR1VnRGg2bGMzUkpEVVQwZAmlrNFNDcjZABOXF5VjBsMTFqRlVwTzBrV2tOZAkhSX05OdE1LNWxiYjd0cGd0VEc4VTJ3",
      "after": "QVFIUjZAMZAkUtcGNJZAWx1U25VUmF2LWgtRHRwVDZAYckRLaTg1UWpsTVRlOExENUF2NWNRdDl6WWJYVUFjRnpxMVRRaEs3bXI2amxIcVZAnWEpyT3k1THZA5bmxB"
    }
  }
}

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


## NOVOS

- PUT /campaign/active/:id
- GET /campaign/list/businessid/:userId
- GET /campaign/listAll/:userId
```JSON
{
  "name": "Business Name",
  "id": "Business ID"
}


```




## CLONAR CAMPANHA



```BASH
/campaign/{act}/{userId}/copy

``


BODY



```BASH
{
  "adlabels": [
    {
      "id": "1234567890",
      "name": "Sample Label"
    }
  ],
  "bid_strategy": "MENOR_CUSTO_SEM_CAP",
  "budget_schedule_specs": [
    {
      "time_start": 1625097600,
      "time_end": 1625184000,
      "budget_value": 1000,
      "budget_value_type": "DAILY"
    }
  ],
  "buying_type": "AUCTION",
  "campaign_optimization_type": "NENHUM",
  "daily_budget": 5000,
  "execution_options": [
    "validate_only",
    "include_recommendations"
  ],
  "is_skadnetwork_attribution": false,
  "is_using_l3_schedule": false,
  "iterative_split_test_configs": [],
  "lifetime_budget": 100000,
  "name": "Sample Campaign",
  "objective": "INSTALACOES_DE_APLICATIVOS",
  "promoted_object": {},
  "source_campaign_id": 9876543210,
  "special_ad_categories": [],
  "special_ad_category_country": ["US"],
  "spend_cap": 200000,
  "start_time": "2024-07-29T00:00:00Z",
  "status": "ACTIVE",
  "stop_time": "2024-08-29T00:00:00Z",
  "topline_id": 54321
}



```

## Rota que retornar os acts


```bash

/campaign/{userId}/act
```


- Response


```json

{
  "owned_ad_accounts": {
    "data": [
      {
        "account_id": "1234567890",
        "name": "Minha Conta de Anúncios",
        "account_status": 1,
        "currency": "USD",
        "timezone_name": "America/Los_Angeles",
        "timezone_offset_hours_utc": -7,
        "spend_cap": 1000,
        "amount_spent": 500
        // Outros campos podem estar presentes aqui
      },
      // Outras contas de anúncios
    ],
    "paging": {
      "cursors": {
        "before": "abc123",
        "after": "xyz789"
      }
    }
  },
  "id": "businessID"
}



```