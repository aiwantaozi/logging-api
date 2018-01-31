logging-api
========

A microservice that does micro things.

## Building

`make`


## Running

`./bin/logging-api`

## API Design

There are two type logging config, cluster level and project level, under cluster level logging user could select one of targets below: Embedded Elasticsearch, Elasticsearch, Splunk, Kafka, Syslog. Under project level, Elasticsearch, Splunk, Kafka, Syslog targets are available, Embedded is not available under project.

### Cluster Level
#### CREATE
```
HTTP/1.1 POST /v3/clusters/local/loggings
Content-Type: application/json

{
    "currentTarget": "elasticsearch",
    "elasticsearchConfig": {
        "host": "192.168.1.10",
        "port": 9200,
        "indexPrefix": "local-", //format is cluster_id-
        "dateformat": "YYYY.MM.DD",
        "authUser": "admin",
        "authPassword": "admin"
    },
    "splunkConfig: {
        "host": "192.168.1.10",
        "port": 9200,
        "protocol": "https",
        "source": "rancher",
        "timeFormat": "unixtime",
        "token": "jdhfajhdfaldfhalkfhjl"
    },
    "embeddedConfig": {
        "kibanaURL": "192.168.1.10:5601"
    },
    "kafkaConfig": { 
        "brokerType": "broker",
        "brokers": ["192.168.1.10:2188", "192.168.1.11:2188"],
        "topic": "rancher",
        "dataType": "json",
        "maxSendRetries": 3
    },
    "syslogConfig": {
        "host": "192.168.1.12",
        "port": 514,rancher
        "severity": "notice",
        "program": "fluentd"
    },
    "secrets":[
        {"type":"secretReference","uid":"0","gid":"0","mode":"444","name":"elasticsearch-username","secretId":"ns2:username"},
        {"type":"secretReference","uid":"0","gid":"0","mode":"444","name":"elasticsearch-password","secretId":"ns2:password"}
    ]
}
```

Attention: the kafka also can use zookeep address to send data, the kafkaConfig will be:
    "kafkaConfig": { 
        "brokerType": "zookeeper",
        "zookeeperHost": "192.168.1.10",
        "zookeeperPort": 2199,
        "topic": "rancher",
        "outputDataType": "json",
        "maxSendRetries": 3
    }

#### DELETE
```
HTTP/1.1 DELETE /v3/clusters/local/loggings/cluster-logging
```
#### UPDATE
```
HTTP/1.1 PUT /v3/clusters/local/loggings/cluster-logging
```
### Project Level

