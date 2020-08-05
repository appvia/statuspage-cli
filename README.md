# Statuspage CLI

Statuspage CLI allows you to create, read, update and delete Statuspage resources.

## Prerequisite
It is assumed that you have created an Atlassian Statuspage account, a page and an API key.

## Getting Started

Statuspage resources exist within the scope of a `page`, therefore `--page-id` or `-p` is a required global flag with all commands.

The Statuspage CLI only supports key-based authentication, therefore `API_KEY` environment variable (default) or `--key-id` or `-k` is a required global flag with all commands.

### Installing

```
$ go build -o statuspage
```

### Usage

```
$ ./statuspage
A command line interface for Atlassian statuspage.

Usage:
  statuspage [command]

Available Commands:
  create      Creates one of more resources in statuspage
  delete      Allows you to delete one of more resources in statuspage
  get         Allows you to get one of more resources in statuspage
  help        Help about any command
  update      Allows you to update one of more resources in statuspage

Flags:
      --config string   config file (default is $HOME/.statuspage.yaml)
  -h, --help            help for statuspage
  -t, --toggle          Help message for toggle

Use "statuspage [command] --help" for more information about a command.
```

### Get page identifier
```
PAGE_ID=$(./statuspage get page -k <API_KEY> | jq -r .[].id)
```

### Get component identifier
```
COMPONENT_ID=$(./statuspage get component -k <API_KEY> -p $PAGE_ID | jq '.[] | select(.name=="<COMPONENT_NAME") | .id')
```

### Create incident and associate a component
```
./statuspage create incident -k <API_KEY> -n 'Example incident' -b 'created by Statuspage CLI' -p $PAGE_ID -s investigating -c $COMPONENT_ID=<COMPONENT_STATUS>
```

### Create incident and associate more than one component
```
COMPONENT_1_ID=$(./statuspage get component -k <API_KEY> -p $PAGE_ID | jq '.[] | select(.name=="<COMPONENT_1_NAME") | .id')

COMPONENT_2_ID=$(./statuspage get component -k <API_KEY> -p $PAGE_ID | jq '.[] | select(.name=="<COMPONENT_2_NAME") | .id')

./statuspage create incident -k <API_KEY> -n 'Example incident' -b 'created by Statuspage CLI' -p $PAGE_ID -s investigating -c $COMPONENT_1_ID=<COMPONENT_1_STATUS> -c $COMPONENT_2_ID=<COMPONENT_2_STATUS>
```

### Update incident and associated component(s)
```
INCIDENT_ID=$(./statuspage get incident -k <API_KEY> -p $PAGE_ID | jq '.[] | select(.name=="<INCIDENT_NAME>") | .id')

./statuspage update incident -k <API_KEY> -b 'created by the statuspage CLI' -i $INCIDENT_ID -p $PAGE_ID -s identified -c $COMPONENT_1_ID=<COMPONENT_1_STATUS> -c $COMPONENT_2_ID=<COMPONENT_2_STATUS>
```
