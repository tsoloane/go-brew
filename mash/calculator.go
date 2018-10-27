package mash

const DEFAULT_BOILOFF_PER_HOUR 		float64 	= 3  		//Litres per hour
const DEFAULT_TRUB_LOSS 			float64		= 2			//Litres
const DEFAULT_GRAIN_ABSORPTION 		float64		= 0.5  		//Litres per Kg
const DEFAULT_HOP_ABSORPTION_RATE 	float64		= 0.15		//Litres per gram
const DEFAULT_GRAIN_TEMPERATURE	  	float64 	= 21		//Degrees Celcius
const DEFAULT_BOIL_TIME				float64		= 60		//Minutes

type Mash struct {
	expectedVolume 		float64 	//Expected volume for the fermenter
	grainWeight 		float64		//Total weight of the grain bill
	hopWeight			float64		//Total weight of hops
	boilTime			float64		//Boil Time in minutes
	trubVolume			float64		//Volume of the trub left after boil off
	grainAbsorption 	float64		//Grain absorption rate in Litres/kgram
	hopAbsorption 		float64		//Hop Absorption in Liters/gram
	grainTemperature	float64		//Temperature of your grain in degrees celcius
	mashTemperature		float64		//Desired Mash Temperature
	kettleDiameter		float64		//Diameter of the kettle
	boilOffRate			float64		//BoilOff Rate
}

func (m Mash) ExpectedBoilLossGivenRate(rate float64) float64 {
	m.boilOffRate=rate

	return m.ExpectedBoilLoss()
}

/**
 * Calculate the expected Mash loss for this mash
 */
func (m Mash) ExpectedBoilLoss() float64 {
	if m.boilOffRate == 0 {
		m.boilOffRate = DEFAULT_BOILOFF_PER_HOUR
	}
	if m.boilTime == 0 {
		m.boilTime = DEFAULT_BOIL_TIME
	}

	return (m.boilTime/60)*m.boilOffRate
}

