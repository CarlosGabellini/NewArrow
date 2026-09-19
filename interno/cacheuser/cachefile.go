package cacheuser

import "path/filepath"

//Aqui estamos criando os caminhos ja pre-prontos para as pastas que vamos prescisar colocar
// no cache provavelmente, essa funcao eh importante para criar a pasta que por sua vez vai
// abrigar o arquivo .json das musicas que criaremos.
func WayJSON_file() string {
	pastaJSON := CriandoA_pastaJSON()
	const nomeFile string = "minhasMusicas.json"

	if pastaJSON == "" {
		return ""
	}

	CaminhoJSON := filepath.Join(pastaJSON, nomeFile)

	return CaminhoJSON
}