

# CLI IPs & Servers

![image](https://github.com/user-attachments/assets/e3274fa6-269e-42ca-97dc-e0f7c28f3739)


Este é um aplicativo de linha de comando (CLI) desenvolvido em Go que permite buscar IPs e nomes de servidores na internet. O projeto utiliza a biblioteca [urfave/cli](https://github.com/urfave/cli) para facilitar a criação de comandos e flags.

## Funcionalidades

- **Busca de IPs**: Permite buscar os endereços IP de um determinado host.
- **Busca de Servidores**: Permite buscar os nomes dos servidores de um determinado host.

## Estrutura do Projeto

O projeto possui a seguinte estrutura:

```
.
├── app/
│   └── app.go
└── main.go
```

- **app/app.go**: Contém a lógica principal para a geração da aplicação CLI e a definição dos comandos.
- **main.go**: Ponto de entrada da aplicação, que inicializa e executa o CLI.

## Uso

### Pré-requisitos

- [Go](https://golang.org/) instalado na sua máquina.

### Instalação

Clone o repositório:

```bash
git clone https://github.com/dev-araujo/go-server-ip.git
```

### Execução

Você pode compilar e executar o CLI usando os comandos abaixo:

1. **Compilação**

   ```bash
   go build 
   ```

2. **Execução**

   ```bash
   ./cli-ips-servers [comando] --host [nome_do_host]
   ```

### Comandos Disponíveis

1. **ip**

   Busca os endereços IP de um determinado host.

   ```bash
   ./cli-ips-servers ip --host dev-araujo.com.br
   ```

   **Flags**:
   - `--host`: Define o host para o qual os IPs serão buscados (padrão: `dev-araujo.com.br`).

2. **servers**

   Busca os nomes dos servidores de um determinado host.

   ```bash
   ./cli-ips-servers servers --host dev-araujo.com.br
   ```

   **Flags**:
   - `--host`: Define o host para o qual os nomes dos servidores serão buscados (padrão: `dev-araujo.com.br`).

### Exemplo de Uso

Para buscar os IPs do host `google.com`, você pode utilizar o seguinte comando:

```bash
./cli-ips-servers ip --host google.com
```

Para buscar os servidores do host `google.com`, você pode utilizar:

```bash
./cli-ips-servers servers --host google.com
```


## Licença

Este projeto está licenciado sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---

#### Autor 👷

<img src="https://user-images.githubusercontent.com/97068163/149033991-781bf8b6-4beb-445a-913c-f05a76a28bfc.png" width="5%" alt="caricatura do autor desse repositório"/>

[![linkedin](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/araujocode/)
