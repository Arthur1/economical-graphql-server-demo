# economical-graphql-server-demo

## run local server

```sh
make db.sqlite
go run cmd/server
```

### deploy Lambda server

```console
$ make server-lambda.zip
$ export AWS_PROFILE="hoge"
$ (cd terraform; terraform init; terraform apply)
```

## sample query

```graphql
{
  songs(first: 3) {
    edges {
      node {
        title
        releasedYear
        artist {
          name
        }
      }
    }
    pageInfo {
      startCursor
      endCursor
    }
    totalCount
  }
}
```

```json
{
  "data": {
    "songs": {
      "edges": [
        {
          "node": {
            "title": "北ウイング",
            "releasedYear": 1984,
            "artist": {
              "name": "中森明菜"
            }
          }
        },
        {
          "node": {
            "title": "悲しみがとまらない",
            "releasedYear": 1983,
            "artist": {
              "name": "杏里"
            }
          }
        },
        {
          "node": {
            "title": "シンデレラ・ハネムーン",
            "releasedYear": 1978,
            "artist": {
              "name": "岩崎宏美"
            }
          }
        }
      ],
      "pageInfo": {
        "startCursor": "gaFp0wAAAAAAAAAB",
        "endCursor": "gaFp0wAAAAAAAAAD"
      },
      "totalCount": 19
    }
  }
}
```
