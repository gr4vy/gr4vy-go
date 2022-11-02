var1="return v.value.MarshalJSON()"
rep1="if v.value == nil { \n\t\t return []byte("'"null"'"), nil \n\t } \n\t return v.value.MarshalJSON()"
sed -i '' "s/$var1/$rep1/g" ./api/utils.go
