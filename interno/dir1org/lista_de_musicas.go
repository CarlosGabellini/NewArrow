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

func Pasta_de_musica() (string, error) {
	diretorioHome, err := os.UserHomeDir()

	if err != nil {
		return "", nil
	}
	
	return filepath.Join(diretorioHome, "Music"), nil
}