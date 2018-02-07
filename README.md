####安装步骤：

- go get govendor

    ` $ go get github.com/kardianos/govendor`
    
- get project

    `$ git clone https://github.com/xiaomeng79/gin_log.git`
    
- govendor sync

    `$ cd src/github.com/xiaomeng79/gin_log/ && govendor sync`
    
- go install

    `cd ../../../ && go install github.com/xiaomeng79/gin_log/`
    
- config project

    `cp bin/conf/server.json.bak bin/conf/server.json`
    
    `vi server.json`
    
- start project

    `nohup gin_log >> output.log 2>&1 &`

