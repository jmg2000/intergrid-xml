package main

import (
	"encoding/xml"
	"reflect"
	"time"
)

type dateTime struct {
	time.Time
}

func (dt *dateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	parse, _ := time.Parse("2006-01-02T03:04:05", v)
	*dt = dateTime{parse}
	return nil
}

type XMLBank struct {
	XMLName     xml.Name `xml:"CatalogObject.Банки"`
	Ref         string
	IsFolder    bool
	Parent      string
	Code        string
	Description string
	CorrAcc     string `xml:"КоррСчет"`
	City        string `xml:"Город"`
	Address     string `xml:"Адрес"`
	Phones      string `xml:"Телефоны"`
}

type XMLBankAccounts struct {
	XMLName     xml.Name `xml:"CatalogObject.БанковскиеСчета"`
	Ref         string
	Owner       XMLBankAccountOwner `xml:"Owner"`
	Code        string
	Description string
	AccountNum  string `xml:"НомерСчета"`
	Bank        string `xml:"Банк"`
}

type XMLBankAccountOwner struct {
	OwnerObject string `xml:"type,attr"`
	Ref         string `xml:",chardata"`
}

type XMLOrganization struct {
	XMLName     xml.Name `xml:"CatalogObject.Организации"`
	Ref         string
	Code        string
	Description string
	INN         string `xml:"ИНН"`
	OKATO       string `xml:"КодПоОКАТО"`
	OKPO        string `xml:"КодПоОКПО"`
	KPP         string `xml:"КПП"`
	FullName    string `xml:"НаименованиеПолное"`
	OGRN        string `xml:"ОГРН"`
	BankAccount string `xml:"ОсновнойБанковскийСчет"`
	Prefix      string `xml:"Префикс"`
	Type        string `xml:"ЮрФизЛицо"`
}

type XMLWRhouse struct {
	XMLName     xml.Name `xml:"CatalogObject.Склады"`
	Ref         string
	IsFolder    bool
	Parent      string
	Code        string
	Description string
}

// контрагенты
type XMLCounterparties struct {
	XMLName     xml.Name `xml:"CatalogObject.Контрагенты"`
	Ref         string
	IsFolder    bool
	Parent      string
	Code        string
	Description string
	INN         string   `xml:"ИНН"`
	OKATO       string   `xml:"КодПоОКАТО"`
	OKPO        string   `xml:"КодПоОКПО"`
	KPP         string   `xml:"КПП"`
	FullName    string   `xml:"НаименованиеПолное"`
	OGRN        string   `xml:"ОГРН"`
	BankAccount string   `xml:"ОсновнойБанковскийСчет"`
	Buyer       bool     `xml:"Покупатель"`
	Traider     bool     `xml:"Поставщик"`
	Type        string   `xml:"ЮрФизЛицо"`
	LastName    string   `xml:"Фамилия"`
	FirstName   string   `xml:"Имя"`
	Patronymic  string   `xml:"Отчество"`
	OrgName     string   `xml:"НазваниеОрганизации"`
	DateLast    dateTime `xml:"ДатаПоследнегоОбращения"`
	Our         bool     `xml:"Свои"`
	PassNum     string   `xml:"ПаспортНомер"`
	PassSer     string   `xml:"ПаспортСерия"`
	PassIssued  string   `xml:"ПаспортВыдан"`
	PassDate    dateTime `xml:"ПаспортДатаВыдачи"`
	Birthday    dateTime `xml:"ДатаРождения"`
	Sex         string   `xml:"Пол"`
	Director    string   `xml:"ДиректорОрганизации"`
	Bookeeper   string   `xml:"ГлавныйБухгалтер"`
	Comment     string   `xml:"Комментарий"`
}

type XMLNomKind struct {
	XMLName     xml.Name `xml:"CatalogObject.ВидыНоменклатуры"`
	Ref         string
	Code        string
	Description string
	Type        string `xml:"ТипНоменклатуры"`
}

func (p *XMLNomKind) Descriptor() string {
	return p.Description
}

