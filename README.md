# Resumo

Recentemente, arquiteturas baseadas em microsserviços ganharam popularidade devido
ao modelo de programação modular, fraco acoplamento entre partes do sistema
e também ao suporte de plataformas de orquestração de contêineres, que oferecem
gerenciamento automático e eficiente de recursos. Ordenação de mensagens é uma
estratégia que busca garantir que todas as réplicas evoluam igualmente, no tempo,
aumentando os níveis de disponibilidade de serviços. Um tema amplo que aborda
diversos algoritmos para tratar problemas de comunicação entre as réplicas, e o presente
trabalho faz menção a alguns algoritmos que implementam difusão atômica ou
baseado em consenso. Este trabalho propõe uma implementação de protocolo HTTP
para um ordenador de mensagens. A implementação segue padrões de projeto em
arquiteturas de microsserviços, tirando assim, a necessidade de haver ordenação de
mensagens inserida no código da aplicação servidora. Sabe-se que orquestradores
oferecem replicação de forma automática, porém o serviço oferecido por orquestradores
serve bem para aplicações stateless. A proposta visa cobrir os casos onde o
serviço a ser replicado é stateful, e a ordem na entrega de mensagens entre réplicas
deve ser a mesma. O objetivo é oferecer a funcionalidade de ordenação de mensagens,
de forma transparente ao usuário, onde os detalhes sobre ordenação não precisam
ser explicitamente abordados por desenvolvedores. Este trabalho estende um trabalho
iniciado pelo grupo, que propõe o Hermes, uma arquitetura para um serviço de ordenação
de mensagens baseada em Proxy. Para aprimorar o serviço de ordenação de
mensagens, implementou-se um novo protocolo de comunicação, o HTTP. Além disso,
investigou-se o desempenho da implementação em casos específicos de vazão e latência.
Após a investigação foram observados que as latências médias do ordenador
de mensagens Hermes ficaram relativamente pouco acima da média em relação ao
sistema sem replicação. Já para o cenário onde é há uma mistura de 50% leitura e
50% escrita, foi mais expressivo o distanciamento das latências médias do Hermes em
relação ao sistema sem replicação.

# Palavras-chave

Protocolo de consenso. Proxy. Interceptação de Mensagens.
Orquestração de contêineres. Ordenação total.
Arquitetura de microsserviços. Kubernetes.
Docker.
