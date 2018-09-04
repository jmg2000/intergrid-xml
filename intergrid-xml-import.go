package main

import (
	_ "database/sql"
	"encoding/xml"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/nakagami/firebirdsql"
	"log"
	"os"
	"time"
)

type XMLData struct {
	XMLName         xml.Name             `xml:"Data"`
	Banks           []*XMLBank           `xml:"CatalogObject.Банки"`
	BankAccounts    []*XMLBankAccounts   `xml:"CatalogObject.БанковскиеСчета"`
	Organizations   []*XMLOrganization   `xml:"CatalogObject.Организации"`
	WRhouses        []*XMLWRhouse        `xml:"CatalogObject.Склады"`
	Counterparties  []*XMLCounterparties `xml:"CatalogObject.Контрагенты"`
	Nomenclatures   []*XMLNomenclature   `xml:"CatalogObject.Номенклатура"`
	NomProps        []*XMLNomProps       `xml:"CatalogObject.ХарактеристикиНоменклатуры"`
	Units           []*XMLUnits          `xml:"CatalogObject.ЕдиницыИзмерения"`
	CarModels       []*XMLCarModel       `xml:"CatalogObject.бит_ус_Модели"`
	BodyColors      []*XMLBodyColor      `xml:"CatalogObject.бит_ус_Цвета"`
	ColorCodes      []*XMLColorCode      `xml:"CatalogObject.бит_КодЦвета"`
	PTSPlaces       []*XMLPTSPlace       `xml:"CatalogObject.бит_МестоХраненияПТС"`
	SaleTypes       []*XMLSaleType       `xml:"CatalogObject.бит_ТипПродажи"`
	Manufacturers   []*XMLManufacturer   `xml:"CatalogObject.Бит_ОрганизацияИзготовительТССтрана"`
	TypesOfApproval []*XMLTypeOfApproval `xml:"CatalogObject.Бит_ОдобренияТипаТС"`
	Countries       []*XMLCountry        `xml:"CatalogObject.КлассификаторСтранМира"`
	GTDNums         []*XMLGTDNum         `xml:"CatalogObject.Бит_СерияНомерГТД"`
	GearTypes       []*XMLGearType       `xml:"CatalogObject.бит_ус_ТипыКПП"`
	Cars            []*XMLCar            `xml:"CatalogObject.бит_ус_Изделия"`
	Realizations    []*XMLRealization    `xml:"DocumentObject.РеализацияТоваровУслуг"`
	OrderStatus     []*XMLOrderStatus    `xml:"CatalogObject.бит_ус_СтатусыЗаказНарядов"`
	KindOfService   []*XMLKindOfService  `xml:"CatalogObject.бит_ус_ВидыОбслуживания"`
	OffServices     []*XMLOffService     `xml:"DocumentObject.бит_ус_ЗаявкаНаРемонт"`
	Inwparts        []*XMLInwpart        `xml:"DocumentObject.ПоступлениеТоваровУслуг"`
	Works           []*XMLDictWorks      `xml:"CatalogObject.бит_ус_Работы"`
	Orders          []*XMLOrder          `xml:"DocumentObject.бит_ус_ЗаказНаряд"`
	NomKinds        []*XMLNomKind        `xml:"CatalogObject.ВидыНоменклатуры"`
	EngineTypes     []*XMLEngineType     `xml:"CatalogObject.бит_ус_МоделиДвигателей"`
}

type XML1CDt struct {
	XMLName xml.Name `xml:"_1CV8DtUD"`
	Data    *XMLData `xml:"Data"`
}

var xmlData *XMLData