type XMLNomenclature struct {
	XMLName     xml.Name `xml:"CatalogObject.Номенклатура"`
	Ref         string
	IsFolder    bool
	Parent      string
	Code        string
	Description string
	Artikul     string `xml:"Артикул"`
	FullName    string `xml:"НаименованиеПолное"`
	GTDNum      string `xml:"НомерГТД"`
	NDSRate     string `xml:"СтавкаНДС"`
	Country     string `xml:"СтранаПроисхождения"`
	Service     bool   `xml:"Услуга"`
	BaseUnit    string `xml:"БазоваяЕдиницаИзмерения"` // guid
	Kind        string `xml:"ВидНоменклатуры"`
}

type XMLNomProps struct {
	XMLName     xml.Name `xml:"CatalogObject.ХарактеристикиНоменклатуры"`
	Ref         string
	Owner       string
	Description string
}

func (p *XMLNomProps) Descriptor() string {
	return p.Description
}

type XMLUnits struct {
	XMLName     xml.Name `xml:"CatalogObject.ЕдиницыИзмерения"`
	Ref         string
	Code        string
	Description string
}

func (p *XMLUnits) Descriptor() string {
	return p.Description
}

type XMLCarModel struct {
	XMLName               xml.Name `xml:"CatalogObject.бит_ус_Модели"`
	Ref                   string
	IsFolder              bool
	Parent                string
	Code                  string
	Description           string
	Nomenclature          string `xml:"Номенклатура"`
	ConditionsOfGuarantee string `xml:"УсловияГарантии"`
}

type XMLBodyColor struct {
	XMLName     xml.Name `xml:"CatalogObject.бит_ус_Цвета"`
	Ref         string
	Code        string
	Description string
}

func (p *XMLBodyColor) Descriptor() string {
	return p.Description
}

type XMLColorCode struct {
	XMLName     xml.Name `xml:"CatalogObject.бит_КодЦвета"`
	Ref         string
	Code        string
	Description string
}

func (p *XMLColorCode) Descriptor() string {
	return p.Description
}

type XMLPTSPlace struct {
	XMLName     xml.Name `xml:"CatalogObject.бит_МестоХраненияПТС"`
	Ref         string
	Code        string
	Description string
}

func (p *XMLPTSPlace) Descriptor() string {
	return p.Description
}

type XMLSaleType struct {
	XMLName     xml.Name `xml:"CatalogObject.бит_ТипПродажи"`
	Ref         string
	Code        string
	Description string
}

type XMLManufacturer struct {
	XMLName     xml.Name `xml:"CatalogObject.Бит_ОрганизацияИзготовительТССтрана"`
	Ref         string
	Code        string
	Description string
}

func (p *XMLManufacturer) Descriptor() string {
	return p.Description
}

type XMLTypeOfApproval struct {
	XMLName     xml.Name `xml:"CatalogObject.Бит_ОдобренияТипаТС"`
	Ref         string
	Code        string
	Description string
}

type XMLCountry struct {
	XMLName     xml.Name `xml:"CatalogObject.КлассификаторСтранМира"`
	Ref         string
	Code        string
	Description string
	FullName    string `xml:"НаименованиеПолное"`
	Alpha2Code  string `xml:"КодАльфа2"`
}

func (p *XMLCountry) Descriptor() string {
	return p.Description
}

type XMLGTDNum struct {
	XMLName     xml.Name `xml:"CatalogObject.Бит_СерияНомерГТД"`
	Ref         string
	Code        string
	Description string
}

func (p *XMLGTDNum) Descriptor() string {
	return p.Description
}

type XMLEngineType struct {
	XMLName     xml.Name `xml:"CatalogObject.бит_ус_МоделиДвигателей"`
	Ref         string
	Code        string
	Description string
}

func (p *XMLEngineType) Descriptor() string {
	return p.Description
}

type XMLGearType struct {
	XMLName     xml.Name `xml:"CatalogObject.бит_ус_ТипыКПП"`
	Ref         string
	Code        string
	Description string
}

func (p *XMLGearType) Descriptor() string {
	return p.Description
}

