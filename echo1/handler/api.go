package handler

import (
	"echo/server"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

type menu struct {
	Id_menu     string
	Nama_menu   string
	Deskripsi   string
	Jenis       string
	Harga       string
	Url_gambar  string
	Total_order string
}

var data []menu

func BacaData(c echo.Context) error {
	menu_makanan()

	return c.JSON(http.StatusOK, data)
}

func BacaPopuler(c echo.Context) error {
	menu_populer()

	return c.JSON(http.StatusOK, data)
}

func menu_makanan() {
	data = nil
	db, err := server.Koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tbl_menu")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer rows.Close()

	for rows.Next() {
		var each = menu{}
		var err = rows.Scan(&each.Id_menu, &each.Nama_menu, &each.Deskripsi, &each.Url_gambar, &each.Jenis, &each.Harga)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data = append(data, each)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}
}

func AddData(c echo.Context) error {
	db, err := server.Koneksi()

	defer db.Close()

	var nama = c.FormValue("Nama_menu")
	var deskripsi = c.FormValue("Deskripsi")
	var harga = c.FormValue("Harga")
	var jenis = c.FormValue("Jenis")
	var url_gambar = c.FormValue("Url_gambar")

	_, err = db.Exec("INSERT INTO tbl_menu VALUES (?,?,?,?,?,?)", nil, nama, deskripsi, url_gambar, jenis, harga)
	if err != nil {
		fmt.Println("Gagal Input Data")
		return c.JSON(http.StatusOK, "Gagal Menambahkan Menu")
	} else {
		fmt.Println("Berhasil di Tambahkan")
		return c.JSON(http.StatusOK, "Berhasil Menambahkan Menu")
	}
}

func UpdateData(c echo.Context) error {
	db, err := server.Koneksi()

	defer db.Close()

	var id = c.FormValue("Id_menu")
	var nama = c.FormValue("Nama_menu")
	var deskripsi = c.FormValue("Deskripsi")
	var harga = c.FormValue("Harga")
	var jenis = c.FormValue("Jenis")
	var url_gambar = c.FormValue("Url_gambar")

	_, err = db.Exec("UPDATE tbl_menu SET nama_menu = ?, deskripsi = ?, harga = ?, jenis = ?, url_gambar = ? WHERE id_menu = ?", nama, deskripsi, harga, jenis, url_gambar, id)
	if err != nil {
		fmt.Println("Gagal Update Data")
		return c.JSON(http.StatusOK, "Gagal Update Menu")
	} else {
		fmt.Println("Berhasil di Update")
		return c.JSON(http.StatusOK, "Berhasil Update Menu")
	}
}

func DeleteData(c echo.Context) error {
	db, err := server.Koneksi()

	defer db.Close()

	var id = c.FormValue("Id_menu")

	_, err = db.Exec("DELETE FROM tbl_menu WHERE id_menu = ?", id)
	if err != nil {
		fmt.Println("Gagal Hapus Data")
		return c.JSON(http.StatusOK, "Gagal Hapus Menu")
	} else {
		fmt.Println("Berhasil di Hapus")
		return c.JSON(http.StatusOK, "Berhasil Hapus Menu")
	}
}

func AddOrder(c echo.Context) error {
	db, err := server.Koneksi()

	defer db.Close()

	var id = c.FormValue("id")
	var nama_pemesan = c.FormValue("nama_pemesan")
	var nomor_telepon = c.FormValue("nomor_telepon")
	var jumlah = c.FormValue("jumlah")
	var alamat = c.FormValue("alamat")

	_, err = db.Exec("INSERT INTO tbl_order VALUES (?, ?, ?, ?, ?, ?)", nil, id, nama_pemesan, nomor_telepon, alamat, jumlah)
	if err != nil {
		fmt.Println("Pesanan Gagal Dibuat")
		return c.HTML(http.StatusOK, "<strong>Gagal Melakukan Pemesanan</strong>")
	} else {
		fmt.Println("Pesanan Berhasil Dibuat")
		return c.HTML(http.StatusOK, "<script>alert('Berhasil Melakukan Pemesanan'); window.location = 'http://localhost:1323';</script>")
	}
	//return c.Redirect(http.StatusSeeOther, "/")
}

func menu_populer() {
	data = nil
	db, err := server.Koneksi()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM vw_totalorder ORDER BY total_order DESC LIMIT 8")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer rows.Close()

	for rows.Next() {
		var each = menu{}
		var err = rows.Scan(&each.Id_menu, &each.Nama_menu, &each.Deskripsi, &each.Url_gambar, &each.Jenis, &each.Harga, &each.Total_order)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data = append(data, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

}