#### CREATE
```
HTTP/1.1 POST /v3/projects/local:project-id/projectlogging
{
    "currentTarget": "elasticsearch",
    "elasticsearchConfig": {
        "host": "192.168.1.10",
        "port": 9200,
        "indexPrefix": "local-", //format is cluster_id-
        "dateformat": "YYYY.MM.DD",
        "authUser": "admin",
        "authPassword": "admin"
    },
    "splunkConfig: {
        "host": "192.168.1.10",
        "port": 9200,
        "protocol": "https",
        "source": "rancher",
        "timeFormat": "unixtime",
        "token": "jdhfajhdfaldfhalkfhjl"
    },
    "kafkaConfig": { 
        "brokerType": "broker",
        "brokers": ["192.168.1.10:2188", "192.168.1.11:2188"],
        "topic": "rancher",
        "dataType": "json",
        "maxSendRetries": 3
    },
    "syslogConfig": {
        "host": "192.168.1.12",
        "port": 514,rancher
        "severity": "notice",
    },
    "secrets":[
        {"type":"secretReference","uid":"0","gid":"0","mode":"444","name":"elasticsearch-username","secretId":"ns2:username"},
        {"type":"secretReference","uid":"0","gid":"0","mode":"444","name":"elasticsearch-password","secretId":"ns2:password"}
    ]
}
```
#### DELETE
```
HTTP/1.1 DELETE /v3/projects/local:project-id/projectlogging/project-logging
```
#### UPDATE
```
HTTP/1.1 PUT /v3/projects/local:project-id/projectlogging/project-logging
```
### Schemas
```
{
    "type": "collection",
    "links": {
        "self": "https://192.168.1.16:8443/v3/schemas"
    },
    "actions": {},
    "pagination": {
        "limit": 1000
    },
    "sort": {
        "order": "asc",
        "reverse": "https://192.168.1.16:8443/v3/schemas?order=desc"
    },
    "resourceType": "schema",
    "data": [
        {
            "actions": {},
            "baseType": "schema",
            "id": "action",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/action"
            },
            "pluralName": "actions",
            "resourceFields": {
                "input": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "output": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "alertCondition",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/alertCondition"
            },
            "pluralName": "alertConditions",
            "resourceFields": {
                "lastTransitionTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "lastUpdateTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "message": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "reason": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "status": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "type": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "alertStatus",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/alertStatus"
            },
            "pluralName": "alertStatuses",
            "resourceFields": {
                "conditions": {
                    "create": false,
                    "nullable": true,
                    "type": "array[alertCondition]",
                    "update": false
                },
                "startedAt": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "state": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "allowedHostPath",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/allowedHostPath"
            },
            "pluralName": "allowedHostPaths",
            "resourceFields": {
                "pathPrefix": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "amazonec2config",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/amazonec2config"
            },
            "pluralName": "amazonec2configs",
            "resourceFields": {
                "accessKey": {
                    "create": true,
                    "default": "",
                    "description": "AWS Access Key",
                    "type": "string",
                    "update": false
                },
                "ami": {
                    "create": true,
                    "default": "",
                    "description": "AWS machine image",
                    "type": "string",
                    "update": false
                },
                "blockDurationMinutes": {
                    "create": true,
                    "default": "0",
                    "description": "AWS spot instance duration in minutes (60, 120, 180, 240, 300, or 360)",
                    "type": "string",
                    "update": false
                },
                "deviceName": {
                    "create": true,
                    "default": "/dev/sda1",
                    "description": "AWS root device name",
                    "type": "string",
                    "update": false
                },
                "endpoint": {
                    "create": true,
                    "default": "",
                    "description": "Optional endpoint URL (hostname only or fully qualified URI)",
                    "type": "string",
                    "update": false
                },
                "iamInstanceProfile": {
                    "create": true,
                    "default": "",
                    "description": "AWS IAM Instance Profile",
                    "type": "string",
                    "update": false
                },
                "insecureTransport": {
                    "create": true,
                    "default": false,
                    "description": "Disable SSL when sending requests",
                    "type": "boolean",
                    "update": false
                },
                "instanceType": {
                    "create": true,
                    "default": "t2.micro",
                    "description": "AWS instance type",
                    "type": "string",
                    "update": false
                },
                "keypairName": {
                    "create": true,
                    "default": "",
                    "description": "AWS keypair to use; requires --amazonec2-ssh-keypath",
                    "type": "string",
                    "update": false
                },
                "monitoring": {
                    "create": true,
                    "default": false,
                    "description": "Set this flag to enable CloudWatch monitoring",
                    "type": "boolean",
                    "update": false
                },
                "openPort": {
                    "create": true,
                    "description": "Make the specified port number accessible from the Internet",
                    "type": "array[string]",
                    "update": false
                },
                "privateAddressOnly": {
                    "create": true,
                    "default": false,
                    "description": "Only use a private IP address",
                    "type": "boolean",
                    "update": false
                },
                "region": {
                    "create": true,
                    "default": "us-east-1",
                    "description": "AWS region",
                    "type": "string",
                    "update": false
                },
                "requestSpotInstance": {
                    "create": true,
                    "default": false,
                    "description": "Set this flag to request spot instance",
                    "type": "boolean",
                    "update": false
                },
                "retries": {
                    "create": true,
                    "default": "5",
                    "description": "Set retry count for recoverable failures (use -1 to disable)",
                    "type": "string",
                    "update": false
                },
                "rootSize": {
                    "create": true,
                    "default": "16",
                    "description": "AWS root disk size (in GB)",
                    "type": "string",
                    "update": false
                },
                "secretKey": {
                    "create": true,
                    "default": "",
                    "description": "AWS Secret Key",
                    "type": "string",
                    "update": false
                },
                "securityGroup": {
                    "create": true,
                    "default": [
                        "docker-machine"
                    ],
                    "description": "AWS VPC security group",
                    "type": "array[string]",
                    "update": false
                },
                "sessionToken": {
                    "create": true,
                    "default": "",
                    "description": "AWS Session Token",
                    "type": "string",
                    "update": false
                },
                "spotPrice": {
                    "create": true,
                    "default": "0.50",
                    "description": "AWS spot instance bid price (in dollar)",
                    "type": "string",
                    "update": false
                },
                "sshKeypath": {
                    "create": true,
                    "default": "",
                    "description": "SSH Key for Instance",
                    "type": "string",
                    "update": false
                },
                "sshUser": {
                    "create": true,
                    "default": "ubuntu",
                    "description": "Set the name of the ssh user",
                    "type": "string",
                    "update": false
                },
                "subnetId": {
                    "create": true,
                    "default": "",
                    "description": "AWS VPC subnet id",
                    "type": "string",
                    "update": false
                },
                "tags": {
                    "create": true,
                    "default": "",
                    "description": "AWS Tags (e.g. key1,value1,key2,value2)",
                    "type": "string",
                    "update": false
                },
                "useEbsOptimizedInstance": {
                    "create": true,
                    "default": false,
                    "description": "Create an EBS optimized instance",
                    "type": "boolean",
                    "update": false
                },
                "usePrivateAddress": {
                    "create": true,
                    "default": false,
                    "description": "Force the usage of private IP address",
                    "type": "boolean",
                    "update": false
                },
                "userdata": {
                    "create": true,
                    "default": "",
                    "description": "path to file with cloud-init user data",
                    "type": "string",
                    "update": false
                },
                "volumeType": {
                    "create": true,
                    "default": "gp2",
                    "description": "Amazon EBS volume type",
                    "type": "string",
                    "update": false
                },
                "vpcId": {
                    "create": true,
                    "default": "",
                    "description": "AWS VPC id",
                    "type": "string",
                    "update": false
                },
                "zone": {
                    "create": true,
                    "default": "a",
                    "description": "AWS zone for instance (i.e. a,b,c,d,e)",
                    "type": "string",
                    "update": false
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "attachedVolume",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/attachedVolume"
            },
            "pluralName": "attachedVolumes",
            "resourceFields": {
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "authnConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/authnConfig"
            },
            "pluralName": "authnConfigs",
            "resourceFields": {
                "options": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "strategy": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "authzConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/authzConfig"
            },
            "pluralName": "authzConfigs",
            "resourceFields": {
                "mode": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "options": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "azureKubernetesServiceConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/azureKubernetesServiceConfig"
            },
            "pluralName": "azureKubernetesServiceConfigs",
            "resourceFields": {},
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "azureconfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/azureconfig"
            },
            "pluralName": "azureconfigs",
            "resourceFields": {
                "availabilitySet": {
                    "create": true,
                    "default": "docker-machine",
                    "description": "Azure Availability Set to place the virtual machine into",
                    "type": "string",
                    "update": false
                },
                "clientId": {
                    "create": true,
                    "default": "",
                    "description": "Azure Service Principal Account ID (optional, browser auth is used if not specified)",
                    "type": "string",
                    "update": false
                },
                "clientSecret": {
                    "create": true,
                    "default": "",
                    "description": "Azure Service Principal Account password (optional, browser auth is used if not specified)",
                    "type": "string",
                    "update": false
                },
                "customData": {
                    "create": true,
                    "default": "",
                    "description": "Path to file with custom-data",
                    "type": "string",
                    "update": false
                },
                "dns": {
                    "create": true,
                    "default": "",
                    "description": "A unique DNS label for the public IP adddress",
                    "type": "string",
                    "update": false
                },
                "dockerPort": {
                    "create": true,
                    "default": "2376",
                    "description": "Port number for Docker engine",
                    "type": "string",
                    "update": false
                },
                "environment": {
                    "create": true,
                    "default": "AzurePublicCloud",
                    "description": "Azure environment (e.g. AzurePublicCloud, AzureChinaCloud)",
                    "type": "string",
                    "update": false
                },
                "image": {
                    "create": true,
                    "default": "canonical:UbuntuServer:16.04.0-LTS:latest",
                    "description": "Azure virtual machine OS image",
                    "type": "string",
                    "update": false
                },
                "location": {
                    "create": true,
                    "default": "westus",
                    "description": "Azure region to create the virtual machine",
                    "type": "string",
                    "update": false
                },
                "noPublicIp": {
                    "create": true,
                    "default": false,
                    "description": "Do not create a public IP address for the machine",
                    "type": "boolean",
                    "update": false
                },
                "openPort": {
                    "create": true,
                    "description": "Make the specified port number accessible from the Internet",
                    "type": "array[string]",
                    "update": false
                },
                "privateIpAddress": {
                    "create": true,
                    "default": "",
                    "description": "Specify a static private IP address for the machine",
                    "type": "string",
                    "update": false
                },
                "resourceGroup": {
                    "create": true,
                    "default": "docker-machine",
                    "description": "Azure Resource Group name (will be created if missing)",
                    "type": "string",
                    "update": false
                },
                "size": {
                    "create": true,
                    "default": "Standard_A2",
                    "description": "Size for Azure Virtual Machine",
                    "type": "string",
                    "update": false
                },
                "sshUser": {
                    "create": true,
                    "default": "docker-user",
                    "description": "Username for SSH login",
                    "type": "string",
                    "update": false
                },
                "staticPublicIp": {
                    "create": true,
                    "default": false,
                    "description": "Assign a static public IP address to the machine",
                    "type": "boolean",
                    "update": false
                },
                "storageType": {
                    "create": true,
                    "default": "Standard_LRS",
                    "description": "Type of Storage Account to host the OS Disk for the machine",
                    "type": "string",
                    "update": false
                },
                "subnet": {
                    "create": true,
                    "default": "docker-machine",
                    "description": "Azure Subnet Name to be used within the Virtual Network",
                    "type": "string",
                    "update": false
                },
                "subnetPrefix": {
                    "create": true,
                    "default": "192.168.0.0/16",
                    "description": "Private CIDR block to be used for the new subnet, should comply RFC 1918",
                    "type": "string",
                    "update": false
                },
                "subscriptionId": {
                    "create": true,
                    "default": "",
                    "description": "Azure Subscription ID",
                    "type": "string",
                    "update": false
                },
                "usePrivateIp": {
                    "create": true,
                    "default": false,
                    "description": "Use private IP address of the machine to connect",
                    "type": "boolean",
                    "update": false
                },
                "vnet": {
                    "create": true,
                    "default": "docker-machine-vnet",
                    "description": "Azure Virtual Network name to connect the virtual machine (in [resourcegroup:]name format)",
                    "type": "string",
                    "update": false
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "brokerList",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/brokerList"
            },
            "pluralName": "brokerLists",
            "resourceFields": {
                "brokerList": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "branch": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "catalogKind": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "url": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "catalog",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/catalogs",
                "self": "https://192.168.1.16:8443/v3/schemas/catalog"
            },
            "pluralName": "catalogs",
            "resourceActions": {
                "refresh": {}
            },
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "branch": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "catalogKind": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "status": {
                    "create": false,
                    "nullable": true,
                    "type": "catalogStatus",
                    "update": false
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "url": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "catalogStatus",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/catalogStatus"
            },
            "pluralName": "catalogStatuses",
            "resourceFields": {
                "commit": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "lastRefreshTimestamp": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "changePasswordInput",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/changePasswordInput"
            },
            "pluralName": "changePasswordInputs",
            "resourceFields": {
                "currentPassword": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "newPassword": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "apiEndpoint": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "caCert": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "defaultClusterRoleForProjectMembers": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "defaultPodSecurityPolicyTemplateId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "driver": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "internal": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "name": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "cluster",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/clusters",
                "self": "https://192.168.1.16:8443/v3/schemas/cluster"
            },
            "pluralName": "clusters",
            "resourceFields": {
                "allocatable": {
                    "create": false,
                    "nullable": true,
                    "type": "map[string]",
                    "update": false
                },
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "apiEndpoint": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "azureKubernetesServiceConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "azureKubernetesServiceConfig",
                    "update": true
                },
                "caCert": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "capacity": {
                    "create": false,
                    "nullable": true,
                    "type": "map[string]",
                    "update": false
                },
                "componentStatuses": {
                    "create": false,
                    "nullable": true,
                    "type": "array[clusterComponentStatus]",
                    "update": false
                },
                "conditions": {
                    "create": false,
                    "nullable": true,
                    "type": "array[clusterCondition]",
                    "update": false
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "defaultClusterRoleForProjectMembers": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[roleTemplate]",
                    "update": true
                },
                "defaultPodSecurityPolicyTemplateId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[podSecurityPolicyTemplate]",
                    "update": true
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "driver": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "embeddedConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "k8sServerConfig",
                    "update": false
                },
                "googleKubernetesEngineConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "googleKubernetesEngineConfig",
                    "update": true
                },
                "id": {
                    "create": false,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "importedConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "importedConfig",
                    "update": false
                },
                "internal": {
                    "create": false,
                    "nullable": true,
                    "type": "boolean",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "limits": {
                    "create": false,
                    "nullable": true,
                    "type": "map[string]",
                    "update": false
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": true
                },
                "nodes": {
                    "create": true,
                    "nullable": true,
                    "type": "array[machineConfig]",
                    "update": true
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "rancherKubernetesEngineConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "rancherKubernetesEngineConfig",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "requested": {
                    "create": false,
                    "nullable": true,
                    "type": "map[string]",
                    "update": false
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "clusterId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "displayName": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "initialWaitSeconds": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "namespaceId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "repeatIntervalSeconds": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "severity": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "clusterAlert",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/clusteralerts",
                "self": "https://192.168.1.16:8443/v3/schemas/clusterAlert"
            },
            "pluralName": "clusterAlerts",
            "resourceActions": {
                "activate": {},
                "deactivate": {},
                "mute": {},
                "unmute": {}
            },
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "clusterId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[cluster]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "displayName": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "initialWaitSeconds": {
                    "create": true,
                    "default": "180",
                    "min": 0,
                    "nullable": true,
                    "required": true,
                    "type": "int",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "namespaceId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[namespace]",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "recipientList": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "array[recipient]",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "repeatIntervalSeconds": {
                    "create": true,
                    "default": "3600",
                    "min": 0,
                    "nullable": true,
                    "required": true,
                    "type": "int",
                    "update": true
                },
                "severity": {
                    "create": true,
                    "default": "critical",
                    "nullable": true,
                    "options": [
                        "info",
                        "critical",
                        "warning"
                    ],
                    "required": true,
                    "type": "enum",
                    "update": true
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "status": {
                    "create": false,
                    "nullable": true,
                    "type": "alertStatus",
                    "update": false
                },
                "targetNode": {
                    "create": true,
                    "nullable": true,
                    "type": "targetNode",
                    "update": true
                },
                "targetSystemService": {
                    "create": true,
                    "nullable": true,
                    "type": "targetSystemService",
                    "update": true
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "clusterComponentStatus",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/clusterComponentStatus"
            },
            "pluralName": "clusterComponentStatuses",
            "resourceFields": {
                "conditions": {
                    "create": true,
                    "nullable": true,
                    "type": "array[componentCondition]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "clusterCondition",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/clusterCondition"
            },
            "pluralName": "clusterConditions",
            "resourceFields": {
                "lastTransitionTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "lastUpdateTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "message": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "reason": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "status": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "type": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "clusterId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "count": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "eventType": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "message": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "namespaceId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "reason": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "clusterEvent",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/clusterevents",
                "self": "https://192.168.1.16:8443/v3/schemas/clusterEvent"
            },
            "pluralName": "clusterEvents",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "clusterId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[cluster]",
                    "update": true
                },
                "count": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "eventType": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "firstTimestamp": {
                    "create": true,
                    "nullable": true,
                    "type": "date",
                    "update": true
                },
                "involvedObject": {
                    "create": true,
                    "nullable": true,
                    "type": "objectReference",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "lastTimestamp": {
                    "create": true,
                    "nullable": true,
                    "type": "date",
                    "update": true
                },
                "message": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "namespaceId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[namespace]",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "reason": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "source": {
                    "create": true,
                    "nullable": true,
                    "type": "eventSource",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "clusterId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "displayName": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "namespaceId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "outputFlushInterval": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "clusterLogging",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/clusterloggings",
                "self": "https://192.168.1.16:8443/v3/schemas/clusterLogging"
            },
            "pluralName": "clusterLoggings",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "clusterId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[cluster]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "displayName": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "elasticsearchConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "elasticsearchConfig",
                    "update": true
                },
                "embeddedConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "embeddedConfig",
                    "update": true
                },
                "kafkaConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "kafkaConfig",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "namespaceId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[namespace]",
                    "update": false
                },
                "outputFlushInterval": {
                    "create": true,
                    "default": "3",
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "outputTags": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "splunkConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "splunkConfig",
                    "update": true
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "status": {
                    "create": false,
                    "nullable": true,
                    "type": "loggingStatus",
                    "update": false
                },
                "syslogConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "syslogConfig",
                    "update": true
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "clusterId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "namespaceId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "clusterRegistrationToken",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/clusterregistrationtokens",
                "self": "https://192.168.1.16:8443/v3/schemas/clusterRegistrationToken"
            },
            "pluralName": "clusterRegistrationTokens",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "clusterId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[cluster]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "namespaceId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[namespace]",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "status": {
                    "create": false,
                    "nullable": true,
                    "type": "clusterRegistrationTokenStatus",
                    "update": false
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "clusterRegistrationTokenStatus",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/clusterRegistrationTokenStatus"
            },
            "pluralName": "clusterRegistrationTokenStatuses",
            "resourceFields": {
                "command": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "manifestUrl": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "token": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "clusterId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "namespaceId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "roleTemplateId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "userId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "clusterRoleTemplateBinding",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/clusterroletemplatebindings",
                "self": "https://192.168.1.16:8443/v3/schemas/clusterRoleTemplateBinding"
            },
            "pluralName": "clusterRoleTemplateBindings",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "clusterId": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "reference[cluster]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "namespaceId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[namespace]",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "roleTemplateId": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "reference[roleTemplate]",
                    "update": true
                },
                "userId": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "reference[user]",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "componentCondition",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/componentCondition"
            },
            "pluralName": "componentConditions",
            "resourceFields": {
                "error": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "message": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "status": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "type": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "cpuInfo",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/cpuInfo"
            },
            "pluralName": "cpuInfos",
            "resourceFields": {
                "count": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "customConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/customConfig"
            },
            "pluralName": "customConfigs",
            "resourceFields": {
                "address": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "dockerSocket": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "internalAddress": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "sshKey": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "user": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "digitaloceanconfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/digitaloceanconfig"
            },
            "pluralName": "digitaloceanconfigs",
            "resourceFields": {
                "accessToken": {
                    "create": true,
                    "default": "",
                    "description": "Digital Ocean access token",
                    "type": "string",
                    "update": false
                },
                "backups": {
                    "create": true,
                    "default": false,
                    "description": "enable backups for droplet",
                    "type": "boolean",
                    "update": false
                },
                "image": {
                    "create": true,
                    "default": "ubuntu-16-04-x64",
                    "description": "Digital Ocean Image",
                    "type": "string",
                    "update": false
                },
                "ipv6": {
                    "create": true,
                    "default": false,
                    "description": "enable ipv6 for droplet",
                    "type": "boolean",
                    "update": false
                },
                "privateNetworking": {
                    "create": true,
                    "default": false,
                    "description": "enable private networking for droplet",
                    "type": "boolean",
                    "update": false
                },
                "region": {
                    "create": true,
                    "default": "nyc3",
                    "description": "Digital Ocean region",
                    "type": "string",
                    "update": false
                },
                "size": {
                    "create": true,
                    "default": "512mb",
                    "description": "Digital Ocean size",
                    "type": "string",
                    "update": false
                },
                "sshKeyFingerprint": {
                    "create": true,
                    "default": "",
                    "description": "SSH key fingerprint",
                    "type": "string",
                    "update": false
                },
                "sshKeyPath": {
                    "create": true,
                    "default": "",
                    "description": "SSH private key path ",
                    "type": "string",
                    "update": false
                },
                "sshPort": {
                    "create": true,
                    "default": "22",
                    "description": "SSH port",
                    "type": "string",
                    "update": false
                },
                "sshUser": {
                    "create": true,
                    "default": "root",
                    "description": "SSH username",
                    "type": "string",
                    "update": false
                },
                "tags": {
                    "create": true,
                    "default": "",
                    "description": "comma-separated list of tags to apply to the Droplet",
                    "type": "string",
                    "update": false
                },
                "userdata": {
                    "create": true,
                    "default": "",
                    "description": "path to file with cloud-init user-data",
                    "type": "string",
                    "update": false
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "embed": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "embedType": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "pluralName": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "dynamicSchema",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/dynamicschemas",
                "self": "https://192.168.1.16:8443/v3/schemas/dynamicSchema"
            },
            "pluralName": "dynamicSchemas",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "collectionActions": {
                    "create": true,
                    "nullable": true,
                    "type": "map[action]",
                    "update": true
                },
                "collectionFields": {
                    "create": true,
                    "nullable": true,
                    "type": "map[field]",
                    "update": true
                },
                "collectionFilters": {
                    "create": true,
                    "nullable": true,
                    "type": "map[filter]",
                    "update": true
                },
                "collectionMethods": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "embed": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "embedType": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "includeableLinks": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "pluralName": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "resourceActions": {
                    "create": true,
                    "nullable": true,
                    "type": "map[action]",
                    "update": true
                },
                "resourceFields": {
                    "create": true,
                    "nullable": true,
                    "type": "map[field]",
                    "update": true
                },
                "resourceMethods": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "status": {
                    "create": false,
                    "nullable": true,
                    "type": "dynamicSchemaStatus",
                    "update": false
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "dynamicSchemaStatus",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/dynamicSchemaStatus"
            },
            "pluralName": "dynamicSchemaStatuses",
            "resourceFields": {
                "fake": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "elasticsearchConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/elasticsearchConfig"
            },
            "pluralName": "elasticsearchConfigs",
            "resourceFields": {
                "authPassword": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "authUsername": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "dateFormat": {
                    "create": true,
                    "default": "YYYY-MM-DD",
                    "nullable": true,
                    "options": [
                        "YYYY-MM-DD",
                        "YYYY-MM",
                        "YYYY"
                    ],
                    "required": true,
                    "type": "enum",
                    "update": true
                },
                "endpoint": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "indexPrefix": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "embeddedConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/embeddedConfig"
            },
            "pluralName": "embeddedConfigs",
            "resourceFields": {
                "dateFormat": {
                    "create": true,
                    "default": "YYYY-MM-DD",
                    "nullable": true,
                    "options": [
                        "YYYY-MM-DD",
                        "YYYY-MM",
                        "YYYY"
                    ],
                    "required": true,
                    "type": "enum",
                    "update": true
                },
                "indexPrefix": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "etcdService",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/etcdService"
            },
            "pluralName": "etcdServices",
            "resourceFields": {
                "extraArgs": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "image": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "eventSource",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/eventSource"
            },
            "pluralName": "eventSources",
            "resourceFields": {
                "component": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "host": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "field",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/field"
            },
            "pluralName": "fields",
            "resourceFields": {
                "create": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "default": {
                    "create": true,
                    "nullable": true,
                    "type": "values",
                    "update": true
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "invalidChars": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "max": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "maxLength": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "min": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "minLength": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "nullable": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "options": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "required": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "type": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "unique": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "update": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "validChars": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "file",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/file"
            },
            "pluralName": "files",
            "resourceFields": {
                "contents": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "filter",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/filter"
            },
            "pluralName": "filters",
            "resourceFields": {
                "modifiers": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "fsGroupStrategyOptions",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/fsGroupStrategyOptions"
            },
            "pluralName": "fsGroupStrategyOptionses",
            "resourceFields": {
                "ranges": {
                    "create": true,
                    "nullable": true,
                    "type": "array[idRange]",
                    "update": true
                },
                "rule": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "githubCredential",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/githubCredential"
            },
            "pluralName": "githubCredentials",
            "resourceFields": {
                "code": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "builtin": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "name": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "globalRole",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/globalroles",
                "self": "https://192.168.1.16:8443/v3/schemas/globalRole"
            },
            "pluralName": "globalRoles",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "builtin": {
                    "create": false,
                    "nullable": true,
                    "type": "boolean",
                    "update": false
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "id": {
                    "create": false,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "rules": {
                    "create": true,
                    "nullable": true,
                    "type": "array[policyRule]",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "globalRoleId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "userId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "globalRoleBinding",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/globalrolebindings",
                "self": "https://192.168.1.16:8443/v3/schemas/globalRoleBinding"
            },
            "pluralName": "globalRoleBindings",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "globalRoleId": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "reference[globalRole]",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "userId": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "reference[user]",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "googleKubernetesEngineConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/googleKubernetesEngineConfig"
            },
            "pluralName": "googleKubernetesEngineConfigs",
            "resourceFields": {
                "clusterIpv4Cidr": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "credential": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "diskSizeGb": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "enableAlphaFeature": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "horizontalPodAutoscaling": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "httpLoadBalancing": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "imageType": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "kubernetesDashboard": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "legacyAbac": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "locations": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "machineType": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "masterVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "network": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "networkPolicyConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "nodeCount": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "int",
                    "update": true
                },
                "nodeVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "projectId": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "subNetwork": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "zone": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "name": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "group",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/groups",
                "self": "https://192.168.1.16:8443/v3/schemas/group"
            },
            "pluralName": "groups",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "id": {
                    "create": false,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "groupId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "principalId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "groupMember",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/groupmembers",
                "self": "https://192.168.1.16:8443/v3/schemas/groupMember"
            },
            "pluralName": "groupMembers",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "groupId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[group]",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "principalId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[Principal]",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "hostPortRange",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/hostPortRange"
            },
            "pluralName": "hostPortRanges",
            "resourceFields": {
                "max": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "min": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "idRange",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/idRange"
            },
            "pluralName": "idRanges",
            "resourceFields": {
                "max": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "min": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "importedConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/importedConfig"
            },
            "pluralName": "importedConfigs",
            "resourceFields": {
                "kubeConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "k8sServerConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/k8sServerConfig"
            },
            "pluralName": "k8sServerConfigs",
            "resourceFields": {
                "admissionControllers": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "serviceNetCidr": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "kafkaConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/kafkaConfig"
            },
            "pluralName": "kafkaConfigs",
            "resourceFields": {
                "broker": {
                    "create": true,
                    "nullable": true,
                    "type": "brokerList",
                    "update": true
                },
                "maxSendRetries": {
                    "create": true,
                    "default": "3",
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "topic": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "zookeeper": {
                    "create": true,
                    "nullable": true,
                    "type": "zookeeper",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "kubeAPIService",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/kubeAPIService"
            },
            "pluralName": "kubeAPIServices",
            "resourceFields": {
                "extraArgs": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "image": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "podSecurityPolicy": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "serviceClusterIpRange": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "kubeControllerService",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/kubeControllerService"
            },
            "pluralName": "kubeControllerServices",
            "resourceFields": {
                "clusterCidr": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "extraArgs": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "image": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "serviceClusterIpRange": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "kubeletService",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/kubeletService"
            },
            "pluralName": "kubeletServices",
            "resourceFields": {
                "clusterDnsServer": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "clusterDomain": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "extraArgs": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "image": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "infraContainerImage": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "kubeproxyService",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/kubeproxyService"
            },
            "pluralName": "kubeproxyServices",
            "resourceFields": {
                "extraArgs": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "image": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "kubernetesInfo",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/kubernetesInfo"
            },
            "pluralName": "kubernetesInfos",
            "resourceFields": {
                "kubeProxyVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "kubeletVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "algorithm": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "caCerts": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "cert": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "certFingerprint": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "cn": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "enabled": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "expiresAt": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "issuedAt": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "issuer": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "key": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "keySize": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "mode": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "name": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "serialNumber": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "version": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "listenConfig",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/listenconfigs",
                "self": "https://192.168.1.16:8443/v3/schemas/listenConfig"
            },
            "pluralName": "listenConfigs",
            "resourceFields": {
                "algorithm": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "caCerts": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "cert": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "certFingerprint": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "cn": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "domains": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "enabled": {
                    "create": true,
                    "default": "true",
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "expiresAt": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "id": {
                    "create": false,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "issuedAt": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "issuer": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "key": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true,
                    "writeOnly": true
                },
                "keySize": {
                    "create": false,
                    "nullable": true,
                    "type": "int",
                    "update": false
                },
                "knownIps": {
                    "create": false,
                    "nullable": true,
                    "type": "array[string]",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "mode": {
                    "create": true,
                    "nullable": true,
                    "options": [
                        "https",
                        "http",
                        "acme"
                    ],
                    "type": "enum",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "serialNumber": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "subjectAlternativeNames": {
                    "create": false,
                    "nullable": true,
                    "type": "array[string]",
                    "update": false
                },
                "tos": {
                    "create": true,
                    "default": "auto",
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "version": {
                    "create": false,
                    "nullable": true,
                    "type": "int",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "localCredential",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/localCredential"
            },
            "pluralName": "localCredentials",
            "resourceFields": {
                "password": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "username": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "loggingCondition",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/loggingCondition"
            },
            "pluralName": "loggingConditions",
            "resourceFields": {
                "lastTransitionTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "lastUpdateTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "message": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "reason": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "status": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "type": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "loggingStatus",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/loggingStatus"
            },
            "pluralName": "loggingStatuses",
            "resourceFields": {
                "conditions": {
                    "create": false,
                    "nullable": true,
                    "type": "array[loggingCondition]",
                    "update": false
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "loginInput",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/loginInput"
            },
            "pluralName": "loginInputs",
            "resourceFields": {
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "githubCredential": {
                    "create": true,
                    "nullable": true,
                    "type": "githubCredential",
                    "update": true
                },
                "localCredential": {
                    "create": true,
                    "nullable": true,
                    "type": "localCredential",
                    "update": true
                },
                "responseType": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "ttl": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "clusterId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "hostname": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "ipAddress": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "machineTemplateId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "name": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "namespaceId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "nodeName": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "podCidr": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "providerId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "sshUser": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "unschedulable": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "useInternalIpAddress": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "machine",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/machines",
                "self": "https://192.168.1.16:8443/v3/schemas/machine"
            },
            "pluralName": "machines",
            "resourceFields": {
                "allocatable": {
                    "create": false,
                    "nullable": true,
                    "type": "map[string]",
                    "update": false
                },
                "amazonec2Config": {
                    "create": true,
                    "nullable": true,
                    "type": "amazonec2config",
                    "update": false
                },
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "azureConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "azureconfig",
                    "update": false
                },
                "capacity": {
                    "create": false,
                    "nullable": true,
                    "type": "map[string]",
                    "update": false
                },
                "clusterId": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "reference[cluster]",
                    "update": false
                },
                "conditions": {
                    "create": false,
                    "nullable": true,
                    "type": "array[machineCondition]",
                    "update": false
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "customConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "customConfig",
                    "update": true
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "digitaloceanConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "digitaloceanconfig",
                    "update": false
                },
                "hostname": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "id": {
                    "create": false,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "info": {
                    "create": false,
                    "nullable": true,
                    "type": "nodeInfo",
                    "update": false
                },
                "ipAddress": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "limits": {
                    "create": false,
                    "nullable": true,
                    "type": "map[string]",
                    "update": false
                },
                "machineTemplateId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[machineTemplate]",
                    "update": false
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "namespaceId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[namespace]",
                    "update": false
                },
                "nodeAnnotations": {
                    "create": false,
                    "nullable": true,
                    "type": "map[string]",
                    "update": false
                },
                "nodeLabels": {
                    "create": false,
                    "nullable": true,
                    "type": "map[string]",
                    "update": false
                },
                "nodeName": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "nodeTaints": {
                    "create": false,
                    "nullable": true,
                    "type": "array[taint]",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "podCidr": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "providerId": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "requested": {
                    "create": false,
                    "nullable": true,
                    "type": "map[string]",
                    "update": false
                },
                "requestedHostname": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "role": {
                    "create": true,
                    "nullable": true,
                    "options": [
                        "etcd",
                        "worker",
                        "controlplane"
                    ],
                    "type": "array[enum]",
                    "update": false
                },
                "sshUser": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "taints": {
                    "create": false,
                    "nullable": true,
                    "type": "array[taint]",
                    "update": true
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "unschedulable": {
                    "create": false,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "useInternalIpAddress": {
                    "create": true,
                    "default": "true",
                    "nullable": true,
                    "type": "boolean",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "vmwarevsphereConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "vmwarevsphereconfig",
                    "update": false
                },
                "volumesAttached": {
                    "create": false,
                    "nullable": true,
                    "type": "map[attachedVolume]",
                    "update": false
                },
                "volumesInUse": {
                    "create": false,
                    "nullable": true,
                    "type": "array[string]",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "machineCondition",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/machineCondition"
            },
            "pluralName": "machineConditions",
            "resourceFields": {
                "lastTransitionTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "lastUpdateTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "message": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "reason": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "status": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "type": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "machineConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/machineConfig"
            },
            "pluralName": "machineConfigs",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "customConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "customConfig",
                    "update": true
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "displayName": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "machineTemplateId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[machineTemplate]",
                    "update": false
                },
                "nodeSpec": {
                    "create": true,
                    "nullable": true,
                    "type": "nodeSpec",
                    "update": true
                },
                "requestedHostname": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "role": {
                    "create": true,
                    "nullable": true,
                    "options": [
                        "etcd",
                        "worker",
                        "controlplane"
                    ],
                    "type": "array[enum]",
                    "update": false
                },
                "useInternalIpAddress": {
                    "create": true,
                    "default": "true",
                    "nullable": true,
                    "type": "boolean",
                    "update": false
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "active": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "builtin": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "checksum": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "externalId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "name": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uiUrl": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "url": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "machineDriver",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/machinedrivers",
                "self": "https://192.168.1.16:8443/v3/schemas/machineDriver"
            },
            "pluralName": "machineDrivers",
            "resourceActions": {
                "activate": {
                    "output": "machineDriver"
                },
                "deactivate": {
                    "output": "machineDriver"
                }
            },
            "resourceFields": {
                "active": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "builtin": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "checksum": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "externalId": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "id": {
                    "create": false,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "status": {
                    "create": false,
                    "nullable": true,
                    "type": "machineDriverStatus",
                    "update": false
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "uiUrl": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "url": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "machineDriverCondition",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/machineDriverCondition"
            },
            "pluralName": "machineDriverConditions",
            "resourceFields": {
                "lastTransitionTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "lastUpdateTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "message": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "reason": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "status": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "type": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "machineDriverStatus",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/machineDriverStatus"
            },
            "pluralName": "machineDriverStatuses",
            "resourceFields": {
                "conditions": {
                    "create": false,
                    "nullable": true,
                    "type": "array[machineDriverCondition]",
                    "update": false
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "authCertificateAuthority": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "authKey": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "dockerVersion": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "driver": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "engineInstallURL": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "engineStorageDriver": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "name": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "machineTemplate",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/machinetemplates",
                "self": "https://192.168.1.16:8443/v3/schemas/machineTemplate"
            },
            "pluralName": "machineTemplates",
            "resourceFields": {
                "amazonec2Config": {
                    "create": true,
                    "nullable": true,
                    "type": "amazonec2config",
                    "update": false
                },
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "authCertificateAuthority": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "authKey": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "azureConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "azureconfig",
                    "update": false
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "digitaloceanConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "digitaloceanconfig",
                    "update": false
                },
                "dockerVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "driver": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "engineEnv": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "engineInsecureRegistry": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "engineInstallURL": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "engineLabel": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "engineOpt": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "engineRegistryMirror": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "engineStorageDriver": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "id": {
                    "create": false,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "status": {
                    "create": false,
                    "nullable": true,
                    "type": "machineTemplateStatus",
                    "update": false
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "vmwarevsphereConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "vmwarevsphereconfig",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "machineTemplateCondition",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/machineTemplateCondition"
            },
            "pluralName": "machineTemplateConditions",
            "resourceFields": {
                "lastTransitionTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "lastUpdateTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "reason": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "status": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "type": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "machineTemplateStatus",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/machineTemplateStatus"
            },
            "pluralName": "machineTemplateStatuses",
            "resourceFields": {
                "conditions": {
                    "create": false,
                    "nullable": true,
                    "type": "array[machineTemplateCondition]",
                    "update": false
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "memoryInfo",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/memoryInfo"
            },
            "pluralName": "memoryInfos",
            "resourceFields": {
                "memTotalKiB": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "networkConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/networkConfig"
            },
            "pluralName": "networkConfigs",
            "resourceFields": {
                "options": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "plugin": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "hostname": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "ipAddress": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "podCidr": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "providerId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "unschedulable": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "node",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/nodes",
                "self": "https://192.168.1.16:8443/v3/schemas/node"
            },
            "pluralName": "nodes",
            "resourceFields": {
                "allocatable": {
                    "create": false,
                    "nullable": true,
                    "type": "map[string]",
                    "update": false
                },
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "capacity": {
                    "create": false,
                    "nullable": true,
                    "type": "map[string]",
                    "update": false
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "hostname": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "info": {
                    "create": false,
                    "nullable": true,
                    "type": "nodeInfo",
                    "update": false
                },
                "ipAddress": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "nodeConditions": {
                    "create": false,
                    "nullable": true,
                    "type": "array[nodeCondition]",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "podCidr": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "providerId": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "taints": {
                    "create": false,
                    "nullable": true,
                    "type": "array[taint]",
                    "update": true
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "unschedulable": {
                    "create": false,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "volumesAttached": {
                    "create": false,
                    "nullable": true,
                    "type": "map[attachedVolume]",
                    "update": false
                },
                "volumesInUse": {
                    "create": false,
                    "nullable": true,
                    "type": "array[string]",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "nodeCondition",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/nodeCondition"
            },
            "pluralName": "nodeConditions",
            "resourceFields": {
                "lastHeartbeatTime": {
                    "create": true,
                    "nullable": true,
                    "type": "date",
                    "update": true
                },
                "lastTransitionTime": {
                    "create": true,
                    "nullable": true,
                    "type": "date",
                    "update": true
                },
                "message": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "reason": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "status": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "type": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "nodeInfo",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/nodeInfo"
            },
            "pluralName": "nodeInfos",
            "resourceFields": {
                "cpu": {
                    "create": true,
                    "nullable": true,
                    "type": "cpuInfo",
                    "update": true
                },
                "kubernetes": {
                    "create": true,
                    "nullable": true,
                    "type": "kubernetesInfo",
                    "update": true
                },
                "memory": {
                    "create": true,
                    "nullable": true,
                    "type": "memoryInfo",
                    "update": true
                },
                "os": {
                    "create": true,
                    "nullable": true,
                    "type": "osInfo",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "nodeSpec",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/nodeSpec"
            },
            "pluralName": "nodeSpecs",
            "resourceFields": {
                "podCidr": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "providerId": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "taints": {
                    "create": false,
                    "nullable": true,
                    "type": "array[taint]",
                    "update": true
                },
                "unschedulable": {
                    "create": false,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionActions": {
                "send": {}
            },
            "collectionFilters": {
                "clusterId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "displayName": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "namespaceId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "notifier",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/notifiers",
                "self": "https://192.168.1.16:8443/v3/schemas/notifier"
            },
            "pluralName": "notifiers",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "clusterId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[cluster]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "displayName": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "namespaceId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[namespace]",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "pagerdutyConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "pagerdutyConfig",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "slackConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "slackConfig",
                    "update": true
                },
                "smtpConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "smtpConfig",
                    "update": true
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "status": {
                    "create": false,
                    "nullable": true,
                    "type": "notifierStatus",
                    "update": false
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "webhookConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "webhookConfig",
                    "update": true
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "notifierStatus",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/notifierStatus"
            },
            "pluralName": "notifierStatuses",
            "resourceFields": {},
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "objectReference",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/objectReference"
            },
            "pluralName": "objectReferences",
            "resourceFields": {
                "apiVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "fieldPath": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "kind": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "namespace": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "resourceVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "uid": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "osInfo",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/osInfo"
            },
            "pluralName": "osInfos",
            "resourceFields": {
                "dockerVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "kernelVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "operatingSystem": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "ownerReference",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/ownerReference"
            },
            "pluralName": "ownerReferences",
            "resourceFields": {
                "apiVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "blockOwnerDeletion": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "controller": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "kind": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "uid": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "pagerdutyConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/pagerdutyConfig"
            },
            "pluralName": "pagerdutyConfigs",
            "resourceFields": {
                "serviceKey": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "allowPrivilegeEscalation": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "defaultAllowPrivilegeEscalation": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "hostIPC": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "hostNetwork": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "hostPID": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "privileged": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "readOnlyRootFilesystem": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "podSecurityPolicyTemplate",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/podsecuritypolicytemplates",
                "self": "https://192.168.1.16:8443/v3/schemas/podSecurityPolicyTemplate"
            },
            "pluralName": "podSecurityPolicyTemplates",
            "resourceFields": {
                "allowPrivilegeEscalation": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "allowedCapabilities": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "allowedHostPaths": {
                    "create": true,
                    "nullable": true,
                    "type": "array[allowedHostPath]",
                    "update": true
                },
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "defaultAddCapabilities": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "defaultAllowPrivilegeEscalation": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "fsGroup": {
                    "create": true,
                    "nullable": true,
                    "type": "fsGroupStrategyOptions",
                    "update": true
                },
                "hostIPC": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "hostNetwork": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "hostPID": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "hostPorts": {
                    "create": true,
                    "nullable": true,
                    "type": "array[hostPortRange]",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "privileged": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "readOnlyRootFilesystem": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "requiredDropCapabilities": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "runAsUser": {
                    "create": true,
                    "nullable": true,
                    "type": "runAsUserStrategyOptions",
                    "update": true
                },
                "seLinux": {
                    "create": true,
                    "nullable": true,
                    "type": "seLinuxStrategyOptions",
                    "update": true
                },
                "supplementalGroups": {
                    "create": true,
                    "nullable": true,
                    "type": "supplementalGroupsStrategyOptions",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "volumes": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "policyRule",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/policyRule"
            },
            "pluralName": "policyRules",
            "resourceFields": {
                "apiGroups": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "nonResourceURLs": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "resourceNames": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "resources": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "verbs": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "namespaceId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "value": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "preference",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/preferences",
                "self": "https://192.168.1.16:8443/v3/schemas/preference"
            },
            "pluralName": "preferences",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "namespaceId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[namespace]",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "value": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "displayName": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "loginName": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "me": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "memberOf": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "profilePicture": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "profileURL": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "principal",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/principals",
                "self": "https://192.168.1.16:8443/v3/schemas/principal"
            },
            "pluralName": "principals",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "displayName": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "extraInfo": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "loginName": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "me": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "memberOf": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "profilePicture": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "profileURL": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "clusterId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "name": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "namespaceId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "podSecurityPolicyTemplateId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "project",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/projects",
                "self": "https://192.168.1.16:8443/v3/schemas/project"
            },
            "pluralName": "projects",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "clusterId": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "reference[cluster]",
                    "update": true
                },
                "conditions": {
                    "create": false,
                    "nullable": true,
                    "type": "array[projectCondition]",
                    "update": false
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "id": {
                    "create": false,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "namespaceId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[namespace]",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "podSecurityPolicyTemplateId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[podSecurityPolicyTemplate]",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "displayName": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "initialWaitSeconds": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "namespaceId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "projectId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "repeatIntervalSeconds": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "severity": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "projectAlert",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/projectalerts",
                "self": "https://192.168.1.16:8443/v3/schemas/projectAlert"
            },
            "pluralName": "projectAlerts",
            "resourceActions": {
                "activate": {},
                "deactivate": {},
                "mute": {},
                "unmute": {}
            },
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "displayName": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "initialWaitSeconds": {
                    "create": true,
                    "default": "180",
                    "min": 0,
                    "nullable": true,
                    "required": true,
                    "type": "int",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "namespaceId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[namespace]",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "projectId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[project]",
                    "update": true
                },
                "recipientList": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "array[recipient]",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "repeatIntervalSeconds": {
                    "create": true,
                    "default": "3600",
                    "min": 0,
                    "nullable": true,
                    "required": true,
                    "type": "int",
                    "update": true
                },
                "severity": {
                    "create": true,
                    "default": "critical",
                    "nullable": true,
                    "options": [
                        "info",
                        "critical",
                        "warning"
                    ],
                    "required": true,
                    "type": "enum",
                    "update": true
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "status": {
                    "create": false,
                    "nullable": true,
                    "type": "alertStatus",
                    "update": false
                },
                "targetPod": {
                    "create": true,
                    "nullable": true,
                    "type": "targetPod",
                    "update": true
                },
                "targetWorkload": {
                    "create": true,
                    "nullable": true,
                    "type": "targetWorkload",
                    "update": true
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "projectCondition",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/projectCondition"
            },
            "pluralName": "projectConditions",
            "resourceFields": {
                "lastTransitionTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "lastUpdateTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "message": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "reason": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "status": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "type": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "displayName": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "namespaceId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "outputFlushInterval": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "projectId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "projectLogging",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/projectloggings",
                "self": "https://192.168.1.16:8443/v3/schemas/projectLogging"
            },
            "pluralName": "projectLoggings",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "displayName": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "elasticsearchConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "elasticsearchConfig",
                    "update": true
                },
                "kafkaConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "kafkaConfig",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "namespaceId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[namespace]",
                    "update": false
                },
                "outputFlushInterval": {
                    "create": true,
                    "default": "3",
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "outputTags": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "projectId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[project]",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "splunkConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "splunkConfig",
                    "update": true
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "status": {
                    "create": false,
                    "nullable": true,
                    "type": "loggingStatus",
                    "update": false
                },
                "syslogConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "syslogConfig",
                    "update": true
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "namespaceId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "projectId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "roleTemplateId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "userId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "projectRoleTemplateBinding",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/projectroletemplatebindings",
                "self": "https://192.168.1.16:8443/v3/schemas/projectRoleTemplateBinding"
            },
            "pluralName": "projectRoleTemplateBindings",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "namespaceId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[namespace]",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "projectId": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "reference[project]",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "roleTemplateId": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "reference[roleTemplate]",
                    "update": true
                },
                "userId": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "reference[user]",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "question",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/question"
            },
            "pluralName": "questions",
            "resourceFields": {
                "default": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "group": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "invalidChars": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "label": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "max": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "maxLength": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "min": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "minLength": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "options": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "required": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "type": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "validChars": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "variable": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "rancherKubernetesEngineConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/rancherKubernetesEngineConfig"
            },
            "pluralName": "rancherKubernetesEngineConfigs",
            "resourceFields": {
                "addons": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "authentication": {
                    "create": true,
                    "nullable": true,
                    "type": "authnConfig",
                    "update": true
                },
                "authorization": {
                    "create": true,
                    "nullable": true,
                    "type": "authzConfig",
                    "update": true
                },
                "ignoreDockerVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "network": {
                    "create": true,
                    "nullable": true,
                    "type": "networkConfig",
                    "update": true
                },
                "nodes": {
                    "create": true,
                    "nullable": true,
                    "type": "array[rkeConfigNode]",
                    "update": true
                },
                "services": {
                    "create": true,
                    "nullable": true,
                    "type": "rkeConfigServices",
                    "update": true
                },
                "sshKeyPath": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "systemImages": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "version": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "recipient",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/recipient"
            },
            "pluralName": "recipients",
            "resourceFields": {
                "customPagerdutyConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "pagerdutyConfig",
                    "update": true
                },
                "customWebhookConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "webhookConfig",
                    "update": true
                },
                "notifierId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[notifier]",
                    "update": true
                },
                "recipient": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "releaseInfo",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/releaseInfo"
            },
            "pluralName": "releaseInfos",
            "resourceFields": {
                "createTimestamp": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "modifiedAt": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "templateVersionId": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "version": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "rkeConfigNode",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/rkeConfigNode"
            },
            "pluralName": "rkeConfigNodes",
            "resourceFields": {
                "address": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "dockerSocket": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "hostnameOverride": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "internalAddress": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "machineId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[machine]",
                    "update": true
                },
                "role": {
                    "create": true,
                    "nullable": true,
                    "options": [
                        "etcd",
                        "worker",
                        "controlplane"
                    ],
                    "type": "array[enum]",
                    "update": true
                },
                "sshKey": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "sshKeyPath": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "user": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "rkeConfigServices",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/rkeConfigServices"
            },
            "pluralName": "rkeConfigServiceses",
            "resourceFields": {
                "etcd": {
                    "create": true,
                    "nullable": true,
                    "type": "etcdService",
                    "update": true
                },
                "kubeApi": {
                    "create": true,
                    "nullable": true,
                    "type": "kubeAPIService",
                    "update": true
                },
                "kubeController": {
                    "create": true,
                    "nullable": true,
                    "type": "kubeControllerService",
                    "update": true
                },
                "kubelet": {
                    "create": true,
                    "nullable": true,
                    "type": "kubeletService",
                    "update": true
                },
                "kubeproxy": {
                    "create": true,
                    "nullable": true,
                    "type": "kubeproxyService",
                    "update": true
                },
                "scheduler": {
                    "create": true,
                    "nullable": true,
                    "type": "schedulerService",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "builtin": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "context": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "external": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "hidden": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "name": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "roleTemplate",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/roletemplates",
                "self": "https://192.168.1.16:8443/v3/schemas/roleTemplate"
            },
            "pluralName": "roleTemplates",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "builtin": {
                    "create": false,
                    "nullable": true,
                    "type": "boolean",
                    "update": false
                },
                "context": {
                    "create": true,
                    "nullable": true,
                    "options": [
                        "project",
                        "cluster"
                    ],
                    "type": "string",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "external": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "hidden": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "id": {
                    "create": false,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "roleTemplateIds": {
                    "create": true,
                    "nullable": true,
                    "type": "array[reference[roleTemplate]]",
                    "update": true
                },
                "rules": {
                    "create": true,
                    "nullable": true,
                    "type": "array[policyRule]",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "runAsUserStrategyOptions",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/runAsUserStrategyOptions"
            },
            "pluralName": "runAsUserStrategyOptionses",
            "resourceFields": {
                "ranges": {
                    "create": true,
                    "nullable": true,
                    "type": "array[idRange]",
                    "update": true
                },
                "rule": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "schedulerService",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/schedulerService"
            },
            "pluralName": "schedulerServices",
            "resourceFields": {
                "extraArgs": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "image": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "seLinuxOptions",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/seLinuxOptions"
            },
            "pluralName": "seLinuxOptionses",
            "resourceFields": {
                "level": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "role": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "type": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "user": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "seLinuxStrategyOptions",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/seLinuxStrategyOptions"
            },
            "pluralName": "seLinuxStrategyOptionses",
            "resourceFields": {
                "rule": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "seLinuxOptions": {
                    "create": true,
                    "nullable": true,
                    "type": "seLinuxOptions",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "setPasswordInput",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/setPasswordInput"
            },
            "pluralName": "setPasswordInputs",
            "resourceFields": {
                "newPassword": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "customized": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "default": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "value": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "setting",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/settings",
                "self": "https://192.168.1.16:8443/v3/schemas/setting"
            },
            "pluralName": "settings",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "customized": {
                    "create": false,
                    "nullable": true,
                    "type": "boolean",
                    "update": false
                },
                "default": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "value": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "slackConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/slackConfig"
            },
            "pluralName": "slackConfigs",
            "resourceFields": {
                "url": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "smtpConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/smtpConfig"
            },
            "pluralName": "smtpConfigs",
            "resourceFields": {
                "defaultRecipient": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "host": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "dnsLabel",
                    "update": true
                },
                "password": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "masked",
                    "update": true
                },
                "port": {
                    "create": true,
                    "max": 65535,
                    "min": 1,
                    "nullable": true,
                    "required": true,
                    "type": "int",
                    "update": true
                },
                "tls": {
                    "create": true,
                    "default": "true",
                    "nullable": true,
                    "required": true,
                    "type": "boolean",
                    "update": true
                },
                "username": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "splunkConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/splunkConfig"
            },
            "pluralName": "splunkConfigs",
            "resourceFields": {
                "endpoint": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "source": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "token": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "externalId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "installNamespace": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "namespaceId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "projectId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "prune": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "user": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "stack",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/stacks",
                "self": "https://192.168.1.16:8443/v3/schemas/stack"
            },
            "pluralName": "stacks",
            "resourceActions": {
                "rollback": {
                    "input": "revision"
                },
                "upgrade": {
                    "input": "templateVersionId"
                }
            },
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "answers": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "externalId": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "groups": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "installNamespace": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "namespaceId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[namespace]",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "projectId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[project]",
                    "update": true
                },
                "prune": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "status": {
                    "create": false,
                    "nullable": true,
                    "type": "stackStatus",
                    "update": false
                },
                "tag": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "templates": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "user": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "stackStatus",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/stackStatus"
            },
            "pluralName": "stackStatuses",
            "resourceFields": {
                "releases": {
                    "create": false,
                    "nullable": true,
                    "type": "array[releaseInfo]",
                    "update": false
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "supplementalGroupsStrategyOptions",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/supplementalGroupsStrategyOptions"
            },
            "pluralName": "supplementalGroupsStrategyOptionses",
            "resourceFields": {
                "ranges": {
                    "create": true,
                    "nullable": true,
                    "type": "array[idRange]",
                    "update": true
                },
                "rule": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "syslogConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/syslogConfig"
            },
            "pluralName": "syslogConfigs",
            "resourceFields": {
                "endpoint": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "program": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "severity": {
                    "create": true,
                    "default": "notice",
                    "nullable": true,
                    "options": [
                        "emerg",
                        "alert",
                        "crit",
                        "err",
                        "warning",
                        "notice",
                        "info",
                        "debug"
                    ],
                    "type": "enum",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "taint",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/taint"
            },
            "pluralName": "taints",
            "resourceFields": {
                "effect": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "key": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "timeAdded": {
                    "create": true,
                    "nullable": true,
                    "type": "date",
                    "update": true
                },
                "value": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "targetNode",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/targetNode"
            },
            "pluralName": "targetNodes",
            "resourceFields": {
                "cpuThreshold": {
                    "create": true,
                    "min": 1,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "diskThreshold": {
                    "create": true,
                    "max": 100,
                    "min": 1,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "id": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "isReady": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "memThreshold": {
                    "create": true,
                    "max": 100,
                    "min": 1,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "selector": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "targetPod",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/targetPod"
            },
            "pluralName": "targetPods",
            "resourceFields": {
                "id": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "isRunning": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "isScheduled": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "restartTimes": {
                    "create": true,
                    "min": 1,
                    "nullable": true,
                    "type": "int",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "targetSystemService",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/targetSystemService"
            },
            "pluralName": "targetSystemServices",
            "resourceFields": {
                "type": {
                    "create": true,
                    "default": "scheduler",
                    "nullable": true,
                    "options": [
                        "dns",
                        "etcd",
                        "controller",
                        "manager",
                        "network",
                        "scheduler"
                    ],
                    "required": true,
                    "type": "enum",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "targetWorkload",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/targetWorkload"
            },
            "pluralName": "targetWorkloads",
            "resourceFields": {
                "id": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "selector": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "unavailablePercentage": {
                    "create": true,
                    "max": 100,
                    "min": 1,
                    "nullable": true,
                    "required": true,
                    "type": "int",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "catalogId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "category": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "defaultTemplateVersionId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "defaultVersion": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "folderName": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "icon": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "iconFilename": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "isSystem": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "license": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "maintainer": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "path": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "projectURL": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "readme": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "templateBase": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "upgradeFrom": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "template",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/templates",
                "self": "https://192.168.1.16:8443/v3/schemas/template"
            },
            "pluralName": "templates",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "catalogId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[catalog]",
                    "update": true
                },
                "categories": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "category": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "defaultTemplateVersionId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[templateVersion]",
                    "update": true
                },
                "defaultVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "folderName": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "icon": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "iconFilename": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "isSystem": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "license": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "maintainer": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "path": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "projectURL": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "readme": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "status": {
                    "create": false,
                    "nullable": true,
                    "type": "templateStatus",
                    "update": false
                },
                "templateBase": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "upgradeFrom": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "versions": {
                    "create": true,
                    "nullable": true,
                    "type": "array[templateVersionSpec]",
                    "update": true
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "templateStatus",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/templateStatus"
            },
            "pluralName": "templateStatuses",
            "resourceFields": {},
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "externalId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "maximumRancherVersion": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "minimumRancherVersion": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "readme": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "revision": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "upgradeFrom": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "version": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "templateVersion",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/templateversions",
                "self": "https://192.168.1.16:8443/v3/schemas/templateVersion"
            },
            "pluralName": "templateVersions",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "externalId": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "files": {
                    "create": true,
                    "nullable": true,
                    "type": "array[file]",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "maximumRancherVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "minimumRancherVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "questions": {
                    "create": true,
                    "nullable": true,
                    "type": "array[question]",
                    "update": true
                },
                "readme": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "revision": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "status": {
                    "create": false,
                    "nullable": true,
                    "type": "templateVersionStatus",
                    "update": false
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "upgradeFrom": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "upgradeVersionLinks": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "version": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "templateVersionSpec",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/templateVersionSpec"
            },
            "pluralName": "templateVersionSpecs",
            "resourceFields": {
                "externalId": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "files": {
                    "create": true,
                    "nullable": true,
                    "type": "array[file]",
                    "update": true
                },
                "maximumRancherVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "minimumRancherVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "questions": {
                    "create": true,
                    "nullable": true,
                    "type": "array[question]",
                    "update": true
                },
                "readme": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "revision": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "upgradeFrom": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "upgradeVersionLinks": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "version": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "templateVersionStatus",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/templateVersionStatus"
            },
            "pluralName": "templateVersionStatuses",
            "resourceFields": {},
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionActions": {
                "login": {
                    "input": "loginInput",
                    "output": "token"
                },
                "logout": {}
            },
            "collectionFilters": {
                "authProvider": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "isDerived": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "lastUpdateTime": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "token": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "ttl": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "userId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "userPrincipal": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "token",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/tokens",
                "self": "https://192.168.1.16:8443/v3/schemas/token"
            },
            "pluralName": "tokens",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "authProvider": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "groupPrincipals": {
                    "create": true,
                    "nullable": true,
                    "type": "array[reference[Principal]]",
                    "update": true
                },
                "isDerived": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "lastUpdateTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "providerInfo": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "token": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": false,
                    "writeOnly": true
                },
                "ttl": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "userId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[User]",
                    "update": true
                },
                "userPrincipal": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[Principal]",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionActions": {
                "changepassword": {
                    "input": "changePasswordInput"
                }
            },
            "collectionFilters": {
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "me": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "mustChangePassword": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "name": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "password": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "username": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "user",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/users",
                "self": "https://192.168.1.16:8443/v3/schemas/user"
            },
            "pluralName": "users",
            "resourceActions": {
                "setpassword": {
                    "input": "setPasswordInput",
                    "output": "user"
                }
            },
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "id": {
                    "create": false,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "me": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "mustChangePassword": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "password": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": false,
                    "writeOnly": true
                },
                "principalIds": {
                    "create": true,
                    "nullable": true,
                    "type": "array[reference[Principal]]",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "username": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "values",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/values"
            },
            "pluralName": "valueses",
            "resourceFields": {
                "boolValue": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "intValue": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "stringSliceValue": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "stringValue": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "vmwarevsphereconfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/vmwarevsphereconfig"
            },
            "pluralName": "vmwarevsphereconfigs",
            "resourceFields": {
                "boot2dockerUrl": {
                    "create": true,
                    "default": "",
                    "description": "vSphere URL for boot2docker image",
                    "type": "string",
                    "update": false
                },
                "cfgparam": {
                    "create": true,
                    "description": "vSphere vm configuration parameters (used for guestinfo)",
                    "type": "array[string]",
                    "update": false
                },
                "cloudinit": {
                    "create": true,
                    "default": "",
                    "description": "vSphere cloud-init file or url to set in the guestinfo",
                    "type": "string",
                    "update": false
                },
                "cpuCount": {
                    "create": true,
                    "default": "2",
                    "description": "vSphere CPU number for docker VM",
                    "type": "string",
                    "update": false
                },
                "datacenter": {
                    "create": true,
                    "default": "",
                    "description": "vSphere datacenter for docker VM",
                    "type": "string",
                    "update": false
                },
                "datastore": {
                    "create": true,
                    "default": "",
                    "description": "vSphere datastore for docker VM",
                    "type": "string",
                    "update": false
                },
                "diskSize": {
                    "create": true,
                    "default": "20480",
                    "description": "vSphere size of disk for docker VM (in MB)",
                    "type": "string",
                    "update": false
                },
                "hostsystem": {
                    "create": true,
                    "default": "",
                    "description": "vSphere compute resource where the docker VM will be instantiated. This can be omitted if using a cluster with DRS.",
                    "type": "string",
                    "update": false
                },
                "memorySize": {
                    "create": true,
                    "default": "2048",
                    "description": "vSphere size of memory for docker VM (in MB)",
                    "type": "string",
                    "update": false
                },
                "network": {
                    "create": true,
                    "description": "vSphere network where the docker VM will be attached",
                    "type": "array[string]",
                    "update": false
                },
                "password": {
                    "create": true,
                    "default": "",
                    "description": "vSphere password",
                    "type": "string",
                    "update": false
                },
                "pool": {
                    "create": true,
                    "default": "",
                    "description": "vSphere resource pool for docker VM",
                    "type": "string",
                    "update": false
                },
                "username": {
                    "create": true,
                    "default": "",
                    "description": "vSphere username",
                    "type": "string",
                    "update": false
                },
                "vcenter": {
                    "create": true,
                    "default": "",
                    "description": "vSphere IP/hostname for vCenter",
                    "type": "string",
                    "update": false
                },
                "vcenterPort": {
                    "create": true,
                    "default": "443",
                    "description": "vSphere Port for vCenter",
                    "type": "string",
                    "update": false
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "webhookConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/webhookConfig"
            },
            "pluralName": "webhookConfigs",
            "resourceFields": {
                "url": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "zookeeper",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/zookeeper"
            },
            "pluralName": "zookeepers",
            "resourceFields": {
                "endpoint": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        }
    ]
}
```

