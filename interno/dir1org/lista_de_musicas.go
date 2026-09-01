package dir1org

import (
	"os"
	"path/filepath"
)

type ListaMusicas struct {
	Nome_album string `json:"nome_album"`
	Nome_da_musica string `json:"nome_musica"`
	Artista string `json:"artista"`
	Ano int `json:"ano_musica"`
	Caminho_path string `json:"caminho_path"`
	ModTime int64 `json:"ModTime"`
	Size int64 `json:"Size"`
}

func EncontreHome() (string, error) {
	diretorioHome, err := os.UserHomeDir()

	if err != nil {
		return "", nil
	}
	
	return filepath.Join(diretorioHome, "Music"), nil
}

/*-----------------------------------------Sobre VerOSdiretorios_musics-----------------------------------
	Basicamente essa funcao vai me ajudar a colocar os nomes no frontEnd de maneira correta, antes de realmente
	abrir os diretorios e ver as musicas no FrontEnd, serve basicamente para me ajudar a listar de forma correta
	sem depender da funcao principal.

	------------------------------------------------------------------------------------------------------
 */

func VerOsdiretorios_musics(home string) ([]string, error) {
	diretorios, err := os.ReadDir(home)
	if err != nil {
		return nil, err
	}

	var diretoriosMusicas []string

	for _, diretorio := range diretorios {
		if diretorio.IsDir() {
			diretoriosMusicas = append(diretoriosMusicas, diretorio.Name())
		}
	}

	return diretoriosMusicas, nil
}