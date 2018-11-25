package scale

import (
  "strings"
  //"unicode"
)

func Scale(tonic, interval string) []string {
  sharp := [12]string {
     "A","A#","B","C","C#","D","D#","E","F","F#","G","G#",
  }
  flat := [12]string {
    "A","Bb","B","C","Db","D","Eb","E","F","Gb","G","Ab",
  }

  s := map[string][12]string{
    "A": sharp,
    "Bb": flat,
    "B": sharp,
    "C": sharp,
    "Db": flat,
    "D": sharp,
    "Eb": flat,
    "E": sharp,
    "F": flat,
    "Gb": flat,
    "G": sharp,
    "Ab": flat,
    "e": sharp,
    "b": sharp,
    "f#": sharp,
    "c#": sharp,
    "g#": sharp,
    "d#": sharp,
    "d": flat,
    "g": flat,
    "c": flat,
    "f": flat,
    "bb": flat,
    "eb": flat,
    "a": sharp,

    }
  step2Number := map[string]int{
    "M": 2,
    "m": 1,
    "A": 3,
  }
    formula := s[tonic]
    i := 0
    for _, v := range formula {
      if v == strings.Title(tonic) {
        break
      }
      i++
    }

    if interval == "" {
      chromeScale := make([]string, 0, 12)
      for j := i; j < i + 12; j++ {
         // below adds the tone at the current position (j) then begins from the beginning
        chromeScale = append(chromeScale, formula[j % 12])
      }
      return chromeScale
    }
    steps := strings.Split(interval, "")
    scale := make([]string, 0, len(steps) + 1)
    for _, step := range steps {
      scale = append(scale, formula[i % 12])
      numericalStep := step2Number[step]
      i += numericalStep
    }
    return scale
}
