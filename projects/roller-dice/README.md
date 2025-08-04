
# 🎲 Roller Dice API

[](https://opensource.org/licenses/MIT)

Uma API simples em Go para rolar dados de RPG. Este projeto foi criado como um sandbox para praticar e entender melhor a estrutura de projetos em Go.

## 🌟 Funcionalidades

  - Rolar um único dado de RPG (D4, D6, D8, D10, D20).
  - Rolar múltiplos dados de diferentes tipos em uma única requisição.
  - Definir um valor customizado para sucesso e falha crítica em rolagens de um único dado.
  - Validação para prevenir a rolagem de dados com lados não convencionais.

## 🚀 Rodando Localmente

1.  **Clone o repositório:**
    ```bash
    git clone https://github.com/dev-araujo/golang__sandbox.git
    ```
2.  **Navegue até o diretório:**
    ```bash
    cd golang__sandbox/roller-dice
    ```
3.  **Execute o servidor:**
    ```bash
    go run ./cmd/main.go
    ```

O servidor estará disponível em `http://localhost:8080`.

## ▶️ Uso

Você pode interagir com a API enviando requisições `POST` para o endpoint `/roll`.

### Exemplos de Requisição

**Rolar um D20:**

```json
{
    "dice_type": 20
}
```

**Rolar múltiplos dados:**

```json
{
    "dice_types": [6, 10, 20]
}
```

**Rolar um D20 com falha/sucesso crítico:**

```json
{
    "dice_type": 20,
    "critical_success": 20,
    "critical_failure": 1
}
```

