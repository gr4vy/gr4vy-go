#!/bin/bash  
file="api/model_card.go"
var1='var environment ENVIRONMENT = "production"'
rep1='var environment string = "production"'
sed -i '.bak' "s/$var1/$rep1/g" "$file"
rm "$file".bak

file="api/model_pay_pal.go"
sed -i '.bak' "s/$var1/$rep1/g" "$file"
rm "$file".bak