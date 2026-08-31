package dir1org

import (
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/dhowden/tag"
)

const (
	tarefas = int(10)
	WorkersOnline = int(8)
)

func Liste_a_pasta(diretorio string) ([]ListeMusicas, error) {

	_caminho := make(chan string, tarefas)				//Canal com 10 buffers;
	_saida_pasta := make(chan ListeMusicas, tarefas)	//10 buffers da struct que eu coloquei;
	var _lista_de_musicas []ListeMusicas
	var _waitG sync.WaitGroup

	for i := 0; i < WorkersOnline; i++ {
		_waitG.Add(1)
		go processarArquivo(_caminho, _saida_pasta, &_waitG)		//Lembre-se que nao eh procedural;
	}

	//Enquanto processa aqui, a funcao principal esta livre para fazer outras coisas;

	/*Aqui vamos criar uma funcao anonima que vai escanear os diretorios das goroutines que encontrarmos, 
	 entao isso vai para o canal de cima para processar o diretorio que encontrarmos, depois fechamos o caminho
	 da funcao e retornamos com a nossa lista de musicas;
	 */
	
	go func() {

		//Aqui fizemos uma funcao anonima em go para escanear o diretorio totalmente de uma vez, na qual assim
		// que ele comeca a nossa outra funcao processarArquivo ja vai comecar a escanear a musica tambem a 
		// partir do mesmo canal.
		filepath.WalkDir(diretorio, func(path string, d os.DirEntry, err error) error {

			if err != nil {
				return err
			}

			if !d.IsDir() && strings.ToLower(filepath.Ext(path)) == ".mp3" {
				_caminho <- path		//Canal que a gente vai colocar o nosso codigo;
			}

			return nil
		})

		close(_caminho)			//fechando o canal para parar de receber dados;
	}()

	go func() {
		_waitG.Wait()
		close(_saida_pasta)
	}()

	for musicas := range _saida_pasta {
		_lista_de_musicas = append(_lista_de_musicas, musicas)
	}

	return _lista_de_musicas, nil
}

func processarArquivo(wayPath <-chan string, saidas chan <- ListeMusicas, wg *sync.WaitGroup) {
	defer wg.Done()

	for path := range wayPath {
		meta, err := ColocarMetadados(path)

		if err != nil {
			continue			//Caso isso possa dar um erro;
		}

		saidas <- meta
	}
}

func ColocarMetadados(WayPath string) (ListeMusicas, error) {

	f, err := os.Open(WayPath)
	var _saidasMetadados ListeMusicas

	if err != nil {
		return _saidasMetadados, err
	}

	defer f.Close()

	m, err := tag.ReadFrom(f);

	if err != nil {
		return ListeMusicas{}, err
	}

	_saidasMetadados.Ano = m.Year()
	_saidasMetadados.Artista = m.Artist()
	_saidasMetadados.Nome_album = m.Album()
	_saidasMetadados.Nome_da_musica = m.Title()

	return _saidasMetadados, nil
}