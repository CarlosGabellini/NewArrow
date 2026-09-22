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

	err = os.Mkdir(MeuPerfilCache, 0755)

	if err != nil && !os.IsExist(err) {
		return "", err
	}

	// Se chegou aqui, ou criou com sucesso, ou já existia — ambos são OK.
	return MeuPerfilCache, nil
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
		return ""
	}

	DiretorioJSON := CrieSubdiretorios(CaminhoCache, meuJSON)

	return DiretorioJSON
}