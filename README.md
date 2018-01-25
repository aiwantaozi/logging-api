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
   "type":"collection",
   "links":{
      "self":"http://localhost:1235/v3/schemas"
   },
   "actions":{

   },
   "pagination":{
      "limit":1000
   },
   "sort":{
      "order":"asc",
      "reverse":"http://localhost:1235/v3/schemas?order=desc"
   },
   "resourceType":"schema",
   "data":[
      {
         "actions":{

         },
         "baseType":"schema",
         "id":"elasticsearchConfig",
         "links":{
            "self":"http://localhost:1235/v3/schemas/elasticsearchConfig"
         },
         "pluralName":"elasticsearchConfigs",
         "resourceFields":{
            "authPassword":{
               "create":true,
               "nullable":true,
               "type":"string",
               "update":true
            },
            "authUser":{
               "create":true,
               "nullable":true,
               "type":"string",
               "update":true
            },
            "dateformat":{
               "create":true,
               "default":"YYYY.MM.DD",
               "nullable":true,
               "options":[
                  "YYYY.MM.DD",
                  "YYYY.MM",
                  "YYYY"
               ],
               "type":"enum",
               "update":true
            },
            "host":{
               "create":true,
               "nullable":true,
               "type":"string",
               "update":true
            },
            "indexPrefix":{
               "create":true,
               "nullable":true,
               "type":"string",
               "update":true
            },
            "port":{
               "create":true,
               "default":9200,
               "nullable":true,
               "type":"int",
               "update":true
            }
         },
         "type":"schema",
         "version":{
            "group":"logging.cattle.io",
            "path":"/v3",
            "version":"v3"
         }
      },
      {
         "actions":{

         },
         "baseType":"schema",
         "id":"embeddedConfig",
         "links":{
            "self":"http://localhost:1235/v3/schemas/embeddedConfig"
         },
         "pluralName":"embeddedConfigs",
         "resourceFields":{
            "elasticsearchAccessURL":{
               "create":true,
               "nullable":true,
               "type":"string",
               "update":true
            },
            "kibanaAccessURL":{
               "create":true,
               "nullable":true,
               "type":"string",
               "update":true
            }
         },
         "type":"schema",
         "version":{
            "group":"logging.cattle.io",
            "path":"/v3",
            "version":"v3"
         }
      },
      {
         "actions":{

         },
         "baseType":"schema",
         "id":"kafkaConfig",
         "links":{
            "self":"http://localhost:1235/v3/schemas/kafkaConfig"
         },
         "pluralName":"kafkaConfigs",
         "resourceFields":{
            "brokerType":{
               "create":true,
               "default":"zookeeper",
               "nullable":true,
               "options":[
                  "broker",
                  "zookeeper"
               ],
               "type":"enum",
               "update":true
            },
            "brokers":{
               "create":true,
               "nullable":true,
               "type":"array[string]",
               "update":true
            },
            "dataType":{
               "create":true,
               "default":"json",
               "nullable":true,
               "options":[
                  "json",
                  "ltsv",
                  "msgpack"
               ],
               "type":"enum",
               "update":true
            },
            "maxSendRetries":{
               "create":true,
               "default":1,
               "nullable":true,
               "type":"int",
               "update":true
            },
            "topic":{
               "create":true,
               "nullable":true,
               "type":"string",
               "update":true
            },
            "zookeeperHost":{
               "create":true,
               "nullable":true,
               "type":"string",
               "update":true
            },
            "zookeeperPort":{
               "create":true,
               "default":2181,
               "nullable":true,
               "type":"int",
               "update":true
            }
         },
         "type":"schema",
         "version":{
            "group":"logging.cattle.io",
            "path":"/v3",
            "version":"v3"
         }
      },
      {
         "actions":{

         },
         "baseType":"schema",
         "collectionFilters":{
            "clusterId":{
               "modifiers":[
                  "eq",
                  "ne",
                  "in",
                  "notin"
               ]
            },
            "creatorId":{
               "modifiers":[
                  "eq",
                  "ne",
                  "in",
                  "notin"
               ]
            },
            "currentTarget":{
               "modifiers":[
                  "eq",
                  "ne",
                  "in",
                  "notin"
               ]
            },
            "displayName":{
               "modifiers":[
                  "eq",
                  "ne",
                  "in",
                  "notin"
               ]
            },
            "outputFlushInterval":{
               "modifiers":[
                  "eq",
                  "ne",
                  "in",
                  "notin"
               ]
            },
            "state":{
               "modifiers":[
                  "eq",
                  "ne",
                  "in",
                  "notin"
               ]
            },
            "transitioning":{
               "modifiers":[
                  "eq",
                  "ne",
                  "in",
                  "notin"
               ]
            },
            "transitioningMessage":{
               "modifiers":[
                  "eq",
                  "ne",
                  "in",
                  "notin"
               ]
            },
            "uuid":{
               "modifiers":[
                  "eq",
                  "ne",
                  "in",
                  "notin"
               ]
            }
         },
         "collectionMethods":[
            "POST",
            "GET"
         ],
         "id":"logging",
         "links":{
            "collection":"http://localhost:1235/v3/loggings",
            "self":"http://localhost:1235/v3/schemas/logging"
         },
         "pluralName":"loggings",
         "resourceFields":{
            "annotations":{
               "create":true,
               "nullable":true,
               "type":"map[string]",
               "update":true
            },
            "clusterId":{
               "create":true,
               "nullable":true,
               "type":"reference[cluster]",
               "update":true
            },
            "created":{
               "create":false,
               "nullable":true,
               "type":"date",
               "update":false
            },
            "creatorId":{
               "create":false,
               "type":"reference[user]",
               "update":false
            },
            "currentTarget":{
               "create":true,
               "nullable":true,
               "type":"string",
               "update":true
            },
            "displayName":{
               "create":true,
               "nullable":true,
               "type":"string",
               "update":true
            },
            "elasticsearchConfig":{
               "create":true,
               "nullable":true,
               "type":"elasticsearchConfig",
               "update":true
            },
            "embeddedConfig":{
               "create":true,
               "nullable":true,
               "type":"embeddedConfig",
               "update":true
            },
            "kafkaConfig":{
               "create":true,
               "nullable":true,
               "type":"kafkaConfig",
               "update":true
            },
            "labels":{
               "create":true,
               "nullable":true,
               "type":"map[string]",
               "update":true
            },
            "name":{
               "create":true,
               "nullable":true,
               "type":"dnsLabel",
               "update":false
            },
            "outputFlushInterval":{
               "create":true,
               "nullable":true,
               "type":"int",
               "update":true
            },
            "outputTags":{
               "create":true,
               "nullable":true,
               "type":"map[string]",
               "update":true
            },
            "ownerReferences":{
               "create":false,
               "nullable":true,
               "type":"array[ownerReference]",
               "update":false
            },
            "removed":{
               "create":false,
               "nullable":true,
               "type":"date",
               "update":false
            },
            "splunkConfig":{
               "create":true,
               "nullable":true,
               "type":"splunkConfig",
               "update":true
            },
            "state":{
               "create":false,
               "type":"string",
               "update":false
            },
            "status":{
               "create":false,
               "nullable":true,
               "type":"loggingStatus",
               "update":false
            },
            "syslogConfig":{
               "create":true,
               "nullable":true,
               "type":"syslogConfig",
               "update":true
            },
            "transitioning":{
               "create":false,
               "options":[
                  "yes",
                  "no",
                  "error"
               ],
               "type":"enum",
               "update":false
            },
            "transitioningMessage":{
               "create":false,
               "type":"string",
               "update":false
            },
            "uuid":{
               "create":false,
               "nullable":true,
               "type":"string",
               "update":false
            }
         },
         "resourceMethods":[
            "PUT",
            "DELETE"
         ],
         "type":"schema",
         "version":{
            "group":"logging.cattle.io",
            "path":"/v3",
            "version":"v3"
         }
      },
      {
         "actions":{

         },
         "baseType":"schema",
         "id":"loggingStatus",
         "links":{
            "self":"http://localhost:1235/v3/schemas/loggingStatus"
         },
         "pluralName":"loggingStatuses",
         "resourceFields":{
            "currentTarget":{
               "create":false,
               "nullable":true,
               "type":"string",
               "update":false
            }
         },
         "type":"schema",
         "version":{
            "group":"logging.cattle.io",
            "path":"/v3",
            "version":"v3"
         }
      },
      {
         "actions":{

         },
         "baseType":"schema",
         "id":"ownerReference",
         "links":{
            "self":"http://localhost:1235/v3/schemas/ownerReference"
         },
         "pluralName":"ownerReferences",
         "resourceFields":{
            "apiVersion":{
               "create":true,
               "nullable":true,
               "type":"string",
               "update":true
            },
            "blockOwnerDeletion":{
               "create":true,
               "nullable":true,
               "type":"boolean",
               "update":true
            },
            "controller":{
               "create":true,
               "nullable":true,
               "type":"boolean",
               "update":true
            },
            "kind":{
               "create":true,
               "nullable":true,
               "type":"string",
               "update":true
            },
            "name":{
               "create":true,
               "nullable":true,
               "type":"string",
               "update":true
            },
            "uid":{
               "create":true,
               "nullable":true,
               "type":"string",
               "update":true
            }
         },
         "type":"schema",
         "version":{
            "group":"logging.cattle.io",
            "path":"/v3",
            "version":"v3"
         }
      },
      {
         "actions":{

         },
         "baseType":"schema",
         "collectionFilters":{
            "creatorId":{
               "modifiers":[
                  "eq",
                  "ne",
                  "in",
                  "notin"
               ]
            },
            "currentTarget":{
               "modifiers":[
                  "eq",
                  "ne",
                  "in",
                  "notin"
               ]
            },
            "displayName":{
               "modifiers":[
                  "eq",
                  "ne",
                  "in",
                  "notin"
               ]
            },
            "namespaceId":{
               "modifiers":[
                  "eq",
                  "ne",
                  "in",
                  "notin"
               ]
            },
            "outputFlushInterval":{
               "modifiers":[
                  "eq",
                  "ne",
                  "in",
                  "notin"
               ]
            },
            "projectId":{
               "modifiers":[
                  "eq",
                  "ne",
                  "in",
                  "notin"
               ]
            },
            "state":{
               "modifiers":[
                  "eq",
                  "ne",
                  "in",
                  "notin"
               ]
            },
            "transitioning":{
               "modifiers":[
                  "eq",
                  "ne",
                  "in",
                  "notin"
               ]
            },
            "transitioningMessage":{
               "modifiers":[
                  "eq",
                  "ne",
                  "in",
                  "notin"
               ]
            },
            "uuid":{
               "modifiers":[
                  "eq",
                  "ne",
                  "in",
                  "notin"
               ]
            }
         },
         "collectionMethods":[
            "POST",
            "GET"
         ],
         "id":"projectLogging",
         "links":{
            "collection":"http://localhost:1235/v3/projectloggings",
            "self":"http://localhost:1235/v3/schemas/projectLogging"
         },
         "pluralName":"projectLoggings",
         "resourceFields":{
            "annotations":{
               "create":true,
               "nullable":true,
               "type":"map[string]",
               "update":true
            },
            "created":{
               "create":false,
               "nullable":true,
               "type":"date",
               "update":false
            },
            "creatorId":{
               "create":false,
               "type":"reference[user]",
               "update":false
            },
            "currentTarget":{
               "create":true,
               "nullable":true,
               "type":"string",
               "update":true
            },
            "displayName":{
               "create":true,
               "nullable":true,
               "type":"string",
               "update":true
            },
            "elasticsearchConfig":{
               "create":true,
               "nullable":true,
               "type":"elasticsearchConfig",
               "update":true
            },
            "kafkaConfig":{
               "create":true,
               "nullable":true,
               "type":"kafkaConfig",
               "update":true
            },
            "labels":{
               "create":true,
               "nullable":true,
               "type":"map[string]",
               "update":true
            },
            "name":{
               "create":true,
               "nullable":true,
               "type":"dnsLabel",
               "update":false
            },
            "namespaceId":{
               "create":true,
               "nullable":true,
               "required":true,
               "type":"reference[namespace]",
               "update":false
            },
            "outputFlushInterval":{
               "create":true,
               "nullable":true,
               "type":"int",
               "update":true
            },
            "outputTags":{
               "create":true,
               "nullable":true,
               "type":"map[string]",
               "update":true
            },
            "ownerReferences":{
               "create":false,
               "nullable":true,
               "type":"array[ownerReference]",
               "update":false
            },
            "projectId":{
               "create":true,
               "nullable":true,
               "type":"reference[project]",
               "update":true
            },
            "removed":{
               "create":false,
               "nullable":true,
               "type":"date",
               "update":false
            },
            "splunkConfig":{
               "create":true,
               "nullable":true,
               "type":"splunkConfig",
               "update":true
            },
            "state":{
               "create":false,
               "type":"string",
               "update":false
            },
            "status":{
               "create":false,
               "nullable":true,
               "type":"loggingStatus",
               "update":false
            },
            "syslogConfig":{
               "create":true,
               "nullable":true,
               "type":"syslogConfig",
               "update":true
            },
            "transitioning":{
               "create":false,
               "options":[
                  "yes",
                  "no",
                  "error"
               ],
               "type":"enum",
               "update":false
            },
            "transitioningMessage":{
               "create":false,
               "type":"string",
               "update":false
            },
            "uuid":{
               "create":false,
               "nullable":true,
               "type":"string",
               "update":false
            }
         },
         "resourceMethods":[
            "PUT",
            "DELETE"
         ],
         "type":"schema",
         "version":{
            "group":"logging.cattle.io",
            "path":"/v3",
            "version":"v3"
         }
      },
      {
         "actions":{

         },
         "baseType":"schema",
         "id":"splunkConfig",
         "links":{
            "self":"http://localhost:1235/v3/schemas/splunkConfig"
         },
         "pluralName":"splunkConfigs",
         "resourceFields":{
            "host":{
               "create":true,
               "nullable":true,
               "type":"string",
               "update":true
            },
            "port":{
               "create":true,
               "default":8088,
               "nullable":true,
               "type":"int",
               "update":true
            },
            "protocol":{
               "create":true,
               "default":"http",
               "nullable":true,
               "options":[
                  "http",
                  "https"
               ],
               "type":"enum",
               "update":true
            },
            "source":{
               "create":true,
               "nullable":true,
               "type":"string",
               "update":true
            },
            "timeFormat":{
               "create":true,
               "default":"unixtime",
               "nullable":true,
               "options":[
                  "unixtime",
                  "localtime",
                  "none"
               ],
               "type":"enum",
               "update":true
            },
            "token":{
               "create":true,
               "nullable":true,
               "type":"string",
               "update":true
            }
         },
         "type":"schema",
         "version":{
            "group":"logging.cattle.io",
            "path":"/v3",
            "version":"v3"
         }
      },
      {
         "actions":{

         },
         "baseType":"schema",
         "id":"syslogConfig",
         "links":{
            "self":"http://localhost:1235/v3/schemas/syslogConfig"
         },
         "pluralName":"syslogConfigs",
         "resourceFields":{
            "host":{
               "create":true,
               "nullable":true,
               "type":"string",
               "update":true
            },
            "port":{
               "create":true,
               "default":51400,
               "nullable":true,
               "type":"int",
               "update":true
            },
            "program":{
               "create":true,
               "nullable":true,
               "type":"string",
               "update":true
            },
            "severity":{
               "create":true,
               "default":"notice",
               "nullable":true,
               "options":[
                  "emerg",
                  "alert",
                  "crit",
                  "err",
                  "warning",
                  "notice",
                  "info",
                  "debug"
               ],
               "type":"enum",
               "update":true
            }
         },
         "type":"schema",
         "version":{
            "group":"logging.cattle.io",
            "path":"/v3",
            "version":"v3"
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