type XMLCar struct {
	XMLName         xml.Name `xml:"CatalogObject.бит_ус_Изделия"`
	Ref             string
	DeletionMark    bool
	Code            string
	Description     string
	VIN             string
	BaseEquipment   string   `xml:"БазоваяКомплектация"`
	CarOwner        string   `xml:"ВладелецАвтомобиля"`
	YearOfIssue     int      `xml:"ГодВыпуска"`
	RegNum          string   `xml:"ГосударственныйНомер"`
	PTSDate         dateTime `xml:"ДатаВыдачиПТС"`
	Comment         string   `xml:"Комментарий"`
	Equipment       string   `xml:"Комплектация"`
	Model           string   `xml:"Модель"`
	Engine          string   `xml:"МодельДвигателя"`
	ModelPTS        string   `xml:"МодельПоПТС"`
	EngineNum       string   `xml:"НомерДвигателя"`
	BodyNum         string   `xml:"НомерКузова"`
	PTSNum          string   `xml:"НомерПТС"`
	ChassisNum      string   `xml:"НомерШасси"`
	PTSSer          string   `xml:"СерияПТС"`
	EngineType      string   `xml:"ТипДвигателя"`
	GearType        string   `xml:"ТипКПП"`
	VehicleType     string   `xml:"ТипТС"`
	BodyColor       string   `xml:"ЦветКузова"`
	PTSColor        string   `xml:"ЦветПоПТС"`
	FullName        string   `xml:"НаименованиеПолное"`
	Traider         string   `xml:"битПродавец"`
	ColorCode       string   `xml:"битКодЦвета"`
	PTSPlace        string   `xml:"битМестоХраненияПТС"`
	SaleType        string   `xml:"битТипПродажи"`
	InPricePlan     float64  `xml:"битРасчетнаяПриходнаяЦена"`
	InPriceFact     float64  `xml:"битРеальнаяПриходнаяЦена"`
	ConstopDate     dateTime `xml:"битДатаОкончанияКонсигнации"`
	SupPayDate      dateTime `xml:"битДатаОплатыПоставщику"`
	SupSaleDate     dateTime `xml:"битДатаСписанияВСистемеПоставщика"`
	FinStatus       string   `xml:"битФинансовыйСтатусАвтомобиля"`
	InDocNum        string   `xml:"битПриходНомер"`
	VehicleCat      string   `xml:"битКатегорияТС"`
	EngPower        string   `xml:"битМощностьДвигателя"`
	EngPowerUnit    string   `xml:"битЕдиницаМощностиДвигателя"`
	EngVolume       float64  `xml:"битРабочийОбъемДвигателя"`
	EcoClass        string   `xml:"битЭкологическийКласс"`
	PermissibleMass float64  `xml:"битРазрешеннаяМасса"`
	WieghtWOutLoad  float64  `xml:"битМассаБезНагрузки"`
	Manufacturer    string   `xml:"битОрганизацияИзготовительТССтрана"`
	TypeOfApproval  string   `xml:"битОдобренияТипаТС"`
	Country         string   `xml:"битСтранаВывозаТС"`
	GTDNum          string   `xml:"битСерияНомерГТД"`
	SalonColor      string   `xml:"ЦветСалона"`
	FirstSaleDate   dateTime `xml:"ДатаПервичнойПродажи"`
	SaleDate        dateTime `xml:"ДатаПродажи"`
	Date            dateTime `xml:"ДатаСозданияКарточки"`
	LastVisit       dateTime `xml:"ДатаПоследнегоЗаезда"`
	LastTO          dateTime `xml:"ДатаПоследнегоТО"`
	TODate          dateTime `xml:"РекомендуемаяДатаСледующегоЗаезда"`
}

