//Aqui vamos comecar a listar as musicas e colocar elas no html;

/*Importando as listas e colocar ela em um valor constante para tocar ela depois.*/

import { RetornandoLista } from '../../wailsjs/go/main/App.js';
import { Player } from './web-audio.js';

//Atribuindo a uma variavel uma classe para objetos, este eh a classe do meu player de musica que eu 
// fiz no JS, ela eh responsavel por tocar as musicas.
const _meu_player_music = new Player;

//Retornando a lista de musica que o go fez;
const _lista_de_musicas = await RetornandoLista();

//Os query Selector ficam nessa parte aqui!
const music_painel = document.querySelector(".music-painel");
const barra_progresso = document.querySelector(".barra-progresso");
const barra_progresso_interna = document.querySelector(".barra-progresso-preenchida");

//A lista do for implementa a lista de musica que o go fez em um loop e adiciona ao html.

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
  let numero_da_musica = parseInt(elemento.dataset.index);

  _meu_player_music.load(`/music/${encodeURIComponent(_lista_de_musicas[numero_da_musica].caminho_path)}`);
  _meu_player_music.play();
});

barra_progresso.addEventListener("click", (event) => {
  
  let retangulo = barra_progresso.getBoundingClientRect();
  
  let clique_x = event.clientX - retangulo.left; // posição do clique relativa à barra
  let porcentagem_clicada = clique_x / retangulo.width;
  
  let novo_tempo = porcentagem_clicada * _meu_player_music.audio.duration;
    _meu_player_music.seek(novo_tempo);
})


_meu_player_music.onProgress = (currentTime, duration) => {
  let porcentagem = currentTime / duration * 100;
  barra_progresso_interna.style.width = `${porcentagem}%`;
}