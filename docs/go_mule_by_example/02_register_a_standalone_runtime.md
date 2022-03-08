# Register a go-mule

### 1. Register a go-mule

```
go-mule create -H <hybrid-token> <server-name>
```

Expected output:

```
Standalone created successfully [Id: '1003071' Certificate: './1003071/1003071.pem' Private Key: './1003071/1003071.key' CA Certificate: './1003071/ca.pem']
```

### 2. Inspect your current directory

```
ls -ld
```

Expected output:

```
drwxr-xr-x  5 foobar  staff  160 Mar  7 09:11 1003071
```

### 3. Inspect the runtime directory

```
cd 1003071
ls -la
```

Expected output:

```
-rw-r--r--   1 foobar  staff  1667 Mar  7 09:11 1003071.key
-rw-r--r--   1 foobar  staff  1155 Mar  7 09:11 1003071.pem
-rw-r--r--   1 foobar  staff  1383 Mar  7 09:11 ca.pem
```
