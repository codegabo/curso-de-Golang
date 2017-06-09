package main

import (
	"fmt"
	//driver para mysql
	// se usa el _ para que solo se ejecute la funcion init
	_ "github.com/go-sql-driver/mysql"
	//paquete para hacer conexciones a base de datos
	"github.com/jinzhu/gorm"
)

// Producto contiene los datos de un artículo
type Producto struct {
	gorm.Model   //ID, CreatedAt, UpdatedAt, DeletedAt
	CodigoBarras string
	Precio       uint
}

func main() {
	// tipo de gestor de base de datos, usuario:contraseña
	db, err := gorm.Open("mysql", "ehnqhlcm_juangabrielmm:31GaD3y#il06er%^7s8@tcp(mysql.anaxanet.com:3306)/ehnqhlcm_pruebasappqui")
	if err != nil {
		panic("Error en la conexión a la base de datos")
	}
	defer db.Close()
	fmt.Println("Se conectó a la base de datos")

	//crear tablas
	/*
		db.CreateTable(&Producto{})
		fmt.Println("Se ha creado la tabla productos")
	*/
	// crear registros
	/*
		db.Create(&Producto{CodigoBarras: "001", Precio: 45000})
		db.Create(&Producto{CodigoBarras: "002", Precio: 50000})
		db.Create(&Producto{CodigoBarras: "003", Precio: 12000})
		fmt.Println("Se ha registrado el prodcuto")
	*/
	// consulta de datos
	//consulta del primer registro solamente
	var p Producto
	db.First(&p)
	fmt.Println("Este es el primer registro: ", p)

	//consula del primer registro que coincida con el id 2
	/*var p Producto
	db.First(&p)
	fmt.Println("Estos son los productos actuales: ", p)
	//actualizar registro
	p.Precio = 4500
	db.Save(&p)
	fmt.Println("El nuevo precio del producto es: ", p.Precio)
	*/
}
