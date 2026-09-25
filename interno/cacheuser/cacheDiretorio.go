package cacheuser

import (
	"fmt"
	"os"
	"path/filepath"
)

/*----------------------------------------- CacheDiretorio ------------------------------------------- 
	As funcoes abaixo tem como objetivo principal achar a pasta de Cache do usuario e manipular ela, crian
do o diretorio NewArrow que vai juntar e guardar dados essenciais para as musicas rodarem, as funcoes tem 
como total objetivo a organizacao, manipulacao e suporte para arquivos e diretorios em Cache.
------------------------------------------------------------------------------------------------------
 */

func OrganizandoCache(setName string) (string, error) {
	MeuCache, err := os.UserCacheDir()
	
	if err != nil {
		return "", err
	}

	/*------------------------------------------------------------------------------------------------
		Ao inves de colocar setName para a funcao, eu poderia fazer diretamente o nome da pasta e criar
	ela numa unica funcao unificadora, tomei a descisao por convecao, nao vou alterar totalmente o codi-
	-go para isso, tenho que tomar cuidado com o tratamento de erros para nao encontrar novos bugs.
	--------------------------------------------------------------------------------------------------
	 */
	
	MeuPerfilCache := filepath.Join(MeuCache, setName)
	err = os.Mkdir(MeuPerfilCache, 0755)

	if err != nil && !os.IsExist(err) {
		return "", err
	}

	//Se chegou aqui, ou criou com sucesso, ou já existia — ambos são OK.
	return MeuPerfilCache, nil
}

/* Retorna o caminho da pasta do diretorio do NewArrow. */
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
		return ""
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

//Criando a pasta album para colocar as fotos das musicas depois;
func CriandoAPastaAlbum() string {
	CaminhoCache, err := CriandoMeuDiretorioCache()
	const meusAlbuns string = "photos-album"

	if err != nil {
		return ""
	}

	DiretorioAlbum := CrieSubdiretorios(CaminhoCache, meusAlbuns)

	if DiretorioAlbum == "" {
		return ""
	}

	return DiretorioAlbum
}

//Cria e retorna a pasta filho a partir da pasta album, organizando as fotos das musicas
// por diretorio;
func PastaAlbumDir(setNome string) string {
	PastaAlbum := CriandoAPastaAlbum()
	CaminhoDaPasta := CrieSubdiretorios(PastaAlbum, setNome)

	//Validando o erro corretamente!
	if CaminhoDaPasta == "" || PastaAlbum == "" {
		return ""
	}

	return CaminhoDaPasta
}