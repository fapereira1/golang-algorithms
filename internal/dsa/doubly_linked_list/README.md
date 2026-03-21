# Doubly Linked List

## 1. Problema
Como abordado na Singly Linked List, listas encadeadas resolvem muito bem a alocação dinâmica e os custos de operações no ponteiro de acesso (como na Cabeça da lista).
Contudo, temos um problema arquitetônico com `Singly Linked List`: nós não podemos olhar "para trás". Por exemplo, ao tentar remover o próprio ponteiro final `tail`, temos a obrigação escalar de passar por toda a base de nós do começo ao fim ($O(n)$) pois precisamos conhecer quem é o nó antecessor àquele que vamos desconectar. Esse cenário degrada severamente a eficiência em casos onde o acesso/reversão nas extremidades em tempo real é crítico.

## 2. Abordagem
A solução de Data Structures para mitigar este problema é a expansão para a **Doubly Linked List** (Lista Duplamente Encadeada). O funcionamento base continua idêntico à *Singly* mas, agora, todo Nó da nossa rede guardará não apenas o ponteiro do **Próximo** componente (`Next`), mas também o ponteiro do **Anterior** componente (`Prev`).

Com isso, podemos andar e manipular a rede interligada de forma limpa indo e voltando, conseguindo assim: 
1. Apagar Nós do fim em um custo constante $\bm{O(1)}$.
2. Apagar qualquer Nó contanto que tenhamos apenas a sua referência, sem precisar saber de onde ele veio.

*O "trade-off" (custo de compensação) se encontra de forma estrita no espaço alocado, que aumenta com as referências-ponteiro extras de `Prev` somando-se à memória primária do sistema.*

## 3. Implementação
Nossa implementação em Go (1.21+) adota uma `DoublyLinkedList[T]` controlando os ponteiros extras. 
Assim como a anterior, possuimos os apontamentos `head` (início), e `tail` (final):
- Inserções em `$O(1)$`: A estrutura coordena perfeitamente os links entre de `Prev`/`Next` nas pontes diretas de `$O(1)$` (`Prepend` e `Append`).
- **`Reverse()` In-Place**: Reversão da ordem numa doubly-linked-list é muito natural no algoritmo (diferente da *Singly*), focando apenas num mini loop trocando o atributo de Next e Prev de lugar à medida que se avança. 

## 4. Complexidade

| Operação                          | Tempo             | Espaço Auxiliar | Observações                               |
|-----------------------------------|-------------------|-----------------|-------------------------------------------|
| Inserção no Início (`Prepend`)    | $O(1)$            | $O(1)$          | Duplo link estabelecido em tempo constante|
| Inserção no Fim (`Append`)        | $O(1)$            | $O(1)$          | Acessado indiretamente pelo ponteiro da Cauda|
| Remoção no Fim (`RemoveAt(^sz)`)  | $\bm{O(1)}$       | $O(1)$          | **VANTAGEM**. Usa `tail.Prev` para eliminar o fim em custo limpo O(1)|
| Acesso posicional                 | $O(n)$            | $O(1)$          | Ainda não possuímos índices imediatos de Array|
| Espaço da estrutura de Nó        | $-$               | $O(N)$          | Estrutura de memória custa um ponteiro de 8 bytes adicional por nó|

## 5. Verificação
- Casos complexos de borda (remoção quando só há 1 elemento, injeção em indexadores limitados, varredura com `Prev` da direita para a esquerda) com cobertura massiva via `doubly_linked_list_test.go`.
