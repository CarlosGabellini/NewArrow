import { GerarNumero } from '../../wailsjs/go/main/App.js';

const numerosGerados = document.getElementById("numeros-gerados")
const botaoGerador = document.querySelector(".botao1")


botaoGerador.addEventListener("click", async () => { 
  let Number1 = await GerarNumero()
  numerosGerados.textContent = Number1
})