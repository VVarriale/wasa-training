# wasa-training
# WASA — Web and Software Architecture (Exercises Repository)

Repository delle esercitazioni del corso WASA (Web and Software Architecture).  
Contiene esercizi in:

- HTTP
- JSON / YAML
- API Design
- Go (interfaces, packages, modules)
- HTML, CSS
- JavaScript
- Vue.js

Tutti gli esercizi sono organizzati in cartelle per modulo e includono codice, note e piccoli progetti incrementali.

## Struttura

- `http/` — richieste, risposte, parsing
- `json-yaml/` — lettura, scrittura, conversione
- `api-design/` — design REST e OpenAPI
- `go/` — interfacce, pacchetti, moduli, testing
- `html-css/` — layout e stili
- `javascript/` — basi JS, DOM, async
- `vue/` — mini app e SPA di base

## Struttura directory professionale
wasa-exercises/
    │
    ├── README.md
    ├── .gitignore
    ├── .editorconfig
    ├── go.work
    │
├── http/
│   ├── request-parsing.go
│   ├── simple-http-client.go
│   └── README.md
│
├── json-yaml/
│   ├── json-read.go
│   ├── json-write.go
│   ├── yaml-to-json.go
│   └── README.md
│
├── api-design/
│   ├── rest-api-basics.md
│   ├── openapi-example.yaml
│   └── README.md
│
├── go/
│   ├── interfaces/
│   │   ├── shapes.go
│   │   ├── main.go
│   │   └── README.md
│   ├── packages/
│   │   ├── mathutils/
│   │   │   ├── add.go
│   │   │   └── README.md
│   │   ├── main.go
│   ├── errors/
│   ├── concurrency/
│   ├── testing/
│
├── html-css/
│   ├── index.html
│   ├── styles.css
│   └── README.md
│
├── javascript/
│   ├── basics.js
│   ├── dom.js
│   ├── async.js
│   └── README.md
│
└── vue/
    ├── hello-vue/
    │   ├── index.html
    │   ├── app.js
    │   └── README.md
    ├── mini-spa/
    │   ├── index.html
    │   ├── router.js
    │   └── README.md

## Strumenti consigliati
- VSCode + Go, Copilot, CodeGPT
- Git / GitHub
- Node.js (per esercizi Vue.js)
