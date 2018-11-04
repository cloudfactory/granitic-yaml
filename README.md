# granitic-yaml
Extension to Granitic to support YAML configuration and component definition files instead of JSON.

## Pre-requisites

You must be using Granitic 1.3 or later


## Instructions for use

Open a terminal and run

```
    go get github.com/graniticio/granitic-yaml
```

### Install YAML specific versions of Granitic tools

```
    go install github.com/graniticio/granitic-yaml/cmd/grnc-yaml-bind
    go install github.com/graniticio/granitic-yaml/cmd/grnc-yaml-bind
```

If you will only be using YAML for your Granitic projects, you might want to remove the JSON versions of the
tools to avoid confusion.

```
    rm $GOPATH/bin/grnc-project
    rm $GOPATH/bin/grnc-bind
```

## Modifying your existing projects

### Go files

Your existing JSON-based projects will have a `main` file where Granitic is invoked (usually in the root of your project and called `service.go`)

You will need to change the import in this file from:

```go
    import "github.com/graniticio/granitic"
```

to

```go
    import "github.com/graniticio/granitic_yaml"    
```


and the line that invokes Granitic from

```go
    granitic.StartGranitic(bindings.Components())
```

to

```go
    granitic_yaml.StartGraniticWithYaml(bindings.Components())
```

### JSON files

Your can use an online tool to convert your JSON files to YAML or re-create them by hand. They must be saved with
either a `yml` or `yaml` extension. If you are serving your configuration files with a web-server, the content type
should be set to one of:

```
    text/x-yaml
    application/yaml
    text/yaml
    application/x-yaml
```

If you are using an online tool you may find that string lists are converted to the verbose YAML syntax which may make
Granitic validation rules unreadable. See the FAQ and examples section below for more examples.

## FAQ

### Why isn't this feature part of the Granitic core?

The Granitic core has a firm design principle of not requiring any downstream dependencies. As YAML support
is not included in Go and requires a third-party library, YAML support is included as a separate package.

### Can I include comments in component defintion and configuration files?

Yes

### Can I mix and match YAML and JSON component definition files?

Not at this time.

### Can I mix and match YAML and JSON configuration files?

Yes - because Granitic's internal _facilities_ configuration is JSON based, this was a required feature.

### What list syntax should I use?

YAML has two ways of representing lists:

```yaml
MyList:
  - Element1
  - Element2
```

```yaml
MyList: [Element1, Element2]
```

The first is generally considered more idiomatic, but the second can be more readable, especially for Granitic's 
validation rules. See the examples below.

### Which YAML parser is being used?

[https://github.com/go-yaml/yaml](https://github.com/go-yaml/yaml)

### Why don't empty projects created with grnc-yaml-project start?

The YAML parser removes empty structures, because the template YAML files created by the tool don't specify any
components or packages, the sections are removed by the parser and Granitic won't start. As soon you add at least
one package and component your project will start.

## Examples

The following examples shows what the component definition files and configuration files from Granitic's 
[Advanced validation tutorial](http://www.granitic.io/tutorials/advanced-validation) 


### Component definition file

```yaml
packages:
  - github.com/graniticio/granitic/ws/handler
  - granitic-tutorial/recordstore/endpoint
  - github.com/graniticio/granitic/validate
  - github.com/go-sql-driver/mysql
  - granitic-tutorial/recordstore/db

components:

  # GET Artist components
  artistLogic:
    type: endpoint.ArtistLogic
    EnvLabel: conf:environment.label

  artistHandler:
    type: handler.WsHandler
    HttpMethod: GET
    Logic: ref:artistLogic
    PathPattern: "^/artist/([\\d]+)[/]?$" #Complex strings with escapes should be quoted
    BindPathParams:
      - Id
    FieldQueryParam:
      NormaliseName: normalise

  # POST Artist components
  submitArtistLogic:
    type: endpoint.SubmitArtistLogic

  submitArtistHandler:
    type: handler.WsHandler
    HttpMethod: POST
    Logic: ref:submitArtistLogic
    PathPattern: "^/artist[/]?$"
    AutoValidator: ref:submitArtistValidator

  submitArtistValidator:
    type: validate.RuleValidator
    DefaultErrorCode: INVALID_ARTIST
    Rules: conf:submitArtistRules

  # Database access components
  dbProvider:
    type: db.MySqlProvider
    Config: ref:dbConnection

  dbConnection:
    type: mysql.Config
    User: grnc
    Passwd: "OKnasd8!k"
    Addr: localhost
    DBName: recordstore

```

### Configuration file

```yaml
# Framework configuration
Facilities:
  HttpServer: true
  JsonWs: true
  RuntimeCtl: true
  ServiceErrorManager: true
  QueryManager: true
  RdbmsAccess: true

ApplicationLogger:
  GlobalLogLevel: INFO

QueryManager:
  ProcessorName: sql

# Application configuration
environment:
  label: DEV

# Rules are more readable using the terse form of YAML lists - any YAML control characters (colons) need to be in quotes
submitArtistRules:
  - [Name, STR, 'REQ:NAME_MISSING', TRIM, STOPALL, 'LEN:5-50:NAME_BAD_LENGTH', BREAK, 'REG:^[A-Z]| +$:NAME_BAD_CONTENT']
  - [FirstYearActive, INT,  'RANGE:1700|2100:FIRST_ACTIVE_INVALID']

# As are service errors, but as messages are likely to contain commas, put them in quotes
serviceErrors:
  - [C, INVALID_ARTIST, "Cannot create an artist with the information provided."]
  - [C, NAME_MISSING, "You must supply the Name field on your submission."]
  - [C, NAME_BAD_LENGTH, "Names must be 5-50 characters in length."]
  - [C, NAME_BAD_CONTENT, "Names can only contain letters and spaces."]
  - [C, FIRST_ACTIVE_INVALID, "FirstYearActive must be in the range 1700-2100"]
```

