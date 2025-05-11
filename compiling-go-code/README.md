#compiling-go-code

1 - Go é uma linguagem Compilada

CÓDIGO -> COMPILADOR DO CÓDIGO -> LINGUAGEM DE MÁQUINA -> LINKING (LINKAGEM) -> EXECUTÁVEL (PROGRAMA)
main.go -> compilador do Go -> executável main

- Go é uma linguagem multiplataforma!
- Linguagens de programação podem ser interpretadas ou compiladas!
- (não é sempre) linguagens de programção compiladas tendem a ser mais performáticas
- 'go run main.go' -> compila e roda o arquivo main.go
- 'gobuild' ->

# Questões
- Ele compila apenas a nível de OS ou de arquitetura também? Exemplo: compilei em um Linux x86, eu consigo executar em um Linux arm? OS: Windows, Linux, Max. Arquitetura: 64 bits, 32 bits, ARM.
R: Não, não consegue. Ele compila em nível de arquitetura.
A parte legal é que posso compilar para qualquer arquitetura, em qualquer arquitetura. Ou seja, eu posso compilar para um ARM e um linux x86.
Como fazemos isso?
--- GOOS
-- GOARCH
(linha de comando no terminal: GOOS=windows GOARCH=amd64 go build main.go)

#Casos de uso úteis:
- COmpartilhar um binário com alguém: main num pendrive.
- CI/CD em Go:
    -------CI: compilar
    -------CD: colocar o binário pra rodar na maquina
- A máquina que vai rodar o código não precisa ter acesso a nada do código fonte!

# Por que a forma que o Go compila o código é "melhor" do que em C?
Pros nossos casos de uso:
--- Binário Go é autocontido, ou seja, qualquer biblioteca de sistema etc já vem dentro do binário
--- Código compilado em C nem sempre tem dentro desses códigos as bibliotecas (.h) que o seu código precisa pra rodar dentro do executável

# Ressalves:
--- Compatibilidade de Bibliotecas: Nem todas as bibliotecas Go suportam compilação cruzada.
---- Testes e Validação: Testes em plataformas alvo são cruciais, pois o comportamento do código pode variar entre sistemas operacionais.
---- Dependências Nativas: Para projetos mais complexos, pode ser necessário configurar ferramentas adicionais para lidar com dependências nativas, especialmente ao compilar para plataformas diferentes.

# Flags de compilação
- Ajudam a customizar o binário (executável) ferado
- E as vezes ajudam a diminuir o tamanho desse executável
- Vocês vao ver o Dockerfile customizando flags pra gerar um binário menor pra rodar em ambientes produtivos

<!-- JS - Interpretada?
Java - Compilação -> Máquina virtual (JVM) -> Interpretação
(entao quem entende a compilaçao é a maquina virtual) -->