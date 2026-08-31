export namespace dir1org {
	
	export class ListaMusicas {
	    nome_album: string;
	    nome_musica: string;
	    artista: string;
	    ano_musica: number;
	    caminho_path: string;
	    ModTime: number;
	    Size: number;
	
	    static createFrom(source: any = {}) {
	        return new ListaMusicas(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.nome_album = source["nome_album"];
	        this.nome_musica = source["nome_musica"];
	        this.artista = source["artista"];
	        this.ano_musica = source["ano_musica"];
	        this.caminho_path = source["caminho_path"];
	        this.ModTime = source["ModTime"];
	        this.Size = source["Size"];
	    }
	}

}

