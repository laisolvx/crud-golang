package rest_err

import "net/http"

type RestErr struct { // é uma estrutura que encapsula informações sobre um erro, como mensagem, código de erro, uma descrição do erro (Err), e causas específicas do erro.
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct { // é uma estrutura que representa uma causa específica de erro, indicando o campo (Field) e uma mensagem de erro associada.
	Fiel    string `json:"field"`
	Message string `json:"message"`
}

func (r *RestErr) Error() string { // o método Error() implementa a interface error e retorna a mensagem de erro, tornando possível usar instâncias de RestErr como objetos de erro padrão em Go.
	return r.Message
}

// funções NewRestErr e funções construtoras específicas para criar instâncias de RestErr para diferentes tipos de erros comuns em serviços RESTful, como erros de requisição inválida

func NewRestErr(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadResquetError(message string) *RestErr { // erros de validação de requisição inválida
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadResquetValidateError(message string, causes []Causes) *RestErr { // // erros internos do servidor
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestErr { // erros internos do servidor
	return &RestErr{
		Message: message,
		Err:     "internal server error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *RestErr { // erros de acesso proibido
	return &RestErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *RestErr { // erros de acesso proibido
	return &RestErr{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}

// essas funções construtoras ajudam a criar instâncias de RestErr com valores padrão para facilitar a identificação e manipulação de erros em um serviço RESTful, fornecendo uma estrutura consistente para representar e comunicar informações sobre erros.
