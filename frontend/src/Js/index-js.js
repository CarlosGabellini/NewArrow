import { PlayTrack } from '../../wailsjs/go/main/App.js';

//Os imports devem ter EXATAMENTE o nome da funcao exata (Eh case sensitive); e tambem colocar
//o caminho exato, na hora de gerar o binario tambem vou ter que dar uma olhada sobre como isso
//funciona para distribuicao;

const tocarMusica = document.getElementById("btn-player");

async function tocar() {
    try {
      await PlayTrack();
      console.log('Tocando!');
    } catch (err) {
    console.log("erro ao tocar: ", err);
    }
}

//Este eh um exemplo de como a funcao no JS deve funcionar, sempre acompanhada do Async e await;
tocarMusica.addEventListener("click", tocar);