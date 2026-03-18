# Plano sequencial de implementação — DSA Expert Skill

> **Objetivo:** implementar todas as estruturas de dados e algoritmos cobertos pela skill `dsa-expert`,
> seguindo uma ordem pedagógica onde cada módulo constrói sobre os anteriores.
> Cada item inclui o que implementar, os pré-requisitos, e um problema de validação.

---

## Fase 1 — Fundações (estruturas lineares)

Estas são as estruturas que sustentam quase tudo que vem depois. Sem domínio total aqui, o restante desmorona.

### 1.1 Dynamic Array

- **Implementar:** alocação, append com resize (fator 2), acesso por índice, inserção/remoção em posição arbitrária
- **Conceitos-chave:** análise amortizada (por que N appends custam O(N) total)
- **Pré-requisitos:** nenhum
- **Validação:** implementar e provar que o custo amortizado de append é O(1)

### 1.2 Linked list (singly + doubly)

- **Implementar:** inserção (head, tail, posição), remoção, reversão iterativa, detecção de ciclo (Floyd's)
- **Conceitos-chave:** manipulação de ponteiros, sentinela nodes
- **Pré-requisitos:** 1.1 (entender trade-offs vs array)
- **Validação:** dado uma lista com ciclo, encontrar o nó de início do ciclo e provar por que o algoritmo funciona (relação a = kc - b)

### 1.3 Stack

- **Implementar:** push, pop, peek sobre array dinâmico
- **Conceitos-chave:** LIFO, monotonic stack pattern
- **Pré-requisitos:** 1.1
- **Validação:** resolver "next greater element" com monotonic stack em O(n)

### 1.4 Queue e Deque

- **Implementar:** circular queue sobre array fixo, deque com operações nas duas pontas
- **Conceitos-chave:** FIFO, aritmética modular para índices circulares
- **Pré-requisitos:** 1.1
- **Validação:** resolver "sliding window maximum" com deque monotônica em O(n)

---

## Fase 2 — Hashing e busca

### 2.1 Hash table

- **Implementar:** versão com chaining (linked list por bucket) e versão com open addressing (linear probing)
- **Conceitos-chave:** função hash, load factor, rehashing, trade-offs chaining vs open addressing
- **Pré-requisitos:** 1.1, 1.2
- **Validação:** implementar ambas as estratégias de colisão e comparar performance com load factor variando de 0.3 a 0.9

### 2.2 Binary search (todas as variantes)

- **Implementar:** busca padrão, lower_bound, upper_bound, search on answer
- **Conceitos-chave:** invariante do loop, por que `mid = lo + (hi - lo) / 2` evita overflow, quando usar `lo < hi` vs `lo <= hi`
- **Pré-requisitos:** 1.1
- **Validação:** resolver "minimum capacity to ship packages within D days" usando binary search on answer

---

## Fase 3 — Árvores

### 3.1 Binary search tree (BST)

- **Implementar:** inserção, busca, remoção (3 casos), 4 traversals (in-order, pre-order, post-order, level-order), successor/predecessor
- **Conceitos-chave:** invariante BST, por que árvores degeneradas viram listas ligadas
- **Pré-requisitos:** 1.2 (ponteiros), 2.2 (conceito de busca)
- **Validação:** construir BST a partir de sequência ordenada, observar degeneração O(n), motivar balanceamento

### 3.2 Heap / priority queue

- **Implementar:** min-heap sobre array, sift_up, sift_down, heapify (bottom-up), extract-min, insert, decrease-key
- **Conceitos-chave:** propriedade de heap, representação em array (parent/left/right), prova de que heapify é O(n)
- **Pré-requisitos:** 1.1, 3.1 (conceito de árvore)
- **Validação:** implementar heapify e demonstrar com a série geométrica que o custo é O(n) e não O(n log n)

### 3.3 Trie (prefix tree)

- **Implementar:** inserção, busca exata, busca por prefixo, remoção com poda
- **Conceitos-chave:** trade-off memória vs velocidade, otimização com array[26] vs hash map
- **Pré-requisitos:** 2.1 (hash map para children), 3.1 (navegação em árvore)
- **Validação:** implementar autocomplete — dado um prefixo, retornar todas as palavras que o completam

### 3.4 Balanced BST (conceitual)

- **Implementar:** AVL tree com rotações (LL, RR, LR, RL)
- **Conceitos-chave:** fator de balanceamento, por que AVL garante O(log n), quando preferir AVL vs Red-Black
- **Pré-requisitos:** 3.1
- **Validação:** inserir sequência 1..n e verificar que a altura permanece ≤ 1.44 log n

---

## Fase 4 — Grafos (representação + travessias)

### 4.1 Representação de grafos

- **Implementar:** adjacency list, adjacency matrix, edge list
- **Conceitos-chave:** trade-offs de espaço e tempo entre as três representações, quando usar cada uma
- **Pré-requisitos:** 1.1, 1.2, 2.1
- **Validação:** modelar um grafo real (ex: mapa de cidades) nas três representações e comparar uso de memória

### 4.2 BFS

- **Implementar:** travessia com fila, shortest path em grafo não-ponderado, multi-source BFS
- **Conceitos-chave:** por que BFS garante menor caminho em grafos sem peso, level-order
- **Pré-requisitos:** 1.4 (queue), 4.1
- **Validação:** resolver "rotting oranges" (multi-source BFS em grid)

### 4.3 DFS

- **Implementar:** versão recursiva e iterativa (com stack), classificação de arestas (tree, back, forward, cross)
- **Conceitos-chave:** tempo de descoberta/finalização, detecção de ciclo em grafo dirigido (coloração WHITE/GRAY/BLACK)
- **Pré-requisitos:** 1.3 (stack), 4.1
- **Validação:** detectar ciclo em grafo dirigido e retornar o ciclo encontrado

### 4.4 Topological sort

- **Implementar:** Kahn's algorithm (BFS com in-degree) e versão DFS (post-order reverso)
- **Conceitos-chave:** por que só funciona em DAGs, relação com detecção de ciclo
- **Pré-requisitos:** 4.2, 4.3
- **Validação:** dado um conjunto de dependências de build (como um Makefile), determinar ordem de compilação

---

## Fase 5 — Grafos (caminhos e MST)

### 5.1 Dijkstra

- **Implementar:** com min-heap (priority queue), reconstrução do caminho
- **Conceitos-chave:** por que pesos negativos quebram o algoritmo, complexidade O((V+E) log V)
- **Pré-requisitos:** 3.2 (heap), 4.2 (BFS como caso especial)
- **Validação:** encontrar menor caminho em mapa rodoviário ponderado e reconstruir a rota

### 5.2 Bellman-Ford

- **Implementar:** relaxamento em V-1 passes, detecção de ciclo negativo
- **Conceitos-chave:** por que V-1 passes são suficientes, quando preferir sobre Dijkstra
- **Pré-requisitos:** 5.1 (entender limitação de Dijkstra)
- **Validação:** detectar oportunidade de arbitragem em câmbio (ciclo negativo no grafo de log das taxas)

### 5.3 Floyd-Warshall

- **Implementar:** all-pairs shortest path com DP 3D (otimizado para 2D)
- **Conceitos-chave:** significado do "vértice intermediário k", transitive closure
- **Pré-requisitos:** 4.1 (matrix), conceitos de DP (fase 7)
- **Validação:** dado grafo de N cidades, calcular distância mínima entre todos os pares

### 5.4 Kruskal (MST)

- **Implementar:** ordenar arestas + union-find para construir MST
- **Conceitos-chave:** propriedade de corte, por que a aresta mais leve que cruza um corte está no MST
- **Pré-requisitos:** 6.1 (union-find — implementar junto ou antes)
- **Validação:** encontrar custo mínimo para conectar N servidores em rede

---

## Fase 6 — Estruturas avançadas

### 6.1 Union-Find (disjoint set)

- **Implementar:** find com path compression, union by rank
- **Conceitos-chave:** inverse Ackermann (α(n) ≈ O(1) na prática), por que as duas otimizações juntas são essenciais
- **Pré-requisitos:** 1.1
- **Validação:** resolver "number of connected components" em grafo dinâmico (arestas chegando online)

### 6.2 Segment tree

- **Implementar:** build, point update, range query (versão iterativa), lazy propagation para range update
- **Conceitos-chave:** por que O(log n) para queries, como lazy propagation adia o trabalho
- **Pré-requisitos:** 3.1 (árvore), conceito de divide and conquer
- **Validação:** resolver "range sum query with updates" — suportar update(i, val) e sum(l, r) ambos em O(log n)

### 6.3 Fenwick tree (BIT)

- **Implementar:** point update, prefix query, range query
- **Conceitos-chave:** truque do `i & (-i)`, por que funciona apenas para operações com inverso (soma sim, min não)
- **Pré-requisitos:** 6.2 (entender o problema que resolve, como alternativa mais simples)
- **Validação:** implementar count of inversions usando Fenwick tree em O(n log n)

### 6.4 Bloom filter

- **Implementar:** inserção (set k bits), query (check k bits), cálculo de false positive rate
- **Conceitos-chave:** probabilístico (false positives ok, false negatives impossíveis), parametrização (m, k, n)
- **Pré-requisitos:** 2.1 (hashing)
- **Validação:** construir um filtro para 1M URLs e verificar empiricamente que a taxa de falso positivo converge para o valor teórico

---

## Fase 7 — Sorting

### 7.1 Merge sort

- **Implementar:** versão recursiva com merge, versão para linked list (sem array auxiliar)
- **Conceitos-chave:** estabilidade, contagem de inversões como subproduto
- **Pré-requisitos:** 1.1, 1.2
- **Validação:** contar inversões em array de 10⁵ elementos em O(n log n)

### 7.2 Quick sort

- **Implementar:** partição Hoare e Lomuto, pivot randomizado, 3-way partition (Dutch National Flag)
- **Conceitos-chave:** por que O(n²) worst case, por que randomização faz worst case astronomicamente improvável
- **Pré-requisitos:** 1.1
- **Validação:** implementar as 3 variantes e comparar número de swaps em arrays com muitos duplicados

### 7.3 Counting sort e radix sort

- **Implementar:** counting sort, radix sort (LSD)
- **Conceitos-chave:** por que quebram o lower bound O(n log n) — não são baseados em comparação
- **Pré-requisitos:** 1.1, 2.1 (conceito de bucket)
- **Validação:** ordenar 10⁶ inteiros no range [0, 10⁶] e comparar tempo com quick sort

---

## Fase 8 — Técnicas algorítmicas

### 8.1 Sliding window e two pointers

- **Implementar:** fixed window (max sum of k), variable window (longest unique substring), minimum window substring, two pointers em array sorted (two sum, 3sum)
- **Conceitos-chave:** quando expandir vs contrair a janela, invariante mantida pelo ponteiro esquerdo
- **Pré-requisitos:** 1.1, 2.1 (hash map para contagem)
- **Validação:** resolver "minimum window substring" em O(n)

### 8.2 Backtracking

- **Implementar:** template choose → explore → unchoose, permutações, combinações, N-Queens, Sudoku solver
- **Conceitos-chave:** árvore de decisão, pruning para cortar branches inválidos cedo
- **Pré-requisitos:** 4.3 (DFS como conceito), recursão
- **Validação:** resolver N-Queens para N=12 e contar soluções, comparar com e sem pruning

### 8.3 Divide and conquer

- **Implementar:** merge sort (já feito), closest pair of points, Karatsuba multiplication
- **Conceitos-chave:** master theorem para análise de recorrências
- **Pré-requisitos:** 7.1
- **Validação:** dado N pontos no plano, encontrar o par mais próximo em O(n log n) — provar que a strip contém no máximo 6 pontos relevantes

---

## Fase 9 — Dynamic programming

A fase mais densa. Cada sub-item é uma família inteira de problemas.

### 9.1 DP 1D

- **Implementar:** climbing stairs, house robber, coin change (min coins), longest increasing subsequence (O(n log n) com binary search)
- **Conceitos-chave:** definir estado, escrever recorrência, base cases, otimizar espaço
- **Pré-requisitos:** 2.2 (binary search para LIS otimizado)
- **Validação:** implementar LIS em O(n log n) e traçar o array `tails` passo a passo

### 9.2 DP 2D

- **Implementar:** longest common subsequence, edit distance, grid paths (unique paths com obstáculos), 0/1 knapsack
- **Conceitos-chave:** tabela 2D, reconstrução da solução (backtrack na tabela), otimização para 1 row
- **Pré-requisitos:** 9.1
- **Validação:** implementar edit distance com reconstrução — dado "kitten" → "sitting", mostrar as operações

### 9.3 DP intervalar

- **Implementar:** matrix chain multiplication, burst balloons
- **Conceitos-chave:** iterar por tamanho de intervalo, não por posição
- **Pré-requisitos:** 9.2
- **Validação:** resolver burst balloons e reconstruir a ordem ótima de estourar os balões

### 9.4 Bitmask DP

- **Implementar:** TSP (travelling salesman), contagem de subsets com propriedade
- **Conceitos-chave:** estado = (máscara de bits, posição atual), iteração sobre subsets de uma máscara
- **Pré-requisitos:** 9.2, manipulação de bits
- **Validação:** resolver TSP para N=15 cidades e verificar que a complexidade é O(2^n × n²)

---

## Fase 10 — Greedy

### 10.1 Interval scheduling

- **Implementar:** activity selection (max non-overlapping), minimum meeting rooms (min-heap de end times), merge overlapping intervals
- **Conceitos-chave:** prova greedy-stays-ahead, por que ordenar por end time e não por start time
- **Pré-requisitos:** 3.2 (heap para meeting rooms), 7.x (sorting)
- **Validação:** provar formalmente que activity selection por end time é ótimo via exchange argument

### 10.2 Huffman coding

- **Implementar:** construir árvore de Huffman com min-heap, gerar códigos, encode/decode
- **Conceitos-chave:** prefix-free codes, optimalidade do Huffman para códigos sem contexto
- **Pré-requisitos:** 3.2 (heap), 3.1 (árvore)
- **Validação:** comprimir e descomprimir uma string, verificar que a compressão é menor que 8 bits/char

---

## Fase 11 — String algorithms

### 11.1 KMP

- **Implementar:** construção da failure function (LPS array), pattern matching
- **Conceitos-chave:** por que o ponteiro do texto nunca retrocede, por que O(n + m)
- **Pré-requisitos:** 1.1
- **Validação:** buscar padrão em texto de 10⁶ chars e verificar complexidade linear

### 11.2 Rabin-Karp

- **Implementar:** rolling hash, matching com verificação de colisão
- **Conceitos-chave:** escolha de base e módulo, por que verificar match após hash collision
- **Pré-requisitos:** 2.1 (hashing)
- **Validação:** multi-pattern matching — buscar 100 padrões simultaneamente em texto grande

### 11.3 Z-algorithm e Manacher's

- **Implementar:** construção do Z-array, longest palindromic substring via Manacher
- **Conceitos-chave:** Z-box, por que Manacher é O(n) apesar dos loops aninhados
- **Pré-requisitos:** 11.1 (entender string matching)
- **Validação:** encontrar o maior palíndromo em string de 10⁵ chars em O(n)

---

## Fase 12 — Bit manipulation e math

### 12.1 Bit manipulation

- **Implementar:** count set bits (Brian Kernighan), power of two check, isolate/clear/toggle bits, XOR tricks (find unique element)
- **Conceitos-chave:** two's complement, por que `x & (x-1)` limpa o lowest set bit
- **Pré-requisitos:** nenhum (pode ser feito em qualquer ponto)
- **Validação:** dado array onde todos aparecem 2x exceto um, encontrar o único em O(n) tempo e O(1) espaço

### 12.2 Math e number theory

- **Implementar:** GCD (Euclides), modular exponentiation (fast power), sieve of Eratosthenes
- **Conceitos-chave:** por que GCD é O(log(min(a,b))), Fermat's little theorem para inverso modular
- **Pré-requisitos:** nenhum
- **Validação:** fatorar todos os primos até 10⁷ em menos de 1 segundo usando sieve

---

## Resumo de dependências

```
Fase 1 (linear)
  └─→ Fase 2 (hashing + busca)
       └─→ Fase 3 (árvores)
            ├─→ Fase 4 (grafos: travessias)
            │    └─→ Fase 5 (grafos: caminhos + MST)
            └─→ Fase 6 (estruturas avançadas)
  └─→ Fase 7 (sorting)
       └─→ Fase 8 (técnicas: window, backtracking, D&C)
            └─→ Fase 9 (dynamic programming)
                 └─→ Fase 10 (greedy)

Fases independentes (podem ser feitas a qualquer momento):
  • Fase 11 (strings) — após Fase 2
  • Fase 12 (bits + math) — sem pré-requisitos
```

---

## Como usar este plano com a skill

Para cada item do plano, peça ao Antigravity usando a skill `dsa-expert`:

> "Implemente [estrutura/algoritmo] em [linguagem]. Siga o workflow completo:
> problema → abordagem → implementação → complexidade → verificação."

A skill vai consultar automaticamente os arquivos de referência (`data-structures.md` e `algorithm-patterns.md`)
e entregar a implementação com análise de complexidade, invariantes e edge cases.

Se você tiver uma skill de linguagem (ex: `go-expert`, `ruby-expert`, `python-expert`),
a skill DSA compõe com ela para gerar código idiomático na linguagem escolhida.