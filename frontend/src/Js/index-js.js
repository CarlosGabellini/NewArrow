//Aqui vamos comecar a listar as musicas e colocar elas no html;

import { RetornandoLista } from '../../wailsjs/go/main/App.js';

//Atribuindo a uma variavel o array de objetos que fiz;

const _lista_de_musicas = await RetornandoLista();

let nome_music = document.querySelector(".nome-musica");

nome_music.textContent = _lista_de_musicas[0].nome_musica;
