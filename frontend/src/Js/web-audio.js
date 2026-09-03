//Basicamente desisti do beep, muito complicado de ficar mechendo e bastante dor de cabeca, 
// deixa o JS cuidar disso, o projeto foi feito para ser simples de entender, e nao algo complexo;

export class Player {
  constructor() {
    this.audio = new Audio();
    this.audio.addEventListener('timeupdate', () => this.onProgress?.(this.audio.currentTime, this.audio.duration));
    this.audio.addEventListener('ended', () => this.onEnded?.());
  }

  load(filePath) {
    // Wails: exponha um binding que sirva o arquivo, ex. via asset server ou file:// 
    this.audio.src = filePath;
  }

  play() { this.audio.play().catch(err => console.error("Erro ao tocar:", err)); }
  pause() { this.audio.pause(); }
  toggle() { this.audio.paused ? this.play() : this.pause(); }

  setVolume(v) { this.audio.volume = v; } // 0.0 a 1.0
  seek(seconds) { this.audio.currentTime = seconds; }
}