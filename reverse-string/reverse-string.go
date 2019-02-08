package reverse

func String(initial string) string {
  reversed := []byte{}

  for i := len(initial) -1 ; i >= 0; i-- {
    v := initial[i]
    reversed = append(reversed, v)
  }
    return string(reversed)
}
