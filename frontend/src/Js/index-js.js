import { ToqueMusica } from '../../wailsjs/go/main/App.js';

const tocarMusica = document.getElementById("btn-player");

async function tocar() {
    try {
      await ToqueMusica();
      console.log('Tocando!');
    } catch (err) {
    console.log("erro ao tocar: ", err);
    }
}


tocarMusica.addEventListener("click", tocar);