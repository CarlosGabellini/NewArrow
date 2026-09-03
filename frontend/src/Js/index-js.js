//Aqui vamos comecar a listar as musicas e colocar elas no html;

/*Importando as listas e colocar ela em um valor constante para tocar ela depois.*/

import { RetornandoLista } from '../../wailsjs/go/main/App.js';
import { PlayTrack } from '../../wailsjs/go/main/App.js';

let estaTocando = false;
let informacao_console = true;

async function TocarTrack(caminho) {
  try {
    PlayTrack(caminho)
  } catch (error) {
    console.log(`error: ${error}`);
  }
}

//Atribuindo a uma variavel o array de objetos que fiz;

const _lista_de_musicas = await RetornandoLista();
const music_painel = document.querySelector(".music-painel");


for (let i = 0; i < _lista_de_musicas.length; i++) {
  let faixa_de_musica = document.createElement("div");
  faixa_de_musica.className = "faixa-de-musica";
  
  let _nome_musica = document.createElement("div");
  _nome_musica.className = "nome-musica";

  let _nome_album = document.createElement("div");
  _nome_album.className = "nome-album";
  _nome_album.textContent = _lista_de_musicas[i].nome_album;

  let _artista = document.createElement("div");
  _artista.className = "artista";
  _artista.textContent = _lista_de_musicas[i].artista;

  _nome_musica.textContent = _lista_de_musicas[i].nome_musica;
  faixa_de_musica.appendChild(_nome_musica);
  faixa_de_musica.appendChild(_nome_album);
  faixa_de_musica.appendChild(_artista);

  faixa_de_musica.dataset.index = i;

  music_painel.appendChild(faixa_de_musica);
}

music_painel.addEventListener("click", (event) => {
  let elemento = event.target.closest(".faixa-de-musica");
  if (!elemento) return;
  
  let numero_da_musica = elemento.dataset.index;
  parseInt(numero_da_musica);

  TocarTrack(_lista_de_musicas[numero_da_musica].caminho_path);
});