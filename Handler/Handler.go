package Handler

import (
	"Nuevo_go/JCesarBat/Nuevo_go/model"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var Comida = regexp.MustCompile(`^/Comida/*$`)
var ComidaID = regexp.MustCompile(`^/Comida/([0-9]+)$`)

type Manejador_Comida struct {
}

func (c *Manejador_Comida) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch {
	case r.Method == http.MethodPost && Comida.MatchString(r.URL.Path):
		Crear(w, r)
		return

	case r.Method == http.MethodGet && Comida.MatchString(r.URL.Path):

		Lista(w, r)
		return
	case r.Method == http.MethodPut && ComidaID.MatchString(r.URL.Path):

		Actualizar(w, r)
		return
	case r.Method == http.MethodGet && ComidaID.MatchString(r.URL.Path):

		Obtener(w, r)
		return
	case r.Method == http.MethodDelete && ComidaID.MatchString(r.URL.Path):

		Borrar(w, r)
		return
	default:
		Default(w, r)
		return
	}

}
func Default(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "la direccion es incorrecta o no coincide con los metodos CRUD implementados")
	GenericaError(w, r)
}

func Obtener(w http.ResponseWriter, r *http.Request) {
	number_path := strings.Split(r.URL.Path, "/")
	ID, _ := strconv.ParseUint(number_path[2], 10, 32)

	comida, err := model.Reed(uint(ID))
	Info_comida := map[string]string{}
	if err != nil {
		io.WriteString(w, err.Error())
		GenericaError(w, r)
	} else {
		Info_comida["Nombre"] = comida.Nombre
		Info_comida["Ingrediente1"] = comida.Ingrediente1
		Info_comida["Ingrediente2"] = comida.Ingrediente2

		json_comida, _ := json.Marshal(Info_comida)

		w.Write(json_comida)
	}
}

func Crear(w http.ResponseWriter, r *http.Request) {
	var comida model.Comida

	if err := json.NewDecoder(r.Body).Decode(&comida); err != nil {

		GenericaError(w, r)
	}
	if comida.Nombre == "" || comida.Ingrediente1 == "" || comida.Ingrediente2 == "" {

		io.WriteString(w, "no puede haber datos nulos o con el identificador Nombre, Ingrediente1 ,ingrediente2 incorrecto ")
		GenericaError(w, r)
	} else {

		err := model.Save(comida)
		if err != nil {

			GenericaError(w, r)
		}

		io.WriteString(w, "se guardo en la base de datos ")
		w.WriteHeader(http.StatusOK)
	}
}

func Lista(w http.ResponseWriter, r *http.Request) {

	comidas, error := model.Listar()
	if error != nil {
		fmt.Println(error)
		GenericaError(w, r)
	}

	json_comidas, _ := json.Marshal(comidas)

	w.Write(json_comidas)

	w.WriteHeader(http.StatusOK)

}

func Actualizar(w http.ResponseWriter, r *http.Request) {

	var comida model.Comida

	if err := json.NewDecoder(r.Body).Decode(&comida); err != nil {

		io.WriteString(w, "los datos ingresados son errones   :"+err.Error())
		GenericaError(w, r)
	}
	number_path := strings.Split(r.URL.Path, "/")
	ID, _ := strconv.ParseUint(number_path[2], 10, 32)

	err := model.Update(uint(ID), comida)
	if err != nil {
		io.WriteString(w, err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}

func Borrar(w http.ResponseWriter, r *http.Request) {
	number_path := strings.Split(r.URL.Path, "/")
	ID, _ := strconv.ParseUint(number_path[2], 10, 32)

	err := model.Delete(uint(ID))
	if err != nil {
		io.WriteString(w, "No existe la comida con ese ID   ")
		GenericaError(w, r)

	} else {
		io.WriteString(w, "se elimino correctamente ")
		w.WriteHeader(http.StatusOK)
	}

}

func GenericaError(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusInternalServerError)

}
