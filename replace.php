<?php
$FILENAME = 'openapi.v1.json';

$FIND = <<<EOF
            "enum": [
              "==",
              "!=",
              "<=",
              ">=",
              ">",
              "<",
              "IN",
              "NOT IN"
            ]
EOF;

$str=file_get_contents($FILENAME);
$str=str_replace($FIND, "",$str);

$FIND = <<<EOF
            "enum": [
              "==",
              "!=",
              "<=",
              ">=",
              ">",
              "<"
            ]
EOF;

$str=str_replace($FIND, "",$str);

$FIND = <<<EOF
                  "enum": [
                    "text",
                    "multiline",
                    "number"
                  ]
EOF;

$str=str_replace($FIND, "",$str);

//write the entire string
file_put_contents($FILENAME, $str);

?>