package ratings

type Rating string

const (
	empty Rating = ""

	/* Positive ratings */
	buy              Rating = "Buy"               //recomendacion de compra
	speculativeBuy   Rating = "Speculative Buy"   //compra especulativa, alto riesgo pero con gran potencial de crecimiento
	strongBuy        Rating = "Strong-Buy"        //compra fuerte, excelente desempeño y suba significativa
	outperform       Rating = "Outperform"        //se espera que la accion supere el rendimiento de otras en el mercado o de su sector
	outperformer     Rating = "Outperformer"      //similar a "Outperform", se espera que la accion supere en rendimiento a otras en el mercado
	marketOutperform Rating = "Market Outperform" //se espera que la accion tenga un rendimiento superior al mercado en general
	positive         Rating = "Positive"          //el analista tiene una perspectiva positiva sobre la accion, aunque puede no ser tan fuerte como un "Buy" explícito
	overweight       Rating = "Overweight"        //el analista recomienda sobreponderar la accion en un portafolio, lo que significa que esperan un buen rendimiento
	sectorOutperform Rating = "Sector Outperform" //se espera que la accion tenga un rendimiento superior al promedio de su sector

	/* Neutral ratings */
	neutral       Rating = "Neutral"        //no se espera que la accion tenga un desempeño significativamente bueno o malo, se mantiene en una posicion neutral
	hold          Rating = "Hold"           //mantener la accion en el portafolio actual, no se recomienda ni comprar ni vender
	equalWeight   Rating = "Equal Weight"   //similar a "Hold", indica que no se espera que la accion tenga un rendimiento superior o inferior a su sector o al mercado en general
	inLine        Rating = "In-Line"        //la accion debería tener un rendimiento similar al mercado o a su sector
	marketPerform Rating = "Market Perform" //se espera que la accion tenga un rendimiento igual al del mercado
	peerPerform   Rating = "Peer Perform"   //se espera que la accion se comporte de manera similar a otras empresas en el mismo sector
	sectorWeight  Rating = "Sector Weight"  //deberia tener un rendimiento similar al promedio de su sector o al mercado en general

	/* Cautious ratings */
	cautious Rating = "Cautious" //el analista muestra una actitud precavida, indicando que hay incertidumbre o riesgo, pero sin una recomendacion clara de venta
	negative Rating = "Negative" //se espera que la accion tenga un desempeño negativo o deficiente en el futuro cercano

	/* Negative ratings */
	underperform       Rating = "Underperform"        //se espera que la accion tenga un rendimiento inferior al de otras en el mercado
	underweight        Rating = "Underweight"         //el analista recomienda reducir la exposicion a esta accion en el portafolio, ya que se espera un rendimiento bajo
	sectorUnderperform Rating = "Sector Underperform" //se espera que la accion tenga un rendimiento inferior al promedio de su sector
	reduce             Rating = "Reduce"              //similar a "Underweight", se recomienda vender o reducir la exposicion a esta accion
	sectorPerform      Rating = "Sector Perform"      //se espera que la accion se desempeñe de manera similar al sector, sin un rendimiento destacado
	sell               Rating = "Sell"                //recomendacion de venta

)

var stringToRating = map[string]Rating{
	"":                    empty,
	"Buy":                 buy,
	"Cautious":            cautious,
	"Equal Weight":        equalWeight,
	"Hold":                hold,
	"In-Line":             inLine,
	"Market Outperform":   marketOutperform,
	"Market Perform":      marketPerform,
	"Negative":            negative,
	"Neutral":             neutral,
	"Outperform":          outperform,
	"Outperformer":        outperformer,
	"Overweight":          overweight,
	"Peer Perform":        peerPerform,
	"Positive":            positive,
	"Reduce":              reduce,
	"Sector Outperform":   sectorOutperform,
	"Sector Perform":      sectorPerform,
	"Sector Underperform": sectorUnderperform,
	"Sector Weight":       sectorWeight,
	"Sell":                sell,
	"Speculative Buy":     speculativeBuy,
	"Strong-Buy":          strongBuy,
	"Underperform":        underperform,
	"Underweight":         underweight,
}

var positiveRatings = []Rating{
	buy,
	speculativeBuy,
	strongBuy,
	outperform,
	outperformer,
	marketOutperform,
	positive,
	overweight,
	sectorOutperform,
}

var cautiousRatings = []Rating{
	cautious,
	negative,
}

var negativeRatings = []Rating{
	underperform,
	underweight,
	sectorUnderperform,
	reduce,
	sectorPerform,
	sell,
}
