# Guia de Estilo para o Projeto Basic Blockchain Go

Este guia estabelece as convenções de código para projetos Go em nossa organização, complementando as práticas recomendadas da linguagem e adaptando-se às nossas necessidades específicas.

## 🎯 Princípios Fundamentais
- **Clareza e Legibilidade**: O código deve ser fácil de entender e manter.
- **Simplicidade**: Prefira soluções diretas e evit complexidade desnecessária.
- **Consistência**: Mantenha um estilo coerente em todo o código.
- **Desempenho**: Escreva código eficiente, mas não à custa da legibilidade.

## 📝 Convenções de Código

### Formatação
- Utilize `gofmt` para formatação automática.
- Linhas devem ter no máximo **100 caracteres** (um desvio comum do padrão Go de 80 para melhor legibilidade em monitores modernos) :cite[2].

### Nomenclatura
- **Variáveis**: Use `camelCase` (ex: `userName`).
- **Constantes**: Use `UpperCase` com underscores (ex: `MAX_CONNECTIONS`).
- **Funções**: Use `camelCase` em funções internas (ex: `calculateTotal()`) e PascalCase para funções que serão exportadas(ex: `CalculateTotal`).
- **Interfaces**: Use `er` no final se apropriado (ex: `Reader`).
- **Pacotes**: Nomes curtos e em minúsculas (ex: `util`).