type XMLRealization struct {
	XMLName       xml.Name `xml:"DocumentObject.РеализацияТоваровУслуг"`
	Ref           string
	DeletionMark  bool
	Code          string
	Description   string
	Date          dateTime
	Number        string
	Organization  string     `xml:"Организация"`
	BankAccount   string     `xml:"БанковскийСчетОрганизации"`
	Wrhouse       string     `xml:"Склад"`
	Counterpartie string     `xml:"Контрагент"`
	DocSum        float64    `xml:"СуммаДокумента"`
	ProxyNum      string     `xml:"ДоверенностьНомер"`
	ProxyDate     dateTime   `xml:"ДоверенностьДата"`
	ProxyIssued   string     `xml:"ДоверенностьВыдана"`
	ProxyWho      string     `xml:"ДоверенностьЧерезКого"`
	Parts         XMLPart    `xml:"Товары"`
	Services      XMLService `xml:"Услуги"`
}

type XMLPart struct {
	XMLName xml.Name  `xml:"Товары"`
	Rows    []*XMLRow `xml:"Row"`
}

type XMLService struct {
	XMLName xml.Name  `xml:"Услуги"`
	Rows    []*XMLRow `xml:"Row"`
}

type XMLRow struct {
	XMLName      xml.Name `xml:"Row"`
	Content      string   `xml:"Содержание"`
	Unit         string   `xml:"ЕдиницаИзмерения"`
	Qty          float64  `xml:"Количество"`
	Nomenclature string   `xml:"Номенклатура"`
	Discount     float64  `xml:"ПроцентСкидкиНаценки"`
	NDSRate      string   `xml:"СтавкаНДС"`
	Summa        float64  `xml:"Сумма"`
	NDSSum       float64  `xml:"СуммаНДС"`
	Price        float64  `xml:"Цена"`
	Wrhouse      string   `xml:"Склад"`
	DiscSum      float64  `xml:"СуммаСкидки"`
}

type XMLOrderStatus struct {
	XMLName     xml.Name `xml:"CatalogObject.бит_ус_СтатусыЗаказНарядов"`
	Ref         string
	Code        string
	Description string
	Type        string `xml:"Тип"`
	StatusNum   int    `xml:"НомерСтатуса"`
}

type XMLKindOfService struct {
	XMLName     xml.Name `xml:"CatalogObject.бит_ус_ВидыОбслуживания"`
	Ref         string
	Code        string
	Description string
}

type XMLOffService struct {
	XMLName     xml.Name `xml:"DocumentObject.бит_ус_ЗаявкаНаРемонт"`
	Ref         string
	Code        string
	Description string
	Date        dateTime
	Number      string
}

type XMLInwpart struct {
	XMLName     xml.Name `xml:"DocumentObject.ПоступлениеТоваровУслуг"`
	Ref         string
	Code        string
	Description string
	Date        dateTime
	Number      string
}

type XMLDictWorks struct {
	XMLName     xml.Name `xml:"CatalogObject.бит_ус_Работы"`
	Ref         string
	IsFolder    bool
	Parent      string
	Code        string
	Description string
	FullName    string `xml:"НаименованиеПолное"`
}

func (p *XMLDictWorks) Descriptor() string {
	return p.Description
}

type XMLOrder struct {
	XMLName      xml.Name `xml:"DocumentObject.бит_ус_ЗаказНаряд"`
	Ref          string
	DeletionMark bool
	Code         string
	Description  string
	Date         dateTime
	Number       string
	Car          string   `xml:"Автомобиль"`
	DateClose    dateTime `xml:"ДатаЗакрытия"`
	DateOpen     dateTime `xml:"ДатаОткрытия"`
	Traider      string   `xml:"Организация"`
	Buyer        string   `xml:"Плательщик"`
	Probeg       int      `xml:"Пробег"`
	DocSum       float64  `xml:"СуммаДокумента"`
	VIN          string   `xml:"VIN"`
	FedNum       string   `xml:"ГосударственныйНомер"`
	Retail       bool     `xml:"Розница"`
	InData       dateTime `xml:"Доп_ВремяЗаездаВРемЗону"`
	OutData      dateTime `xml:"Доп_ВремяВыездаИзРемЗоны"`
	Owner        string   `xml:"Владелец"`
	Recomend     string   `xml:"РекомендацииКлиенту"`
	Works        XMLWorks `xml:"Работы"`
}

