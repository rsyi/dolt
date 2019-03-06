package diff

import (
	"github.com/fatih/color"
	"github.com/liquidata-inc/ld/dolt/go/libraries/doltcore/row"
	"github.com/liquidata-inc/ld/dolt/go/libraries/doltcore/table/pipeline"
)

var greenTextProp = map[string]interface{}{ColorRowProp: color.GreenString}
var redTextProp = map[string]interface{}{ColorRowProp: color.RedString}
var yellowTextProp = map[string]interface{}{ColorRowProp: color.YellowString}

func ColoringTransform(r row.Row, props pipeline.ReadableMap) ([]*pipeline.TransformedRowResult, string) {
	var updatedProps map[string]interface{}
	diffType, ok := props.Get(DiffTypeProp)

	if ok {
		ct, ok := diffType.(DiffChType)

		if ok {
			switch ct {
			case DiffAdded:
				updatedProps = greenTextProp
			case DiffRemoved:
				updatedProps = redTextProp
			case DiffModifiedOld:
				updatedProps = yellowTextProp
			case DiffModifiedNew:
				updatedProps = yellowTextProp
			}
		}
	}

	return []*pipeline.TransformedRowResult{{r, updatedProps}}, ""
}