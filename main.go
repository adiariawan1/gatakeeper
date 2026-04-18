package main

import (
	"BACK-END-ONLINESHOP/src/config"
	"BACK-END-ONLINESHOP/src/controllers"
	"BACK-END-ONLINESHOP/src/model"
	"BACK-END-ONLINESHOP/src/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)


type Roles struct {
	role_id int
	role_name string
}
func main(){
	db, err := config.ConnectDB()

	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	http.HandleFunc("/login", controllers.LoginHeader)
	// create
    
	http.HandleFunc("/users/create", controllers.Middleware(func(w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodPost{
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var newUsers model.User

		err := json.NewDecoder(r.Body).Decode(&newUsers)

		if err != nil {
			http.Error(w,"format json not valide"+ err.Error(), http.StatusBadRequest)
			return
		}

		err = repository.InsertUser(db,newUsers)

		if err != nil {
			http.Error(w,"failed insert users ", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message":"selamat! users baru di tambahkan " }`))
	}))


	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request){
		dataUser, err := repository.GetAllUsers(db)

		if err != nil {
			fmt.Println("not found data", err.Error())
			http.Error(w, "failed get data from database", 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dataUser)
	})
	// end create

	

	// update
	http.HandleFunc("/users/update", controllers.Middleware(func(w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodPut{
			http.Error(w,"Method not permission", http.StatusMethodNotAllowed)
			return
		}

		iddata:= r.URL.Query().Get("id")

		if iddata == "" {
			http.Error(w,"id dont null", http.StatusBadRequest)
			return
		}

		idAngka,err := strconv.Atoi(iddata) 

		if err!= nil {
			http.Error(w, "id harus angka", http.StatusBadRequest)
			return
		}

		var dataUser model.User


		dataUser.User_id = idAngka

		err = json.NewDecoder(r.Body).Decode(&dataUser)

		if err != nil {
			fmt.Println("decode error", err.Error())
			http.Error(w,"format json not valide"+ err.Error(), http.StatusBadRequest)
			return
		}

		err = repository.UpdateUser(db, dataUser)

		if err != nil {
			fmt.Println("failed post data")
			http.Error(w,"error 500", http.StatusInternalServerError)
			return
		}

		w.Write([]byte(`{"message":"Data berhasil di update " }`))

	}))
	// end update


	// delete
	http.HandleFunc("/users/delete", controllers.Middleware(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w,"Method harus delete", http.StatusMethodNotAllowed)
			return
		}

		idData := r.URL.Query().Get("id")
		 
		if idData == "" {
			http.Error(w, "id tidak boleh kosong", http.StatusBadRequest)
			return
		}

		idAngka, err := strconv.Atoi(idData)

		if err != nil {
			http.Error(w,"id harus berupa angka", http.StatusBadRequest)
			return
		}


		err = repository.DeleteUser(db, idAngka)

		if err != nil {
			fmt.Println("gagal hapus data")
			http.Error(w, "failed 500", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(`{"message":"Data berhasil di hapus" }`))
	}))
	// end delete

	fmt.Println("server running on port 8080")
	http.ListenAndServe(":8080", nil)
}