type XMLWorks struct {
	XMLName xml.Name      `xml:"Работы"`
	Rows    []*XMLWorkRow `xml:"Row"`
}

type XMLWorkRow struct {
	XMLName   xml.Name `xml:"Row"`
	Work      string   `xml:"Работа"` // XMLDictWorks
	Qty       float64  `xml:"Количество"`
	Price     float64  `xml:"Цена"`
	NDSRate   string   `xml:"СтавкаНДС"`
	Summa     float64  `xml:"Сумма"`
	NDSSum    float64  `xml:"СуммаНДС"`
	Discount  float64  `xml:"ПроцентСкидкиНаценки"`
	HPrice    float64  `xml:"ЦенаНормоЧаса"`
	TotalTime float64  `xml:"ОбщееВремя"`
}

type Descriptor interface {
	Descriptor() string
}

func getDescription(dict Descriptor) string {
	if reflect.ValueOf(dict).IsNil() == false {
		return dict.Descriptor()
	}
	return ""
}

func getObject(xmlObj []*Descriptor, objRef string) *Descriptor {
	for _, obj := range xmlObj {
		
	}
	return nil
}

func getBank(bankRef string) *XMLBank {
	for _, bank := range xmlData.Banks {
		if bank.Ref == bankRef {
			return bank
		}
	}
	return nil
}

func getBankAcc(bankAccRef string) *XMLBankAccounts {
	for _, bankAcc := range xmlData.BankAccounts {
		if bankAcc.Ref == bankAccRef {
			return bankAcc
		}
	}
	return nil
}

func getNomProp(nomPropRef string) *XMLNomProps {
	for _, nomProp := range xmlData.NomProps {
		if nomProp.Ref == nomPropRef {
			return nomProp
		}
	}
	return nil
}

func getCountry(countryRef string) *XMLCountry {
	for _, country := range xmlData.Countries {
		if country.Ref == countryRef {
			return country
		}
	}
	return nil
}

func getColor(colorRef string) *XMLBodyColor {
	for _, color := range xmlData.BodyColors {
		if color.Ref == colorRef {
			return color
		}
	}
	return nil
}

func getGTD(gtdRef string) *XMLGTDNum {
	for _, gtd := range xmlData.GTDNums {
		if gtd.Ref == gtdRef {
			return gtd
		}
	}
	return nil
}

func getNomKind(kindRef string) *XMLNomKind {
	for _, kind := range xmlData.NomKinds {
		if kind.Ref == kindRef {
			return kind
		}
	}
	return nil
}

func getUnit(unitRef string) *XMLUnits {
	for _, unit := range xmlData.Units {
		if unit.Ref == unitRef {
			return unit
		}
	}
	return nil
}

func getEngineType(engineRef string) *XMLEngineType {
	for _, engine := range xmlData.EngineTypes {
		if engine.Ref == engineRef {
			return engine
		}
	}
	return nil
}

func getGearType(gearRef string) *XMLGearType {
	for _, gear := range xmlData.GearTypes {
		if gear.Ref == gearRef {
			return gear
		}
	}
	return nil
}

func getColorCode(colorRef string) *XMLColorCode {
	for _, color := range xmlData.ColorCodes {
		if color.Ref == colorRef {
			return color
		}
	}
	return nil
}

func getPTSPlace(ptsRef string) *XMLPTSPlace {
	for _, pts := range xmlData.PTSPlaces {
		if pts.Ref == ptsRef {
			return pts
		}
	}
	return nil
}

func getManufacturer(manufacturerRef string) *XMLManufacturer {
	for _, manufacturer := range xmlData.Manufacturers {
		if manufacturer.Ref == manufacturerRef {
			return manufacturer
		}
	}
	return nil
}

func getOrganization(orgRef string) *XMLOrganization {
	for _, org := range xmlData.Organizations {
		if org.Ref == orgRef {
			return org
		}
	}
	return nil
}

func getWork(workRef string) *XMLDictWorks {
	for _, work := range xmlData.Works {
		if work.Ref == workRef {
			return work
		}
	}
	return nil
}
