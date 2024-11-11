import http from 'k6/http';
import { check } from 'k6';
import { sleep } from 'k6';

export let options = {
  vus: 10, // Número de usuários virtuais simulados
  duration: '15s', // Duração do teste
  thresholds: {
    'http_req_duration': ['p(95)<500'], // Defina limites para a duração das solicitações
  },
};

export default function () {
  // Acesse a página principal da aplicação
  let response = http.get('http://goapp.127.0.0.1.nip.io/');

  // Verifique se a solicitação foi bem-sucedida
  check(response, {
    'status is 200': (r) => r.status === 200,
  });

  // Espere por um curto período
  sleep(2);

  // Simule cliques nos links de redes sociais
  ['instagram', 'linkedin', 'github'].forEach((network) => {
    response = http.get(`http://goapp.127.0.0.1.nip.io/increment-${network}-counter`);

    // Verifique se a solicitação foi bem-sucedida
    check(response, {
      [`status is 200 for ${network}`]: (r) => r.status === 200,
    });

    // Espere por um curto período antes do próximo clique
    sleep(1);
  });

  // Acesse a página de métricas prometheus
  response = http.get('http://goapp.127.0.0.1.nip.io/metrics');

  // Verifique se a solicitação foi bem-sucedida
  check(response, {
    'status is 200 for metrics': (r) => r.status === 200,
  });

  // Espere por um curto período antes do próximo passo
  sleep(1);
}
