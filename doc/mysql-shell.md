# MYSQL-SHELL

## Connect to mysql server
```text
mysql-shell.mysqlsh --mysqlx -u docmguser -h 172.17.0.2 -P 33060
```
## Common command
```text
mysql-py []> session
<Session:docmguser@172.17.0.2:33060>
mysql-py []> shell
shell
mysql-py []> shell.status()
MySQL Shell version 8.0.20

Connection Id:                9
Default schema:               
Current schema:               
Current user:                 docmguser@172.17.0.1
SSL:                          Cipher in use: ECDHE-RSA-AES128-GCM-SHA256 TLSv1.2
Using delimiter:              ;
Server version:               8.0.19 MySQL Community Server - GPL
Protocol version:             X protocol
Client library:               8.0.20
Connection:                   172.17.0.2 via TCP/IP
TCP port:                     33060
Server characterset:          utf8mb4
Schema characterset:          utf8mb4
Client characterset:          utf8mb4
Conn. characterset:           utf8mb4
Compression:                  Enabled
Uptime:                       14 min 13.0000 sec
```
## Connect to global session
```text
mysql-py []> \connect mysqlx://docmguser@172.17.0.2:33060
```
OR create new scheme
```text
mysql-js> shell.connect( {scheme:'mysqlx', user:'user', host:'localhost', port:33060} )
```

## Use mysql shell to execute the content of file 
```text
mysqlsh --sql < exec_file.sql   
mysqlsh --py < exec_file.py
mysqlsh --js < exec_file.js
```

# Store JSON in mysql
## References:
```text
https://www.sitepoint.com/use-json-data-fields-mysql-databases/
```
## Create JSON field
```text
CREATE TABLE `book` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(200) NOT NULL,
  `tags` json DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;
```

## Aad JSON data
```text
INSERT INTO `book` (`title`, `tags`)
VALUES (
  'ECMAScript 2015: A SitePoint Anthology',
  '["JavaScript", "ES2015", "JSON"]'
);
```

## Create json array
```text
mysql> SELECT JSON_ARRAY(1, 2, 'abc');
+-------------------------+
| JSON_ARRAY(1, 2, 'abc') |
+-------------------------+
| [1, 2, "abc"]           |
+-------------------------+
```

## Create json object
```text
mysql> SELECT JSON_OBJECT('a', 1, 'b', 2);
+-----------------------------+
| JSON_OBJECT('a', 1, 'b', 2) |
+-----------------------------+
| {"a": 1, "b": 2}            |
+-----------------------------+
```

## Create merge document
```text
mysql> SELECT JSON_MERGE('["a", 1]', '{"key": "value"}');
+--------------------------------------------+
| JSON_MERGE('["a", 1]', '{"key": "value"}') |
+--------------------------------------------+
| ["a", 1, {"key": "value"}]                 |
+--------------------------------------------+
```

## Check type in json
```mysql> SELECT JSON_TYPE('[1, 2, "abc"]');
+----------------------------+
| JSON_TYPE('[1, 2, "abc"]') |
+----------------------------+
| ARRAY                      |
+----------------------------+
```

## Validate JSON
```text
mysql> SELECT JSON_VALID('[1, 2, "abc"]');
+-----------------------------+
| JSON_VALID('[1, 2, "abc"]') |
+-----------------------------+
|                           1 |
+-----------------------------+
```

## Search JSON data
```text
SELECT * FROM `book` 
WHERE JSON_CONTAINS(tags, '["JavaScript"]');

SELECT * FROM `book` 
WHERE JSON_SEARCH(tags, 'one', 'Java%') IS NOT NULL;
```

## JSON paths
```text
SELECT JSON_EXTRACT(
  '{"id": 1, "website": "SitePoint"}', 
  '$.website'
);
```

## Extract JSON paths in queries
```text
SELECT
  name,
  tags->"$[0]" AS `tag1`
FROM `book`;

SELECT
  name, profile->"$.twitter" AS `twitter`
FROM `user`;

SELECT
  name, profile->"$.twitter" AS `twitter`
FROM `user`;
```

## Modifying part of a JSON document 
```text
UPDATE `book`
  SET tags = JSON_MERGE(tags, '["technical"]')
WHERE
  JSON_SEARCH(tags, 'one', 'technical') IS NULL AND
  JSON_SEARCH(tags, 'one', 'JavaScript') IS NOT NULL;
```

## Search ChildID in mysql
```text
SELECT 
@path_to_uuid := JSON_UNQUOTE(JSON_SEARCH(directory_tree_work_space, 'one', '34f6a240-5fd0-4c19-b745-2ebcaea18b90')) AS path_to_uuid,
@path_to_parent := TRIM(TRAILING '.uuid' from @path_to_uuid) AS path_to_parent,
JSON_EXTRACT(directory_tree_work_space, @path_to_parent) AS result 
FROM projects where directory_tree_work_space like "%34f6a240-5fd0-4c19-b745-2ebcaea18b90%"\G
```