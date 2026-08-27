import { PlayTrack } from '../../wailsjs/go/main/App.js';
import { TogglePause } from '../../wailsjs/go/main/App.js';

//Os imports devem ter EXATAMENTE o nome da funcao exata (Eh case sensitive); e tambem colocar
//o caminho exato, na hora de gerar o binario tambem vou ter que dar uma olhada sobre como isso
//funciona para distribuicao;

const tocarMusica = document.getElementById("btn-player");
let estaTocando = false;
let informationConsole = true;

async function tocarMusic() {
  
  try { 

    //Se a musica nao esta tocando, entao nao existe a opcao de pausar a musica;
  
    if (estaTocando) {
      await TogglePause();

      if (informationConsole) {
        console.log('Pausado!');
        informationConsole = false;
        
      } else {
        console.log("despausado!");
        informationConsole = true;
      }
      
      estaTocando = true;
    }

    //Aqui comeca a tocar a musica e fazer a comunicacao com o go;

    else {
      await PlayTrack();
      console.log("musica tocando agora!");
      estaTocando = true;
    }
  }
 
  catch (error) {
    console.log("erro: " + error);
  }
}

tocarMusica.addEventListener("click", tocarMusic);

//Este eh um exemplo de como a funcao no JS deve funcionar, sempre acompanhada do Async e await;