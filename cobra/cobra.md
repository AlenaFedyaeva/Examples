# How to use cobra (on MAC)

1) Создаем конфиг файл для Cobra
~/.cobra.yaml file:

```
author: Name SecondName <spf@spf13.com>
license: 
 header: This file is part of application.
  text: |
    {{ .copyright }}
useViper: false
```

2) Создаем новый пустой проект

```
mkdir trycobra
go mod init trycobra 
```

3) Использование Cobra
- Именование пакета рекомендуется делать в CamelCase
- Инициализацию клиента Cobra CLI рекомендуют сделать так

```
cobra init --pkg-name <app>
```
В моем случае делаю init в текущую директорию следующим образом

```
~/go/bin/cobra-cli init  
```
потому что иначе не работает команда 

```
~/go/bin/cobra-cli add hello
```

4) Implementation

- Create main.go & root.go

```
~/go/bin/cobra-cli init 
```
- Add sub commands
In my case (MAC) this command work only if cobra was initialized without package name

```
~/go/bin/cobra-cli add hello
~/go/bin/cobra-cli add bye
```


- **You can use arr args [1,2,3,...n] for command
see example /cmd/arrargs.go**


- **You can change parent for command from ROOT to any other and receive nested commands
see examlple /cmd/even.go - it's parent /cmd/arrargs.go**

5) Run examples

```
go run main.go arrargs -h      
go run main.go arrargs even 1 2 3 3 4 5 6 
go run main.go arrargs 1 2 3 3 4 5 6
go run main.go arrargs 1 2 3 3 4 5 6.4  -f
```