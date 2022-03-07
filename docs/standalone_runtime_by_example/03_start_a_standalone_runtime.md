# Start a standalone-runtime

### 1. Start the runtime

```
go-mule start 
```

Expected output

```
Runtime Configuration - Flows Per Application [total: '10']
Runtime Configuration - Schedulers Per Application [fixedFrequency: '5' cron: '5']
...
```

### 2. Start a runtime with dvery deployed application with one hundred flows

```
go-mule start --flows-per-app 100
```
