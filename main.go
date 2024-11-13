package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)


var (
	instagramClicks = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "social_media_clicks_instagram",
		Help: "Number of clicks on Instagram link",
	})

	linkedinClicks = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "social_media_clicks_linkedin",
		Help: "Number of clicks on LinkedIn link",
	})

	githubClicks = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "social_media_clicks_github",
		Help: "Number of clicks on GitHub link",
	})

	pageViews = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "page_views_total",
		Help: "Total number of page views",
	})

	metricsMtx sync.Mutex
)

func init() {
	// Registra as métricas no Prometheus
	prometheus.MustRegister(instagramClicks)
	prometheus.MustRegister(linkedinClicks)
	prometheus.MustRegister(githubClicks)
	prometheus.MustRegister(pageViews)

	// Inicializa o log para registrar no console
	log.SetOutput(os.Stdout)  // Configura para logar no console
	log.Println("Iniciando o servidor HTTP...")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Captura o IP do cliente
		clientIP := r.RemoteAddr
		if forwardedFor := r.Header.Get("X-Forwarded-For"); forwardedFor != "" {
			clientIP = strings.Split(forwardedFor, ",")[0]  // Pega o primeiro IP na lista, se existir
		}

		// Captura o User-Agent (informações sobre o navegador)
		userAgent := r.UserAgent()

		// Captura o nome do host
		hostname, _ := os.Hostname()

		// Incrementa o contador de visualizações
		metricsMtx.Lock()
		defer metricsMtx.Unlock()
		pageViews.Inc()

		// Log de requisição com IP e User-Agent
		log.Printf("Requisição recebida de IP: %s, User-Agent: %s", clientIP, userAgent)

		// Respondendo com a página HTML
		w.Write([]byte(`
            <!DOCTYPE html>
            <html lang="en">
                <head>
                    <meta charset="UTF-8">
                    <meta name="viewport" content="width=device-width, initial-scale=1.0">
                    <script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>
                    <style>
                        body {
                            display: flex;
                            flex-direction: column;
                            align-items: center;
                            min-height: 100vh;
                            margin: 0;
                            font-family: Arial, sans-serif;
                            background-color: #D8923B; /* #BE7839 Cor de fundo da tela */
                        }

                        h1 {
                            text-align: center;
                            color: white; /* Alterado para a cor branca */
                        }

                        p {
                            text-align: center;
                            color: #333;
                        }

                        .image-container {
                            display: block;
                            margin: auto;
                            cursor: zoom-in;
                            transition: background-color 300ms;
                        }

                        .responsive-image {
                            max-width: 100%;
                            height: auto;
                            width: 100%;
                            border-radius: 10px;
                        }

                        /* Navbar styles */
                        .navbar {
                            display: flex;
                            justify-content: space-between;
                            align-items: center;
                            padding: 10px;
                            background-color: white; /* Cor de fundo da barra de navegação */
                            width: 100%;
                        }

                        .navbar a {
                            margin-right: 10px;
                        }

                        /* Adicione o seguinte estilo para o botão Metrics */
                        .navbar a[href="/metrics"] {
                            text-decoration: none;
                            color: #333;
                            padding: 8px 16px;
                            border: 2px solid #000; /* Adiciona uma borda preta de 2px */
                            border-radius: 10px; /* Adiciona bordas arredondadas */
                            transition: background-color 300ms;
                        }

                        /* Conteúdo principal */
                        .content {
                            background-color: #BE7839; /* #D8923B Cor de fundo do conteúdo principal */
                            padding: 7px;
                            margin: 7px;
                            border-radius: 15px;
                            box-shadow: 0 4px 4px rgba(0, 0, 0, 0.1);
                            width: 30%; /* Ajuste o tamanho conforme necessário */
                        }

                        /* Alteração para a cor branca do texto */
                        .white-text p {
                            color: white;
                        }

                        /* Footer styles */
                        .footer {
                            display: flex;
                            justify-content: space-around;
                            align-items: center;
                            padding: 5px;
                            background-color: white; /* Cor de fundo do rodapé */
                            color: #333;
                            width: 100%;
                        }

                        .social-links img {
                            height: 30px;
                            width: 30px;
                            border-radius: 25%; /* Adiciona borda circular às imagens */
                            margin-right: 5px;
                            cursor: pointer;
                        }

                        /* Media query for smaller screens */
                        @media (max-width: 700px) {
                            .content {
                                width: 90%;
                            }
                        }

                        @media (max-width: 500px) {
                            .content {
                                width: 85%;
                            }
                        }
                    </style>
                </head>
                <body>
                    <div class="navbar">
                        <a href="/metrics" style="text-decoration: none; color: #333;">Metrics</a>
                    </div>
                    <div class="content">
                        <h1>Live DevOps!!</h1>
                        <div class="image-container">
                            <img class="responsive-image" src="https://raw.githubusercontent.com/flowramps/workshop-argo/main/img/flowramps.jpg" alt="Imagem">
                        </div>
                        <div class="white-text">
                            <p>Nome do Pod: ` + hostname + `</p>
                        </div>
                    </div>
                    <div class="footer">
                        <div class="social-links">
                            <a href="https://www.instagram.com/flow.ramps/" title="Instagram" target="_blank" rel="noopener"><img src="https://raw.githubusercontent.com/flowramps/workshop-argo/main/img/instagram.png"></a>
                            <a href="https://www.linkedin.com/in/rafaelrampasso/" title="LinkedIn" target="_blank" rel="noopener"><img src="https://raw.githubusercontent.com/flowramps/workshop-argo/main/img/linkedin.png"></a>
                            <a href="https://github.com/flowramps/" title="Github" target="_blank" rel="noopener"><img src="https://raw.githubusercontent.com/flowramps/workshop-argo/main/img/github.png"></a>
                        </div>
                    </div>
                    <script>
                        $(document).ready(function() {
                            $(".social-links img").click(function() {
                                var socialMedia = $(this).parent().attr("title").toLowerCase();
                                $.get("/increment-" + socialMedia + "-counter", function(data) {
                                    console.log("Contador do " + socialMedia + " incrementado");
                                });
                            });
                        });
                    </script>
                </body>
            </html>
        `))
	})

	http.HandleFunc("/increment-instagram-counter", func(w http.ResponseWriter, r *http.Request) {
		metricsMtx.Lock()
		defer metricsMtx.Unlock()
		instagramClicks.Inc()
		log.Println("Contador do Instagram incrementado.")
		w.Write([]byte("Contador do Instagram incrementado"))
	})

	http.HandleFunc("/increment-linkedin-counter", func(w http.ResponseWriter, r *http.Request) {
		metricsMtx.Lock()
		defer metricsMtx.Unlock()
		linkedinClicks.Inc()
		log.Println("Contador do LinkedIn incrementado.")
		w.Write([]byte("Contador do LinkedIn incrementado"))
	})

	http.HandleFunc("/increment-github-counter", func(w http.ResponseWriter, r *http.Request) {
		metricsMtx.Lock()
		defer metricsMtx.Unlock()
		githubClicks.Inc()
		log.Println("Contador do GitHub incrementado.")
		w.Write([]byte("Contador do GitHub incrementado"))
	})

	http.Handle("/metrics", promhttp.Handler())
	log.Println("Servidor rodando em :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
