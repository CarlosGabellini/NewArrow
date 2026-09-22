package cacheuser

import (
	"fmt"
	"os"
	"path/filepath"
)

func OrganizandoCache(setName string) (string, error) {
	MeuCache, err := os.UserCacheDir()

	if err != nil {
		return "", err
	}

	MeuPerfilCache := filepath.Join(MeuCache, setName)

	//Aqui estou criando a pasta a partir de onde esta o cache do usuario.
	// No Windows isso fica em AppData/
	err = os.Mkdir(MeuPerfilCache, 0755)

	if err != nil {

		if os.IsExist(err) {
			return MeuPerfilCache, err
		}

		return "", err
	}

	//Aqui fazendo os tratamentos de erros de forma adequada para o arquivo.
	return MeuPerfilCache, err
}

func CriandoMeuDiretorioCache() (string, error) {
	caminhoCache, err := OrganizandoCache("NewArrow")

	if err != nil {
		return "", err
	}

	return caminhoCache, err
}

//Essa funcao nao somente serve para o cache, mas sim para criar qualquer subpasta dentro de outra pasta;
func CrieSubdiretorios(pastaPai, pastaFilho string) string {
	//Aqui serve para criar subdiretorios no cache caso eu va prescisar!
	// Lembrando que eh importante deixar isso evidente em algum arquivo para olhar no cache
	// depois!

	if pastaPai == "" {
		return ""
	}
	
	CaminhoSeguido := filepath.Join(pastaPai, pastaFilho)

	//Usando o os.Mkdir para criar as pastas das quais necessito depois no FrontEnd.
	err1 := os.MkdirAll(CaminhoSeguido, 0755)

	//Usando MkdirALl, por que ele eh melhor caso o diretorio ja tenha feito!
	if err1 != nil {
		return fmt.Sprintln(err1)
	}

	return CaminhoSeguido
}

//Pasta aonde vao ficar as musicas do usuario;
func CriandoA_pastaJSON() string {
	CaminhoCache, err := CriandoMeuDiretorioCache()
	const meuJSON string = "MyJSON"

	if err != nil {
		return fmt.Sprintln(err)
	}

	DiretorioJSON := CrieSubdiretorios(CaminhoCache, meuJSON)

	return DiretorioJSON
}