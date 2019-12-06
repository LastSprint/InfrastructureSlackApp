# InfrastructureSlackApp

Конфиг файл должен лежать в папке `config` которая лежит рядом с бинарником

```Json
{
    "jiraLogin": string,
    "jiraPassword": string,
    "slackToken": string,
    "figmaToken": string,
    "mongodb": {
        "connectionString": string,
        "dataBaseString": string,
        "testConnectionString": string,
        "testDataBaseString": string
    },
    "jiraBaseHost": "https://jira.serv.com/browse/",
    "jiraApiSearchUrl": "https://jira.serv.com/rest/api/2/search"
}
```