import http from 'k6/http';
import { check } from 'k6';

export let options = {
  vus: 10, // Número de usuários virtuais simulados
  duration: '60s', // Duração do teste
};

export default function () {
  // Acesse apenas o endpoint principal
  let response = http.get('http://dev.goapp.127.0.0.1.nip.io/metrics');

  // Verifique se a solicitação foi bem-sucedida
  check(response, {
    'status is 200': (r) => r.status === 200,
  });
}

