import http from 'k6/http';
import { check } from 'k6';
import { sleep } from 'k6';

export let options = {
  vus: 50, // Número de usuários virtuais simulados
  duration: '60s', // Duração do teste
};

export default function () {
  // Acesse a página principal da aplicação
  let response = http.get('http://dev.goapp.127.0.0.1.nip.io/');

  // Verifique se a solicitação foi bem-sucedida
  check(response, {
    'status is 200': (r) => r.status === 200,
  });

  // Espere por um curto período
  sleep(2);

  // Simule cliques nos links de redes sociais
  ['instagram', 'linkedin', 'github'].forEach((network) => {
    response = http.get(`http://dev.goapp.127.0.0.1.nip.io/click?network=${network}`);

    // Verifique se a solicitação foi bem-sucedida
    check(response, {
      [`status is 200 for ${network}`]: (r) => r.status === 200,
    });

    // Espere por um curto período antes do próximo clique
    sleep(1);
  });
}

