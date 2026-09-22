//Implementando o JS nessa pagina aqui!

import { Player } from "./web-audio.js";
import { AtualizarCache, RetornandoLista } from "../../wailsjs/go/main/App.js";
import { ListarDiretorios } from "../../wailsjs/go/main/App.js";

//Atribuindo uma variavel a classe do audio que eu mesmo criei;
const MeuPlayer = new Player;

//Os query Selector aqui embaixo;
const scan_button = document.querySelector(".scan-button")
const meus_diretorios = document.querySelector(".directories-list");
const A1_lista_de_diretorios = await ListarDiretorios();
const minha_track_list = document.querySelector(".track-list");

const pausar_retomar = document.getElementById("playBtn");

let barra_progresso = document.querySelector(".progress-bar-fill");
let barra_sem_preencher = document.querySelector(".progress-bar");

/*---------------------------------Async Functions ------------------------------------------
  Aqui colocar a funcao de listar a pasta do golang, para ver se eu consegui criar o arquivo
JSON no cache do usuario, isso vai servir depois na hora de empacotar o binario;

---------------------------------------------------------------------------------------------
*/

async function InjetarJSON() {
  try {
    await AtualizarCache();

  } catch (error) {
    console.log("error: ", error)
  } 
}

scan_button.addEventListener("click", () => {
  let myMusics = InjetarJSON();
  console.log(myMusics)
})

//Retornando a lista de musica;
let _My_list = []

//Sempre use try e catch, se nao usar, pode recusar e quebra o codigo todo!
//(Experiencia propria!)

try {
  _My_list = await RetornandoLista();
  
} catch (error) {
  console.log("Ainda não há cache, ou erro ao ler:", error);
}

//Injetando os diretorios que vou colocar depois;

if (A1_lista_de_diretorios.length < 7) {
  
  for (let i = 0; i < A1_lista_de_diretorios.length; i++) {
    
      const li = document.createElement("li");
      li.textContent = A1_lista_de_diretorios[i];
      meus_diretorios.appendChild(li);
  }
}
//Injetando as musicas aqui nessa parte do for;

for (let a = 0; a < _My_list.length; a++) {
  const _lista = document.createElement("li");

  const numero_lista = document.createElement("span");
  numero_lista.className = "trackNumber";
  numero_lista.textContent = `${a + 1}. `;

  //Agora ficou mais organizado, em _lista, injetamos o nome_musica e artista;
  _lista.appendChild(numero_lista);
  _lista.append(` ${_My_list[a].nome_musica} - ${_My_list[a].artista}`);

  minha_track_list.append(_lista);

  _lista.addEventListener("click", () => {
    MeuPlayer.load(`/music/${encodeURIComponent(_My_list[a].caminho_path)}`);
    MeuPlayer.play();
  })
}

pausar_retomar.addEventListener("click", () => {
  MeuPlayer.toggle();
})

MeuPlayer.onProgress = (currentTime, duration) => {
  let porcentagem = (currentTime / duration) * 100;
  barra_progresso.style.width = `${porcentagem}%`;
}