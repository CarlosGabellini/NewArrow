# NewArrow

## In portuguese

NewArrow eh um player de musica offline simples construido a partir do Go + JS, onde usamos
o framework Wails para construi-lo, ele eh feito para proporcionar conforto e facilidade de
uso.

## Arquitetura

O JS cuida do player e dos estados que ele faz, enquanto o go eh responsavel pela organizacao
de diretorios, o arquivo JSON em que eh armazenado as musicas, e responsavel pelo binario em-
pacotado.
A arvore do projeto pode ser vista abaixo:
```bash
.
├── app.go
├── build
│   ├── appicon.png
│   ├── bin
│   ├── darwin
│   ├── README.md
│   └── windows
├── frontend
│   ├── dist
│   ├── index.html
│   ├── links-para-guardar.txt
│   ├── node_modules
│   ├── package-lock.json
│   ├── package.json
│   ├── package.json.md5
│   ├── src
│   ├── vite.config.js
│   └── wailsjs
├── go.mod
├── go.sum
├── interno
│   ├── cacheuser
│   └── dir1org
├── main.go
├── myJSONfile.json
├── README.md
└── wails.json
```

Ela esta assim __neste recente momento__.