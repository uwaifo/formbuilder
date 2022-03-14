package form

import "reflect"

type field struct {
	Label       string
	Name        string
	Type        string
	Placeholder string
}

func fields(strct interface{}) []field {

	rv := reflect.ValueOf(strct)
	t := rv.Type()

	var ret []field
	for i := 0; i < t.NumField(); i++ {
		tf := t.Field(i)
		f := field{
			Label:       tf.Name,
			Name:        tf.Name,
			Type:        "text",
			Placeholder: tf.Name,
		}

		ret = append(ret, f)

	}
	return ret
	//tf := t.Field(0)

	// return []field{{
	// 	Label:       tf.Name,
	// 	Name:        tf.Name,
	// 	Type:        "text",
	// 	Placeholder: tf.Name,
	// },
	// }

	// return field{
	// 	Label:       "Name",
	// 	Name:        "Name",
	// 	Type:        "text",
	// 	Placeholder: "Name",
	// }

}

// func HTML(strct interface{})  {

// }
