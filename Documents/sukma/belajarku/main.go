package main

import (
	"belajarku/config"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type makananStruct struct {
	Kode    int
	Nama    string
	Menu    string
	Jumlah  int
	Catatan string
}
type minumanStruct struct {
	Kode    int
	Nama    string
	Menu    string
	Jumlah  int
	Catatan string
}
type cemilanStruct struct {
	Kode    int
	Nama    string
	Menu    string
	Jumlah  int
	Catatan string
}

func main() {
	config.ConnectDB()
	http.HandleFunc("/", home)
	http.HandleFunc("/Makanan", Makanan)
	http.HandleFunc("/orderanMakanan", orderanMakanan)
	http.HandleFunc("/getOrderanMakanan", getOrderanMakanan)
	http.HandleFunc("/editOrderanMakanan", editOrderanMakanan)
	http.HandleFunc("/updateOrderanMakanan", updateOrderanMakanan)
	http.HandleFunc("/deleteOrderanMakanan", deleteOrderanMakanan)
	http.HandleFunc("/Minuman", Minuman)
	http.HandleFunc("/orderanMinuman", orderanMinuman)
	http.HandleFunc("/getOrderanMinuman", getOrderanMinuman)
	http.HandleFunc("/Cemilan", Cemilan)
	http.HandleFunc("/orderanCemilan", orderanCemilan)
	http.HandleFunc("/getOrderanCemilan", getOrderanCemilan)
	fmt.Println("---running---")

	http.ListenAndServe(":1234", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("views/home.html")
	if err != nil {
		log.Fatal("gagal memuat halaman utama", err)
	}
	err = templ.Execute(w, nil)
	if err != nil {
		log.Fatal("gagal execute halaman utama", err)
	}
}

func Makanan(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("views/Makanan.html")
	if err != nil {
		log.Fatal("gagal memuat halaman order makanan", err)
	}
	err = templ.Execute(w, nil)
	if err != nil {
		log.Fatal("gagal execute halaman order makanan", err)
	}
	w.Write([]byte("Terima kasih sudah order!"))

}

func Minuman(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("views/Minuman.html")
	if err != nil {
		log.Fatal("gagal memuat halaman order minuman", err)
	}
	err = templ.Execute(w, nil)
	if err != nil {
		log.Fatal("gagal execute halaman order minuman", err)
	}
	w.Write([]byte("Terima kasih sudah order!"))

}

func Cemilan(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("views/Cemilan.html")
	if err != nil {
		log.Fatal("gagal memuat halaman order cemilan", err)
	}
	err = templ.Execute(w, nil)
	if err != nil {
		log.Fatal("gagal execute halaman order cemilan", err)
	}
	w.Write([]byte("Terima kasih sudah order!"))

}

func orderanMakanan(w http.ResponseWriter, r *http.Request) {
	var ma makananStruct
	Kode, err := strconv.Atoi(r.FormValue("inputKode"))
	if err != nil {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		return
	}
	ma.Kode = Kode
	ma.Nama = r.FormValue("inputNama")
	ma.Menu = r.FormValue("inputMenu")
	Jumlah, err := strconv.Atoi(r.FormValue("inputJumlah"))
	if err != nil {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		return
	}
	ma.Jumlah = Jumlah
	ma.Catatan = r.FormValue("inputCatatan")

	if ma.Kode <= 0 {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		w.Write([]byte("Kode harus diisi!"))
		return
	}

	if ma.Nama == "" {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		w.Write([]byte("Nama harus diisi!"))
		return
	}

	if ma.Menu == "" {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		w.Write([]byte("Menu harus diisi!"))
		return
	}

	if ma.Jumlah <= 0 {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		w.Write([]byte("Jumlah harus diisi!"))
		return
	}
	_, err = config.DB.Query("INSERT INTO Makanan (kode, nama, menu, jumlah, Catatan) VALUES(?,?,?,?,?)", ma.Kode, ma.Nama, ma.Menu, ma.Jumlah, ma.Catatan)
	if err != nil {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Terima kasih sudah order!"))

}

func orderanMinuman(w http.ResponseWriter, r *http.Request) {
	var mi minumanStruct
	Kode, err := strconv.Atoi(r.FormValue("inputKode"))
	if err != nil {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		return
	}
	mi.Kode = Kode
	mi.Nama = r.FormValue("inputNama")
	mi.Menu = r.FormValue("inputMenu")
	Jumlah, err := strconv.Atoi(r.FormValue("inputJumlah"))
	if err != nil {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		return
	}
	mi.Jumlah = Jumlah
	mi.Catatan = r.FormValue("inputCatatan")

	if mi.Kode <= 0 {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		w.Write([]byte("Kode harus diisi!"))
		return
	}

	if mi.Nama == "" {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		w.Write([]byte("Nama harus diisi!"))
		return
	}

	if mi.Menu == "" {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		w.Write([]byte("Menu harus diisi!"))
		return
	}

	if mi.Jumlah <= 0 {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		w.Write([]byte("Jumlah harus diisi!"))
		return
	}
	_, err = config.DB.Query("INSERT INTO Minuman (kode, nama, menu, jumlah, catatan) VALUES(?,?,?,?,?)", mi.Kode, mi.Nama, mi.Menu, mi.Jumlah, mi.Catatan)
	if err != nil {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Terima kasih sudah order!"))

}

func orderanCemilan(w http.ResponseWriter, r *http.Request) {
	var ce cemilanStruct
	ce.Nama = r.FormValue("inputNama")
	ce.Menu = r.FormValue("inputMenu")
	ce.Catatan = r.FormValue("inputCatatan")

	Kode, err := strconv.Atoi(r.FormValue("inputKode"))
	if err != nil {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		return
	}
	ce.Kode = Kode

	Jumlah, err := strconv.Atoi(r.FormValue("inputJumlah"))
	if err != nil {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		return
	}
	ce.Jumlah = Jumlah

	if ce.Kode <= 0 {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		w.Write([]byte("Kode harus diisi!"))
		return
	}

	if ce.Nama == "" {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		w.Write([]byte("Nama harus diisi!"))
		return
	}

	if ce.Menu == "" {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		w.Write([]byte("Menu harus diisi!"))
		return
	}

	if ce.Jumlah <= 0 {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		w.Write([]byte("Jumlah harus diisi!"))
		return
	}
	_, err = config.DB.Query("INSERT INTO Cemilan (kode, nama, menu, jumlah, catatan) VALUES(?,?,?,?,?)", ce.Kode, ce.Nama, ce.Menu, ce.Jumlah, ce.Catatan)
	if err != nil {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Terima kasih sudah order!"))

}

func getOrderanMakanan(w http.ResponseWriter, r *http.Request) {
	var listOrderanMakanan []makananStruct
	datas, err := config.DB.Query("SELECT * FROM Makanan")
	if err != nil {
		log.Fatal("cannot display orderan", err)
	}
	defer datas.Close()
	for datas.Next() {
		var ma makananStruct
		datas.Scan(&ma.Kode, &ma.Nama, &ma.Menu, &ma.Jumlah, &ma.Catatan)
		if err != nil {
			log.Fatal("cannot parse orderan", err)
		}
		listOrderanMakanan = append(listOrderanMakanan, ma)
	}

	templ, err := template.ParseFiles("views/tableMakanan.html")
	if err != nil {
		log.Fatal("gagal memuat halaman tableMakanan", err)
	}
	err = templ.Execute(w, listOrderanMakanan)
	if err != nil {
		log.Fatal("gagal execute halaman tableMakanan", err)
	}
}

func getOrderanMinuman(w http.ResponseWriter, r *http.Request) {
	var listOrderanMinuman []minumanStruct
	datas, err := config.DB.Query("SELECT * FROM Minuman")
	if err != nil {
		log.Fatal("cannot display orderan", err)
	}
	defer datas.Close()
	for datas.Next() {
		var mi minumanStruct
		datas.Scan(&mi.Kode, &mi.Nama, &mi.Menu, &mi.Jumlah, &mi.Catatan)
		if err != nil {
			log.Fatal("cannot parse orderan", err)
		}
		listOrderanMinuman = append(listOrderanMinuman, mi)
	}

	templ, err := template.ParseFiles("views/tableMinuman.html")
	if err != nil {
		log.Fatal("gagal memuat halaman tableMinuman", err)
	}
	err = templ.Execute(w, listOrderanMinuman)
	if err != nil {
		log.Fatal("gagal execute halaman tableMinuman", err)
	}
}

func getOrderanCemilan(w http.ResponseWriter, r *http.Request) {
	var listOrderanCemilan []cemilanStruct
	datas, err := config.DB.Query("SELECT * FROM Cemilan")
	if err != nil {
		log.Fatal("cannot display orderan", err)
	}
	defer datas.Close()
	for datas.Next() {
		var ce cemilanStruct
		datas.Scan(&ce.Kode, &ce.Nama, &ce.Menu, &ce.Jumlah, &ce.Catatan)
		if err != nil {
			log.Fatal("cannot parse orderan", err)
		}
		listOrderanCemilan = append(listOrderanCemilan, ce)
	}

	templ, err := template.ParseFiles("views/tableCemilan.html")
	if err != nil {
		log.Fatal("gagal memuat halaman tableCemilan", err)
	}
	err = templ.Execute(w, listOrderanCemilan)
	if err != nil {
		log.Fatal("gagal execute halaman tableCemilan", err)
	}
}

func editOrderanMakanan(w http.ResponseWriter, r *http.Request) {
	Kode := r.URL.Query().Get("Kode")
	data, err := config.DB.Query("SELECT * FROM Makanan WHERE Kode=?", Kode)
	if err != nil {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		return
	}
	defer data.Close()
	var ma makananStruct
	for data.Next() {
		err = data.Scan(&ma.Kode, &ma.Nama, &ma.Menu, &ma.Jumlah, &ma.Catatan)
		if err != nil {
			w.Write([]byte("error \n"))
			w.Write([]byte(err.Error()))
			return
		}
	}
	templ, err := template.ParseFiles("views/updateOrderanMakanan.html")
	if err != nil {
		log.Fatal("gagal memuat halaman updateOrderanMakanan", err)
	}
	err = templ.Execute(w, ma)
	if err != nil {
		log.Fatal("gagal execute halaman updateOrderanMakanan", err)
	}
}

func updateOrderanMakanan(w http.ResponseWriter, r *http.Request) {
	var ma makananStruct
	Kode, err := strconv.Atoi(r.FormValue("inputKode"))
	if err != nil {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		return
	}
	ma.Kode = Kode
	ma.Nama = r.FormValue("inputNama")
	ma.Menu = r.FormValue("inputMenu")
	Jumlah, err := strconv.Atoi(r.FormValue("inputJumlah"))
	if err != nil {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		return
	}
	ma.Jumlah = Jumlah
	ma.Catatan = r.FormValue("inputCatatan")

	if ma.Kode <= 0 {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		w.Write([]byte("Kode harus diisi!"))
		return
	}

	if ma.Nama == "" {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		w.Write([]byte("Nama harus diisi!"))
		return
	}

	if ma.Menu == "" {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		w.Write([]byte("Menu harus diisi!"))
		return
	}

	if ma.Jumlah <= 0 {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		w.Write([]byte("Jumlah harus diisi!"))
		return
	}

	_, err = config.DB.Query("UPDATE Makanan SET Nama=?, Menu=?, Jumlah=?, Catatan=? WHERE Kode=?", ma.Nama, ma.Menu, ma.Jumlah, ma.Catatan, ma.Kode)
	if err != nil {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Orderan anda telah diperbarui!"))

}

func deleteOrderanMakanan(w http.ResponseWriter, r *http.Request) {
	Kode := r.URL.Query().Get("Kode")
	_, err := config.DB.Query("DELETE FROM Makanan WHERE Kode=?", Kode)
	if err != nil {
		w.Write([]byte("error \n"))
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Orderan anda telah dihapus!"))
}
