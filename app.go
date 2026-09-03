package main

import (
	"NewArrow/interno/dir1org"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

//Funcao importante que vai escanear o JSON e retornar a lista de musicas diretamente para
//o JavaScript.

func (a *App) RetornandoLista() ([]dir1org.ListaMusicas, error) {
	var CaminhoJSON string = "./myJSONfile.json"
	lista, err := dir1org.EscanearJSON(CaminhoJSON)

	if err != nil {
		return nil, err
	}

	return lista, nil
}

func (a *App) AtualizarCache() (string, error) {
	diretorioHome, err := dir1org.EncontreHome()

	if err != nil {
		return "", err
	}

	lista2, err := dir1org.Liste_a_pasta(diretorioHome)
	if err != nil {
		return "", err
	}

	var sucessoMensagem string = fmt.Sprintf("O arquivo foi um sucesso! primeira musica %s", 
		lista2[0].Nome_da_musica)

	return sucessoMensagem, nil
}

func (a *App) ListarDiretorios() ([]string, error) {
	caminhoHOME, err := dir1org.EncontreHome()

	if err != nil {
		return nil, err
	}
	
	diretoriosEncontrados, err := dir1org.VerOsdiretorios_musics(caminhoHOME)

	if err != nil {
		return nil, err
	}
	
	return diretoriosEncontrados, err
}

/*-----------------------------Sobre a funcao NewMusicAssetHandler() -------------------------
	Essa funcao retornar um handler http, ou seja, ela nao processa nada sozinha, ela fabrica algo
	capaz de responder requisicoes. Eh tipo uma fabrica de atendentes que sabem responder pedidos
	do tipo "Me de os arquivos de musica".
 */

func NewMusicAssetHandler() http.Handler {
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//Fazendo uma funcao emcapsulada a partir do return;

		fmt.Println("Handler chamado, path:", r.URL.Path) // debug
		
		filePath := strings.TrimPrefix(r.URL.Path, "/music/")
		filePath, err := url.QueryUnescape(filePath)
		
		if err != nil {
			http.Error(w, "caminho invalido", http.StatusBadRequest)
			return
		}

		f, err := os.Open(filePath)
		
		if err != nil {
			http.Error(w, "arquivo nao encontrado", http.StatusNotFound)
			return
		}
		
		defer f.Close()

		w.Header().Set("Content-Type", "audio/mpeg")
		http.ServeContent(w, r, filePath, time.Time{}, f)
	})
}