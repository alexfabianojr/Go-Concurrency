# Go-Concurrency

Concorrência vs Paralelismo

Go é uma linguagem concorrente

A concorrência em Go é algo intrínseco a linguagem 

- Paralelismo: você está fazendo duas coisas exatamente ao mesmo tempo (simultaneamente)
    * só é possível através de muitos núcleos de cpu
    * executar código simultaneamente em processadores físicos diferentes

- Concorrência: é a capacidade de você administrar múltiplas tarefas, podendo ocorrer em apenas um único processador
    * Deve ser cuidado o overhead de controle que pode sobrecarregar desnecessariamente a aplicação
    * intercalar (administrar) vários processos ao mesmo tempo e isso pode occorer mesmo com um único processador físico
    * é a forma de estruturar o seu programa
    * composição de processos independentes

- É possível que a concorrência seja melhor que o paralelismo
- Paralelismo exige muito mais do SO e do hardware



