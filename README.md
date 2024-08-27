# SOLID Principles in Golang

Este projeto demonstra a implementação dos princípios SOLID usando exemplos em Golang. Cada princípio é mostrado em duas versões: antes e depois de aplicar o princípio. O objetivo é ajudar os desenvolvedores a entender como cada princípio pode melhorar o design do código.

## Arquivos

### S.go - Single Responsibility Principle (SRP)
Neste exemplo, a classe original `User` lidava tanto com o salvamento no banco de dados quanto com o envio de emails. Isso viola o SRP porque a classe tem mais de uma razão para mudar. Após aplicar o SRP, as responsabilidades foram divididas em `UserRepository` para operações de banco de dados e `EmailService` para o envio de emails.

### O.go - Open/Closed Principle (OCP)
O código original calculava a área de diferentes formas usando um switch case, o que viola o OCP, pois adicionar novas formas exigiria modificar a função existente. Após aplicar o OCP, usamos uma interface `Shape` com um método `Area`, permitindo adicionar novas formas sem modificar o código existente.

### L.go - Liskov Substitution Principle (LSP)
Inicialmente, o exemplo permitia que uma avestruz herdasse o método `Fly`, mesmo que avestruzes não voem. Isso viola o LSP porque a subclasse não se comporta como esperado com base na classe pai. Após aplicar o LSP, cada pássaro tem um método `Move` que reflete seu movimento natural. Criamos uma interface `Bird` com o método `Move()`, e cada tipo de pássaro (pardal, avestruz e pinguim) implementa `Move()` de maneira diferente (`Fly`, `Run`, `Swim`), respeitando seus comportamentos.
Caso fosse necessário lidar apenas com pássaros que voem, teríamos que ter uma interface `Flyer` na qual teríamos implementações apenas de pássaros que voem, evitando assim, comportamento inesperado pelo sistema.

### I.go - Interface Segregation Principle (ISP)
No primeiro exemplo, a interface `Worker` força tanto humanos quanto robôs a implementar métodos que eles não precisam. Após aplicar o ISP, as interfaces foram divididas em `Worker` e `Eater`, de modo que cada classe implementa apenas os métodos que necessita.

### D.go - Dependency Inversion Principle (DIP)
O primeiro exemplo tem um módulo de alto nível (`Switch`) que depende de um módulo de baixo nível (`LightBulb`), violando o DIP. Após aplicar o DIP, o `Switch` depende de uma abstração (`Switchable`), permitindo que ele funcione com qualquer dispositivo que implemente a interface.

## Como Executar
Você pode executar cada arquivo individualmente para ver as diferenças antes e depois de aplicar os princípios SOLID:

```bash
go run S.go
go run O.go
go run L.go
go run I.go
go run D.go
