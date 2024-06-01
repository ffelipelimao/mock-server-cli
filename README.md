### Description

- This CLI is responsible to create a web server with your requirements

### Installation

`go get -u github.com/ffelipelimao/mock-server-cli`

### Usage

- Create a file json with the <b>port</b> and <b>requests</b> objects

    - <b>ports</b>: responsible to your HTTP port
    - <b>requests</b>: An array with the specification of each request

- In <b>requests</b> arrays must contain an <b>verb</b>,<b>path</b>,<b>responseType</b> <b>response.body</b>
    - <b>verb</b>: an valid HTTP verb
    - <b>path</b>: your path to your HTTP response, use `{}` to variables
    - <b>responseType</b>: only use <b>json</b> or <b>xml</b> in lower case
    - <b>response.body</b>: if your using json just paste your response, if your are using xml your response.body should by STRING. Example: `"<merchant>\n  <id>19b58298-c473-4190-984c-9aca82bbd3a8<id>\n </merchant>"`

```json
{
    "port": "8080",
    "requests": [
        {
            "verb": "GET",
            "path": "/v2/awesome/example/{id}",
            "responseType": "json",
            "response": {
                "body": {
                    "merchant": {
                        "id": "19b58298-c473-4190-984c-9aca82bbd3a8",
                        "name": "Exemplo Loja",
                        "location": "São Paulo",
                        "products": [
                            {
                                "product": {
                                    "name": "Produto 1",
                                    "price": 10.00
                                }
                            },
                            {
                                "product": {
                                    "name": "Produto 2",
                                    "price": 20.00
                                }
                            }
                        ]
                    }
                }
            }
        }
    ]
}
```