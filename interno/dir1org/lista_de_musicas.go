package dir1org

import (
	"NewArrow/interno/cacheuser"
	"fmt"
	"os"
	"path/filepath"

	"github.com/dhowden/tag"
)

//Struct principal para escaneamento de dados.
type ListaMusicas struct {
	Nome_album string `json:"nome_album"`
	Nome_da_musica string `json:"nome_musica"`
	Artista string `json:"artista"`
	Ano int `json:"ano_musica"`
	Caminho_path string `json:"caminho_path"`
	ModTime int64 `json:"ModTime"`			//Obrigatorio ModTime e Size serem int64 especificamente;
	Size int64 `json:"Size"`
	Diretorio string `json:"diretorio"`
}

//Cria a lista de struct para o JS sobre onde fica o caminho correto das fotos de albuns;
type AlbunsFoto struct {
	Caminho string
	Album string
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
	Basicamente essa funcao vai me ajudar a colocar os nomes no frontEnd de maneira correta, antes de 
	realmente abrir os diretorios e ver as musicas no FrontEnd, serve basicamente para me ajudar a lis-
	tar de forma correta sem depender da funcao principal.
	------------------------------------------------------------------------------------------------------
 */

//Fico me perguntando se na funcao abaixo nao seria melhor fazer isso em uma goroutine ou escanear os
//diretorios que ja tenho a partir do JSON, toda vez que abre o aplicativo essa funcao vai ser chamada.

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

/*--------------------------------- Escaneamento da foto do album --------------------------------------- 
	As funcoes a seguir tem haver com o escaneamento da foto do album, as funcoes abaixo serao essenciais
para que o FrontEnd consiga carregar as fotos das musicas depois, vou colocar em um arquivo de cache para
colocar as fotos.
	A funcao EscanearFotosAlbum tem como objetivo criar diretorios e arquivos onde vamos extrair a capa
dos albuns e colocar em uma pasta Cache do usuario, entao o JS vai saber o caminho para carregar a imagem
a partir do caminho que a gente fez.

	Eu criei uma struct propria que retorna exatamente onde esta a pasta com as fotos das musicas que guar-
-damos a partir dessa funcao.
---------------------------------------------------------------------------------------------------------
 */

//Escaneia e cria os diretorios com as fotos dos albuns para usar depois;
func EscanearFotosAlbum() {

	MinhaHome, err1 := os.UserHomeDir()
	CaminhoDasPastas := make([]string, 0, 10)	//Nome dos diretorios de mesmo nome na pasta userCache()
	
	if err1 != nil {
		fmt.Println(err1)
	}
	
	MeusDiretorios, err := VerOsdiretorios_musics(MinhaHome)
	MinhasMusicas, err2 := EscanearJSON(cacheuser.WayJSON_file())

	if err != nil {
		fmt.Println(err)
	}

	if err2 != nil {
		fmt.Println(err2)
	}

	for indice := range MeusDiretorios {
		CaminhoDasPastas = append(CaminhoDasPastas, cacheuser.PastaAlbumDir(MeusDiretorios[indice]))

		if CaminhoDasPastas[indice] == "" {
			fmt.Println("Criacao da pasta deu errado!")
			break
		}
	}

	for _, Musics := range MinhasMusicas {

		f, err := os.Open(Musics.Caminho_path)

		if err != nil {
			fmt.Println(err)
			continue
		}
		
		m, err := tag.ReadFrom(f)

		//Fechando o arquivo aqui para nao sobrecarregar o SO.
		//Ele deve ser colocado aqui para podermos ler com o tag primeiro.
		f.Close()
		
		if err != nil {
			fmt.Println(err)
			continue
		}

		pintura := m.Picture()

		if pintura == nil {
			//Essa musica nao tem capa embutida no arquivo;
			continue
		}

		ext := pintura.Ext

		if ext == "" {
			ext = "jpeg"
		}

		for indice2 := range CaminhoDasPastas {

			//Extrai somente o nome da pasta no cache, ela deve ser exatamente igual ao Musics.Diretorio, 
			//para colocar no lugar corretamente;
			_, Comparacao := filepath.Split(CaminhoDasPastas[indice2])

			if Comparacao == Musics.Diretorio {

				//Criando o nome do arquivo aqui!
				extensaoArquivo := Musics.Nome_da_musica + "." + ext
				CaminhoCapa := filepath.Join(CaminhoDasPastas[indice2], extensaoArquivo)

				if _, err := os.Stat(CaminhoCapa); err == nil {
					continue	//Capa ja existe, pule
				}

				if err := os.WriteFile(CaminhoCapa, pintura.Data, 0644); err != nil {
					fmt.Println(err)
					continue
				}
			}
		}
	}
}