# REST API

My "Lorem Ipsum" implementation of a REST API with authentication, password hashing, and JSON web tokens.

## Project Directory Structure
Here's the directory structure of the project:
```mermaid
graph LR;
    root[root]
    root --> cmd
    root --> db
    root --> middlewares
    root --> models
    root --> routes
    root --> utils
    root --> LICENSE
    root --> README.md
    cmd --> api
    api --> main.go
    db --> db.go
    middlewares --> auth.go
    models --> event.go
    models --> user.go
    routes --> events.go
    routes --> register.go
    routes --> routes.go
    routes --> users.go
    utils --> hash.go
    utils --> jwt.go
```

## License
This project is licensed under the [MIT License](LICENSE).