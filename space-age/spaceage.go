package space

func Age(seconds float64, planet string) float64 {
 years := map[string]float64{
   "Earth": 1.0,
   "Mercury": 0.2408467,
   "Venus": 0.61519726,
   "Mars": 1.8808158,
   "Jupiter": 11.862615,
   "Saturn": 29.447498,
   "Uranus": 84.016846,
   "Neptune": 164.79132,

   }

   formula := years[planet]
   secondsPerYear := 31557600 * formula

  return seconds / secondsPerYear
}
