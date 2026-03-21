# Singly Linked List

## 1. Problema
Arrays dinâmicos, apesar de práticos, requerem que seus elementos fiquem alocados sequencialmente na memória (contíguos). Quando precisamos fazer muitas inserções ou remoções no início ou no meio da lista, precisamos deslocar todo o resto, o que tem um custo de desempenho linear $O(n)$.
Além disso, com memórias muito fragmentadas, alocar um chunk imenso de forma contígua pode ser inviável.

## 2. Abordagem
A solução de Data Structures clássica para isso é a **Singly Linked List** (Lista Simplesmente Encadeada). Em vez de armazenar todos os elementos em posições sequenciais, alocamos os nós ("Nodes") soltos pelo sistema de memória heap, e conectamos cada nó ao seu elemento adjacente.
Cada Nó contém:
1. O valor armazendo (`Value`).
2. O endereço do próximo Nó (`Next`).

Com isso, podemos inserir e remover do topo com custo imediato (apenas alterando os ponteiros). Se mantermos um ponteiro interno adicional referenciando o final (`tail`), também podemos adicionar elementos no final da lista em tempo constante.

## 3. Implementação
Nossa implementação em Go (1.21+) utiliza **Generics** (`[T comparable]`) para que a Lista possa gerenciar qualquer tipo e suportar procuras por valor.

A estrutura possui os campos `head` (início), `tail` (final) e `size`.
Foram implementados métodos nativos fundamentais da lista e técnicas algorítmicas adicionais comumente testadas em problemas de desenvolvimento, como:
- **`Reverse()`**: inverte a lista localmente usando três ponteiros, garantindo complexidade $O(N)$ em tempo e $\bm{O(1)}$ em espaço algorítmico sem a criação de itens extras.
- **`HasCycle()`**: algoritmo para detecção de ciclos. Utilizamos `Floyd's Tortoise and Hare` que avança dois ponteiros (lento/rápido) descobrindo ciclos infinitos de forma extremamente eficiente ($O(N)$ de tempo, $O(1)$ de espaço).

## 4. Complexidade

| Operação                          | Tempo             | Espaço Auxiliar | Observações                               |
|-----------------------------------|-------------------|-----------------|-------------------------------------------|
| Inserção no innício (`Prepend`)   | $O(1)$            | $O(1)$          | Custo de alterar ponteiros imediatos      |
| Inserção no Fim (`Append`)        | $\bm{O(1)}$       | $O(1)$          | Somente via uso de cache do ponteiro `tail` |
| Inserção no Meio (`InsertAt`)     | $O(n)$            | $O(1)$          | Requer atravessar $N$ nós                 |
| Remoção no Início (`RemoveAt(0)`) | $O(1)$            | $O(1)$          | O GC de Go cuidará do nó perdido          |
| Remoção no Fim                      | $O(n)$            | $O(1)$          | Precisamos achar o nó *anterior* ao `tail`|
| Acesso posicional                   | $O(n)$            | $O(1)$          | Necessário varrer o link um por um        |

## 5. Verificação
- Garantimos proteções lógicas para evitar Out Of Range panics.
- Algoritmos de `Reverse` e manipulação de pontas foram integralmente cobertos por testes validados em `singly_linked_list_test.go`.
- Teste contra ciclos infinitos injetados deliberadamente para certificar que o sistema Tortoise & Hare é perfeitamente capturado.