func main() {

	var err error
	var inFileName string

	if len(os.Args) > 1 {
		inFileName = os.Args[1]
	}

	inFile := os.Stdout
	if inFile, err = os.Open(inFileName); err != nil {
		log.Fatal(err)
	}
	defer inFile.Close()

	db, err := sqlx.Connect("firebirdsql", "sysdba:masterkey@192.168.5.206:3050/database/parts/inter_akos.gdb")
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println("Database connection established..")
	defer db.Close()

	start := time.Now()

	xml1CDt := &XML1CDt{}
	decoder := xml.NewDecoder(inFile)
	if err := decoder.Decode(xml1CDt); err != nil {
		log.Println(err)
	}
	xmlData = xml1CDt.Data

	fmt.Println("Processed:")
	fmt.Printf("Banks: %d\n", len(xmlData.Banks))
	fmt.Printf("Bank Accounts: %d\n", len(xmlData.BankAccounts))
	fmt.Printf("Organizations: %d\n", len(xmlData.Organizations))
	fmt.Printf("WRhouses: %d\n", len(xmlData.WRhouses))
	fmt.Printf("Counterparties: %d\n", len(xmlData.Counterparties))
	fmt.Printf("Nomenclatures: %d\n", len(xmlData.Nomenclatures))
	fmt.Printf("NomProps: %d\n", len(xmlData.NomProps))
	fmt.Printf("Units: %d\n", len(xmlData.Units))
	fmt.Printf("CarModels: %d\n", len(xmlData.CarModels))
	fmt.Printf("BodyColors: %d\n", len(xmlData.BodyColors))
	fmt.Printf("ColorCodes: %d\n", len(xmlData.ColorCodes))
	fmt.Printf("PTSPlaces: %d\n", len(xmlData.PTSPlaces))
	fmt.Printf("SaleTypes: %d\n", len(xmlData.SaleTypes))
	fmt.Printf("Manufacturers: %d\n", len(xmlData.Manufacturers))
	fmt.Printf("TypesOfApproval: %d\n", len(xmlData.TypesOfApproval))
	fmt.Printf("Countries: %d\n", len(xmlData.Countries))
	fmt.Printf("GTDNums: %d\n", len(xmlData.GTDNums))
	fmt.Printf("GearTypes: %d\n", len(xmlData.GearTypes))
	fmt.Printf("Cars: %d\n", len(xmlData.Cars))
	fmt.Printf("Realizations: %d\n", len(xmlData.Realizations))
	fmt.Printf("OrderStatus: %d\n", len(xmlData.OrderStatus))
	fmt.Printf("KindOfService: %d\n", len(xmlData.KindOfService))
	fmt.Printf("OffServices: %d\n", len(xmlData.OffServices))
	fmt.Printf("Inwparts: %d\n", len(xmlData.Inwparts))
	fmt.Printf("Works: %d\n", len(xmlData.Works))
	fmt.Printf("Orders: %d\n", len(xmlData.Orders))

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	/*
		for _, nom := range xmlData.Nomenclatures {
			var cnt int
			for _, nom1 := range xmlData.Nomenclatures {
				if nom1.Ref == nom.Ref {
					cnt++
					fmt.Printf("%s - %d\n", nom.Ref, cnt)
				}

			}
		}
		os.Exit(1)
	*/
	var inserted, updated int

	start = time.Now()
	for _, org := range xmlData.Counterparties {
		var id int
		var rows *sqlx.Rows

		if org.IsFolder == true {
			continue
		}

		if org.Type == "ФизЛицо" {
			rows, err = db.Queryx("select first 1 id from orgbase where priv_surname = ? and priv_name = ? and priv_lastname = ?", org.LastName, org.FirstName, org.Patronymic)
			if err != nil {
				log.Printf("query error: %v", err)
			}
			for rows.Next() {
				if err := rows.Scan(&id); err != nil {
					log.Fatal(err)
				}
			}
			rows.Close()
			if id > 0 {
				db.MustExec("update orgbase set guid = ? where id = ?", org.Ref, id)
				updated++
			} else {
				_, err = db.NamedExec("insert into orgbase (id, parent, isfolder, orgmne, orgtyp, orgname, priv_surname, priv_name, priv_lastname, "+
					"priv_passser, priv_passnum, priv_passout, priv_passdate, birthday, lastrefdate, deleted, guid) "+
					"values (gen_id(gen_orgbase_id, 1), :parent, 0, :orgmne, 2, :orgname, :priv_surname, :priv_name, :priv_lastname, "+
					":priv_passser, :priv_passnum, :priv_passout, :priv_passdate, :birthday, :lastrefdate, 'A', :guid)",
					map[string]interface{}{
						"parent":        138771,
						"orgmne":        fmt.Sprintf("%.25s", org.Code),
						"orgname":       fmt.Sprintf("%.100s", org.FullName),
						"priv_surname":  fmt.Sprintf("%.30s", org.LastName),
						"priv_name":     fmt.Sprintf("%.30s", org.FirstName),
						"priv_lastname": fmt.Sprintf("%.30s", org.Patronymic),
						"priv_passser":  org.PassSer,
						"priv_passnum":  org.PassNum,
						"priv_passout":  fmt.Sprintf("%.100s", org.PassIssued),
						"priv_passdate": org.PassDate.Time,
						"birthday":      org.Birthday.Time,
						"lastrefdate":   org.DateLast.Time,
						"guid":          org.Ref,
					})
				if err != nil {
					log.Printf("insert person %v", err)
				}
				inserted++
			}

		} else {
			rows, err = db.Queryx("select first 1 id from orgbase where org_inn = ?", org.INN)
			if err != nil {
				log.Printf("query error: %v", err)
			}
			for rows.Next() {
				if err := rows.Scan(&id); err != nil {
					log.Fatal(err)
				}
			}
			rows.Close()
			if id > 0 {
				db.MustExec("update orgbase set guid = ? where id = ?", org.Ref, id)
				updated++
			} else {
				_, err = db.NamedExec("insert into orgbase (id, parent, isfolder, orgmne, orgtyp, orgname, org_inn, org_kpp, org_director, "+
					"org_bookeeper, lastrefdate, deleted, guid) "+
					"values (gen_id(gen_orgbase_id, 1), :parent, 0, :orgmne, 1, :orgname, :org_inn, :org_kpp, :org_director, "+
					":org_bookeeper, :lastrefdate, 'A', :guid)",
					map[string]interface{}{
						"parent":        138772,
						"orgmne":        fmt.Sprintf("%.25s", org.Code),
						"orgname":       fmt.Sprintf("%.100s", org.FullName),
						"org_inn":       fmt.Sprintf("%.12s", org.INN),
						"org_kpp":       fmt.Sprintf("%.9s", org.KPP),
						"org_director":  fmt.Sprintf("%.30s", org.Director),
						"org_bookeeper": fmt.Sprintf("%.30s", org.Bookeeper),
						"lastrefdate":   org.DateLast.Time,
						"guid":          org.Ref,
					})
				if err != nil {
					log.Printf("insert org %v", err)
				}
				inserted++
			}
		}
	}

	fmt.Printf("Orgbase: inserted %d records, updated %d records, ", inserted, updated)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	inserted, updated = 0, 0

	// Bank Account
	start = time.Now()
	for _, bankAcc := range xmlData.BankAccounts {
		var orgid, bankAccId int

		//fmt.Printf("OwnerObject = %s, Owner = %s\n", bankAcc.Owner.OwnerObject, bankAcc.Owner.Ref)

		if bankAcc.Owner.OwnerObject != "CatalogRef.Контрагенты" {
			continue
		}

		db.QueryRowx("select id from orgbase where guid = ?", bankAcc.Owner.Ref).Scan(&orgid)

		if orgid != 0 {
			db.QueryRowx("select id from orgbase_banks where guid = ?", bankAcc.Ref).Scan(&bankAccId)
			if bankAccId == 0 {
				bank := getBank(bankAcc.Bank)
				db.QueryRowx("select id from orgbase_banks where account = ? and bik = ?", bankAcc.AccountNum, bank.Code).Scan(&bankAccId)
				if bankAccId == 0 {
					_, err := db.NamedExec("insert into orgbase_banks (id, name, orgid, account, bik, subaccount, adress, accountname, deleted, guid) "+
						"values (gen_id(gen_orgbase_banks_id, 1), :name, :orgid, :account, :bik, :subaccount, :adress, :accountname, 'A', :guid)",
						map[string]interface{}{
							"name":        fmt.Sprintf("%.100s", bank.Description),
							"orgid":       orgid, // bankAcc.Owner
							"account":     bankAcc.AccountNum,
							"bik":         bank.Code,
							"subaccount":  bank.CorrAcc,
							"adress":      fmt.Sprintf("%.100s", bank.Address),
							"accountname": fmt.Sprintf("%.100s", bankAcc.Description),
							"guid":        bankAcc.Ref,
						})
					if err != nil {
						log.Printf("insert into orgbase_banks error: %v", err)
					}
					inserted++
				} else {
					db.MustExec("update orgbase_banks set guid = ? where id = ?", bankAcc.Ref, bankAccId)
					updated++
				}
			}
		}

	}

	fmt.Printf("Orgbase_banks: inserted %d records, updated %d records, ", inserted, updated)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	inserted, updated = 0, 0

	// Контрагенты
	start = time.Now()
	for _, org := range xmlData.Counterparties {
		bankAcc := getBankAcc(org.BankAccount)

		if bankAcc != nil {
			bank := getBank(bankAcc.Bank)
			_, err := db.NamedExec("update orgbase set org_account = :accnum, org_corr = :accser, org_bik = :bik, org_bank = :bank where guid = :guid",
				map[string]interface{}{
					"accnum": bankAcc.AccountNum,
					"accser": bank.CorrAcc,
					"bik":    bank.Code,
					"bank":   fmt.Sprintf("%.100s", bank.Description),
					"guid":   org.Ref,
				})
			if err != nil {
				log.Printf("update orgbase error: %v", err)
			}
			updated++
		}
	}

	fmt.Printf("Orgbase main bank: updated %d records, ", updated)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	// Nomenclature
	inserted, updated = 0, 0

	start = time.Now()
	for _, nom := range xmlData.Nomenclatures {
		var plistid, workid, wrhid int

		if nom.Service == true {
			db.QueryRowx("select id from works where rname = ?", nom.Description).Scan(&workid)
			db.QueryRowx("select id from pricelist where name = ? and category = 3", nom.Description).Scan(&plistid)
			// Works
			if workid != 0 {
				db.MustExec("update works set guid = ? where id = ?", nom.Ref, workid)
				updated++
			} else {
				_, err := db.NamedExec("insert into works (id, parent, isfolder, rname, deleted, guid) values (gen_id(gen_works_id, 1), 7582, 0, :rname, 'D', :guid)",
					map[string]interface{}{
						"rname": fmt.Sprintf("%.60s", nom.Description),
						"guid":  nom.Ref,
					})
				if err != nil {
					log.Printf("insert into works error: %v", err)
				}
				inserted++
			}
			db.QueryRowx("select id from works where guid = ?", nom.Ref).Scan(&workid)
			// Pricelist
			if plistid != 0 {
				db.MustExec("update pricelist set guid = ? where id = ?", nom.Ref, plistid)
				updated++
			} else {
				_, err := db.NamedExec("insert into pricelist (id, parent, isfolder, name, category, descript, catalogn, deleted, sourceid, guid) "+
					"values (gen_id(gen_pricelist_id, 1), 115553, 0, :name, 3, :desc, :catalogn, 'A', :sourceid, :guid)",
					map[string]interface{}{
						"name":     fmt.Sprintf("%.100s", nom.Description),
						"desc":     nom.FullName,
						"catalogn": fmt.Sprintf("%.100s", nom.Artikul),
						"sourceid": workid,
						"guid":     nom.Ref,
					})
				if err != nil {
					log.Printf("insert into pricelist error: %v", err)
				}
				inserted++
			}
		} else {
			// if Nomenclature is car or part
			if getDescription(getNomKind(nom.Kind)) != "Автомобиль" {
				// Part
				db.QueryRowx("select id from wrhpart where catalogn = ?", nom.Artikul).Scan(&wrhid)
				db.QueryRowx("select id from pricelist where catalogn = ? and category = 2", nom.Artikul).Scan(&plistid)
				// Wrhpart
				if wrhid != 0 {
					res := db.MustExec("update wrhpart set guid = ? where id = ?", nom.Ref, wrhid)
					_, err := res.RowsAffected()
					if err != nil {
						log.Printf("update wrhpart error: %v", err)
					}
					updated++
				} else {
					country := getCountry(nom.Country)
					gtd := getGTD(nom.GTDNum)
					unit := getUnit(nom.BaseUnit)

					_, err := db.NamedExec("insert into wrhpart (id, parent, isfolder, catalogn, rname, country, edism, gtdnum, deleted, guid) "+
						"values (gen_id(gen_wrhpart_id, 1), 1007604, 0, :catn, :rname, :country, :edism, :gtdnum, 'A', :guid)",
						map[string]interface{}{
							"catn":    fmt.Sprintf("%.30s", nom.Artikul),
							"rname":   fmt.Sprintf("%.100s", nom.Description),
							"country": fmt.Sprintf("%.25s", getDescription(country)),
							"edism":   getDescription(unit), // nom.BaseUnit
							"gtdnum":  fmt.Sprintf("%.27s", getDescription(gtd)),
							"guid":    nom.Ref,
						})
					if err != nil {
						log.Printf("insert into wrhpart error: %v", err)
					}
					inserted++
				}
				db.QueryRowx("select id from wrhpart where guid = ?", nom.Ref).Scan(&wrhid)
				// Pricelist
				if plistid != 0 {
					db.MustExec("update pricelist set guid = ? where id = ?", nom.Ref, plistid)
					updated++
				} else {
					_, err := db.NamedExec("insert into pricelist (id, parent, isfolder, name, category, descript, catalogn, deleted, sourceid, guid) "+
						"values (gen_id(gen_pricelist_id, 1), 115554, 0, :name, 2, :desc, :catalogn, 'A', :sourceid, :guid)",
						map[string]interface{}{
							"name":     nom.Description,
							"desc":     nom.FullName,
							"catalogn": nom.Artikul,
							"sourceid": wrhid,
							"guid":     nom.Ref,
						})
					if err != nil {
						log.Printf("insert into pricelist error: %v", err)
					}
					inserted++
				}
			}
		}
	}

	fmt.Printf("Pricelist: inserted %d records, updated %d records, ", inserted, updated)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	// Cars
	inserted, updated = 0, 0

	start = time.Now()
	for _, car := range xmlData.Cars {
		var carid, orgid, salerid, plistid, carsaleid int

		db.QueryRowx("select id from cars where vin = ?", car.VIN).Scan(&carid)
		if carid != 0 {
			db.MustExec("update cars set guid = ? where id = ?", car.Ref, carid)
			updated++
		} else {
			compl := getNomProp(car.Equipment)
			country := getCountry(car.Country)
			color := getColor(car.BodyColor)
			gtd := getGTD(car.GTDNum)

			db.QueryRowx("select id from orgbase where guid = ?", car.CarOwner).Scan(&orgid)
			db.QueryRowx("select id from orgbase where guid = ?", car.Traider).Scan(&salerid)
			_, err := db.NamedExec("insert into cars (id, parent, isfolder, vin, data, marka, model, complectation, country, color, eng, kus, fednum, ouner, saledate, saler, fyear, deleted, ptsnum, gtdnum, guid)"+
				" values (gen_id(gen_cars_id, 1), 47906, 0, :vin, :data, :marka, :model, :compl, :country, :color, :eng, :kus, :fednum, :ouner, :saledate, :saler, :fyear, 'A', :ptsnum, :gtdnum, :guid)",
				map[string]interface{}{
					"vin":      fmt.Sprintf("%.20s", car.VIN),
					"data":     time.Now(),
					"marka":    "NISSAN",
					"model":    fmt.Sprintf("%.20s", car.ModelPTS),
					"compl":    fmt.Sprintf("%.100s", getDescription(compl)),
					"country":  fmt.Sprintf("%.25s", getDescription(country)),
					"color":    fmt.Sprintf("%.25s", getDescription(color)),
					"eng":      fmt.Sprintf("%.25s", car.EngineNum),
					"kus":      fmt.Sprintf("%.25s", car.BodyNum),
					"fednum":   fmt.Sprintf("%.10s", car.RegNum),
					"ouner":    orgid,
					"saledate": car.FirstSaleDate.Time,
					"saler":    salerid,
					"fyear":    car.YearOfIssue,
					"ptsnum":   fmt.Sprintf("%.20s", car.PTSSer+" "+car.PTSNum),
					"gtdnum":   fmt.Sprintf("%.27s", getDescription(gtd)),
					"guid":     car.Ref,
				})
			if err != nil {
				log.Printf("insert into cars error: %v", err)
			}

			inserted++
		}

		// Pricelist
		equipment := getDescription(getNomProp(car.Equipment))
		country := getCountry(car.Country)

		db.QueryRowx("select id from pricelist where name = ?", equipment).Scan(&plistid)
		if plistid != 0 {
			db.MustExec("update pricelist set guid = ? where id = ?", car.Ref, plistid)
			updated++
		} else {
			_, err := db.NamedExec("insert into pricelist (id, parent, isfolder, name, category, country, catalogn, deleted, guid) "+
				"values (gen_id(gen_pricelist_id, 1), 0, 0, :name, 1, :country, :catalogn, 'A', :guid)",
				map[string]interface{}{
					"name":     fmt.Sprintf("%.100s", equipment),
					"catalogn": car.FullName,
					"country":  fmt.Sprintf("%.25s", getDescription(country)),
					"guid":     car.Ref,
				})
			if err != nil {
				log.Printf("insert into pricelist error: %v", err)
			}
			inserted++
		}

		// Carsale
		color := getColor(car.BodyColor)
		gtd := getGTD(car.GTDNum)
		engine := getEngineType(car.EngineType)
		gear := getGearType(car.GearType)
		colorCode := getColorCode(car.ColorCode)
		pts := getPTSPlace(car.PTSPlace)
		//manufacturer := getManufacturer(car.Manufacturer)

		db.QueryRowx("select id from cars where guid = ?", car.Ref).Scan(&carid)
		db.QueryRowx("select id from carsale where vin = ?", car.VIN).Scan(&carsaleid)
		if carsaleid != 0 {
			db.MustExec("update carsale set guid = ? where id = ?", car.Ref, carsaleid)
			updated++
		} else {
			var ownerid, priceid int
			var carsale = make(map[string]interface{})

			db.QueryRowx("select id from orgbase where guid = ?", car.CarOwner).Scan(&ownerid)

			if ownerid == 0 {
				carsale["byer"] = nil
				carsale["owner"] = nil
			} else {
				carsale["byer"] = ownerid
				carsale["owner"] = ownerid
			}

			db.QueryRowx("select id from price where guid = ?", fmt.Sprintf("%.100s", equipment)).Scan(&priceid)
			carsale["priceid"] = priceid
			carsale["carid"] = carid
			carsale["marka"] = "NISSAN"
			carsale["model"] = fmt.Sprintf("%.20s", car.ModelPTS)
			carsale["compl"] = fmt.Sprintf("%.100s", equipment)
			carsale["engtype"] = fmt.Sprintf("%.15s", getDescription(engine))
			carsale["transtype"] = fmt.Sprintf("%.10s", getDescription(gear))
			carsale["country"] = fmt.Sprintf("%.25s", getDescription(country))
			carsale["color"] = fmt.Sprintf("%.25s", getDescription(color))
			carsale["colorid"] = fmt.Sprintf("%.10s", getDescription(colorCode))
			carsale["vin"] = fmt.Sprintf("%.20s", car.VIN)
			carsale["eng"] = fmt.Sprintf("%.25s", car.EngineNum)
			carsale["kus"] = fmt.Sprintf("%.25s", car.BodyNum)
			carsale["fyear"] = car.YearOfIssue
			carsale["traider"] = 5414
			carsale["data"] = time.Now()
			carsale["inpriceplan"] = car.InPricePlan
			carsale["inpricefact"] = car.InPriceFact
			carsale["pts"] = fmt.Sprintf("%.15s", getDescription(pts))
			carsale["gtd"] = fmt.Sprintf("%.27s", getDescription(gtd))
			carsale["outdata"] = car.SaleDate.Time
			carsale["ptsnum"] = fmt.Sprintf("%.20s", car.PTSSer+" "+car.PTSNum)
			carsale["comment"] = fmt.Sprintf("%.150s", car.Comment)
			carsale["guid"] = car.Ref

			_, err := db.NamedExec("insert into carsale (id, parent, isfolder, priceid, carid, marka, model, complectation, engtype, transtype, "+
				"country, color, colorid, vin, eng, kus, fyear, traider, data, inpriceplan, inpricefact, ptsplace, gtdnum, "+
				"byer, owner, outdata, ptsnum, comment, deleted, guid) "+
				"values (gen_id(gen_carsale_id, 1), 12, 0, :priceid, :carid, :marka, :model, :compl, :engtype, :transtype, "+
				":country, :color, :colorid, :vin, :eng, :kus, :fyear, :traider, :data, :inpriceplan, :inpricefact, :pts, :gtd, "+
				":byer, :owner, :outdata, :ptsnum, :comment, 'A', :guid)", carsale)

			if err != nil {
				log.Printf("insert into carsale error: %v", err)
			}
			inserted++
		}
	}

	fmt.Printf("Cars: inserted %d records, updated %d records, ", inserted, updated)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	// Orders
	inserted, updated = 0, 0
	errorCnt := 0
	carNullCnt := 0

	start = time.Now()
	for _, order := range xmlData.Orders {
		var orderid int

		if order.DeletionMark == true {
			continue
		}

		db.QueryRowx("select id from orders where docnum = ?", order.Number).Scan(&orderid)

		if orderid != 0 {
			db.MustExec("update orders set guid = ? where id = ?", order.Ref, orderid)
			updated++
		} else {
			var id, carid, ownerid int
			var model string

			db.QueryRowx("select id, model from cars where guid = ?", order.Car).Scan(&carid, &model)
			db.QueryRowx("select id from orgbase where guid = ?", order.Owner).Scan(&ownerid)
			db.QueryRowx("select gen_id(gen_ordwork_id, 1) from rdb$database").Scan(&id)

			if carid != 0 {

				_, err := db.NamedExec("insert into orders (id, parent, isfolder, docnum, data, docsum, docnds, carnum, marka, model, "+
					"fednum, ouner, traider, probeg, indata, outdata, comment, state, recomend, docstatus, confid, execsum, deleted, guid) "+
					"values (:id, 8, 0, :docnum, :data, :docsum, :docnds, :car, 'NISSAN', :model, "+
					":fednum, :ouner, 5414, :probeg, :indata, :outdata, 'ПРОДАЖА НАЛ', 'C', :recomend, 'ЗАКРЫТ', 1, :execsum, 'S', :guid)",
					map[string]interface{}{
						"id":       id,
						"docnum":   fmt.Sprintf("%.14s", order.Number),
						"data":     order.Date.Time,
						"docsum":   order.DocSum,
						"docnds":   float64((order.DocSum / 118) * 18),
						"car":      carid,
						"model":    fmt.Sprintf("%.30s", model),
						"fednum":   order.FedNum,
						"ouner":    ownerid,
						"probeg":   order.Probeg,
						"indata":   order.InData.Time,
						"outdata":  order.OutData.Time,
						"recomend": fmt.Sprintf("%.500s", order.Recomend),
						"execsum":  order.DocSum,
						"guid": order.Ref,
					})
				if err != nil {
					fmt.Printf("insert into orders error: %v", err)
					errorCnt++
				}
				inserted++
				for _, ordwork := range order.Works.Rows {
					work := getWork(ordwork.Work)
					_, err := db.NamedExec("insert into ordwork (id, parent, workname, timelimit, hprice, wprice, dsumma, deleted, executed) "+
						"values (gen_id(gen_ordwork_id, 1), :parent, :wname, :timelimit, :hprice, :wprice, :dsumma, 'A', 1)",
						map[string]interface{}{
							"parent":    id,
							"wname":     fmt.Sprintf("%.100s", getDescription(work)),
							"timelimit": ordwork.TotalTime,
							"hprice":    ordwork.HPrice,
							"wprice":    float64(ordwork.TotalTime * ordwork.HPrice),
							"dsumma":    float64(ordwork.TotalTime * ordwork.HPrice),
						})
					if err != nil {
						log.Printf("insert into ordwork error: %v", err)
						errorCnt++
					}
					inserted++
				}
			} else {
				carNullCnt++
			}

		}
	}

	fmt.Printf("Orders: inserted %d records, updated %d records, errors: %d, carid null cnt: %d, ", inserted, updated, errorCnt, carNullCnt)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}
