初始化对象
```golang
type User struct {
    ID uint `json:"id"`
    Name string `json:"name"
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) Init(name string) {
    u.ID = 1
    u.Name = name
    u.CreatedAt = time.Now()
    u.UpdatedAt = time.Now()
}
// 初始化对象
instance := reflection.New(&User{})
```
获取所有字段
```golang
fields := instance.Fields
```
检查是否含有指定字段
```golang
exists := instance.HasField(fieldName)
```
获取指定字段对象
```golang
field := instance.Fields[fieldName]
```
检查字段是否匿名
```golang
field.IsAnonymous()
```
检查字段是否合法
```golang
field.IsValid()
```
检查字段是否可被赋值
```golang
field.CanSet()
```
获取指定字段值 
```golang
fieldValue := instance.Fields[fieldName].Value
fieldValue, err := instance.Get(fieldName)
fieldValue := instance.MustGet(fieldName)
```
设置指定字段值
```golang
err := instance.Set(fieldName, value)
instance.MustSet(fieldName, value)
err := instance.Fields[fieldName].Set(value)
instance.Fields[fieldName].MustSet(value)
```
获取字段标签值
```golang
tagValue, err := instance.GetTag(fieldName, tagName)
tagValue := instance.MustGetTag(fieldName, tagName)
tagValue := instance.Fields[fieldName].GetTag(tagName)
tagValue, ok := instance.Fields[fieldName].LookUpTag(tagName)
```
获取所有方法
```golang
methods := instance.Methods
```
检查方法是否存在
```golang
exists := instance.HasMethod(methodName)
```
获取单个方法对象
```golang
method := instance.Methods[methodName]
```
检查方法是否导出
```golang
method.IsExported()
```
调用指定方法
```golang
results, err := instance.Call(methodName, param1, param2...)
results := instance.MustCall(methodName, param1, param2...)
results, err := instance.Methods[methodName].Call(param1, param2...)
results := instance.Methods[methodName].MustCall(param1, param2...)
```