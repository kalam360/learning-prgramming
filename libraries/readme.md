# Learning path

## Web Fullstack development

```mermaid
graph TD;
    vite_app --> html_css_responsive_ui_tailwind
    -->vuejs_frontend;

    postgresql/mongodb --> sql_alchemy/pymongo 
    --> alembic_migration --> Fastapi_pydantic/express_schema_backend;  

    Fastapi_pydantic/express_schema_backend --> app_jwt_auth;

    vuejs_frontend --> app_jwt_auth;

    app_jwt_auth --> docker_kubernates_orchestration
    --> celeryworker_mher/flower --> rabbitmq/kafka 
    --> traefik_reverse_proxy;

    d3/threejs/echarts-->vuejs_frontend;

    solidity_truffle_smartcontracts --> ethereumjs/web3js/geth --> vuejs_frontend;


```

## Desktop Application

```mermaid
graph TD;
    javaFX-->gradle --> vscode;
    Qt_cpp--> cmake_vcpkg --> vscode;
    c#_windows--> .net --> vscode;
    vscode --> Desktop_Application
```

