package dir1org

import (
	"encoding/json"
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

/*-----------------------------------Biblioteca de escaneamento de diretorio------------------------------//

	Basicamente essa biblioteca aqui serve para criar funcoes para escanear em um JSON e ser consumida pelo
	JavaScript depois, a arquitetura foi desenhada para que o cache e o arquivo nao prescise ficar sendo sobrees-
	-crito toda vez que rodar o aplicativo em busca da biblioteca.
	As pastas de musicas principal deve ficar em Os.HomeDir() onde fica a pasta "Music", assim lemos aquela 
	pasta onde vai ficar as playlists do proprio usuario.

	A funcao Liste_a_pasta:

	Ela eh o coracao, deve retornar um array de lista de musicas ja encaminhados para o JSON, onde o JSON vai
	ser lido pelo JavaScript e entao vai tocar a musica pela outra biblioteca que ja temos.
	Utilizamos 2 canais de buffer para fazer o processamento, alem de uma funcao anomima que vai escanear o 
	diretorio da HOME. Relembrando que utilizamos goroutines aqui para processamento paralelo. (Fica melhor para
	nao sobrecarregar responsabilidades).

	A funcao processarArquivo:
	Basicamente somente faz um for simples dos 10 buffers que temos para escanear os metadados na outra funcao, 
	esse metadados escaneia os nossos arquivos .mp3 (A leitura somente acontece para o .mp3), e volta em um array
	para a funcao principal a partir do buffer que criamos que se chama _saida_pasta.
	
	_____________________________________________________________________________________________________________
 */

func EscanearJSON(arquivo string) ([]ListaMusicas, error) {
	//Aqui vou escanear o JSON.
	var _lista_de_musicas []ListaMusicas

	f, err := os.Open(arquivo)
	if err != nil {
		return _lista_de_musicas, err
	}
	defer f.Close()

	chave_validacao := make(map[string]ListaMusicas)

	/*
		Basicamente o make(map[string]ListaMusicas) cria um caminho para armazenar cada struct 
	que a gente fez no nosso JSON, pense no map[string] como map"C://users/felip/Musics", assim
	a gente pega o caminho do arquivo, coloca em um map e guarda numa lista.
		Depois que a gente guarda nessa lista, somente prescisa fazer um for simples onde atribuimos
	ao nosso array de struct para devolver ao JavaScript depois, mas por que fazer isso? Simplesmente
	para nao ter que ficar atualizando o cache toda hora nos nossos arquivos e nem ter que ficar
	reecrevendo o disco, eh mais rapido para abrir a aplicacao.
	
	 */

	if err := json.NewDecoder(f).Decode(&chave_validacao); err != nil {
		return _lista_de_musicas, err
	}

	for _, v := range chave_validacao {
		_lista_de_musicas = append(_lista_de_musicas, v)
	}

	return _lista_de_musicas, nil
}

func Liste_a_pasta(diretorio string) ([]ListaMusicas, error) {

	//Validacao do JSON para ver se ele nao esta corrompido;
	var _lista_de_musicas []ListaMusicas
	chave_validacao := make(map[string]ListaMusicas)
	caminhos_vistos := make(map[string]bool)		//Levando em consideracao musicas apagadas;

	if f, err := os.Open("NewArrow/myJSONfile.json"); err == nil {
    	json.NewDecoder(f).Decode(&chave_validacao)
     	f.Close()
	}

	//Se o JSON eh nil, o chave validacao vai interromper o fluxo e cagar o programa;
	
	if chave_validacao == nil {
    	chave_validacao = make(map[string]ListaMusicas)
	}
	
	//Caso o nosso diretorio nao tenha sido escaneado, vamos seguir o caminho de sempre;
	
	_caminho := make(chan string, tarefas)				//Canal com 10 buffers;
	_saida_pasta := make(chan ListaMusicas, tarefas)	//10 buffers da struct que eu coloquei;
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
				informacao, err := d.Info()
				entrada, se_existe := chave_validacao[path]
	
				if err != nil {
					return err
				}

				/*Fazendo a comparacao com o map aqui para nao ter que ficar sobreescrevendo o JSON toda
				 santa vez, utilizamos 2 maps para que nao prescise ficar escaneando.
				 */
	
				if se_existe && entrada.ModTime == informacao.ModTime().Unix() && entrada.Size == informacao.Size() {
					_saida_pasta <- entrada
					caminhos_vistos[path] = true

				} else {
					_caminho <- path
					caminhos_vistos[path] = true
				}
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

	for path := range chave_validacao {
		if !caminhos_vistos[path] {
			delete(chave_validacao, path)
		}
	}

	for _, musica := range _lista_de_musicas {
		chave_validacao[musica.Caminho_path] = musica
	}

	saida, err := os.Create("./myJSONfile.json")
	
	if err != nil {
    	return _lista_de_musicas, err
	}

	defer saida.Close()

	if err := json.NewEncoder(saida).Encode(chave_validacao); err != nil {
    	return _lista_de_musicas, err
	}

	return _lista_de_musicas, nil
}

func processarArquivo(wayPath <-chan string, saidas chan <- ListaMusicas, wg *sync.WaitGroup) {
	defer wg.Done()

	for path := range wayPath {
		meta, err := ColocarMetadados(path)

		if err != nil {
			continue			//Caso isso possa dar um erro;
		}

		saidas <- meta
	}
}

/*------------------------------------A funcao ColocarMetadados----------------------------------------------//

	Basicamente importamos uma biblioteca que se chama dhowgen/tag, que le o nosso arquivo para nos sobre qual
	eh o artista, o ano, etc. Colocamos essas informacoes para o player ficar mais completo e utilizar elas depois
	no frontEnd.

_____________________________________________________________________________________________________________
 */

func ColocarMetadados(WayPath string) (ListaMusicas, error) {

	f, err := os.Open(WayPath)
	
	var _saidasMetadados ListaMusicas

	
	if err != nil {
		return _saidasMetadados, err
	}
	
	informacoes, err := f.Stat()

	if err != nil {
		return _saidasMetadados, err
	}

	defer f.Close()

	m, err := tag.ReadFrom(f);

	if err != nil {
		return ListaMusicas{}, err
	}

	_saidasMetadados.Ano = m.Year()
	_saidasMetadados.Artista = m.Artist()
	_saidasMetadados.Nome_album = m.Album()
	_saidasMetadados.Nome_da_musica = m.Title()
	_saidasMetadados.Caminho_path = WayPath;
	_saidasMetadados.ModTime = informacoes.ModTime().Unix()
	_saidasMetadados.Size = informacoes.Size()

	return _saidasMetadados, nil
}