```
{
    "type": "collection",
    "links": {
        "self": "https://192.168.1.16:8443/v3/schemas"
    },
    "actions": {},
    "pagination": {
        "limit": 1000
    },
    "sort": {
        "order": "asc",
        "reverse": "https://192.168.1.16:8443/v3/schemas?order=desc"
    },
    "resourceType": "schema",
    "data": [
        {
            "actions": {},
            "baseType": "schema",
            "id": "action",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/action"
            },
            "pluralName": "actions",
            "resourceFields": {
                "input": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "output": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "alertCondition",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/alertCondition"
            },
            "pluralName": "alertConditions",
            "resourceFields": {
                "lastTransitionTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "lastUpdateTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "message": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "reason": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "status": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "type": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "alertStatus",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/alertStatus"
            },
            "pluralName": "alertStatuses",
            "resourceFields": {
                "conditions": {
                    "create": false,
                    "nullable": true,
                    "type": "array[alertCondition]",
                    "update": false
                },
                "startedAt": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "state": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "allowedHostPath",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/allowedHostPath"
            },
            "pluralName": "allowedHostPaths",
            "resourceFields": {
                "pathPrefix": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "amazonec2config",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/amazonec2config"
            },
            "pluralName": "amazonec2configs",
            "resourceFields": {
                "accessKey": {
                    "create": true,
                    "default": "",
                    "description": "AWS Access Key",
                    "type": "string",
                    "update": false
                },
                "ami": {
                    "create": true,
                    "default": "",
                    "description": "AWS machine image",
                    "type": "string",
                    "update": false
                },
                "blockDurationMinutes": {
                    "create": true,
                    "default": "0",
                    "description": "AWS spot instance duration in minutes (60, 120, 180, 240, 300, or 360)",
                    "type": "string",
                    "update": false
                },
                "deviceName": {
                    "create": true,
                    "default": "/dev/sda1",
                    "description": "AWS root device name",
                    "type": "string",
                    "update": false
                },
                "endpoint": {
                    "create": true,
                    "default": "",
                    "description": "Optional endpoint URL (hostname only or fully qualified URI)",
                    "type": "string",
                    "update": false
                },
                "iamInstanceProfile": {
                    "create": true,
                    "default": "",
                    "description": "AWS IAM Instance Profile",
                    "type": "string",
                    "update": false
                },
                "insecureTransport": {
                    "create": true,
                    "default": false,
                    "description": "Disable SSL when sending requests",
                    "type": "boolean",
                    "update": false
                },
                "instanceType": {
                    "create": true,
                    "default": "t2.micro",
                    "description": "AWS instance type",
                    "type": "string",
                    "update": false
                },
                "keypairName": {
                    "create": true,
                    "default": "",
                    "description": "AWS keypair to use; requires --amazonec2-ssh-keypath",
                    "type": "string",
                    "update": false
                },
                "monitoring": {
                    "create": true,
                    "default": false,
                    "description": "Set this flag to enable CloudWatch monitoring",
                    "type": "boolean",
                    "update": false
                },
                "openPort": {
                    "create": true,
                    "description": "Make the specified port number accessible from the Internet",
                    "type": "array[string]",
                    "update": false
                },
                "privateAddressOnly": {
                    "create": true,
                    "default": false,
                    "description": "Only use a private IP address",
                    "type": "boolean",
                    "update": false
                },
                "region": {
                    "create": true,
                    "default": "us-east-1",
                    "description": "AWS region",
                    "type": "string",
                    "update": false
                },
                "requestSpotInstance": {
                    "create": true,
                    "default": false,
                    "description": "Set this flag to request spot instance",
                    "type": "boolean",
                    "update": false
                },
                "retries": {
                    "create": true,
                    "default": "5",
                    "description": "Set retry count for recoverable failures (use -1 to disable)",
                    "type": "string",
                    "update": false
                },
                "rootSize": {
                    "create": true,
                    "default": "16",
                    "description": "AWS root disk size (in GB)",
                    "type": "string",
                    "update": false
                },
                "secretKey": {
                    "create": true,
                    "default": "",
                    "description": "AWS Secret Key",
                    "type": "string",
                    "update": false
                },
                "securityGroup": {
                    "create": true,
                    "default": [
                        "docker-machine"
                    ],
                    "description": "AWS VPC security group",
                    "type": "array[string]",
                    "update": false
                },
                "sessionToken": {
                    "create": true,
                    "default": "",
                    "description": "AWS Session Token",
                    "type": "string",
                    "update": false
                },
                "spotPrice": {
                    "create": true,
                    "default": "0.50",
                    "description": "AWS spot instance bid price (in dollar)",
                    "type": "string",
                    "update": false
                },
                "sshKeypath": {
                    "create": true,
                    "default": "",
                    "description": "SSH Key for Instance",
                    "type": "string",
                    "update": false
                },
                "sshUser": {
                    "create": true,
                    "default": "ubuntu",
                    "description": "Set the name of the ssh user",
                    "type": "string",
                    "update": false
                },
                "subnetId": {
                    "create": true,
                    "default": "",
                    "description": "AWS VPC subnet id",
                    "type": "string",
                    "update": false
                },
                "tags": {
                    "create": true,
                    "default": "",
                    "description": "AWS Tags (e.g. key1,value1,key2,value2)",
                    "type": "string",
                    "update": false
                },
                "useEbsOptimizedInstance": {
                    "create": true,
                    "default": false,
                    "description": "Create an EBS optimized instance",
                    "type": "boolean",
                    "update": false
                },
                "usePrivateAddress": {
                    "create": true,
                    "default": false,
                    "description": "Force the usage of private IP address",
                    "type": "boolean",
                    "update": false
                },
                "userdata": {
                    "create": true,
                    "default": "",
                    "description": "path to file with cloud-init user data",
                    "type": "string",
                    "update": false
                },
                "volumeType": {
                    "create": true,
                    "default": "gp2",
                    "description": "Amazon EBS volume type",
                    "type": "string",
                    "update": false
                },
                "vpcId": {
                    "create": true,
                    "default": "",
                    "description": "AWS VPC id",
                    "type": "string",
                    "update": false
                },
                "zone": {
                    "create": true,
                    "default": "a",
                    "description": "AWS zone for instance (i.e. a,b,c,d,e)",
                    "type": "string",
                    "update": false
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "attachedVolume",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/attachedVolume"
            },
            "pluralName": "attachedVolumes",
            "resourceFields": {
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "authnConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/authnConfig"
            },
            "pluralName": "authnConfigs",
            "resourceFields": {
                "options": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "strategy": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "authzConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/authzConfig"
            },
            "pluralName": "authzConfigs",
            "resourceFields": {
                "mode": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "options": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "azureKubernetesServiceConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/azureKubernetesServiceConfig"
            },
            "pluralName": "azureKubernetesServiceConfigs",
            "resourceFields": {},
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "azureconfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/azureconfig"
            },
            "pluralName": "azureconfigs",
            "resourceFields": {
                "availabilitySet": {
                    "create": true,
                    "default": "docker-machine",
                    "description": "Azure Availability Set to place the virtual machine into",
                    "type": "string",
                    "update": false
                },
                "clientId": {
                    "create": true,
                    "default": "",
                    "description": "Azure Service Principal Account ID (optional, browser auth is used if not specified)",
                    "type": "string",
                    "update": false
                },
                "clientSecret": {
                    "create": true,
                    "default": "",
                    "description": "Azure Service Principal Account password (optional, browser auth is used if not specified)",
                    "type": "string",
                    "update": false
                },
                "customData": {
                    "create": true,
                    "default": "",
                    "description": "Path to file with custom-data",
                    "type": "string",
                    "update": false
                },
                "dns": {
                    "create": true,
                    "default": "",
                    "description": "A unique DNS label for the public IP adddress",
                    "type": "string",
                    "update": false
                },
                "dockerPort": {
                    "create": true,
                    "default": "2376",
                    "description": "Port number for Docker engine",
                    "type": "string",
                    "update": false
                },
                "environment": {
                    "create": true,
                    "default": "AzurePublicCloud",
                    "description": "Azure environment (e.g. AzurePublicCloud, AzureChinaCloud)",
                    "type": "string",
                    "update": false
                },
                "image": {
                    "create": true,
                    "default": "canonical:UbuntuServer:16.04.0-LTS:latest",
                    "description": "Azure virtual machine OS image",
                    "type": "string",
                    "update": false
                },
                "location": {
                    "create": true,
                    "default": "westus",
                    "description": "Azure region to create the virtual machine",
                    "type": "string",
                    "update": false
                },
                "noPublicIp": {
                    "create": true,
                    "default": false,
                    "description": "Do not create a public IP address for the machine",
                    "type": "boolean",
                    "update": false
                },
                "openPort": {
                    "create": true,
                    "description": "Make the specified port number accessible from the Internet",
                    "type": "array[string]",
                    "update": false
                },
                "privateIpAddress": {
                    "create": true,
                    "default": "",
                    "description": "Specify a static private IP address for the machine",
                    "type": "string",
                    "update": false
                },
                "resourceGroup": {
                    "create": true,
                    "default": "docker-machine",
                    "description": "Azure Resource Group name (will be created if missing)",
                    "type": "string",
                    "update": false
                },
                "size": {
                    "create": true,
                    "default": "Standard_A2",
                    "description": "Size for Azure Virtual Machine",
                    "type": "string",
                    "update": false
                },
                "sshUser": {
                    "create": true,
                    "default": "docker-user",
                    "description": "Username for SSH login",
                    "type": "string",
                    "update": false
                },
                "staticPublicIp": {
                    "create": true,
                    "default": false,
                    "description": "Assign a static public IP address to the machine",
                    "type": "boolean",
                    "update": false
                },
                "storageType": {
                    "create": true,
                    "default": "Standard_LRS",
                    "description": "Type of Storage Account to host the OS Disk for the machine",
                    "type": "string",
                    "update": false
                },
                "subnet": {
                    "create": true,
                    "default": "docker-machine",
                    "description": "Azure Subnet Name to be used within the Virtual Network",
                    "type": "string",
                    "update": false
                },
                "subnetPrefix": {
                    "create": true,
                    "default": "192.168.0.0/16",
                    "description": "Private CIDR block to be used for the new subnet, should comply RFC 1918",
                    "type": "string",
                    "update": false
                },
                "subscriptionId": {
                    "create": true,
                    "default": "",
                    "description": "Azure Subscription ID",
                    "type": "string",
                    "update": false
                },
                "usePrivateIp": {
                    "create": true,
                    "default": false,
                    "description": "Use private IP address of the machine to connect",
                    "type": "boolean",
                    "update": false
                },
                "vnet": {
                    "create": true,
                    "default": "docker-machine-vnet",
                    "description": "Azure Virtual Network name to connect the virtual machine (in [resourcegroup:]name format)",
                    "type": "string",
                    "update": false
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "brokerList",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/brokerList"
            },
            "pluralName": "brokerLists",
            "resourceFields": {
                "brokerList": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "branch": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "catalogKind": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "url": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "catalog",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/catalogs",
                "self": "https://192.168.1.16:8443/v3/schemas/catalog"
            },
            "pluralName": "catalogs",
            "resourceActions": {
                "refresh": {}
            },
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "branch": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "catalogKind": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "status": {
                    "create": false,
                    "nullable": true,
                    "type": "catalogStatus",
                    "update": false
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "url": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "catalogStatus",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/catalogStatus"
            },
            "pluralName": "catalogStatuses",
            "resourceFields": {
                "commit": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "lastRefreshTimestamp": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "changePasswordInput",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/changePasswordInput"
            },
            "pluralName": "changePasswordInputs",
            "resourceFields": {
                "currentPassword": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "newPassword": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "apiEndpoint": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "caCert": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "defaultClusterRoleForProjectMembers": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "defaultPodSecurityPolicyTemplateId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "driver": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "internal": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "name": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "cluster",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/clusters",
                "self": "https://192.168.1.16:8443/v3/schemas/cluster"
            },
            "pluralName": "clusters",
            "resourceFields": {
                "allocatable": {
                    "create": false,
                    "nullable": true,
                    "type": "map[string]",
                    "update": false
                },
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "apiEndpoint": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "azureKubernetesServiceConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "azureKubernetesServiceConfig",
                    "update": true
                },
                "caCert": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "capacity": {
                    "create": false,
                    "nullable": true,
                    "type": "map[string]",
                    "update": false
                },
                "componentStatuses": {
                    "create": false,
                    "nullable": true,
                    "type": "array[clusterComponentStatus]",
                    "update": false
                },
                "conditions": {
                    "create": false,
                    "nullable": true,
                    "type": "array[clusterCondition]",
                    "update": false
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "defaultClusterRoleForProjectMembers": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[roleTemplate]",
                    "update": true
                },
                "defaultPodSecurityPolicyTemplateId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[podSecurityPolicyTemplate]",
                    "update": true
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "driver": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "embeddedConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "k8sServerConfig",
                    "update": false
                },
                "googleKubernetesEngineConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "googleKubernetesEngineConfig",
                    "update": true
                },
                "id": {
                    "create": false,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "importedConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "importedConfig",
                    "update": false
                },
                "internal": {
                    "create": false,
                    "nullable": true,
                    "type": "boolean",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "limits": {
                    "create": false,
                    "nullable": true,
                    "type": "map[string]",
                    "update": false
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": true
                },
                "nodes": {
                    "create": true,
                    "nullable": true,
                    "type": "array[machineConfig]",
                    "update": true
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "rancherKubernetesEngineConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "rancherKubernetesEngineConfig",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "requested": {
                    "create": false,
                    "nullable": true,
                    "type": "map[string]",
                    "update": false
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "clusterId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "displayName": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "initialWaitSeconds": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "namespaceId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "repeatIntervalSeconds": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "severity": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "clusterAlert",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/clusteralerts",
                "self": "https://192.168.1.16:8443/v3/schemas/clusterAlert"
            },
            "pluralName": "clusterAlerts",
            "resourceActions": {
                "activate": {},
                "deactivate": {},
                "mute": {},
                "unmute": {}
            },
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "clusterId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[cluster]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "displayName": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "initialWaitSeconds": {
                    "create": true,
                    "default": "180",
                    "min": 0,
                    "nullable": true,
                    "required": true,
                    "type": "int",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "namespaceId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[namespace]",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "recipientList": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "array[recipient]",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "repeatIntervalSeconds": {
                    "create": true,
                    "default": "3600",
                    "min": 0,
                    "nullable": true,
                    "required": true,
                    "type": "int",
                    "update": true
                },
                "severity": {
                    "create": true,
                    "default": "critical",
                    "nullable": true,
                    "options": [
                        "info",
                        "critical",
                        "warning"
                    ],
                    "required": true,
                    "type": "enum",
                    "update": true
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "status": {
                    "create": false,
                    "nullable": true,
                    "type": "alertStatus",
                    "update": false
                },
                "targetNode": {
                    "create": true,
                    "nullable": true,
                    "type": "targetNode",
                    "update": true
                },
                "targetSystemService": {
                    "create": true,
                    "nullable": true,
                    "type": "targetSystemService",
                    "update": true
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "clusterComponentStatus",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/clusterComponentStatus"
            },
            "pluralName": "clusterComponentStatuses",
            "resourceFields": {
                "conditions": {
                    "create": true,
                    "nullable": true,
                    "type": "array[componentCondition]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "clusterCondition",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/clusterCondition"
            },
            "pluralName": "clusterConditions",
            "resourceFields": {
                "lastTransitionTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "lastUpdateTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "message": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "reason": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "status": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "type": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "clusterId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "count": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "eventType": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "message": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "namespaceId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "reason": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "clusterEvent",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/clusterevents",
                "self": "https://192.168.1.16:8443/v3/schemas/clusterEvent"
            },
            "pluralName": "clusterEvents",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "clusterId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[cluster]",
                    "update": true
                },
                "count": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "eventType": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "firstTimestamp": {
                    "create": true,
                    "nullable": true,
                    "type": "date",
                    "update": true
                },
                "involvedObject": {
                    "create": true,
                    "nullable": true,
                    "type": "objectReference",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "lastTimestamp": {
                    "create": true,
                    "nullable": true,
                    "type": "date",
                    "update": true
                },
                "message": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "namespaceId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[namespace]",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "reason": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "source": {
                    "create": true,
                    "nullable": true,
                    "type": "eventSource",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "clusterId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "displayName": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "namespaceId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "outputFlushInterval": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "clusterLogging",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/clusterloggings",
                "self": "https://192.168.1.16:8443/v3/schemas/clusterLogging"
            },
            "pluralName": "clusterLoggings",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "clusterId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[cluster]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "displayName": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "elasticsearchConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "elasticsearchConfig",
                    "update": true
                },
                "embeddedConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "embeddedConfig",
                    "update": true
                },
                "kafkaConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "kafkaConfig",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "namespaceId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[namespace]",
                    "update": false
                },
                "outputFlushInterval": {
                    "create": true,
                    "default": "3",
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "outputTags": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "splunkConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "splunkConfig",
                    "update": true
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "status": {
                    "create": false,
                    "nullable": true,
                    "type": "loggingStatus",
                    "update": false
                },
                "syslogConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "syslogConfig",
                    "update": true
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "clusterId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "namespaceId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "clusterRegistrationToken",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/clusterregistrationtokens",
                "self": "https://192.168.1.16:8443/v3/schemas/clusterRegistrationToken"
            },
            "pluralName": "clusterRegistrationTokens",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "clusterId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[cluster]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "namespaceId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[namespace]",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "status": {
                    "create": false,
                    "nullable": true,
                    "type": "clusterRegistrationTokenStatus",
                    "update": false
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "clusterRegistrationTokenStatus",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/clusterRegistrationTokenStatus"
            },
            "pluralName": "clusterRegistrationTokenStatuses",
            "resourceFields": {
                "command": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "manifestUrl": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "token": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "clusterId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "namespaceId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "roleTemplateId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "userId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "clusterRoleTemplateBinding",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/clusterroletemplatebindings",
                "self": "https://192.168.1.16:8443/v3/schemas/clusterRoleTemplateBinding"
            },
            "pluralName": "clusterRoleTemplateBindings",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "clusterId": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "reference[cluster]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "namespaceId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[namespace]",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "roleTemplateId": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "reference[roleTemplate]",
                    "update": true
                },
                "userId": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "reference[user]",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "componentCondition",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/componentCondition"
            },
            "pluralName": "componentConditions",
            "resourceFields": {
                "error": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "message": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "status": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "type": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "cpuInfo",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/cpuInfo"
            },
            "pluralName": "cpuInfos",
            "resourceFields": {
                "count": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "customConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/customConfig"
            },
            "pluralName": "customConfigs",
            "resourceFields": {
                "address": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "dockerSocket": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "internalAddress": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "sshKey": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "user": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "digitaloceanconfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/digitaloceanconfig"
            },
            "pluralName": "digitaloceanconfigs",
            "resourceFields": {
                "accessToken": {
                    "create": true,
                    "default": "",
                    "description": "Digital Ocean access token",
                    "type": "string",
                    "update": false
                },
                "backups": {
                    "create": true,
                    "default": false,
                    "description": "enable backups for droplet",
                    "type": "boolean",
                    "update": false
                },
                "image": {
                    "create": true,
                    "default": "ubuntu-16-04-x64",
                    "description": "Digital Ocean Image",
                    "type": "string",
                    "update": false
                },
                "ipv6": {
                    "create": true,
                    "default": false,
                    "description": "enable ipv6 for droplet",
                    "type": "boolean",
                    "update": false
                },
                "privateNetworking": {
                    "create": true,
                    "default": false,
                    "description": "enable private networking for droplet",
                    "type": "boolean",
                    "update": false
                },
                "region": {
                    "create": true,
                    "default": "nyc3",
                    "description": "Digital Ocean region",
                    "type": "string",
                    "update": false
                },
                "size": {
                    "create": true,
                    "default": "512mb",
                    "description": "Digital Ocean size",
                    "type": "string",
                    "update": false
                },
                "sshKeyFingerprint": {
                    "create": true,
                    "default": "",
                    "description": "SSH key fingerprint",
                    "type": "string",
                    "update": false
                },
                "sshKeyPath": {
                    "create": true,
                    "default": "",
                    "description": "SSH private key path ",
                    "type": "string",
                    "update": false
                },
                "sshPort": {
                    "create": true,
                    "default": "22",
                    "description": "SSH port",
                    "type": "string",
                    "update": false
                },
                "sshUser": {
                    "create": true,
                    "default": "root",
                    "description": "SSH username",
                    "type": "string",
                    "update": false
                },
                "tags": {
                    "create": true,
                    "default": "",
                    "description": "comma-separated list of tags to apply to the Droplet",
                    "type": "string",
                    "update": false
                },
                "userdata": {
                    "create": true,
                    "default": "",
                    "description": "path to file with cloud-init user-data",
                    "type": "string",
                    "update": false
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "embed": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "embedType": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "pluralName": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "dynamicSchema",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/dynamicschemas",
                "self": "https://192.168.1.16:8443/v3/schemas/dynamicSchema"
            },
            "pluralName": "dynamicSchemas",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "collectionActions": {
                    "create": true,
                    "nullable": true,
                    "type": "map[action]",
                    "update": true
                },
                "collectionFields": {
                    "create": true,
                    "nullable": true,
                    "type": "map[field]",
                    "update": true
                },
                "collectionFilters": {
                    "create": true,
                    "nullable": true,
                    "type": "map[filter]",
                    "update": true
                },
                "collectionMethods": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "embed": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "embedType": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "includeableLinks": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "pluralName": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "resourceActions": {
                    "create": true,
                    "nullable": true,
                    "type": "map[action]",
                    "update": true
                },
                "resourceFields": {
                    "create": true,
                    "nullable": true,
                    "type": "map[field]",
                    "update": true
                },
                "resourceMethods": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "status": {
                    "create": false,
                    "nullable": true,
                    "type": "dynamicSchemaStatus",
                    "update": false
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "dynamicSchemaStatus",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/dynamicSchemaStatus"
            },
            "pluralName": "dynamicSchemaStatuses",
            "resourceFields": {
                "fake": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "elasticsearchConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/elasticsearchConfig"
            },
            "pluralName": "elasticsearchConfigs",
            "resourceFields": {
                "authPassword": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "authUsername": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "dateFormat": {
                    "create": true,
                    "default": "YYYY-MM-DD",
                    "nullable": true,
                    "options": [
                        "YYYY-MM-DD",
                        "YYYY-MM",
                        "YYYY"
                    ],
                    "required": true,
                    "type": "enum",
                    "update": true
                },
                "endpoint": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "indexPrefix": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "embeddedConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/embeddedConfig"
            },
            "pluralName": "embeddedConfigs",
            "resourceFields": {
                "dateFormat": {
                    "create": true,
                    "default": "YYYY-MM-DD",
                    "nullable": true,
                    "options": [
                        "YYYY-MM-DD",
                        "YYYY-MM",
                        "YYYY"
                    ],
                    "required": true,
                    "type": "enum",
                    "update": true
                },
                "indexPrefix": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "etcdService",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/etcdService"
            },
            "pluralName": "etcdServices",
            "resourceFields": {
                "extraArgs": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "image": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "eventSource",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/eventSource"
            },
            "pluralName": "eventSources",
            "resourceFields": {
                "component": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "host": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "field",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/field"
            },
            "pluralName": "fields",
            "resourceFields": {
                "create": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "default": {
                    "create": true,
                    "nullable": true,
                    "type": "values",
                    "update": true
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "invalidChars": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "max": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "maxLength": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "min": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "minLength": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "nullable": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "options": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "required": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "type": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "unique": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "update": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "validChars": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "file",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/file"
            },
            "pluralName": "files",
            "resourceFields": {
                "contents": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "filter",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/filter"
            },
            "pluralName": "filters",
            "resourceFields": {
                "modifiers": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "fsGroupStrategyOptions",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/fsGroupStrategyOptions"
            },
            "pluralName": "fsGroupStrategyOptionses",
            "resourceFields": {
                "ranges": {
                    "create": true,
                    "nullable": true,
                    "type": "array[idRange]",
                    "update": true
                },
                "rule": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "githubCredential",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/githubCredential"
            },
            "pluralName": "githubCredentials",
            "resourceFields": {
                "code": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "builtin": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "name": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "globalRole",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/globalroles",
                "self": "https://192.168.1.16:8443/v3/schemas/globalRole"
            },
            "pluralName": "globalRoles",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "builtin": {
                    "create": false,
                    "nullable": true,
                    "type": "boolean",
                    "update": false
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "id": {
                    "create": false,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "rules": {
                    "create": true,
                    "nullable": true,
                    "type": "array[policyRule]",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "globalRoleId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "userId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "globalRoleBinding",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/globalrolebindings",
                "self": "https://192.168.1.16:8443/v3/schemas/globalRoleBinding"
            },
            "pluralName": "globalRoleBindings",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "globalRoleId": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "reference[globalRole]",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "userId": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "reference[user]",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "googleKubernetesEngineConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/googleKubernetesEngineConfig"
            },
            "pluralName": "googleKubernetesEngineConfigs",
            "resourceFields": {
                "clusterIpv4Cidr": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "credential": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "diskSizeGb": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "enableAlphaFeature": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "horizontalPodAutoscaling": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "httpLoadBalancing": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "imageType": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "kubernetesDashboard": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "legacyAbac": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "locations": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "machineType": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "masterVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "network": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "networkPolicyConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "nodeCount": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "int",
                    "update": true
                },
                "nodeVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "projectId": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "subNetwork": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "zone": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "name": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "group",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/groups",
                "self": "https://192.168.1.16:8443/v3/schemas/group"
            },
            "pluralName": "groups",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "id": {
                    "create": false,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "groupId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "principalId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "groupMember",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/groupmembers",
                "self": "https://192.168.1.16:8443/v3/schemas/groupMember"
            },
            "pluralName": "groupMembers",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "groupId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[group]",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "principalId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[Principal]",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "hostPortRange",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/hostPortRange"
            },
            "pluralName": "hostPortRanges",
            "resourceFields": {
                "max": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "min": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "idRange",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/idRange"
            },
            "pluralName": "idRanges",
            "resourceFields": {
                "max": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "min": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "importedConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/importedConfig"
            },
            "pluralName": "importedConfigs",
            "resourceFields": {
                "kubeConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "k8sServerConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/k8sServerConfig"
            },
            "pluralName": "k8sServerConfigs",
            "resourceFields": {
                "admissionControllers": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "serviceNetCidr": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "kafkaConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/kafkaConfig"
            },
            "pluralName": "kafkaConfigs",
            "resourceFields": {
                "broker": {
                    "create": true,
                    "nullable": true,
                    "type": "brokerList",
                    "update": true
                },
                "maxSendRetries": {
                    "create": true,
                    "default": "3",
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "topic": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "zookeeper": {
                    "create": true,
                    "nullable": true,
                    "type": "zookeeper",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "kubeAPIService",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/kubeAPIService"
            },
            "pluralName": "kubeAPIServices",
            "resourceFields": {
                "extraArgs": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "image": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "podSecurityPolicy": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "serviceClusterIpRange": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "kubeControllerService",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/kubeControllerService"
            },
            "pluralName": "kubeControllerServices",
            "resourceFields": {
                "clusterCidr": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "extraArgs": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "image": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "serviceClusterIpRange": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "kubeletService",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/kubeletService"
            },
            "pluralName": "kubeletServices",
            "resourceFields": {
                "clusterDnsServer": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "clusterDomain": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "extraArgs": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "image": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "infraContainerImage": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "kubeproxyService",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/kubeproxyService"
            },
            "pluralName": "kubeproxyServices",
            "resourceFields": {
                "extraArgs": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "image": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "kubernetesInfo",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/kubernetesInfo"
            },
            "pluralName": "kubernetesInfos",
            "resourceFields": {
                "kubeProxyVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "kubeletVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "algorithm": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "caCerts": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "cert": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "certFingerprint": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "cn": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "enabled": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "expiresAt": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "issuedAt": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "issuer": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "key": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "keySize": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "mode": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "name": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "serialNumber": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "version": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "listenConfig",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/listenconfigs",
                "self": "https://192.168.1.16:8443/v3/schemas/listenConfig"
            },
            "pluralName": "listenConfigs",
            "resourceFields": {
                "algorithm": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "caCerts": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "cert": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "certFingerprint": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "cn": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "domains": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "enabled": {
                    "create": true,
                    "default": "true",
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "expiresAt": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "id": {
                    "create": false,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "issuedAt": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "issuer": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "key": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true,
                    "writeOnly": true
                },
                "keySize": {
                    "create": false,
                    "nullable": true,
                    "type": "int",
                    "update": false
                },
                "knownIps": {
                    "create": false,
                    "nullable": true,
                    "type": "array[string]",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "mode": {
                    "create": true,
                    "nullable": true,
                    "options": [
                        "https",
                        "http",
                        "acme"
                    ],
                    "type": "enum",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "serialNumber": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "subjectAlternativeNames": {
                    "create": false,
                    "nullable": true,
                    "type": "array[string]",
                    "update": false
                },
                "tos": {
                    "create": true,
                    "default": "auto",
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "version": {
                    "create": false,
                    "nullable": true,
                    "type": "int",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "localCredential",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/localCredential"
            },
            "pluralName": "localCredentials",
            "resourceFields": {
                "password": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "username": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "loggingCondition",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/loggingCondition"
            },
            "pluralName": "loggingConditions",
            "resourceFields": {
                "lastTransitionTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "lastUpdateTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "message": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "reason": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "status": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "type": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "loggingStatus",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/loggingStatus"
            },
            "pluralName": "loggingStatuses",
            "resourceFields": {
                "conditions": {
                    "create": false,
                    "nullable": true,
                    "type": "array[loggingCondition]",
                    "update": false
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "loginInput",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/loginInput"
            },
            "pluralName": "loginInputs",
            "resourceFields": {
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "githubCredential": {
                    "create": true,
                    "nullable": true,
                    "type": "githubCredential",
                    "update": true
                },
                "localCredential": {
                    "create": true,
                    "nullable": true,
                    "type": "localCredential",
                    "update": true
                },
                "responseType": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "ttl": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "clusterId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "hostname": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "ipAddress": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "machineTemplateId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "name": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "namespaceId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "nodeName": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "podCidr": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "providerId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "sshUser": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "unschedulable": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "useInternalIpAddress": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "machine",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/machines",
                "self": "https://192.168.1.16:8443/v3/schemas/machine"
            },
            "pluralName": "machines",
            "resourceFields": {
                "allocatable": {
                    "create": false,
                    "nullable": true,
                    "type": "map[string]",
                    "update": false
                },
                "amazonec2Config": {
                    "create": true,
                    "nullable": true,
                    "type": "amazonec2config",
                    "update": false
                },
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "azureConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "azureconfig",
                    "update": false
                },
                "capacity": {
                    "create": false,
                    "nullable": true,
                    "type": "map[string]",
                    "update": false
                },
                "clusterId": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "reference[cluster]",
                    "update": false
                },
                "conditions": {
                    "create": false,
                    "nullable": true,
                    "type": "array[machineCondition]",
                    "update": false
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "customConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "customConfig",
                    "update": true
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "digitaloceanConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "digitaloceanconfig",
                    "update": false
                },
                "hostname": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "id": {
                    "create": false,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "info": {
                    "create": false,
                    "nullable": true,
                    "type": "nodeInfo",
                    "update": false
                },
                "ipAddress": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "limits": {
                    "create": false,
                    "nullable": true,
                    "type": "map[string]",
                    "update": false
                },
                "machineTemplateId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[machineTemplate]",
                    "update": false
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "namespaceId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[namespace]",
                    "update": false
                },
                "nodeAnnotations": {
                    "create": false,
                    "nullable": true,
                    "type": "map[string]",
                    "update": false
                },
                "nodeLabels": {
                    "create": false,
                    "nullable": true,
                    "type": "map[string]",
                    "update": false
                },
                "nodeName": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "nodeTaints": {
                    "create": false,
                    "nullable": true,
                    "type": "array[taint]",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "podCidr": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "providerId": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "requested": {
                    "create": false,
                    "nullable": true,
                    "type": "map[string]",
                    "update": false
                },
                "requestedHostname": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "role": {
                    "create": true,
                    "nullable": true,
                    "options": [
                        "etcd",
                        "worker",
                        "controlplane"
                    ],
                    "type": "array[enum]",
                    "update": false
                },
                "sshUser": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "taints": {
                    "create": false,
                    "nullable": true,
                    "type": "array[taint]",
                    "update": true
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "unschedulable": {
                    "create": false,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "useInternalIpAddress": {
                    "create": true,
                    "default": "true",
                    "nullable": true,
                    "type": "boolean",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "vmwarevsphereConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "vmwarevsphereconfig",
                    "update": false
                },
                "volumesAttached": {
                    "create": false,
                    "nullable": true,
                    "type": "map[attachedVolume]",
                    "update": false
                },
                "volumesInUse": {
                    "create": false,
                    "nullable": true,
                    "type": "array[string]",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "machineCondition",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/machineCondition"
            },
            "pluralName": "machineConditions",
            "resourceFields": {
                "lastTransitionTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "lastUpdateTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "message": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "reason": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "status": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "type": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "machineConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/machineConfig"
            },
            "pluralName": "machineConfigs",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "customConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "customConfig",
                    "update": true
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "displayName": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "machineTemplateId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[machineTemplate]",
                    "update": false
                },
                "nodeSpec": {
                    "create": true,
                    "nullable": true,
                    "type": "nodeSpec",
                    "update": true
                },
                "requestedHostname": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "role": {
                    "create": true,
                    "nullable": true,
                    "options": [
                        "etcd",
                        "worker",
                        "controlplane"
                    ],
                    "type": "array[enum]",
                    "update": false
                },
                "useInternalIpAddress": {
                    "create": true,
                    "default": "true",
                    "nullable": true,
                    "type": "boolean",
                    "update": false
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "active": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "builtin": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "checksum": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "externalId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "name": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uiUrl": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "url": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "machineDriver",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/machinedrivers",
                "self": "https://192.168.1.16:8443/v3/schemas/machineDriver"
            },
            "pluralName": "machineDrivers",
            "resourceActions": {
                "activate": {
                    "output": "machineDriver"
                },
                "deactivate": {
                    "output": "machineDriver"
                }
            },
            "resourceFields": {
                "active": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "builtin": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "checksum": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "externalId": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "id": {
                    "create": false,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "status": {
                    "create": false,
                    "nullable": true,
                    "type": "machineDriverStatus",
                    "update": false
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "uiUrl": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "url": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "machineDriverCondition",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/machineDriverCondition"
            },
            "pluralName": "machineDriverConditions",
            "resourceFields": {
                "lastTransitionTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "lastUpdateTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "message": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "reason": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "status": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "type": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "machineDriverStatus",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/machineDriverStatus"
            },
            "pluralName": "machineDriverStatuses",
            "resourceFields": {
                "conditions": {
                    "create": false,
                    "nullable": true,
                    "type": "array[machineDriverCondition]",
                    "update": false
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "authCertificateAuthority": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "authKey": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "dockerVersion": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "driver": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "engineInstallURL": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "engineStorageDriver": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "name": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "machineTemplate",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/machinetemplates",
                "self": "https://192.168.1.16:8443/v3/schemas/machineTemplate"
            },
            "pluralName": "machineTemplates",
            "resourceFields": {
                "amazonec2Config": {
                    "create": true,
                    "nullable": true,
                    "type": "amazonec2config",
                    "update": false
                },
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "authCertificateAuthority": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "authKey": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "azureConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "azureconfig",
                    "update": false
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "digitaloceanConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "digitaloceanconfig",
                    "update": false
                },
                "dockerVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "driver": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "engineEnv": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "engineInsecureRegistry": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "engineInstallURL": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "engineLabel": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "engineOpt": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "engineRegistryMirror": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "engineStorageDriver": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "id": {
                    "create": false,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "status": {
                    "create": false,
                    "nullable": true,
                    "type": "machineTemplateStatus",
                    "update": false
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "vmwarevsphereConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "vmwarevsphereconfig",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "machineTemplateCondition",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/machineTemplateCondition"
            },
            "pluralName": "machineTemplateConditions",
            "resourceFields": {
                "lastTransitionTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "lastUpdateTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "reason": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "status": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "type": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "machineTemplateStatus",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/machineTemplateStatus"
            },
            "pluralName": "machineTemplateStatuses",
            "resourceFields": {
                "conditions": {
                    "create": false,
                    "nullable": true,
                    "type": "array[machineTemplateCondition]",
                    "update": false
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "memoryInfo",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/memoryInfo"
            },
            "pluralName": "memoryInfos",
            "resourceFields": {
                "memTotalKiB": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "networkConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/networkConfig"
            },
            "pluralName": "networkConfigs",
            "resourceFields": {
                "options": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "plugin": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "hostname": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "ipAddress": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "podCidr": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "providerId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "unschedulable": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "node",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/nodes",
                "self": "https://192.168.1.16:8443/v3/schemas/node"
            },
            "pluralName": "nodes",
            "resourceFields": {
                "allocatable": {
                    "create": false,
                    "nullable": true,
                    "type": "map[string]",
                    "update": false
                },
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "capacity": {
                    "create": false,
                    "nullable": true,
                    "type": "map[string]",
                    "update": false
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "hostname": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "info": {
                    "create": false,
                    "nullable": true,
                    "type": "nodeInfo",
                    "update": false
                },
                "ipAddress": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "nodeConditions": {
                    "create": false,
                    "nullable": true,
                    "type": "array[nodeCondition]",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "podCidr": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "providerId": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "taints": {
                    "create": false,
                    "nullable": true,
                    "type": "array[taint]",
                    "update": true
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "unschedulable": {
                    "create": false,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "volumesAttached": {
                    "create": false,
                    "nullable": true,
                    "type": "map[attachedVolume]",
                    "update": false
                },
                "volumesInUse": {
                    "create": false,
                    "nullable": true,
                    "type": "array[string]",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "nodeCondition",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/nodeCondition"
            },
            "pluralName": "nodeConditions",
            "resourceFields": {
                "lastHeartbeatTime": {
                    "create": true,
                    "nullable": true,
                    "type": "date",
                    "update": true
                },
                "lastTransitionTime": {
                    "create": true,
                    "nullable": true,
                    "type": "date",
                    "update": true
                },
                "message": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "reason": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "status": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "type": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "nodeInfo",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/nodeInfo"
            },
            "pluralName": "nodeInfos",
            "resourceFields": {
                "cpu": {
                    "create": true,
                    "nullable": true,
                    "type": "cpuInfo",
                    "update": true
                },
                "kubernetes": {
                    "create": true,
                    "nullable": true,
                    "type": "kubernetesInfo",
                    "update": true
                },
                "memory": {
                    "create": true,
                    "nullable": true,
                    "type": "memoryInfo",
                    "update": true
                },
                "os": {
                    "create": true,
                    "nullable": true,
                    "type": "osInfo",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "nodeSpec",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/nodeSpec"
            },
            "pluralName": "nodeSpecs",
            "resourceFields": {
                "podCidr": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "providerId": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "taints": {
                    "create": false,
                    "nullable": true,
                    "type": "array[taint]",
                    "update": true
                },
                "unschedulable": {
                    "create": false,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionActions": {
                "send": {}
            },
            "collectionFilters": {
                "clusterId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "displayName": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "namespaceId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "notifier",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/notifiers",
                "self": "https://192.168.1.16:8443/v3/schemas/notifier"
            },
            "pluralName": "notifiers",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "clusterId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[cluster]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "displayName": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "namespaceId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[namespace]",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "pagerdutyConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "pagerdutyConfig",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "slackConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "slackConfig",
                    "update": true
                },
                "smtpConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "smtpConfig",
                    "update": true
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "status": {
                    "create": false,
                    "nullable": true,
                    "type": "notifierStatus",
                    "update": false
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "webhookConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "webhookConfig",
                    "update": true
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "notifierStatus",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/notifierStatus"
            },
            "pluralName": "notifierStatuses",
            "resourceFields": {},
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "objectReference",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/objectReference"
            },
            "pluralName": "objectReferences",
            "resourceFields": {
                "apiVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "fieldPath": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "kind": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "namespace": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "resourceVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "uid": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "osInfo",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/osInfo"
            },
            "pluralName": "osInfos",
            "resourceFields": {
                "dockerVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "kernelVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "operatingSystem": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "ownerReference",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/ownerReference"
            },
            "pluralName": "ownerReferences",
            "resourceFields": {
                "apiVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "blockOwnerDeletion": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "controller": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "kind": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "uid": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "pagerdutyConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/pagerdutyConfig"
            },
            "pluralName": "pagerdutyConfigs",
            "resourceFields": {
                "serviceKey": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "allowPrivilegeEscalation": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "defaultAllowPrivilegeEscalation": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "hostIPC": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "hostNetwork": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "hostPID": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "privileged": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "readOnlyRootFilesystem": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "podSecurityPolicyTemplate",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/podsecuritypolicytemplates",
                "self": "https://192.168.1.16:8443/v3/schemas/podSecurityPolicyTemplate"
            },
            "pluralName": "podSecurityPolicyTemplates",
            "resourceFields": {
                "allowPrivilegeEscalation": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "allowedCapabilities": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "allowedHostPaths": {
                    "create": true,
                    "nullable": true,
                    "type": "array[allowedHostPath]",
                    "update": true
                },
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "defaultAddCapabilities": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "defaultAllowPrivilegeEscalation": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "fsGroup": {
                    "create": true,
                    "nullable": true,
                    "type": "fsGroupStrategyOptions",
                    "update": true
                },
                "hostIPC": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "hostNetwork": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "hostPID": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "hostPorts": {
                    "create": true,
                    "nullable": true,
                    "type": "array[hostPortRange]",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "privileged": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "readOnlyRootFilesystem": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "requiredDropCapabilities": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "runAsUser": {
                    "create": true,
                    "nullable": true,
                    "type": "runAsUserStrategyOptions",
                    "update": true
                },
                "seLinux": {
                    "create": true,
                    "nullable": true,
                    "type": "seLinuxStrategyOptions",
                    "update": true
                },
                "supplementalGroups": {
                    "create": true,
                    "nullable": true,
                    "type": "supplementalGroupsStrategyOptions",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "volumes": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "policyRule",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/policyRule"
            },
            "pluralName": "policyRules",
            "resourceFields": {
                "apiGroups": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "nonResourceURLs": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "resourceNames": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "resources": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "verbs": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "namespaceId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "value": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "preference",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/preferences",
                "self": "https://192.168.1.16:8443/v3/schemas/preference"
            },
            "pluralName": "preferences",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "namespaceId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[namespace]",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "value": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "displayName": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "loginName": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "me": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "memberOf": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "profilePicture": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "profileURL": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "principal",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/principals",
                "self": "https://192.168.1.16:8443/v3/schemas/principal"
            },
            "pluralName": "principals",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "displayName": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "extraInfo": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "loginName": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "me": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "memberOf": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "profilePicture": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "profileURL": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "clusterId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "name": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "namespaceId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "podSecurityPolicyTemplateId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "project",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/projects",
                "self": "https://192.168.1.16:8443/v3/schemas/project"
            },
            "pluralName": "projects",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "clusterId": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "reference[cluster]",
                    "update": true
                },
                "conditions": {
                    "create": false,
                    "nullable": true,
                    "type": "array[projectCondition]",
                    "update": false
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "id": {
                    "create": false,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "namespaceId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[namespace]",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "podSecurityPolicyTemplateId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[podSecurityPolicyTemplate]",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "displayName": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "initialWaitSeconds": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "namespaceId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "projectId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "repeatIntervalSeconds": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "severity": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "projectAlert",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/projectalerts",
                "self": "https://192.168.1.16:8443/v3/schemas/projectAlert"
            },
            "pluralName": "projectAlerts",
            "resourceActions": {
                "activate": {},
                "deactivate": {},
                "mute": {},
                "unmute": {}
            },
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "displayName": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "initialWaitSeconds": {
                    "create": true,
                    "default": "180",
                    "min": 0,
                    "nullable": true,
                    "required": true,
                    "type": "int",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "namespaceId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[namespace]",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "projectId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[project]",
                    "update": true
                },
                "recipientList": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "array[recipient]",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "repeatIntervalSeconds": {
                    "create": true,
                    "default": "3600",
                    "min": 0,
                    "nullable": true,
                    "required": true,
                    "type": "int",
                    "update": true
                },
                "severity": {
                    "create": true,
                    "default": "critical",
                    "nullable": true,
                    "options": [
                        "info",
                        "critical",
                        "warning"
                    ],
                    "required": true,
                    "type": "enum",
                    "update": true
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "status": {
                    "create": false,
                    "nullable": true,
                    "type": "alertStatus",
                    "update": false
                },
                "targetPod": {
                    "create": true,
                    "nullable": true,
                    "type": "targetPod",
                    "update": true
                },
                "targetWorkload": {
                    "create": true,
                    "nullable": true,
                    "type": "targetWorkload",
                    "update": true
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "projectCondition",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/projectCondition"
            },
            "pluralName": "projectConditions",
            "resourceFields": {
                "lastTransitionTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "lastUpdateTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "message": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "reason": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "status": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "type": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "displayName": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "namespaceId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "outputFlushInterval": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "projectId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "projectLogging",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/projectloggings",
                "self": "https://192.168.1.16:8443/v3/schemas/projectLogging"
            },
            "pluralName": "projectLoggings",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "displayName": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "elasticsearchConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "elasticsearchConfig",
                    "update": true
                },
                "kafkaConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "kafkaConfig",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "namespaceId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[namespace]",
                    "update": false
                },
                "outputFlushInterval": {
                    "create": true,
                    "default": "3",
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "outputTags": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "projectId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[project]",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "splunkConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "splunkConfig",
                    "update": true
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "status": {
                    "create": false,
                    "nullable": true,
                    "type": "loggingStatus",
                    "update": false
                },
                "syslogConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "syslogConfig",
                    "update": true
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "namespaceId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "projectId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "roleTemplateId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "userId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "projectRoleTemplateBinding",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/projectroletemplatebindings",
                "self": "https://192.168.1.16:8443/v3/schemas/projectRoleTemplateBinding"
            },
            "pluralName": "projectRoleTemplateBindings",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "namespaceId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[namespace]",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "projectId": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "reference[project]",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "roleTemplateId": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "reference[roleTemplate]",
                    "update": true
                },
                "userId": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "reference[user]",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "question",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/question"
            },
            "pluralName": "questions",
            "resourceFields": {
                "default": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "group": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "invalidChars": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "label": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "max": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "maxLength": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "min": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "minLength": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "options": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "required": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "type": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "validChars": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "variable": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "rancherKubernetesEngineConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/rancherKubernetesEngineConfig"
            },
            "pluralName": "rancherKubernetesEngineConfigs",
            "resourceFields": {
                "addons": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "authentication": {
                    "create": true,
                    "nullable": true,
                    "type": "authnConfig",
                    "update": true
                },
                "authorization": {
                    "create": true,
                    "nullable": true,
                    "type": "authzConfig",
                    "update": true
                },
                "ignoreDockerVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "network": {
                    "create": true,
                    "nullable": true,
                    "type": "networkConfig",
                    "update": true
                },
                "nodes": {
                    "create": true,
                    "nullable": true,
                    "type": "array[rkeConfigNode]",
                    "update": true
                },
                "services": {
                    "create": true,
                    "nullable": true,
                    "type": "rkeConfigServices",
                    "update": true
                },
                "sshKeyPath": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "systemImages": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "version": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "recipient",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/recipient"
            },
            "pluralName": "recipients",
            "resourceFields": {
                "customPagerdutyConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "pagerdutyConfig",
                    "update": true
                },
                "customWebhookConfig": {
                    "create": true,
                    "nullable": true,
                    "type": "webhookConfig",
                    "update": true
                },
                "notifierId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[notifier]",
                    "update": true
                },
                "recipient": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "releaseInfo",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/releaseInfo"
            },
            "pluralName": "releaseInfos",
            "resourceFields": {
                "createTimestamp": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "modifiedAt": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "templateVersionId": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "version": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "rkeConfigNode",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/rkeConfigNode"
            },
            "pluralName": "rkeConfigNodes",
            "resourceFields": {
                "address": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "dockerSocket": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "hostnameOverride": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "internalAddress": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "machineId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[machine]",
                    "update": true
                },
                "role": {
                    "create": true,
                    "nullable": true,
                    "options": [
                        "etcd",
                        "worker",
                        "controlplane"
                    ],
                    "type": "array[enum]",
                    "update": true
                },
                "sshKey": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "sshKeyPath": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "user": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "rkeConfigServices",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/rkeConfigServices"
            },
            "pluralName": "rkeConfigServiceses",
            "resourceFields": {
                "etcd": {
                    "create": true,
                    "nullable": true,
                    "type": "etcdService",
                    "update": true
                },
                "kubeApi": {
                    "create": true,
                    "nullable": true,
                    "type": "kubeAPIService",
                    "update": true
                },
                "kubeController": {
                    "create": true,
                    "nullable": true,
                    "type": "kubeControllerService",
                    "update": true
                },
                "kubelet": {
                    "create": true,
                    "nullable": true,
                    "type": "kubeletService",
                    "update": true
                },
                "kubeproxy": {
                    "create": true,
                    "nullable": true,
                    "type": "kubeproxyService",
                    "update": true
                },
                "scheduler": {
                    "create": true,
                    "nullable": true,
                    "type": "schedulerService",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "builtin": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "context": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "external": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "hidden": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "name": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "roleTemplate",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/roletemplates",
                "self": "https://192.168.1.16:8443/v3/schemas/roleTemplate"
            },
            "pluralName": "roleTemplates",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "builtin": {
                    "create": false,
                    "nullable": true,
                    "type": "boolean",
                    "update": false
                },
                "context": {
                    "create": true,
                    "nullable": true,
                    "options": [
                        "project",
                        "cluster"
                    ],
                    "type": "string",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "external": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "hidden": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "id": {
                    "create": false,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "roleTemplateIds": {
                    "create": true,
                    "nullable": true,
                    "type": "array[reference[roleTemplate]]",
                    "update": true
                },
                "rules": {
                    "create": true,
                    "nullable": true,
                    "type": "array[policyRule]",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "runAsUserStrategyOptions",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/runAsUserStrategyOptions"
            },
            "pluralName": "runAsUserStrategyOptionses",
            "resourceFields": {
                "ranges": {
                    "create": true,
                    "nullable": true,
                    "type": "array[idRange]",
                    "update": true
                },
                "rule": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "schedulerService",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/schedulerService"
            },
            "pluralName": "schedulerServices",
            "resourceFields": {
                "extraArgs": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "image": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "seLinuxOptions",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/seLinuxOptions"
            },
            "pluralName": "seLinuxOptionses",
            "resourceFields": {
                "level": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "role": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "type": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "user": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "seLinuxStrategyOptions",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/seLinuxStrategyOptions"
            },
            "pluralName": "seLinuxStrategyOptionses",
            "resourceFields": {
                "rule": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "seLinuxOptions": {
                    "create": true,
                    "nullable": true,
                    "type": "seLinuxOptions",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "setPasswordInput",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/setPasswordInput"
            },
            "pluralName": "setPasswordInputs",
            "resourceFields": {
                "newPassword": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "customized": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "default": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "value": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "setting",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/settings",
                "self": "https://192.168.1.16:8443/v3/schemas/setting"
            },
            "pluralName": "settings",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "customized": {
                    "create": false,
                    "nullable": true,
                    "type": "boolean",
                    "update": false
                },
                "default": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "value": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "slackConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/slackConfig"
            },
            "pluralName": "slackConfigs",
            "resourceFields": {
                "url": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "smtpConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/smtpConfig"
            },
            "pluralName": "smtpConfigs",
            "resourceFields": {
                "defaultRecipient": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "host": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "dnsLabel",
                    "update": true
                },
                "password": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "masked",
                    "update": true
                },
                "port": {
                    "create": true,
                    "max": 65535,
                    "min": 1,
                    "nullable": true,
                    "required": true,
                    "type": "int",
                    "update": true
                },
                "tls": {
                    "create": true,
                    "default": "true",
                    "nullable": true,
                    "required": true,
                    "type": "boolean",
                    "update": true
                },
                "username": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "splunkConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/splunkConfig"
            },
            "pluralName": "splunkConfigs",
            "resourceFields": {
                "endpoint": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "source": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "token": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "externalId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "installNamespace": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "namespaceId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "projectId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "prune": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "user": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "stack",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/stacks",
                "self": "https://192.168.1.16:8443/v3/schemas/stack"
            },
            "pluralName": "stacks",
            "resourceActions": {
                "rollback": {
                    "input": "revision"
                },
                "upgrade": {
                    "input": "templateVersionId"
                }
            },
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "answers": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "externalId": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "groups": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "installNamespace": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "namespaceId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[namespace]",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "projectId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[project]",
                    "update": true
                },
                "prune": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "status": {
                    "create": false,
                    "nullable": true,
                    "type": "stackStatus",
                    "update": false
                },
                "tag": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "templates": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "user": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "stackStatus",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/stackStatus"
            },
            "pluralName": "stackStatuses",
            "resourceFields": {
                "releases": {
                    "create": false,
                    "nullable": true,
                    "type": "array[releaseInfo]",
                    "update": false
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "supplementalGroupsStrategyOptions",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/supplementalGroupsStrategyOptions"
            },
            "pluralName": "supplementalGroupsStrategyOptionses",
            "resourceFields": {
                "ranges": {
                    "create": true,
                    "nullable": true,
                    "type": "array[idRange]",
                    "update": true
                },
                "rule": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "syslogConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/syslogConfig"
            },
            "pluralName": "syslogConfigs",
            "resourceFields": {
                "endpoint": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "program": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "severity": {
                    "create": true,
                    "default": "notice",
                    "nullable": true,
                    "options": [
                        "emerg",
                        "alert",
                        "crit",
                        "err",
                        "warning",
                        "notice",
                        "info",
                        "debug"
                    ],
                    "type": "enum",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "taint",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/taint"
            },
            "pluralName": "taints",
            "resourceFields": {
                "effect": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "key": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "timeAdded": {
                    "create": true,
                    "nullable": true,
                    "type": "date",
                    "update": true
                },
                "value": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "targetNode",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/targetNode"
            },
            "pluralName": "targetNodes",
            "resourceFields": {
                "cpuThreshold": {
                    "create": true,
                    "min": 1,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "diskThreshold": {
                    "create": true,
                    "max": 100,
                    "min": 1,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "id": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "isReady": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "memThreshold": {
                    "create": true,
                    "max": 100,
                    "min": 1,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "selector": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "targetPod",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/targetPod"
            },
            "pluralName": "targetPods",
            "resourceFields": {
                "id": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                },
                "isRunning": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "isScheduled": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "restartTimes": {
                    "create": true,
                    "min": 1,
                    "nullable": true,
                    "type": "int",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "targetSystemService",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/targetSystemService"
            },
            "pluralName": "targetSystemServices",
            "resourceFields": {
                "type": {
                    "create": true,
                    "default": "scheduler",
                    "nullable": true,
                    "options": [
                        "dns",
                        "etcd",
                        "controller",
                        "manager",
                        "network",
                        "scheduler"
                    ],
                    "required": true,
                    "type": "enum",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "targetWorkload",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/targetWorkload"
            },
            "pluralName": "targetWorkloads",
            "resourceFields": {
                "id": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "selector": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "unavailablePercentage": {
                    "create": true,
                    "max": 100,
                    "min": 1,
                    "nullable": true,
                    "required": true,
                    "type": "int",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "catalogId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "category": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "defaultTemplateVersionId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "defaultVersion": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "folderName": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "icon": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "iconFilename": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "isSystem": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "license": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "maintainer": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "path": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "projectURL": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "readme": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "templateBase": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "upgradeFrom": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "template",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/templates",
                "self": "https://192.168.1.16:8443/v3/schemas/template"
            },
            "pluralName": "templates",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "catalogId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[catalog]",
                    "update": true
                },
                "categories": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "category": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "defaultTemplateVersionId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[templateVersion]",
                    "update": true
                },
                "defaultVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "folderName": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "icon": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "iconFilename": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "isSystem": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "license": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "maintainer": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "path": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "projectURL": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "readme": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "status": {
                    "create": false,
                    "nullable": true,
                    "type": "templateStatus",
                    "update": false
                },
                "templateBase": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "upgradeFrom": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "versions": {
                    "create": true,
                    "nullable": true,
                    "type": "array[templateVersionSpec]",
                    "update": true
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "templateStatus",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/templateStatus"
            },
            "pluralName": "templateStatuses",
            "resourceFields": {},
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionFilters": {
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "externalId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "maximumRancherVersion": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "minimumRancherVersion": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "readme": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "revision": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "state": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioning": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "transitioningMessage": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "upgradeFrom": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "version": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "templateVersion",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/templateversions",
                "self": "https://192.168.1.16:8443/v3/schemas/templateVersion"
            },
            "pluralName": "templateVersions",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "externalId": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "files": {
                    "create": true,
                    "nullable": true,
                    "type": "array[file]",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "maximumRancherVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "minimumRancherVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "questions": {
                    "create": true,
                    "nullable": true,
                    "type": "array[question]",
                    "update": true
                },
                "readme": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "revision": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "state": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "status": {
                    "create": false,
                    "nullable": true,
                    "type": "templateVersionStatus",
                    "update": false
                },
                "transitioning": {
                    "create": false,
                    "options": [
                        "yes",
                        "no",
                        "error"
                    ],
                    "type": "enum",
                    "update": false
                },
                "transitioningMessage": {
                    "create": false,
                    "type": "string",
                    "update": false
                },
                "upgradeFrom": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "upgradeVersionLinks": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                },
                "version": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "templateVersionSpec",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/templateVersionSpec"
            },
            "pluralName": "templateVersionSpecs",
            "resourceFields": {
                "externalId": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "files": {
                    "create": true,
                    "nullable": true,
                    "type": "array[file]",
                    "update": true
                },
                "maximumRancherVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "minimumRancherVersion": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "questions": {
                    "create": true,
                    "nullable": true,
                    "type": "array[question]",
                    "update": true
                },
                "readme": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "revision": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "upgradeFrom": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "upgradeVersionLinks": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "version": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "templateVersionStatus",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/templateVersionStatus"
            },
            "pluralName": "templateVersionStatuses",
            "resourceFields": {},
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionActions": {
                "login": {
                    "input": "loginInput",
                    "output": "token"
                },
                "logout": {}
            },
            "collectionFilters": {
                "authProvider": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "isDerived": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "lastUpdateTime": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "token": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "ttl": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "userId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "userPrincipal": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "token",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/tokens",
                "self": "https://192.168.1.16:8443/v3/schemas/token"
            },
            "pluralName": "tokens",
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "authProvider": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "groupPrincipals": {
                    "create": true,
                    "nullable": true,
                    "type": "array[reference[Principal]]",
                    "update": true
                },
                "isDerived": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "lastUpdateTime": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "providerInfo": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "token": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": false,
                    "writeOnly": true
                },
                "ttl": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "userId": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[User]",
                    "update": true
                },
                "userPrincipal": {
                    "create": true,
                    "nullable": true,
                    "type": "reference[Principal]",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "collectionActions": {
                "changepassword": {
                    "input": "changePasswordInput"
                }
            },
            "collectionFilters": {
                "creatorId": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "description": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "me": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "mustChangePassword": {
                    "modifiers": [
                        "eq",
                        "ne"
                    ]
                },
                "name": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "password": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "username": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                },
                "uuid": {
                    "modifiers": [
                        "eq",
                        "ne",
                        "in",
                        "notin"
                    ]
                }
            },
            "collectionMethods": [
                "POST",
                "GET"
            ],
            "id": "user",
            "links": {
                "collection": "https://192.168.1.16:8443/v3/users",
                "self": "https://192.168.1.16:8443/v3/schemas/user"
            },
            "pluralName": "users",
            "resourceActions": {
                "setpassword": {
                    "input": "setPasswordInput",
                    "output": "user"
                }
            },
            "resourceFields": {
                "annotations": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "created": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "creatorId": {
                    "create": false,
                    "type": "reference[user]",
                    "update": false
                },
                "description": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "id": {
                    "create": false,
                    "nullable": true,
                    "type": "dnsLabel",
                    "update": false
                },
                "labels": {
                    "create": true,
                    "nullable": true,
                    "type": "map[string]",
                    "update": true
                },
                "me": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "mustChangePassword": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "name": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "ownerReferences": {
                    "create": false,
                    "nullable": true,
                    "type": "array[ownerReference]",
                    "update": false
                },
                "password": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": false,
                    "writeOnly": true
                },
                "principalIds": {
                    "create": true,
                    "nullable": true,
                    "type": "array[reference[Principal]]",
                    "update": true
                },
                "removed": {
                    "create": false,
                    "nullable": true,
                    "type": "date",
                    "update": false
                },
                "username": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                },
                "uuid": {
                    "create": false,
                    "nullable": true,
                    "type": "string",
                    "update": false
                }
            },
            "resourceMethods": [
                "PUT",
                "DELETE"
            ],
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "values",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/values"
            },
            "pluralName": "valueses",
            "resourceFields": {
                "boolValue": {
                    "create": true,
                    "nullable": true,
                    "type": "boolean",
                    "update": true
                },
                "intValue": {
                    "create": true,
                    "nullable": true,
                    "type": "int",
                    "update": true
                },
                "stringSliceValue": {
                    "create": true,
                    "nullable": true,
                    "type": "array[string]",
                    "update": true
                },
                "stringValue": {
                    "create": true,
                    "nullable": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "vmwarevsphereconfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/vmwarevsphereconfig"
            },
            "pluralName": "vmwarevsphereconfigs",
            "resourceFields": {
                "boot2dockerUrl": {
                    "create": true,
                    "default": "",
                    "description": "vSphere URL for boot2docker image",
                    "type": "string",
                    "update": false
                },
                "cfgparam": {
                    "create": true,
                    "description": "vSphere vm configuration parameters (used for guestinfo)",
                    "type": "array[string]",
                    "update": false
                },
                "cloudinit": {
                    "create": true,
                    "default": "",
                    "description": "vSphere cloud-init file or url to set in the guestinfo",
                    "type": "string",
                    "update": false
                },
                "cpuCount": {
                    "create": true,
                    "default": "2",
                    "description": "vSphere CPU number for docker VM",
                    "type": "string",
                    "update": false
                },
                "datacenter": {
                    "create": true,
                    "default": "",
                    "description": "vSphere datacenter for docker VM",
                    "type": "string",
                    "update": false
                },
                "datastore": {
                    "create": true,
                    "default": "",
                    "description": "vSphere datastore for docker VM",
                    "type": "string",
                    "update": false
                },
                "diskSize": {
                    "create": true,
                    "default": "20480",
                    "description": "vSphere size of disk for docker VM (in MB)",
                    "type": "string",
                    "update": false
                },
                "hostsystem": {
                    "create": true,
                    "default": "",
                    "description": "vSphere compute resource where the docker VM will be instantiated. This can be omitted if using a cluster with DRS.",
                    "type": "string",
                    "update": false
                },
                "memorySize": {
                    "create": true,
                    "default": "2048",
                    "description": "vSphere size of memory for docker VM (in MB)",
                    "type": "string",
                    "update": false
                },
                "network": {
                    "create": true,
                    "description": "vSphere network where the docker VM will be attached",
                    "type": "array[string]",
                    "update": false
                },
                "password": {
                    "create": true,
                    "default": "",
                    "description": "vSphere password",
                    "type": "string",
                    "update": false
                },
                "pool": {
                    "create": true,
                    "default": "",
                    "description": "vSphere resource pool for docker VM",
                    "type": "string",
                    "update": false
                },
                "username": {
                    "create": true,
                    "default": "",
                    "description": "vSphere username",
                    "type": "string",
                    "update": false
                },
                "vcenter": {
                    "create": true,
                    "default": "",
                    "description": "vSphere IP/hostname for vCenter",
                    "type": "string",
                    "update": false
                },
                "vcenterPort": {
                    "create": true,
                    "default": "443",
                    "description": "vSphere Port for vCenter",
                    "type": "string",
                    "update": false
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "webhookConfig",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/webhookConfig"
            },
            "pluralName": "webhookConfigs",
            "resourceFields": {
                "url": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        },
        {
            "actions": {},
            "baseType": "schema",
            "id": "zookeeper",
            "links": {
                "self": "https://192.168.1.16:8443/v3/schemas/zookeeper"
            },
            "pluralName": "zookeepers",
            "resourceFields": {
                "endpoint": {
                    "create": true,
                    "nullable": true,
                    "required": true,
                    "type": "string",
                    "update": true
                }
            },
            "type": "schema",
            "version": {
                "group": "management.cattle.io",
                "path": "/v3",
                "subContext": {
                    "clusters": true
                },
                "version": "v3"
            }
        }
    ]
}
```
## License
Copyright (c) 2014-2016 [Rancher Labs, Inc.](http://rancher.com)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

[http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
