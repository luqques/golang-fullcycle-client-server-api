# Desafio Client-Server-API
Desafio da Pós Go Expert pela FullCycle. Neste desafio, você deve aplicar os conhecimentos de HTTP, Contextos, Banco de Dados e Manipulação de Arquivos em Go. Você criará dois sistemas interligados (client.go e server.go) que devem trocar informações respeitando limites estritos de tempo (timeout).

---
<ul>
	<li>
	<p>Ao receber uma requisi&ccedil;&atilde;o em /cotacao, o server deve consumir a API de C&acirc;mbio: https://economia.awesomeapi.com.br/json/last/USD-BRL.</p>
	</li>
	<li>
	<p><strong>Timeout:</strong> O timeout m&aacute;ximo para chamar essa API externa deve ser de <strong>200ms</strong> (usando o pacote context).</p>
	</li>
</ul>
</li>
<li>
<p><strong>Persist&ecirc;ncia (Banco de Dados):</strong></p>

<ul>
	<li>
	<p>O servidor deve registrar cada cota&ccedil;&atilde;o recebida em um banco de dados <strong>SQLite</strong>.</p>
	</li>
	<li>
	<p><strong>Timeout:</strong> O timeout m&aacute;ximo para persistir os dados no banco deve ser de <strong>10ms</strong> (usando o pacote context).</p>
	</li>
</ul>
</li>
<li>
<p><strong>Resposta:</strong></p>

<ul>
	<li>
	<p>O endpoint deve retornar o resultado da cota&ccedil;&atilde;o em formato JSON para o cliente.</p>
	</li>
</ul>
</li>
<li>
<p><strong>Logs:</strong></p>

<ul>
	<li>
	<p>Caso os timeouts (API ou Banco) sejam excedidos, o erro deve ser logado no console do servidor.</p>
	</li>
</ul>
</li>

---
<ul>
	<li>
	<p>Deve solicitar a cota&ccedil;&atilde;o ao endpoint /cotacao do servidor local.</p>
	</li>
	<li>
	<p><strong>Timeout:</strong> O timeout m&aacute;ximo para receber o resultado do servidor deve ser de <strong>300ms</strong> (usando o pacote context).</p>
	</li>
</ul>
</li>
<li>
<p><strong>Processamento e Arquivo:</strong></p>

<ul>
	<li>
	<p>O cliente deve receber apenas o valor atual do c&acirc;mbio (campo bid do JSON).</p>
	</li>
	<li>
	<p>Deve salvar a cota&ccedil;&atilde;o em um arquivo chamado cotacao.txt.</p>
	</li>
	<li>
	<p><strong>Formato do arquivo:</strong> D&oacute;lar: {valor}</p>
	</li>
</ul>
</li>
<li>
<p><strong>Logs:</strong></p>

<ul>
	<li>
	<p>Caso o timeout de 300ms seja excedido, o erro deve ser logado no console do cliente.</p>
	</li>
</ul>
</li>

---
# Como rodar o projeto
Siga este passo a passo para subir e rodar o necessário para testar este projeto.

## Subindo o Banco de dados SQLite
1. Com o Docker instalado e rodando na sua máquina, acesse a raiz deste projeto e rode o comando abaixo:
`docker-compose up -d`


## Rodando o server e client Go
1. Na raiz do projeto, rodar o seguinte comando no cmd para rodar o server.go e em seguida o client.go:
`go run .\cmd\server\main.go`
`go run .\cmd\client\main.go`