## How to start

### By native GO command
- Run terminal command `go run .` at root-level of project directory

### By PM2
- Install PM2 here: https://pm2.keymetrics.io/
- Execute this command
```pm2 start "go run ." --watch --name user -o ./out.log -e ./out.log```
