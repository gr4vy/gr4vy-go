<?php
$FILENAME = 'api/model_card.go';

$FIND = <<<EOF
var environment ENVIRONMENT = "production"
EOF;

$str=file_get_contents($FILENAME);
$str=str_replace($FIND, 'var environment string = "production"',$str);

file_put_contents($FILENAME, $str);

$FILENAME = 'api/model_pay_pal.go';

$str=file_get_contents($FILENAME);
$str=str_replace($FIND, 'var environment string = "production"',$str);

//write the entire string
file_put_contents($FILENAME, $str);

?>