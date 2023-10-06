package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"

	"github.com/zHenriqueGN/BuscaCEPTelegramBot/internal/entity"
)

var (
	ErrInvalidCEP  = errors.New("CEP inválido")
	ErrCEPNotFound = errors.New("CEP não encontrado")
)

func SearchCEP(CEP string) (*entity.Address, error) {
	if !validadeCEP(CEP) {
		return nil, ErrInvalidCEP
	}
	reqUrl := fmt.Sprintf("http://www.viacep.com.br/ws/%s/json/", CEP)
	resp, err := http.Get(reqUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var address entity.Address
	err = json.Unmarshal(respBody, &address)
	if err != nil {
		return nil, err
	}
	if address.CEP == "" {
		return nil, ErrCEPNotFound
	}
	return &address, nil
}

func validadeCEP(CEP string) bool {
	re := regexp.MustCompile(`^\d{5}-?\d{3}$`)
	return re.MatchString(CEP)
}
