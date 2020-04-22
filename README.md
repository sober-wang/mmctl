# mmctl 

You can use mmctl display appid's at Mesos + Marathon

# building mmctl 
ENV
- Go version: 1.13.5


```shell
git clone https://github.com/sober-wang/mmctl.git

cd mmctl 

go build .

chmod +x mmctl 
```

# How to use mmctl

Show help information

```shell
./mmctl --help
```

Show Marathon appid's
```shell
./mmctl get appid --all-appid
```

Show framework on Mesos
```shell
./mmctl get framework 
```

Show appid from a specific framework
```shell
./mmctl get appid -m <YOUR_FRAMEWORK_NAME>

# Maby you want see a appid containers from framework 
./mmctl get appid -m <YOUR_FRAMEWORK_NAME> <APPID_NAME>
```

Show Mesos node's 
```shell
./mmctl get node
```
