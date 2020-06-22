#!/bin/bash

#verifica se o banco de dados existe
mysql -Bse "USE EDP_Botnet" 2> /dev/null

if [ $? -eq 0 ]; then

  query="GRANT ALL PRIVILEGES ON EDP_Botnet.* TO B0tn3t@'localhost' IDENTIFIED BY '45p1r1n4';";

  read -p "Executando: $query, suave? (y/n): " confirmacao;

  if [ "$confirmacao" == 'y' ]; then

    mysql -e "$query"

    mysql -e "flush privileges"

  else

    read -p "suave... aperta qualquer tecla ai"

  fi

else

  echo "O BD não existe, será criado agora"

  mysql -e "CREATE DATABASE EDP_Botnet"

  echo "O BD foi criado, rode o script de novo"

fi

