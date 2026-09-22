package cacheuser

import "path/filepath"

//Cria o arquivo JSON com tudo ja pre_pronto;
func WayJSON_file() string {
	pastaJSON := CriandoA_pastaJSON()
	const nomeFile string = "MyMusics.json"

	if pastaJSON == "" {
		return ""
	}

	CaminhoJSON := filepath.Join(pastaJSON, nomeFile)

	return CaminhoJSON
}