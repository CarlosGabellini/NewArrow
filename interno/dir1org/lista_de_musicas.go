package dir1org

import (
	"os"
	"path/filepath"
)

//Struct principal para escaneamento de dados.
type ListaMusicas struct {
	Nome_album string `json:"nome_album"`
	Nome_da_musica string `json:"nome_musica"`
	Artista string `json:"artista"`
	Ano int `json:"ano_musica"`
	Caminho_path string `json:"caminho_path"`
	ModTime int64 `json:"ModTime"`
	Size int64 `json:"Size"`
	Diretorio string `json:"diretorio"`
}

//Funcao redundante para encontrar a home, provavelmente vou excluir ela alguma hora;
func EncontreHome() (string, error) {
	diretorioHome, err := os.UserHomeDir()

	if err != nil {
		return "", nil
	}
	
	return diretorioHome, err
}

/*-----------------------------------------Sobre VerOSdiretorios_musics-----------------------------------
	Basicamente essa funcao vai me ajudar a colocar os nomes no frontEnd de maneira correta, antes de realmente
	abrir os diretorios e ver as musicas no FrontEnd, serve basicamente para me ajudar a listar de forma correta
	sem depender da funcao principal.

	------------------------------------------------------------------------------------------------------
 */

func VerOsdiretorios_musics(home string) ([]string, error) {
	WayMusics := filepath.Join(home, "Music")
	var MeusDiretorios []string

	//Entrando no filepath para escanear o nome dos diretorios;
	err := filepath.WalkDir(WayMusics, func(path string, d os.DirEntry, err error) error {

		if err != nil {
			return err
		}

		if d.IsDir() && path != WayMusics {
			MeusDiretorios = append(MeusDiretorios, path)
		}

		return nil
	})

	//Aqui eh somente para ficar o nome do diretorio, e nao o caminho completo! para ficar melhor pro
	// FrontEnd.
	for i := 0; i < len(MeusDiretorios); i++ {
		_, NomeDir := filepath.Split(MeusDiretorios[i])
		MeusDiretorios[i] = NomeDir
	}

	return MeusDiretorios, err
}