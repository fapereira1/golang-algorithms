# Dynamic Array

## 1. Problema
Um array estático tem tamanho fixo, definido no momento de sua criação. Quando não sabemos antecipadamente a quantidade total de elementos que precisamos armazenar (ou quando essa quantidade varia muito), uma estrutura de tamanho fixo se torna ineficiente, exigindo que o programador cuide de alocações e realocações de memória constantes, além de controlar o tamanho lógico vs. capacidade total.

## 2. Abordagem
A solução clássica é o **Dynamic Array** (Array Dinâmico), que encapsula um array fixo subjacente juntamente com o controle do seu tamanho (`size`) e de sua capacidade máxima atual (`capacity`).
A grande vantagem é a **duplicação de capacidade** (ou fator de crescimento, geralmente de 1.5x a 2x). 
Quando tentamos inserir um novo elemento e o array está cheio (`size == capacity`), a estrutura:
1. Aloca um novo array estático com o dobro do tamanho (`capacity = capacity * 2`).
2. Copia os elementos antigos para o novo array (o que tem um custo linear $O(n)$ neste exato momento).
3. Insere o novo elemento e faz o ponteiro interno apóntar para este novo array subjacente.

Isso nos permite inserir elementos no final em tempo constante amortizado.

## 3. Implementação
Nossa implementação em Go 1.21+ utiliza **Generics** (`[T any]`) para permitir que o array dinâmico armazene qualquer tipo.

A struct conterá:
- `data []T`: um slice subjacente atuando como array fixo para nossos propósitos teóricos.
- `size int`: o número de elementos presentes (tamanho lógico).
- `capacity int`: o tamanho total do array alocado (tamanho físico).

*Nota: Em Go, arrays dinâmicos já existem na forma de "slices", que usam `append` nativamente. O objetivo aqui é aprender e ilustrar o algoritmo dos bastidores (como a JVM faz com ArrayList no Java, ou C++ com std::vector).*

## 4. Complexidade

| Operação                          | Tempo (Pior Caso) | Tempo (Caso Médio / Amortizado)| Espaço (Auxiliar)|
|-----------------------------------|-------------------|--------------------------------|------------------|
| Acesso (`Get[i]`) / Escrita (`Set[i]`) | $O(1)$            | $O(1)$                         | $O(1)$           |
| Inserção no Fim (`PushBack`)      | $O(n)$            | $\bm{O(1)}$ amortizado         | $O(1)$           |
| Inserção no Início/Meio           | $O(n)$            | $O(n)$                         | $O(1)$           |
| Remoção no Fim (`PopBack`)        | $O(1)$            | $O(1)$                         | $O(1)$           |
| Remoção no Início/Meio            | $O(n)$            | $O(n)$                         | $O(1)$           |

* Para inserção no fim `PushBack`, o custo $O(n)$ acontece apenas quando redimensionamos. Como o array dobra de tamanho, essas cópias ocorrem sucessivamente em momentos esparsos. Pelo método financeiro (análise amortizada), podemos distribuir o peso da cópia ao longo de várias inserções resultando em custo $\bm{O(1)}$.

## 5. Verificação
- Garantimos que elementos inválidos ou *out-of-bounds* retornem panic.
- Controlamos o encolhimento do array (Shrink) caso esteja muito vazio (fator de carga < 25%).
- Todos os cenários são cobertos por testes unitários em `dynamic_array_test.go` validados com o standard tester do Go.
