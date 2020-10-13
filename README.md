# Exemplo mínimo de Go + Web API + Contêiner Docker
Exemplo mínimo de uma imagem de um contêiner Docker rodando uma aplicação Go respondendo a requisições http na porta 80.
Recomendado para testar o deploy em um servidor de nuvem como Microsoft Azure, AWS ou Google cloud platform.
Contém um Dockerfile e uma aplicação simples que podem ser utilizados para estudo, prova de conceito ou como ponto de partida.

Uma aplicação em produção deveria provavelmente seguir uma das duas alternativas:
* Utilizar configurações adicionais de segurança
* Rodar um servidor nginx que serve de proxy para a aplicação

Isso permitirá a aplicação lidar certificados, etc.