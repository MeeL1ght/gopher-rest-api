# REST Api - By Moi

## This is a `REST Api` developed in order to practice with Golang x)

<!-- Banner -->
<img src="./img/banner-rest-api-golang.png" />

## βοΈ Project setup

<div style="display: flex; align-items: center">
  <img
    src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1200px-Go_Logo_Blue.svg.png"
    style="height:28px; width: 75px; cursor: pointer"
  />
  <a
    href="https://go.dev/doc/install"
    target="_blank"
    style="margin-left: 10px"
  >
    βΎ Install Golang
  </a>
</div>

<br>

<div style="display: flex; align-items: center">
  <img
    src="https://umangsoftware.com/wp-content/uploads/2020/05/MongoDB-logo.png"
    style="height:30px; width: 100px; cursor: pointer"
  />
  <a
    href="https://www.mongodb.com/docs/manual/installation/"
    target="_blank"
    style="margin-left: 10px"
  >
    βΎ Install MongoDB
  </a>
</div>

<hr>

## π Run project:

```bash
exec.bat
```
Or

```bash
go run .
```

<hr>

## π¨βπ« { Schemas }

<img
  src="./img/mongodb-schema-gopher.png"
  style="height: 400px; width: 900px"
  title="MongoDB - Schema"
/>

<br>

<img
  src="./img/golang-schema-gopher.png"
  style="height: 400px; width: 900px"
  title="Golang - Schema"
/>

<hr>

## π§ΎοΈ Config file

You can configure the environment variables<br>
in the file `'./config/config.yaml'`

```yaml
# Application
app:
  environment: 'dev'
  host: '127.0.0.1'
  port: 8080
  url: 'http://127.0.0.1:8080'
  endpoint: '/gopher'
db:
  username: ''
  password: ''
  db_name: 'gophersApi'
  db_uri: 'mongodb://127.0.0.1:27017/gophersApi'
```

<hr>

## π¦ Packages used:

<ol>
  <li>
    <a
      href="https://github.com/gofiber/fiber"
      target="_blank"
    >
      π₯ Fiber (Framework)
    </a>
  </li>
  <li>
    <a
      href="https://www.mongodb.com/docs/drivers/go/current/"
      target="_blank"
    >
      π MongoDB Driver
    </a>
  </li>
  <li>
    <a
      href="https://github.com/go-playground/validator"
      target="_blank"
    >
      β Validator
    </a>
  </li>
  <li>
    <a
      href="github.com/spf13/viper"
      target="_blank"
    >
      π Viper
    </a>
  </li>
</ol>
