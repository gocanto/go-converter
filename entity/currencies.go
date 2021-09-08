package converter

type Currencies struct {
	Items *map[string]Currency
}

func Make() Currencies  {
	items := make(map[string]Currency)

	return Currencies{Items: &items}
}

func (current *Currencies) Add(currency Currency)  {
	(*current.Items)[currency.Code] = currency
}

func (current *Currencies) Remove(seed interface{})  {
	//items := *current.Items
	//var code string
	//
	//switch _ := seed.(type) {
	//case string:
	//	code := reflect.TypeOf(seed).String()
	//	//delete(items, reflect.TypeOf(seed).String() )
	//	//delete(items, fmt.Sprintf("%v", seed))
	//	//delete(m, "route")
	//	case Currency:
	//		code := seed.Code
	//	default:
	//
	//}
}
