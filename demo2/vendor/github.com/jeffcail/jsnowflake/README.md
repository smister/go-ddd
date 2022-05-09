# jsnowflake

Go language generates distributed snowflake ID


# Install
```shell
go get github.com/jeffcail/jsnowflake 
```

# Example
```go
m, err := NewMachine(1) // 传入对应的机器ID
	if err != nil {
		panic(err)
	}
id := m.GenerateId()
```


